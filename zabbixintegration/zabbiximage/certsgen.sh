#!/bin/bash

CERT_DIR="certs"

echo "Удаляем из директории $CERT_DIR старые ключи и сертификаты..."
if [ ! -d $CERT_DIR ]; then
    mkdir $CERT_DIR
else 
    rm -r $CERT_DIR/*
fi

echo "Создание директорий для сертификатов..."
mkdir -p $CERT_DIR/postgresql $CERT_DIR/zabbix

echo "Генерация CA..."
openssl req -new -x509 -days 3650 -nodes \
  -out $CERT_DIR/postgresql/ca.crt \
  -keyout $CERT_DIR/postgresql/ca.key \
  -subj "/CN=Zabbix CA"

echo "Генерация сертификата для PostgreSQL сервера..."
openssl req -new -nodes \
  -out $CERT_DIR/postgresql/server.csr \
  -keyout $CERT_DIR/postgresql/server.key \
  -subj "/CN=postgresql"
chown 999:999 $CERT_DIR/postgresql/server.key
chmod 600 $CERT_DIR/postgresql/server.key

openssl x509 -req -in $CERT_DIR/postgresql/server.csr \
  -CA $CERT_DIR/postgresql/ca.crt \
  -CAkey $CERT_DIR/postgresql/ca.key \
  -CAcreateserial \
  -out $CERT_DIR/postgresql/server.crt \
  -days 3650

echo "Генерация сертификата для Zabbix клиента..."
openssl req -new -nodes \
  -out $CERT_DIR/zabbix/client.csr \
  -keyout $CERT_DIR/zabbix/client.key \
  -subj "/CN=zabbix-server"

openssl x509 -req -in $CERT_DIR/zabbix/client.csr \
  -CA $CERT_DIR/postgresql/ca.crt \
  -CAkey $CERT_DIR/postgresql/ca.key \
  -CAcreateserial \
  -out $CERT_DIR/zabbix/client.crt \
  -days 3650

echo "Копирование CA сертификата для Zabbix..."
cp $CERT_DIR/postgresql/ca.crt $CERT_DIR/zabbix/ca.crt

echo "Настройка прав доступа..."
chmod 600 $CERT_DIR/postgresql/*.key
chmod 600 $CERT_DIR/zabbix/*.key
chmod 644 $CERT_DIR/postgresql/*.crt
chmod 644 $CERT_DIR/zabbix/*.crt

echo "Сертификаты успешно сгенерированы!"