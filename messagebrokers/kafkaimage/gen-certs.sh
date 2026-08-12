#!/usr/bin/env bash
set -euo pipefail # «строгий режим»

if [ ! -f .env ]; then
    echo "⚠️  Необходимо создать файл .env со следующими переменными окружения:"
    echo ""
    echo "CERT_DIR=<директория с сертификатами>"
    echo "EXTERNAL_KAFKA_HOST=<внешний хост kafka, по умолчанию localhost>"
    echo "KAFKA_SSL_KEY_PASSWORD=<пароль для приватного хранилища>"
    echo ""
    exit 1
fi

DAYS=3650
EXTERNAL_HOST=${EXTERNAL_KAFKA_HOST:-localhost}

# Экспортируем переменные окружения
export $(grep -v '^#' .env | grep -v '^$' | grep '=' | sed 's/"//g' | xargs)

SAN="DNS:kafka,DNS:kafka-claster,DNS:kafka-claster-test,DNS:localhost,IP:127.0.0.1"

if [[ "$EXTERNAL_HOST" != "localhost" && "$EXTERNAL_HOST" != "127.0.0.1" ]]; then
  if [[ "$EXTERNAL_HOST" =~ ^([0-9]{1,3}\.){3}[0-9]{1,3}$ ]]; then
    SAN="$SAN,IP:$EXTERNAL_HOST"
  else
    SAN="$SAN,DNS:$EXTERNAL_HOST"
  fi
fi

rm -rf "$CERT_DIR"
mkdir -p "$CERT_DIR"

# CA
keytool -genkeypair \
  -alias ca \
  -keyalg RSA \
  -keysize 2048 \
  -dname CN=KafkaCA \
  -validity "$DAYS" \
  -keystore "$CERT_DIR/ca.keystore.jks" \
  -storepass "$KAFKA_SSL_KEY_PASSWORD" \
  -keypass "$KAFKA_SSL_KEY_PASSWORD" \
  -ext bc:c:critical

keytool -exportcert \
  -alias ca \
  -keystore "$CERT_DIR/ca.keystore.jks" \
  -storepass "$KAFKA_SSL_KEY_PASSWORD" \
  -file "$CERT_DIR/ca.crt" \
  -rfc

# Server keystore для Kafka broker
keytool -genkeypair \
  -alias kafka \
  -keyalg RSA \
  -keysize 2048 \
  -dname CN=kafka \
  -validity "$DAYS" \
  -keystore "$CERT_DIR/kafka.server.keystore.jks" \
  -storepass "$KAFKA_SSL_KEY_PASSWORD" \
  -keypass "$KAFKA_SSL_KEY_PASSWORD"

keytool -certreq \
  -alias kafka \
  -keystore "$CERT_DIR/kafka.server.keystore.jks" \
  -storepass "$KAFKA_SSL_KEY_PASSWORD" \
  -file "$CERT_DIR/kafka.csr"

keytool -gencert \
  -alias ca \
  -keystore "$CERT_DIR/ca.keystore.jks" \
  -storepass "$KAFKA_SSL_KEY_PASSWORD" \
  -infile "$CERT_DIR/kafka.csr" \
  -outfile "$CERT_DIR/kafka.crt" \
  -validity "$DAYS" \
  -rfc \
  -ext SAN="$SAN" \
  -ext EKU=serverAuth,clientAuth

keytool -importcert \
  -alias ca \
  -keystore "$CERT_DIR/kafka.server.keystore.jks" \
  -storepass "$KAFKA_SSL_KEY_PASSWORD" \
  -file "$CERT_DIR/ca.crt" \
  -noprompt

keytool -importcert \
  -alias kafka \
  -keystore "$CERT_DIR/kafka.server.keystore.jks" \
  -storepass "$KAFKA_SSL_KEY_PASSWORD" \
  -file "$CERT_DIR/kafka.crt" \
  -noprompt

# Truststore для Kafka broker
keytool -importcert \
  -alias ca \
  -keystore "$CERT_DIR/kafka.server.truststore.jks" \
  -storepass "$KAFKA_SSL_KEY_PASSWORD" \
  -file "$CERT_DIR/ca.crt" \
  -noprompt

# Truststore для клиентов: kafbat-ui и внешних Kafka-клиентов
cp "$CERT_DIR/kafka.server.truststore.jks" "$CERT_DIR/client.truststore.jks"

chmod 755 "$CERT_DIR"
chmod 644 "$CERT_DIR"/*.jks "$CERT_DIR"/*.crt || true

echo "Certificates generated in $CERT_DIR"