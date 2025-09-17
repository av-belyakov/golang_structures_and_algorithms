# Инструкция по запуску Kafka с SSL

1. Сгенерируйте ключи:

```bash
./kafka-keygen.sh <любой_ваш_пароль>
```

2. Запустите docker-compose.ssl.yml

```bash
docker-compose -f docker-compose.ssl.yml up -d
```
