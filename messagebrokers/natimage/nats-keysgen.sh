#!/bin/bash

KEYS_DIR="certs"

echo "Удаляем из директории $CERT_DIR старые ключи и сертификаты..."
if [ ! -d $KEYS_DIR ]; then
    mkdir $KEYS_DIR
else 
    rm $KEYS_DIR/*
fi

# Генерация CA
openssl genrsa -out $KEYS_DIR/rootCA-key.pem 2048
openssl req -x509 -new -nodes -key $KEYS_DIR/rootCA-key.pem -sha256 -days 3650 -out $KEYS_DIR/rootCA.pem -subj "/C=RU/ST=Russia/L=Moscow/O=Org/CN=Local Root CA"

# Генерация серверного сертификата
openssl genrsa -out $KEYS_DIR/server-key.pem 2048
openssl req -new -key $KEYS_DIR/server-key.pem -out $KEYS_DIR/server.csr -subj "/C=RU/ST=Russia/L=Moscow/O=Org/CN=localhost"

cat > $KEYS_DIR/server.ext << EOF
authorityKeyIdentifier=keyid,issuer
basicConstraints=CA:FALSE
keyUsage = digitalSignature, nonRepudiation, keyEncipherment, dataEncipherment
subjectAltName = @alt_names
[alt_names]
DNS.1 = localhost
DNS.2 = nats-server
IP.1 = 127.0.0.1
EOF

openssl x509 -req -in $KEYS_DIR/server.csr -CA $KEYS_DIR/rootCA.pem -CAkey $KEYS_DIR/rootCA-key.pem -CAcreateserial -out $KEYS_DIR/server-cert.pem -days 365 -sha256 -extfile $KEYS_DIR/server.ext

echo "Генерация клиентских сертификатов..."
openssl genrsa -out $KEYS_DIR/client-key.pem 2048
openssl req -new -key $KEYS_DIR/client-key.pem -out $KEYS_DIR/client.csr -subj "/C=RU/ST=Russia/L=Moscow/O=Org/CN=Test Client"

cat > $KEYS_DIR/client.ext << EOF
basicConstraints=CA:FALSE
keyUsage = digitalSignature, keyEncipherment
extendedKeyUsage = clientAuth
EOF

openssl x509 -req -in $KEYS_DIR/client.csr -CA $KEYS_DIR/rootCA.pem -CAkey $KEYS_DIR/rootCA-key.pem -CAcreateserial -out $KEYS_DIR/client-cert.pem -days 365 -sha256 -extfile $KEYS_DIR/client.ext

# Очистка
rm $KEYS_DIR/server.csr $KEYS_DIR/server.ext $KEYS_DIR/rootCA-key.pem $KEYS_DIR/client.csr $KEYS_DIR/client.ext

echo "Сертификаты успешно сгенерированы!"
echo "rootCA.pem - корневой сертификат"
echo "server-key.pem - приватный ключ сервера"
echo "server-cert.pem - сертификат сервера"