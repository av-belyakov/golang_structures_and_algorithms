# Важная информация

## Связка NATS, Prometheus, Grafana (без SSL и авторизации)

Выполняем

```bash
docker-compose -f docker-compose.yml up -d
```

- http://localhost:9090 _доступен графический интерфейс Prometheus_
- http://localhost:9090/targets _все доступные источники данных для Prometheus_

Grafana доступна по http://localhost:3000 admn:admin. После авторизации необходимо поменять пароль администратора. Вся настройка подключения к Prometheus осуществляется через пользовательский интерфейс (подробнее https://grafana.com/docs/grafana/latest/datasources/prometheus/configure/).
Однако, если коротко, Connections->Add new connection->Prometheus->Add new data source, где выполняется полная настройка. После успешной настройки подключённые источники данных можно посмотреть в Connections->Data sources
