#!/bin/bash

KEY_DIR="keys"
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

echo "Создание CA (Certificate Authority)"
openssl req -new -x509 -keyout $KEY_DIR/ca-key -out $KEY_DIR/ca-cert -days 365 -subj "/CN=localhost" -passout pass:$KAFKA_SSL_KEY_PASSWORD
# echo "1. Creating Certificate Authority (CA)..."
# openssl req -new -x509 -keyout ca.key -out ca.crt -days 365 \
#  -subj "/CN=Kafka-KRaft-CA" -passout pass:changeit

echo "Генерация ключа"
keytool -keystore $KEY_DIR/kafka.keystore.jks -alias $ALIAS_KEY_STORE -validity 365 -genkey -keyalg RSA -storepass $KAFKA_SSL_KEYSTORE_PASSWORD -keypass $KAFKA_SSL_KEYSTORE_PASSWORD -dname "CN=Kafka-Test"

echo "Импорт CA в truststore"
keytool -keystore $KEY_DIR/kafka.truststore.jks -alias $ALIAS_CA_ROOT -import -file $KEY_DIR/ca-cert -storepass $KAFKA_SSL_TRUSTSTORE_PASSWORD -keypass $KAFKA_SSL_TRUSTSTORE_PASSWORD -noprompt

echo "Создание CSR"
keytool -keystore $KEY_DIR/kafka.keystore.jks -alias $ALIAS_KEY_STORE -certreq -file $KEY_DIR/cert-file -storepass $KAFKA_SSL_KEYSTORE_PASSWORD -keypass $KAFKA_SSL_KEYSTORE_PASSWORD

echo "Подписание сертификата CA"
openssl x509 -req -CA $KEY_DIR/ca-cert -CAkey $KEY_DIR/ca-key -in $KEY_DIR/cert-file -out $KEY_DIR/cert-signed -days 365 -CAcreateserial -passin pass:$KAFKA_SSL_KEY_PASSWORD

echo "Импорт CA в keystore"
keytool -keystore $KEY_DIR/kafka.keystore.jks -alias $ALIAS_CA_ROOT -import -file $KEY_DIR/ca-cert -storepass $KAFKA_SSL_KEYSTORE_PASSWORD -keypass $KAFKA_SSL_KEYSTORE_PASSWORD -noprompt

echo "Импорт подписанного сертификата"
keytool -keystore $KEY_DIR/kafka.keystore.jks -alias $ALIAS_KEY_STORE -import -file $KEY_DIR/cert-signed -storepass $KAFKA_SSL_KEYSTORE_PASSWORD -keypass $KAFKA_SSL_KEYSTORE_PASSWORD -noprompt
