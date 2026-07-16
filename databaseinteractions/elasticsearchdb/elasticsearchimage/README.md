# Elasticsearch image

## Настройка авторизации

### Генерирование временных паролей для авторизации Kibana в Elasticsearch

Пароли генерируются для переменных окружения

- ELASTICSEARCH_USERNAME=kibana_system
- ELASTICSEARCH_PASSWORD=

Для того что бы сгенерировать в Elasticsearch временные пароли нужно выполнить следующие команды

```bash
docker compose up -d elasticsearch
docker exec -it elasticsearch bash
```

далее внутри контейнера выполнить

```bash
bin/elasticsearch-setup-passwords auto
```

будет сгенерирован список паролей для определённого набора пользователей. Список паролей аналогичен списку в файле _tmp_elasticsearch_credentials_. Для просмотра сгенерированных пароле в docker elasticsearch, выполняем:

```bash
docker inspect <elasticsearch_container_id> | grep -i password
```

Так же более подробную информацию можно получить по ссылке https://www.elastic.co/docs/deploy-manage/security/set-up-minimal-security

### Настройка Elasticvue

Elasticvue представляет собой минималистичный веб интерфейс для управления кластерами Elasticsearch. Пример ниже
![пример-1](./images/image-1.png)

Добавить в настройку переменных окружения Elasticsearch следующие параметры:

- http.cors.enabled=true
- http.cors.allow-origin=http://localhost:8484
- http.cors.allow-headers=X-Requested-With,Content-Type,Content-Length,Authorization
- http.cors.allow-credentials=true

### Настройка снапшотов для создания бекапов в MinIO

Перед запуском контейнеров docker нужно создать публичный ключ и сертификат

```bash
./gencerts.sh
```

Для начала надо настроить Elasticsearch-источника. Добавляем учетные данные MinIO в keystore. Elasticsearch использует защищенный keystore для хранения чувствительных данных. Выполняем следующие команды:

```bash
docker exec -it <source_container_id> bash
# Добавьте Access Key и Secret Key
bin/elasticsearch-keystore add s3.client.default.access_key
bin/elasticsearch-keystore add s3.client.default.secret_key
# Здесь default — это дефолтное имя клиента, которое вы будете использовать при регистрации репозитория
```

То есть если в настройках MinIO есть MINIO_ROOT_USER=minioadmin, то это же имя нужно использовать в строке s3.client.<имя клиента>.access_key и s3.client.<имя клиента>.secret_key.

По терминологии MinIO access_key - username, secret_key - password.

Далее нужно выполнить:

```bash
curl -u elastic:jYQ758IbxEnXxF3SM0T0 -X PUT 'http://localhost:9200/_snapshot/my-minio-repository' -H 'Content-type: application/json' -d '{
  "type": "s3",
  "settings": {
    "bucket": "elasticsearch-backup",
    "client": "default",
    "endpoint": "https://minio:9000",
    "protocol": "https",
    "path_style_access": true,
    "region": "us-east-1",
    "readonly": false,
    "disable_chunked_encoding": true
  }
}'
```

должен быть получен ответ:

```json
{ "acknowledged": true }
```

Создаем снапшот всех индексов и глобального состояния кластера:

```bash
curl -u elastic:jYQ758IbxEnXxF3SM0T0 -X PUT 'http://localhost:9200/_snapshot/my-minio-repository/snapshot_1?wait_for_completion=true'
```

параметр **wait_for_completion=true** позволяет дождатся завершения выполнения задачи.

--------------------------------------------- Всё что ниже для удаления ---------------------------------------------

Changed password for user apm_system
PASSWORD apm_system = a6bXaBVEwbvCmjasl7Vz

Changed password for user kibana_system
PASSWORD kibana_system = MhaRZHuVqrpD8pPcrHeo

Changed password for user kibana
PASSWORD kibana = MhaRZHuVqrpD8pPcrHeo

Changed password for user logstash_system
PASSWORD logstash_system = 5MDvThNxGY6LCLkH8Ejd

Changed password for user beats_system
PASSWORD beats_system = vTbWDsLwcEBOgiZClFoX

Changed password for user remote_monitoring_user
PASSWORD remote_monitoring_user = W1IWVgKdzPXIPlgLqvjU

Changed password for user elastic
PASSWORD elastic = jYQ758IbxEnXxF3SM0T0
