# Netbox image

### Некоторые особенности настройки

После запуска docker compose через команду:

```bash
docker compose up -d
```

После старта

Для того что бы создать пользователя superadmin для Netbox, это первый пользователь который нужен для дальнейшей работы с приложением, нужно выполнить:

```bash
docker compose exec netbox /opt/netbox/netbox/manage.py createsuperuser
```
