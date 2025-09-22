#!/bin/bash

KEY_DIR="certs"
ALIAS_KEY_ROOT=$(cat /dev/urandom | tr -dc 'a-zA-Z0-9' | fold -w 10 | head -n 1)
ALIAS_KEY_STORE=$(cat /dev/urandom | tr -dc 'a-zA-Z0-9' | fold -w 10 | head -n 1)
ALIAS_CLIENT_KEY_STORE=$(cat /dev/urandom | tr -dc 'a-zA-Z0-9' | fold -w 10 | head -n 1)

if [ ! -f .env ]; then
    echo "⚠️  Необходимо создать файл .env со следующими переменными окружения:"
    echo ""
    echo "KAFKA_SSL_KEY_PASSWORD=<пароль для приватного>"
    echo "KAFKA_SSL_KEYSTORE_PASSWORD=<пароль для keystore kafka.keystore.jks>"
    echo "KAFKA_SSL_TRUSTSTORE_PASSWORD=<пароль для truststore kafka.truststore.jks>"
    echo ""
    exit 1
fi

# Экспортируем переменные окружения
export $(grep -v '^#' .env | grep -v '^$' | grep '=' | sed 's/"//g' | xargs)

echo "Удаляем старые ключи и сертификаты из директории $KEY_DIR..."
rm -R $KEY_DIR
if [ ! -d $KEY_DIR ]; then
    mkdir $KEY_DIR
fi

echo "Генерация CA..."
openssl req -new -x509 -keyout $KEY_DIR/ca-key -out $KEY_DIR/ca-cert -days 365 -subj "/CN=localhost" -passout pass:$KAFKA_SSL_KEY_PASSWORD

echo "Генерация Kafka keystore..."
keytool -keystore $KEY_DIR/kafka.keystore.jks -alias $ALIAS_KEY_STORE -validity 365 -genkey -keyalg RSA \
  -storepass $KAFKA_SSL_KEY_PASSWORD -keypass $KAFKA_SSL_KEY_PASSWORD -dname "CN=localhost"

keytool -keystore $KEY_DIR/kafka.keystore.jks -alias $ALIAS_KEY_STORE -certreq -file $KEY_DIR/cert-file \
  -storepass $KAFKA_SSL_KEY_PASSWORD
openssl x509 -req -CA $KEY_DIR/ca-cert -CAkey $KEY_DIR/ca-key -in $KEY_DIR/cert-file -out $KEY_DIR/cert-signed \
  -days 365 -CAcreateserial -passin pass:$KAFKA_SSL_KEY_PASSWORD

keytool -keystore $KEY_DIR/kafka.keystore.jks -alias $ALIAS_KEY_ROOT -import -file $KEY_DIR/ca-cert \
  -storepass $KAFKA_SSL_KEY_PASSWORD -noprompt
keytool -keystore $KEY_DIR/kafka.keystore.jks -alias $ALIAS_KEY_STORE -import -file $KEY_DIR/cert-signed \
  -storepass $KAFKA_SSL_KEY_PASSWORD -noprompt

echo "Генерация Kafka truststore..."
keytool -keystore $KEY_DIR/kafka.truststore.jks -alias $ALIAS_KEY_ROOT -import -file $KEY_DIR/ca-cert \
  -storepass $KAFKA_SSL_KEY_PASSWORD -noprompt

echo "Генерация клиент keystore для kafka-ui..."
keytool -keystore $KEY_DIR/client.keystore.jks -alias $ALIAS_CLIENT_KEY_STORE -validity 365 -genkey -keyalg RSA \
  -storepass $KAFKA_SSL_KEY_PASSWORD -keypass $KAFKA_SSL_KEY_PASSWORD -dname "CN=client"

keytool -keystore $KEY_DIR/client.keystore.jks -alias $ALIAS_CLIENT_KEY_STORE -certreq -file $KEY_DIR/client-cert-file \
  -storepass $KAFKA_SSL_KEY_PASSWORD
openssl x509 -req -CA $KEY_DIR/ca-cert -CAkey $KEY_DIR/ca-key -in $KEY_DIR/client-cert-file -out $KEY_DIR/client-cert-signed \
  -days 365 -CAcreateserial -passin pass:$KAFKA_SSL_KEY_PASSWORD

keytool -keystore $KEY_DIR/client.keystore.jks -alias $ALIAS_KEY_ROOT -import -file $KEY_DIR/ca-cert \
  -storepass $KAFKA_SSL_KEY_PASSWORD -noprompt
keytool -keystore $KEY_DIR/client.keystore.jks -alias $ALIAS_CLIENT_KEY_STORE -import -file $KEY_DIR/client-cert-signed \
  -storepass $KAFKA_SSL_KEY_PASSWORD -noprompt

cp $KEY_DIR/kafka.truststore.jks $KEY_DIR/client.truststore.jks

echo "Преобразование truststore JKS в PEM"
keytool -exportcert -alias $ALIAS_KEY_ROOT -keystore $KEY_DIR/client.truststore.jks -rfc -file $KEY_DIR/ca-cert -storepass $KAFKA_SSL_KEY_PASSWORD

echo "Преобразование keystore JKS в PEM"
keytool -importkeystore -srckeystore $KEY_DIR/client.keystore.jks \
  -destkeystore $KEY_DIR/client.p12 -deststoretype PKCS12 \
  -srcstorepass $KAFKA_SSL_KEY_PASSWORD -deststorepass $KAFKA_SSL_KEY_PASSWORD

openssl pkcs12 -in $KEY_DIR/client.p12 -out $KEY_DIR/client-cert.pem -clcerts -nokeys -passin pass:$KAFKA_SSL_KEY_PASSWORD
openssl pkcs12 -in $KEY_DIR/client.p12 -out $KEY_DIR/client-key.pem -nocerts -nodes -passin pass:$KAFKA_SSL_KEY_PASSWORD