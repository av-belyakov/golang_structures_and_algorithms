# Важная информация

![Go Version](https://img.shields.io/badge/Go-1.25+-00ADD8?style=flat&logo=go)](https://golang.org/)
[![Docker](https://img.shields.io/badge/Docker-Ready-2496ED?style=flat&logo=docker)](https://www.docker.com/)

## Связка NATS, Prometheus, Grafana, без SSL и авторизации

Выполняем:

```bash
docker-compose -f docker-compose.yml up -d
```

Графический интерфейс Prometheus доступен по http://localhost:9090. Статусы подключенных endpoints доступны по _Status->Target Health_, где статусы всех подключенных endpoints должны быть _green_ и _UP_. Для того что бы проверить поступают ли какие либо метрики нужно перейти _Query->Grath_ и в строке запроса написать запрос со словом _nats_, так как мы будем смотреть метрики по nats-server.
Например, можно написать **gnatsd_connz_total**, **gnatsd_connz_subscriptions** и т.д. Следует обратить внимание, что для накопления каких либо данных в Prometheus от nats-server, необходимо выполнить некоторые go тесты.

Grafana доступна по http://localhost:3000 admn:admin. После авторизации необходимо поменять пароль администратора. Вся настройка подключения к Prometheus осуществляется через пользовательский интерфейс (подробнее https://grafana.com/docs/grafana/latest/datasources/prometheus/configure/).
Однако, если коротко, _Connections->Add new connection->Prometheus->Add new data source_, где выполняется полная настройка. После успешной настройки подключённые источники данных можно посмотреть в _Connections->Data sources_.
Для создание дашборда нужно перейти _Home->Dashboards->New dashboard->Edit panel_. выбрать метрики, аналогичные метрикам из Prometheus и нажать кнопку _Run queries_.

Можно скачать официальный дашборд для Nats от Grafana по этой ссылке https://grafana.com/grafana/dashboards/2279-nats-servers/, кстати, авторизовыватся не надо, просто нажимаем _Download JSON_. Или взять из этого файла **grafana_dashboard_nats_rev1.json** (его загрузил по ссылке указанной ранее). Что бы установить дашборд выполняем _Dashbord->New->Import_ и загружаем скачанный файл json.

Большая подборка дашбордов для Grafana https://grafana.com/grafana/dashboards/.

## Связка NATS, Prometheus, Grafana без SSL и с использованием технологии JetStream

Выполнить:

```bash
docker-compose -f docker-compose.jetstream.yml up -d
```

для того чтобы использовать конфигурационный файл **nats-server-jetstream.yml** с минимальными настройками NATS реализующими технологию JetStream.

## Связка NATS, Prometheus, Grafana без SSL, но с авторизацией и аутентификацией пользователя по имени пользователя и паролю

При настройке конфигурационного файла **nats-server-auth.conf** может возникнуть необходимость зашифровать пароли пользователей добавляемые в файл. Для этой цели сначало необходимо установить вспомогательный инструмент [nats](https://github.com/nats-io/natscli?tab=readme-ov-file#installation).

```bash
go install github.com/nats-io/natscli/nats@latest
```

после чего создать зашифрованный пароль

```bash
nats server passwd
```

Для запуска docker контейнера с настройками авторизации пользователей по логину и паролю выполнить:

```bash
docker-compose -f docker-compose.auth.yml up -d
```

## Связка NATS, Prometheus, Grafana с SSL

В начале необходимо создать TSL сертификаты, для этого выполняем:

```bash
./nats-keysgen.sh
```

полсе чего в директории **certs** будут созданы необходимые сертификаты для клиента и сервера. Далее, можно запускать docker контейнеры для тестирования.

```bash
docker-compose -f docker-compose.tls.yml
```
