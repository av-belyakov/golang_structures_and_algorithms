#!/bin/bash

KEYS_DIR="certs"

if [ ! -d $KEYS_DIR ]; then
  mkdir $KEYS_DIR
fi

openssl genrsa -out $KEYS_DIR/private.key 2048
openssl req -new -x509 -days 365 -key $KEYS_DIR/private.key -out $KEYS_DIR/public.crt \
  -subj "/C=RU/ST=Moscow/L=Moscow/O=MyCompany/CN=localhost" \
  -addext "subjectAltName=DNS:localhost,IP:127.0.0.1,DNS:minio.local"
