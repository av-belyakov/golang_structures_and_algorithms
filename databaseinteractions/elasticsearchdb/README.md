# Некоторые особенности настройки доступа к Elasticsearch

Подробнее можно узнать по ссылке https://www.elastic.co/docs/reference/elasticsearch/clients/go/getting-started#_connecting

Ставим go пакет для взаимодействия с Elasticsearch

```bash
go get github.com/elastic/go-elasticsearch/v9@latest
```

Можно выполнить авторизацию в СУБД Elasticsearch с использованием логина и пароля или с помощью ранее созданного API-Key. Пример создания API-Key ниже
![пример-1](./images/image-1.png)
В результате будет создан ключ
![пример-2](./images/image-2.png)

Подробное описание формирования запросов к Elasticsearch можно посмотреть по ссылке https://www.elastic.co/docs/reference/elasticsearch/rest-apis

Здесь более подробено рассказано о языке запросов **Query DSL** https://www.elastic.co/docs/explore-analyze/query-filter/languages/querydsl
