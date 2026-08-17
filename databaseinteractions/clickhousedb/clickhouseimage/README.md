# Настройка доступа к СУБД Clickhouse

### Развёртывание простого docker контейнера Clickhouse

Для развертывания простого docker контейнера Clickhouse, без поддержки шардирования и репликации, выполнить:

```bash
docker compose -f docker-compose.simple.yml up -d
```

### Развертывание Clickhouse с поддержкой шардирования и репликации

Для координации шардирования и репликации используются координаторы Clickhouse Keeper или ZooKeeper. В данном случае для координации будет использоватся Clickhouse Keeper, так как именно его настоятельно рекомендуют использовать в официальной документации Clickhouse.
ClickHouse Keeper — это система координации, которая управляет состоянием кластера, отслеживает, какие реплики активны, и координирует процесс репликации данных между ними. Он пришел на смену ZooKeeper и рекомендуется к использованию, так как обеспечивает лучшую надежность и потребляет меньше ресурсов.
Для production-окружения настоятельно рекомендуется запускать кластер как минимум из трех узлов ClickHouse Keeper. Это необходимо для поддержания кворума — механизма, который гарантирует, что даже при выходе одного узла из строя кластер продолжит работать и принимать решения.

Также в продакшене ClickHouse Keeper лучше выносить на выделенные хосты, отдельно от серверов базы данных, чтобы можно было масштабировать и управлять ими независимо.

В продакшене настоятельно рекомендуется запускать ClickHouse Keeper на выделенных узлах (или контейнерах), отдельно от серверов баз данных.

Порядок запуска Clickhouse БД и Clickhouse-Keepre, сначала формируем кластер координаторов:

```bash
docker compose -f docker-compose.keeper.yml up -d
```

Проверить состояние кластера Clickhouse Keeper можно с помощью команды:

```bash
docker exec clickhouse-keeper-01 clickhouse-keeper-client --host 127.0.0.1 --port 9181 --query "stat"
```

пример вывода:

```bash
ClickHouse Keeper version: v26.7.3.19-stable-c3c4a420478c7e8fbbde863bba1b0dc4ae0eec6b
Clients:
 127.0.0.1:56924(recved=0,sent=0)
 127.0.0.1:56916(recved=1,sent=1)

Latency min/avg/max: 0/13/18
Received: 3
Sent: 2
Connections: 1
Outstanding: 0
Zxid: 0x3
Mode: leader
Node count: 4
```

ClickHouse Keeper успешно запустился и выбрал лидера (Mode: leader). Вот что означают эти строки:

- _Mode: leader_ — ваш узел стал лидером кластера ✅
- _Node count: 4_ — в кластере 4 узла (3 сервера Keeper + 1 корневой)
- _Clients: 127.0.0.1:56924_ — есть активное клиентское подключение
- _Zxid: 0x3_ — текущий идентификатор транзакции (ZooKeeper Transaction ID)

следом нужно запустить Clickhouse БД:

```bash
docker compose -f docker-compose.claster.yml up -d
```

Оба файла docker-compose.keeper.yml и docker-compose.claster.yml используют одну и ту же docker сеть, разница в том что docker-compose.claster.yml который должен запускатся вторым имеет настройки позволяющие использовать уже созданную docker сеть:

```yml
networks:
  clickhouse-network:
    # Используйте 'external', если сеть уже создана вашим Keeper-кластером
    # В противном случае, она будет создана
    external: true
```

Красивый вывод списка запущенных контейнеров:

```bash
docker ps --format 'table {{.ID}}\t{{.Image}}\t{{.Names}}\t{{.RunningFor}}\t{{.Status}}'
```

#### Настройка файла users.xml и добавление пользователей

Обязательно, при настроке файла users.xml добавить пользователю default пароль что бы исключить несанкционированный доступ к БД от имени этого пользователя. Для большей безопасности лучьше хранить его в виде хеша sha256. Команда для генерации хеша пароля:

```bash
echo -n "ваш_надежный_пароль" | sha256sum | tr -d '-'
```

Документация ClickHouse рекомендует использовать SQL-управление доступом через команды CREATE USER и GRANT как более гибкий способ. Для его включения нужно добавить параметр

```yml
<access management>1</access management>
```

для пользователя **default** в _users.xml_, а затем создавать и настраивать пользователей через SQL-запросы. Этот способ позволяет управлять пользователями без перезагрузки сервера.

#### Создание базы данных на всех нодах кластера

Выполняем подключение к БД через клиента clickhouse-client:

```bash
docker exec -it clickhouse-01 sh
clickhouse-client --user default --password <password>
```

Создаём базы данных на всех нодах кластера:

```bash
CREATE DATABASE IF NOT EXISTS example_database ON CLUSTER example_cluster;
```

имя кластера берём из файла настроек _cluster.xml_. Или на каждой ноде вручную:

```bash
CREATE DATABASE IF NOT EXISTS example_database example_cluster;
```

или можно создать реплицируемую базу данных:

```sql
/* это не работает, не создается на всех узлах кластера */
CREATE DATABASE IF NOT EXISTS example_database
ENGINE = Replicated(
    '/clickhouse/databases/example_database',
    '{shard}',
    '{replica}'
);

/* лучше использовать эту команду, она гарантированно создаёт БД в кластере */
CREATE DATABASE IF NOT EXISTS example_database ON CLUSTER example_cluster;
/* ключевым является параметр ON CLUSTER */
```

Просмотр списка баз данных:

```sql
SHOW DATABASES;
```

Удаление базы данных:

```sql
DROP DATABASE example_database ON CLUSTER example_cluster;
```

если используется **ON CLUSTER <имя_кластера>** то удаление базы данных будет выполнятся на всех узлах кластера.

Просмотр всех имен кластеров:

```sql
SELECT DISTINCT cluster FROM system.clusters;
```

#### Создание реплицированной таблицы

```sql
CREATE TABLE IF NOT EXISTS example_database.example_table
ON CLUSTER example_cluster
(
    order_id    UInt64,
    created_at  DateTime,
    user_id     UInt64,
    amount      Decimal(18, 2),
    status      String
)
ENGINE = ReplicatedMergeTree(
    '/clickhouse/tables/{shard}/example_database/example_table',
    '{replica}'
)
PARTITION BY toYYYYMM(created_at)
ORDER BY order_id;
```

проверить репликацию на таблице:

```sql
SELECT
    database,
    table,
    replica_name,
    is_leader,
    total_replicas,
    active_replicas
FROM system.replicas
WHERE database = 'example_table';
```

#### Добавление пользователя через запросы к БД

Для продолжения необходимо создать ещё одного, помимо default, пользователя через запрос к БД.
Выполняем подключение к БД через клиента clickhouse-client:

```bash
docker exec -it clickhouse-01 sh
clickhouse-client --user default --password <password>
```

создаём нового пользователя:

```sql
CREATE USER user_example_click IDENTIFIED BY '$CLICKHOUSE_SERVER_EXAMPLE_USER_PASSWORD';
```

проверяем:

```sql
SHOW USERS;
/* или подробнее */
SELECT * FROM system.users;
/* проверяем права доступа */
SHOW GRANTS FOR user_example_click;
/* проверяем наличие пользователя на всех узлах кластера */
SELECT
    hostName() AS host,
    name
FROM clusterAllReplicas('example_cluster', system.users)
WHERE name = 'user_example'
ORDER BY host;
```

добавим права пользователю на созданную базу данных:

```sql
GRANT SELECT, INSERT, CREATE TABLE ON example_database.* TO user_example_click ON CLUSTER example_cluster;
```

CLICKHOUSE_SERVER_EXAMPLE_USER_PASSWORD заменить на пароль из .env

Идеологически верным решением будет добавить пользователя на все узлы кластера. Это можно сделать путём добавления пользователя на каждый узел в ручную или альтернативное решение, выполнить команду с использованием параметра **ON CLUSTER <имя_кластера>**

```sql
CREATE USER IF NOT EXISTS user_example_click ON CLUSTER example_cluster IDENTIFIED BY 'mJk1-pY52-dA13-Ghx9';
```

если пользователь уже есть на одном из узлов кластера то выше указанная команда поможет добавить его на другие узлы если его там нет.

Для того что бы поменять пароль пользователя на всём кластере:

```sql
ALTER USER user_example_click ON CLUSTER example_cluster IDENTIFIED BY 'NewMySecretPassword123';
```

### Доступ к интерфейсу SQL UI

Для доступа к Веб-интерфейсу пользователя с SQL UI выполнить в браузере:

```bash
http://localhost:8123/play
```

или

```bash
http://localhost:18123/dashboard
```

для доступа к дашбордам.
