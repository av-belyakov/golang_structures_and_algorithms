#!/bin/bash

KEYS_DIR="minio/certs/"

if [ ! -d $KEYS_DIR ]; then
  mkdir $KEYS_DIR
fi

openssl genrsa -out $KEYS_DIR/private.key 2048
openssl req -new -x509 -days 365 -key $KEYS_DIR/private.key -out $KEYS_DIR/public.crt \
  -subj "/C=RU/ST=Moscow/L=Moscow/O=MyCompany/CN=localhost" \
  -addext "subjectAltName=DNS:localhost,DNS:minio,DNS:minio.local,IP:127.0.0.1,IP:192.168.9.53"
