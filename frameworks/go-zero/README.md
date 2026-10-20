# Примеры создания веб-приложений и RPC-сервисов.

Подробная информация доступна по ссылке [go install github.com/zeromicro/go-zero/tools/goctl@latest](https://go-zero.dev/getting-started/)

### Установка и настройка

Установка фреймворка и инструмента goctl:

```bash
go get -u github.com/zeromicro/go-zero
go install github.com/zeromicro/go-zero/tools/goctl@latest
```

Если при вызове goctl получаем ошибку типа 'Команда "goctl" не найдена' то причина в том что директория **$GOPATH/bin** не добавлена в вашу переменную окружения **PATH**. Для того что бы это сделать нужно выполнить следующее. Узнайте путь установки:

```bash
go env GOPATH
```

После установке goctl бинарный файл лежит в подпапке этой директории. Нужно добавить строку в конец файла ~/.bashrc (или ~/.zshrc):

```bash
export PATH="$PATH:$(go env GOPATH)/bin"
```

Затем применить изменения:

```bash
source ~/.bashrc  # или source ~/.zshrc
```

Важно: После изменения переменных окружения обязательно перезапустите терминал. Далее проверка:

```bash
goctl --version
```

Далее, для того что бы сгенерировать полный макет, готовый к использованию выполните:

```bash
goctl api new <имя_проекта>
```

В результате получим полный макет проекта:

```bash
greet/
├── etc/
│   └── greet-api.yaml        # config: port, logging, etc.
├── internal/
│   ├── config/
│   │   └── config.go          # config struct
│   ├── handler/
│   │   ├── greethandler.go    # HTTP handler (auto-generated)
│   │   └── routes.go          # route registration
│   ├── logic/
│   │   └── greetlogic.go      # ← edit this: your business logic
│   └── svc/
│       └── servicecontext.go  # shared dependencies (DB, cache, etc.)
└── greet.go                   # main entrypoint
```

Нужно редактировать только файлы в директории **internal/logic/**. Все остальное генерируется из файла **.api** при каждом запуске 'goctl api go'.

#### Что можно сгенерировать goctl

```bash
goctl api new <name>	#Полный каркас сервиса HTTP API
goctl api go -api f.api -dir .	#Код Go из существующего .api файла
goctl rpc new <name>	#Полный каркас сервиса gRPC
goctl rpc protoc f.proto ...	#Код Go из существующего .proto файла
goctl model mysql ddl -src f.sql -dir .	#Уровень модели БД из схемы SQL
goctl docker -go main.go	#Dockerfile
goctl kube deploy ...	#Манифест развертывания Kubernetes
```

После генерации нового сервиса, например user-api

```bash
goctl api go -api user.api -dir .
go mod init user-api
go mod tidy
```

Возможно понадобиться изменить пути импорта в файлах на название модуля в файле **go.mod**. Вероятность этого высокая если имя модуля в **go.mod** начинается на _github.com/_, а новая версия пакета ещё не была импортирована в GitHub.

В результате будет сгенерирован ряд директорий и файлов представляющий собой макет будущего приложения. **Важно**, радактировать (изменять код методов и функций) можно только файлы находящиеся в директории **logic**.

```bash
user-api/
├── etc/
│   └── user-api.yaml
├── internal/
│   ├── config/config.go
│   ├── handler/
│   │   ├── loginhandler.go         # auto-generated, do not edit
│   │   ├── getuserinfohandler.go   # auto-generated, do not edit
│   │   └── routes.go
│   ├── logic/
│   │   ├── loginlogic.go           # ← implement this
│   │   └── getuserinfologic.go     # ← implement this
│   ├── svc/servicecontext.go
│   └── types/types.go             # auto-generated from DSL
└── user.go
```

Запускаем user-api:

```bash
go run user.go
```

Пытаемся авторизоватся:

```bash
curl -s -X POST http://localhost:8888/user/login \
  -H 'Content-Type: application/json' \
  -d '{"username":"admin","password":"secret"}'
```

Получаем тестовую информацию:

```bash
curl http://localhost:8888/user/1
```

#### Восстановление после изменений в DSL

При каждом изменении user.api выполняйте восстановление без перезаписи файлов логики:

```bash
goctl api go -api user.api -dir .
```

goctl перезаписывает только свои файлы (handler/, types/, routes.go). Ваши logic/ файлы сохраняются.
