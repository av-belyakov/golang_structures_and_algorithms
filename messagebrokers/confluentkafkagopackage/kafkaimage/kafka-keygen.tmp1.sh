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

echo "1. Creating Root Certificate Authority..."
openssl genrsa -out $KEY_DIR/ca.key 2048
openssl req -new -x509 -key $KEY_DIR/ca.key -out $KEY_DIR/ca.crt -days 365 \
  -subj "/C=RU/ST=State/L=Moscow/O=MyCompany/CN=Kafka-Root-CA"

echo "2. Creating Server Certificate..."
openssl genrsa -out $KEY_DIR/server.key 2048

# Конфиг для сервера с SAN
cat > $KEY_DIR/server.cnf << EOF
[req]
distinguished_name = req_distinguished_name
req_extensions = v3_req
prompt = no

[req_distinguished_name]
C = RU
ST = State
L = Moscow
O = MyCompany
CN = kafka

[v3_req]
keyUsage = keyEncipherment, dataEncipherment, digitalSignature
extendedKeyUsage = serverAuth
subjectAltName = @alt_names

[alt_names]
DNS.1 = kafka
DNS.2 = localhost
DNS.3 = 127.0.0.1
IP.1 = 127.0.0.1
EOF

openssl req -new -key $KEY_DIR/server.key -out $KEY_DIR/server.csr -config $KEY_DIR/server.cnf
openssl x509 -req -in $KEY_DIR/server.csr -CA $KEY_DIR/ca.crt -CAkey $KEY_DIR/ca.key \
  -CAcreateserial -out $KEY_DIR/server.crt -days 365 -extensions v3_req -extfile $KEY_DIR/server.cnf

echo "3. Creating Client Certificate..."
openssl genrsa -out $KEY_DIR/client.key 2048

# Конфиг для клиента
cat > $KEY_DIR/client.cnf << EOF
[req]
distinguished_name = req_distinguished_name
req_extensions = v3_req
prompt = no

[req_distinguished_name]
C = RU
ST = State
L = Moscow
O = MyCompany
CN = kafka-client

[v3_req]
keyUsage = digitalSignature, keyEncipherment
extendedKeyUsage = clientAuth
EOF

openssl req -new -key $KEY_DIR/client.key -out $KEY_DIR/client.csr -config $KEY_DIR/client.cnf
openssl x509 -req -in $KEY_DIR/client.csr -CA $KEY_DIR/ca.crt -CAkey $KEY_DIR/ca.key \
  -CAcreateserial -out $KEY_DIR/client.crt -days 365 -extensions v3_req -extfile $KEY_DIR/client.cnf

echo "4. Creating JKS stores for Kafka broker..."
# Серверный keystore (содержит server certificate + private key)
openssl pkcs12 -export -in $KEY_DIR/server.crt -inkey $KEY_DIR/server.key \
  -out $KEY_DIR/server.p12 -name server -password pass:$KAFKA_SSL_KEY_PASSWORD
keytool -importkeystore -srckeystore $KEY_DIR/server.p12 -srcstoretype PKCS12 \
  -srcstorepass $KAFKA_SSL_KEY_PASSWORD -destkeystore $KEY_DIR/kafka.keystore.jks \
  -deststorepass $KAFKA_SSL_KEY_PASSWORD -destkeypass $KAFKA_SSL_KEY_PASSWORD

# Серверный truststore (содержит CA для проверки клиентов)
keytool -keystore $KEY_DIR/kafka.truststore.jks -alias $ALIAS_KEY_STORE -import \
  -file $KEY_DIR/ca.crt -storepass $KAFKA_SSL_KEY_PASSWORD -noprompt

echo "5. Verifying certificate chain..."
echo "Server certificate:"
openssl verify -CAfile $KEY_DIR/ca.crt $KEY_DIR/server.crt

echo "Client certificate:"
openssl verify -CAfile $KEY_DIR/ca.crt $KEY_DIR/client.crt

echo "=== PKI infrastructure created successfully! ==="
echo "Root CA:          certs/ca.crt"
echo "Server private:   certs/server.key"
echo "Server public:    certs/server.crt"
echo "Client private:   certs/client.key"
echo "Client public:    certs/client.crt"
echo "Server keystore:  certs/kafka.keystore.jks (password: $KAFKA_SSL_KEY_PASSWORD)"
echo "Server truststore: certs/kafka.truststore.jks (password: $KAFKA_SSL_KEY_PASSWORD)"