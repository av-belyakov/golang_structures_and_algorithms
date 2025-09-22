# Инструкция по запуску Kafka с SSL

### Генерация ключей и сертификатов

Заменить в .env файле KAFKA_SSL_KEY_PASSWORD на свой.
Выполнить:

```bash
./kafka-certsgen.sh
```

Из всех созданных, в результате работы скрипта kafka-keygen.sh, файлов для Kafka нужны kafka.keystore.jks и kafka.truststore.jks, для kafka-go ca.crt
