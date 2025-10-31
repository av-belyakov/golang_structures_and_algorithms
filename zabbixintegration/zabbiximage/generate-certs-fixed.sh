#!/bin/bash

CERT_DIR="certs"

echo "Удаляем из директории $CERT_DIR старые ключи и сертификаты..."
if [ ! -d $CERT_DIR ]; then
    mkdir $CERT_DIR $CERT_DIR/postgresql $CERT_DIR/zabbix
else 
    rm $CERT_DIR/*
fi

echo "Генерация CA..."
openssl genrsa -out $CERT_DIR/postgresql/ca.key 2048
openssl req -new -x509 -days 3650 -key $CERT_DIR/postgresql/ca.key \
  -out $CERT_DIR/postgresql/ca.crt \
  -subj "/C=US/ST=State/L=City/O=IT/CN=Zabbix CA"

echo "Генерация сертификата для PostgreSQL сервера..."
openssl genrsa -out $CERT_DIR/postgresql/server.key 2048
openssl req -new -key $CERT_DIR/postgresql/server.key \
  -out $CERT_DIR/postgresql/server.csr \
  -subj "/C=US/ST=State/L=City/O=IT/CN=postgresql"
chown 999:999 $CERT_DIR/postgresql/server.key
chmod 600 $CERT_DIR/postgresql/server.key

openssl x509 -req -in $CERT_DIR/postgresql/server.csr \
  -CA $CERT_DIR/postgresql/ca.crt \
  -CAkey $CERT_DIR/postgresql/ca.key \
  -CAcreateserial \
  -out $CERT_DIR/postgresql/server.crt \
  -days 3650 -sha256


echo "Генерация сертификата для Zabbix клиента..."
openssl genrsa -out $CERT_DIR/zabbix/client.key 2048
openssl req -new -key $CERT_DIR/zabbix/client.key \
  -out $CERT_DIR/zabbix/client.csr \
  -subj "/C=US/ST=State/L=City/O=IT/CN=zabbix-server"

openssl x509 -req -in $CERT_DIR/zabbix/client.csr \
  -CA $CERT_DIR/postgresql/ca.crt \
  -CAkey $CERT_DIR/postgresql/ca.key \
  -CAcreateserial \
  -out $CERT_DIR/zabbix/client.crt \
  -days 3650 -sha256

echo "Копирование CA сертификата для Zabbix..."
cp $CERT_DIR/postgresql/ca.crt $CERT_DIR/zabbix/ca.crt

echo "Создание объединенных файлов сертификатов..."
cat $CERT_DIR/zabbix/client.crt $CERT_DIR/zabbix/client.key > $CERT_DIR/zabbix/client-combined.pem

echo "Настройка прав доступа..."
chmod 600 $CERT_DIR/postgresql/*.key
chmod 600 $CERT_DIR/zabbix/*.key
chmod 644 $CERT_DIR/postgresql/*.crt
chmod 644 $CERT_DIR/zabbix/*.crt
chmod 600 $CERT_DIR/zabbix/client-combined.pem

echo "Проверка сертификатов..."
openssl verify -CAfile $CERT_DIR/postgresql/ca.crt $CERT_DIR/postgresql/server.crt
openssl verify -CAfile $CERT_DIR/postgresql/ca.crt $CERT_DIR/zabbix/client.crt

echo "Сертификаты успешно сгенерированы и проверены!"