# Подсказки по работе с SSL сертификатами

Для того что бы проверить CA (Certificate Authority) выполняем

```bash
openssl s_client -connect localhost:9093 -CAfile ./certs/ca.crt </dev/null
```

так же можно посмотреть сертификат отдельно

```bash
openssl verify -CAfile ./certs/ca.crt ./certs/server.crt
```

посмотреть список допустимых IP-адресов для сертификата

```bash
openssl x509 -in ./certs/server.crt -text -noout | grep -A5 "Subject Alternative Name"
```
