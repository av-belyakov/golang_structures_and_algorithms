#!/bin/bash

openssl genrsa -out certs/private.key 2048
openssl req -new -x509 -days 365 -key certs/private.key -out certs/public.crt \
  -subj "/C=RU/ST=Moscow/L=Moscow/O=MyCompany/CN=localhost" \
  -addext "subjectAltName=DNS:localhost,IP:127.0.0.1,DNS:minio.local"
