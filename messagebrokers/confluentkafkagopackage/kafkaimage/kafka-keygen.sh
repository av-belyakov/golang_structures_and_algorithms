#!/bin/bash

KEY_DIR="keys"

mr -R $KEY_DIR

if [ -z "$1" ]; then
    echo "Задайте пароль для формирования ключей"
    exit 1
fi

if [ ! -d $KEY_DIR ]; then
    mkdir $KEY_DIR
fi

echo "Создание CA (Certificate Authority)"
openssl req -new -x509 -keyout $KEY_DIR/ca-key -out $KEY_DIR/ca-cert -days 365 -subj "/CN=localhost" -passout pass:$1

echo "Генерация ключа"
keytool -keystore $KEY_DIR/kafka.keystore.jks -alias localhost -validity 365 -genkey -keyalg RSA \
  -storepass $1 -keypass $1 -dname "CN=Kafka-Test"

echo "Импорт CA в truststore"
keytool -keystore $KEY_DIR/kafka.truststore.jks -alias CARoot -import -file $KEY_DIR/ca-cert \
  -storepass $1 -keypass $1 -noprompt

echo "Создание CSR"
keytool -keystore $KEY_DIR/kafka.keystore.jks -alias localhost -certreq -file $KEY_DIR/cert-file \
  -storepass $1 -keypass changeit

echo "Подписание сертификата CA"
openssl x509 -req -CA $KEY_DIR/ca-cert -CAkey $KEY_DIR/ca-key -in $KEY_DIR/cert-file \ 
  -out $KEY_DIR/cert-signed -days 365 -CAcreateserial -passin pass:$1
#openssl x509 -req -CA ca-cert -CAkey ca-key -in cert-file -out cert-signed \
#  -days 365 -CAcreateserial -passin pass:changeit

echo "Импорт CA в keystore"
keytool -keystore $KEY_DIR/kafka.keystore.jks -alias CARoot -import -file $KEY_DIR/ca-cert \
  -storepass $1 -keypass $1 -noprompt

echo "Импорт подписанного сертификата"
keytool -keystore $KEY_DIR/kafka.keystore.jks -alias localhost -import -file $KEY_DIR/cert-signed \
  -storepass $1 -keypass $1 -noprompt