#!/bin/bash

ALIAS_KEY_STORE=$(cat /dev/urandom | tr -dc 'a-zA-Z0-9' | fold -w 10 | head -n 1)
ALIAS_CLIENT_KEY_STORE=$(cat /dev/urandom | tr -dc 'a-zA-Z0-9' | fold -w 10 | head -n 1)

if [ ! -f .env ]; then
    echo "⚠️  Необходимо создать файл .env со следующими переменными окружения:"
    echo ""
    echo "CERT_DIR=<директория с сертификатами>"
    echo "KAFKA_SSL_KEY_PASSWORD=<пароль для приватного>"
    echo "KAFKA_SSL_KEYSTORE_PASSWORD=<пароль для keystore kafka.keystore.jks>"
    echo "KAFKA_SSL_TRUSTSTORE_PASSWORD=<пароль для truststore kafka.truststore.jks>"
    echo ""
    exit 1
fi

# Экспортируем переменные окружения
export $(grep -v '^#' .env | grep -v '^$' | grep '=' | sed 's/"//g' | xargs)

echo "Удаляем из директории $CERT_DIR старые ключи и сертификаты..."
rm -R $CERT_DIR
if [ ! -d $CERT_DIR ]; then
    mkdir $CERT_DIR
fi

echo "ALIAS_KEY_STORE: $ALIAS_KEY_STORE"
echo "ALIAS_CLIENT_KEY_STORE: $ALIAS_CLIENT_KEY_STORE"

echo "Создание конфигурационного файла для OpenSSL..."
cat > $CERT_DIR/openssl.cnf << EOF
[req]
distinguished_name = req_distinguished_name
x509_extensions = v3_req
prompt = no

[req_distinguished_name]
C = RU
ST = State
L = Moscow
O = Organization
OU = Organizational Unit
CN = localhost

[v3_req]
keyUsage = keyEncipherment, dataEncipherment, digitalSignature
extendedKeyUsage = serverAuth, clientAuth
subjectAltName = @alt_names

[alt_names]
DNS.1 = localhost
DNS.2 = kafka
IP.1 = 127.0.0.1
IP.2 = 172.19.0.2
EOF

echo "Генерация CA..."
openssl req -new -x509 -keyout $CERT_DIR/ca.key -out $CERT_DIR/ca.crt -days 365 \
  -subj "/C=RU/ST=State/L=Moscow/O=Organization/CN=CA" \
  -passout pass:$KAFKA_SSL_KEY_PASSWORD

echo "Генерация сертификата Kafka с SAN..."
openssl req -new -keyout $CERT_DIR/kafka-key -out $CERT_DIR/kafka-cert-req -days 365 \
  -config $CERT_DIR/openssl.cnf -passout pass:$KAFKA_SSL_KEY_PASSWORD

openssl x509 -req -in $CERT_DIR/kafka-cert-req -CA $CERT_DIR/ca.crt -CAkey $CERT_DIR/ca.key \
  -out $CERT_DIR/kafka-cert -days 365 -CAcreateserial -extensions v3_req \
  -extfile $CERT_DIR/openssl.cnf -passin pass:$KAFKA_SSL_KEY_PASSWORD

echo "Конвертация в JKS (для Kafka)..."
openssl pkcs12 -export -out $CERT_DIR/kafka.p12 -inkey $CERT_DIR/kafka-key -in $CERT_DIR/kafka-cert \
  -passin pass:$KAFKA_SSL_KEY_PASSWORD -passout pass:$KAFKA_SSL_KEY_PASSWORD -name kafka

keytool -importkeystore -srckeystore $CERT_DIR/kafka.p12 -srcstoretype PKCS12 \
  -destkeystore $CERT_DIR/kafka.keystore.jks -deststorepass $KAFKA_SSL_KEY_PASSWORD \
  -srcstorepass $KAFKA_SSL_KEY_PASSWORD

keytool -keystore $CERT_DIR/kafka.truststore.jks -alias $ALIAS_KEY_STORE -import -file $CERT_DIR/ca.crt \
  -storepass $KAFKA_SSL_KEY_PASSWORD -noprompt

echo "Клиентские сертификаты в JKS..."
openssl req -new -keyout $CERT_DIR/client-key -out $CERT_DIR/client-cert-req -days 365 \
  -subj "/C=RU/ST=State/L=Moscow/O=Organization/CN=client" -passout pass:$KAFKA_SSL_KEY_PASSWORD

openssl x509 -req -in $CERT_DIR/client-cert-req -CA $CERT_DIR/ca.crt -CAkey $CERT_DIR/ca.key \
  -out $CERT_DIR/client-cert -days 365 -CAcreateserial -passin pass:$KAFKA_SSL_KEY_PASSWORD

openssl pkcs12 -export -out $CERT_DIR/client.p12 -inkey $CERT_DIR/client-key -in $CERT_DIR/client-cert \
  -passin pass:$KAFKA_SSL_KEY_PASSWORD -passout pass:$KAFKA_SSL_KEY_PASSWORD -name client

keytool -importkeystore -srckeystore $CERT_DIR/client.p12 -srcstoretype PKCS12 \
  -destkeystore $CERT_DIR/client.keystore.jks -deststorepass $KAFKA_SSL_KEY_PASSWORD \
  -srcstorepass $KAFKA_SSL_KEY_PASSWORD

cp $CERT_DIR/kafka.truststore.jks $CERT_DIR/client.truststore.jks

openssl pkcs12 -in $CERT_DIR/client.p12 -out $CERT_DIR/client-cert.pem -clcerts -nokeys -passin pass:$KAFKA_SSL_KEY_PASSWORD
openssl pkcs12 -in $CERT_DIR/client.p12 -out $CERT_DIR/client-key.pem -nocerts -nodes -passin pass:$KAFKA_SSL_KEY_PASSWORD

echo ""
echo "=== Инфраструктура открытых ключей (PKI) успешно создана ==="
echo "Root CA:          $CERT_DIR/ca.crt"
echo "Client private:   $CERT_DIR/client-key.pem"
echo "Client public:    $CERT_DIR/client-cert.pem"
echo "Client keystore: $CERT_DIR/client.keystore.jks (password: $KAFKA_SSL_KEY_PASSWORD)"
echo "Client truststore: $CERT_DIR/client.truststore.jks (password: $KAFKA_SSL_KEY_PASSWORD)"
echo "Server keystore:  $CERT_DIR/kafka.keystore.jks (password: $KAFKA_SSL_KEY_PASSWORD)"
echo "Server truststore: $CERT_DIR/kafka.truststore.jks (password: $KAFKA_SSL_KEY_PASSWORD)"