#!/bin/bash

KEY_DIR="certs"
ALIAS_KEY_STORE=$(cat /dev/urandom | tr -dc 'a-zA-Z0-9' | fold -w 10 | head -n 1)
ALIAS_CA_ROOT=$(cat /dev/urandom | tr -dc 'a-zA-Z0-9' | fold -w 10 | head -n 1)

if [ ! -f .env ]; then
    echo "⚠️  Необходимо создать файл .env со следующими переменными окружения:"
    echo ""
    echo "KAFKA_SSL_KEYSTORE_PASSWORD=<пароль для keystore kafka.keystore.jks>"
    echo "KAFKA_SSL_KEY_PASSWORD=<пароль для приватного ключа в keystore>"
    echo "KAFKA_SSL_TRUSTSTORE_PASSWORD=<пароль для truststore kafka.truststore.jks>"
    echo ""
    exit 1
fi

# Экспортируем переменные окружения
export $(grep -v '^#' .env | grep -v '^$' | grep '=' | sed 's/"//g' | xargs)

rm -R $KEY_DIR
if [ ! -d $KEY_DIR ]; then
    mkdir $KEY_DIR
fi

echo "ALIAS_KEY_STORE: $ALIAS_KEY_STORE"
echo "ALIAS_CA_ROOT: $ALIAS_CA_ROOT"

echo "1. Creating Certificate Authority (CA)..."
openssl req -new -x509 -keyout $KEY_DIR/ca.key -out $KEY_DIR/ca.crt -days 365 \
  -subj "/CN=Kafka-KRaft-CA" -passout pass:$KAFKA_SSL_KEY_PASSWORD

echo "2. Creating server certificate..."
# Генерация приватного ключа для сервера
openssl genrsa -out $KEY_DIR/server.key 2048

# Создание CSR для сервера
openssl req -new -key $KEY_DIR/server.key -out $KEY_DIR/server.csr \
  -subj "/CN=kafka-server"

# Подписание сертификата сервера
openssl x509 -req -in $KEY_DIR/server.csr -CA $KEY_DIR/ca.crt -CAkey $KEY_DIR/ca.key \
  -CAcreateserial -out $KEY_DIR/server.crt -days 365 -passin pass:$KAFKA_SSL_KEY_PASSWORD

echo "3. Creating client certificate for Go application..."
# Генерация приватного ключа для клиента
openssl genrsa -out $KEY_DIR/client.key 2048

# Создание CSR для клиента
openssl req -new -key $KEY_DIR/client.key -out $KEY_DIR/client.csr \
  -subj "/CN=kafka-client"

# Подписание сертификата клиента
openssl x509 -req -in $KEY_DIR/client.csr -CA $KEY_DIR/ca.crt -CAkey $KEY_DIR/ca.key \
  -CAcreateserial -out $KEY_DIR/client.crt -days 365 -passin pass:$KAFKA_SSL_KEY_PASSWORD

echo "4. Creating PKCS12 bundle for Java clients (optional)..."
# Создание PKCS12 для обратной совместимости
openssl pkcs12 -export -in $KEY_DIR/server.crt -inkey $KEY_DIR/server.key \
  -out $KEY_DIR/server.p12 -name server -password pass:$KAFKA_SSL_KEY_PASSWORD

openssl pkcs12 -export -in $KEY_DIR/client.crt -inkey $KEY_DIR/client.key \
  -out $KEY_DIR/client.p12 -name client -password pass:$KAFKA_SSL_KEY_PASSWORD

echo "5. Creating JKS stores for Kafka broker..."
# Создание keystore для Kafka broker
keytool -importkeystore -srckeystore $KEY_DIR/server.p12 -srcstoretype PKCS12 \
  -srcstorepass $KAFKA_SSL_KEYSTORE_PASSWORD -destkeystore $KEY_DIR/kafka.keystore.jks \
  -deststorepass $KAFKA_SSL_KEYSTORE_PASSWORD -destkeypass $KAFKA_SSL_KEYSTORE_PASSWORD

# Создание truststore
keytool -keystore $KEY_DIR/kafka.truststore.jks -alias $ALIAS_KEY_STORE -import \
  -file $KEY_DIR/ca.crt -storepass $KAFKA_SSL_TRUSTSTORE_PASSWORD -keypass $KAFKA_SSL_TRUSTSTORE_PASSWORD -noprompt

echo "6. Creating client truststore..."
keytool -keystore $KEY_DIR/client.truststore.jks -alias $ALIAS_KEY_STORE -import \
  -file $KEY_DIR/ca.crt -storepass $KAFKA_SSL_TRUSTSTORE_PASSWORD -keypass $KAFKA_SSL_TRUSTSTORE_PASSWORD -noprompt