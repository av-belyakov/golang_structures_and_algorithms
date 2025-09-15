# golang-migrate

### Пример команды для создания новой миграции

Команда выполняется из директории с файлом docker-compose.yml

```bash
docker run --rm -u $(id -u):$(id -g) -v $PWD/golang-migrate/postgres/migrations:/db/migrations:rw golang-migrate:latest create -ext sql -dir /db/migrations -ext sql -tz UTC -seq create_new_user_table
```

### Пример команды для смены версии миграции в БД

```bash
docker run -v "$PWD/postgres/migrations/":/migrations --network host golang-migrate:latest -path=/migrations/ -database postgresql://soc:yTa89-nMsf-GAbs2@localhost:5432/socdb?sslmode=disable force <номер версии>
```
