# Инструкция по запуску Kafka с SSL из Docker образа

### Генерация ключей и сертификатов

Заменить в .env файле KAFKA_SSL_KEY_PASSWORD на свой.
Изменить, при необходимости, ID кластера Kafka, параметр KAFKA_CLUSTER_ID.
Выполнить:

```bash
./kafka-keysgen.sh
```

Из всех созданных, в результате работы скрипта 'kafka-keysgen.sh', файлов нужны следующие:

1. Для Kafka:

- kafka.keystore.jks;
- kafka.truststore.jks.

2. Для Kafbat-UI:

- client.keystore.jks;
- client.truststore.jks.

3. Для пакета github.com/segmentio/kafka-go:

- если в docker-compose параметр для настройки Kafka KAFKA_CFG_SSL_CLIENT_AUTH: none, то нужен только файл ca.crt;
- если в docker-compose параметр для настройки Kafka KAFKA_CFG_SSL_CLIENT_AUTH: required, то нужны следующие файлы: ca.crt, client-cert.pem, client-key.pem.

### Запуск контейнеров с Kafka и Kafbat-UI

Необходимо выполнить:

```bash
docker-compose -f docker-compose.ssl.yml up -d
```

если нужны контейнеры с SSL без дополнительной аутентификации.

Или выполнить:

```bash
docker-compose -f docker-compose.sslaut.yml up -d
```

если нужны контейнеры с SSL и дополнительной аутентификаций.

### Подсказки по работе с SSL сертификатами

Для того что бы проверить CA (Certificate Authority) выполняем

```bash
openssl s_client -connect localhost:9093 -CAfile ./certs/ca.crt </dev/null
```

так же можно посмотреть сертификат отдельно

```bash
openssl verify -CAfile ./certs/ca.crt ./certs/server.crt
```

посмотреть список допустимых IP-адресов для сертификата

```bash
openssl x509 -in ./certs/server.crt -text -noout | grep -A5 "Subject Alternative Name"
```
