#!/bin/bash

CERT_DIR="certs"

echo "Удаляем из директории $CERT_DIR старые ключи и сертификаты..."
if [ ! -d $CERT_DIR ]; then
    mkdir $CERT_DIR
else 
    rm $CERT_DIR/*
fi

# Генерация CA (корневого сертификата)
openssl genrsa -out $CERT_DIR/rootCA-key.pem 2048
openssl req -x509 -new -nodes -key $CERT_DIR/rootCA-key.pem -sha256 -days 3650 -out $CERT_DIR/rootCA.pem -subj "/C=RU/ST=Russia/L=Moscow/O=Org/CN=Local Root CA"

# Генерация серверного сертификата
openssl genrsa -out $CERT_DIR/server-key.pem 2048
openssl req -new -key $CERT_DIR/server-key.pem -out $CERT_DIR/server.csr -subj "/C=RU/ST=Russia/L=Moscow/O=Org/CN=localhost"

cat > $CERT_DIR/server.ext << EOF
authorityKeyIdentifier=keyid,issuer
basicConstraints=CA:FALSE
keyUsage = digitalSignature, nonRepudiation, keyEncipherment, dataEncipherment
subjectAltName = @alt_names
[alt_names]
DNS.1 = localhost
DNS.2 = nats-server
IP.1 = 127.0.0.1
EOF

openssl x509 -req -in $CERT_DIR/server.csr -CA $CERT_DIR/rootCA.pem -CAkey $CERT_DIR/rootCA-key.pem -CAcreateserial -out $CERT_DIR/server-cert.pem -days 365 -sha256 -extfile $CERT_DIR/server.ext

echo "Генерация клиентских сертификатов..."
openssl genrsa -out $CERT_DIR/client-key.pem 2048
openssl req -new -key $CERT_DIR/client-key.pem -out $CERT_DIR/client.csr -subj "/C=RU/ST=Russia/L=Moscow/O=Org/CN=Test Client"

cat > $CERT_DIR/client.ext << EOF
basicConstraints=CA:FALSE
keyUsage = digitalSignature, keyEncipherment
extendedKeyUsage = clientAuth
EOF

openssl x509 -req -in $CERT_DIR/client.csr -CA $CERT_DIR/rootCA.pem -CAkey $CERT_DIR/rootCA-key.pem -CAcreateserial -out $CERT_DIR/client-cert.pem -days 365 -sha256 -extfile $CERT_DIR/client.ext

# Очистка
rm $CERT_DIR/server.csr $CERT_DIR/server.ext $CERT_DIR/rootCA-key.pem $CERT_DIR/client.csr $CERT_DIR/client.ext

echo "Сертификаты успешно сгенерированы!"
echo "rootCA.pem - корневой сертификат"
echo "server-key.pem - приватный ключ сервера"
echo "server-cert.pem - сертификат сервера"