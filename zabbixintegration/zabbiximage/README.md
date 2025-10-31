# Zabbix Image

### Порядок разворачивания контейнеров с zabbix-server, zabbix-web, zabbix-agent и postgres СУБД

Для того что бы создать TLS сертификаты, выполнть:

```bash
./gencerts.sh
```

После этого появится директория _certs_ с TLS сертификатами и ключами для сервера и клиентов.

Выполнить команду:

```bash
docker-compose up -d
```
