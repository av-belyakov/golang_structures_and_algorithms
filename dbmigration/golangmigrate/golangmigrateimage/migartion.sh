#!/bin/bash

ACTION="up"
#ACTION="down 1"

DIR_POSTGRES_MGR="postgres/migrations"
DIR_CLICKHOUSE_MGR="clickhouse/migrations"
DIR_CLICKHOUSE_TEMP="clickhouse/templates"
FILE_CLICKHOUSE_TEMPLATE="template_1.sql"

dir_array=("$DIR_POSTGRES_MGR" "$DIR_CLICKHOUSE_MGR"  "$DIR_CLICKHOUSE_TEMP")

check_postgres() {
    echo "⏳ Проверяем подключение к PostgreSQL..."
    echo "Хост: $PG_HOST, Порт: $PG_PORT, Пользователь: $PG_USER, БД: $PG_DB"

    local pg_timeout=${5:-"5"}
    
    if command -v pg_isready &> /dev/null; then
        if pg_isready -h "$PG_HOST" -p "$PG_PORT" -U "$PG_USER" -d "$PG_DB" -t "$pg_timeout" &> /dev/null; then
            echo "✅ PostgreSQL доступен"
            return 0
        else
            echo "❌ Не удалось подключиться к PostgreSQL"
            return 1
        fi
    fi

    echo "⚠️  Утилита pg_isready не найдена"
}

check_clickhouse() {
    echo "⏳ Проверяем подключение к Clickhouse..."
    echo "Хост: $CH_HOST, Порт: $CH_HTTP__PORT, Пользователь: $CH_USER, БД: $CH_DB"

    local ch_timeout=10
    local curl_cmd="curl -s -o /dev/null -w \"%{http_code}\" --connect-timeout $ch_timeout"
    
    if [ -n "$CH_PASSWORD" ]; then
        curl_cmd="$curl_cmd -u $CH_USER:$CH_PASSWORD"
    fi
    
    local response=$(eval "$curl_cmd \"http://$CH_HOST:$CH_HTTP_PORT/\" --data \"SELECT 1\"")
    
    if [ "$response" = "200" ]; then
        echo "✅ Clickhouse доступен"
        return 0
    else
        echo "❌ Не удалось подключиться к Clickhouse"
        return 1
    fi	
}

check() {
	if check_postgres && check_clickhouse; then
        return 0
    fi

	return 1 
}

echo "⏳ Проверяем наличие директорий" 
# Проверяем наличие директорий
for item in "${dir_array[@]}"; do
    if [ ! -d $item ]; then
	    echo "directory $item not found"
	    exit 1
    fi

    echo "✅ Директория $item доступна"
done

echo "⏳ Проверяем наличие файла с шаблонами для Clickhouse"
# Проверяем наличие файла с шаблонами для Clickhouse
if [ ! -f "$DIR_CLICKHOUSE_TEMP/$FILE_CLICKHOUSE_TEMPLATE" ]; then
    echo ".env file not found"
    exit 1
fi
echo "✅ Файл с шаблонами для Clickhouse доступен"

echo "⏳ Проверяем наличие директории под миграцию в Clickhouse"
# Проверяем наличие директории под миграцию в Clickhouse
if [ ! -f "$DIR_CLICKHOUSE_MGR" ]; then
    mkdir -p "$DIR_CLICKHOUSE_MGR"
fi
echo "✅ Директория под миграцию в Clickhouse доступна"

echo "⏳ Проверяем наличие файла с переменными окружения"
# Проверяем наличие файла с переменными окружения
if [ ! -f .env ]; then
    echo ".env file not found"
    exit 1
fi
echo "✅ Файл с переменными окружения найден"

echo "⏳ Экспортируем переменные окружения"
# Экспортируем переменные окружения
export $(grep -v '^#' .env | grep -v '^$' | grep '=' | sed 's/"//g' | xargs)

echo "⏳ Проверяем доступность БД"
# Проверяем доступность БД
while true; do
    if check; then
        break
    fi

    echo "БД не доступна, ждем..."
    sleep 3
done


# Выполняем миграцию для Postgres #
###################################
echo "⏳ Выполняем миграцию для Postgres"
migrate -path=$PWD/postgres/migrations/ -database postgresql://$PG_USER:$PG_PASSWORD@$PG_HOST:$PG_PORT/$PG_DB?sslmode=disable $ACTION
migrate -path=$PWD/postgres/migrations/ -database postgresql://$PG_USER:$PG_PASSWORD@$PG_HOST:$PG_PORT/$PG_DB?sslmode=disable version
#docker run -v "$PWD/$DIR_POSTGRES_MGR/":/migrations --network host golang-migrate:latest -path=/migrations/ -database postgresql://$PG_USER:$PG_PASSWORD@$PG_HOST:$PG_PORT/$PG_DB?sslmode=disable $ACTION

# Создаём новую миграцию для Clickhouse на основе шаблона от Postgres #
#######################################################################
next_num=$(ls -1 "$DIR_CLICKHOUSE_MGR"/*.up.sql 2>/dev/null | wc -l | awk '{printf "%06d", $1 + 1}')
if [ -z "$next_num" ] || [ "$next_num" = "000000" ]; then
    next_num="000001"
fi

UP_FILE="$DIR_CLICKHOUSE_MGR/${next_num}_schema.up.sql"
DOWN_FILE="$DIR_CLICKHOUSE_MGR/${next_num}_schema.down.sql"

export PG_HOST PG_PORT PG_USER PG_PASSWORD PG_DB
envsubst < "$DIR_CLICKHOUSE_TEMP/$FILE_CLICKHOUSE_TEMPLATE" > "$UP_FILE"
echo "Generated up migration: $UP_FILE (with environment variables applied)"

cat > "$DOWN_FILE" << EOF
-- ClickHouse migration: ${NEXT_NUM}_${MIGRATION_NAME}.down.sql
-- Rollback logic for ${MIGRATION_NAME}
-- This file will be populated with rollback commands

-- Example rollback commands:
-- DROP TABLE IF EXISTS alerts;
-- DROP TABLE IF EXISTS alerts_assigned;
-- DROP TABLE IF EXISTS alerts_exported;
-- DROP TABLE IF EXISTS alerts_viewed;
-- DROP TABLE IF EXISTS classtypes;
-- DROP TABLE IF EXISTS directions;
-- DROP TABLE IF EXISTS event_rules;
-- DROP TABLE IF EXISTS event_ttps;
-- DROP TABLE IF EXISTS events;
-- DROP TABLE IF EXISTS eventsources;
-- DROP TABLE IF EXISTS iocs;
-- DROP TABLE IF EXISTS iocsources;
-- DROP TABLE IF EXISTS geoip;
-- DROP TABLE IF EXISTS localclasstypes;
-- DROP TABLE IF EXISTS nccciclasses;
-- DROP TABLE IF EXISTS objects;
-- DROP TABLE IF EXISTS objects_homenets;
-- DROP TABLE IF EXISTS objects_resources;
-- DROP TABLE IF EXISTS regions;
-- DROP TABLE IF EXISTS scopes;
-- DROP TABLE IF EXISTS signatures;
-- DROP TABLE IF EXISTS tasks;
-- DROP TABLE IF EXISTS tasks2;
-- DROP TABLE IF EXISTS users;
-- DROP TABLE IF EXISTS users_histories;

-- Drop dictionaries
-- DROP DICTIONARY IF EXISTS siem2_iocsources;
-- DROP DICTIONARY IF EXISTS siem2_geoip;
-- DROP DICTIONARY IF EXISTS siem2_objects;
-- DROP DICTIONARY IF EXISTS siem2_externalsystems;
-- DROP DICTIONARY IF EXISTS siem2_priorities;
-- DROP DICTIONARY IF EXISTS siem2_regions;
-- DROP DICTIONARY IF EXISTS siem2_iocs;
-- DROP DICTIONARY IF EXISTS siem2_eventsources;
-- DROP DICTIONARY IF EXISTS siem2_localclasstypes;
-- DROP DICTIONARY IF EXISTS siem2_countries;
-- DROP DICTIONARY IF EXISTS siem2_signatures;
-- DROP DICTIONARY IF EXISTS siem2_directions;
-- DROP DICTIONARY IF EXISTS siem2_homenets;
-- DROP DICTIONARY IF EXISTS siem2_resources;
-- DROP DICTIONARY IF EXISTS siem2_classtypes;
-- DROP DICTIONARY IF EXISTS siem2_nccciclasses;
-- DROP DICTIONARY IF EXISTS siem2_scopes;

EOF

# Выполняем миграцию для Clickhouse #
#####################################
echo "⏳ Выполняем миграцию для Clickhouse"
./migrate -path=$PWD/clickhouse/migrations/ -database="clickhouse://$CH_HOST:$CH_PORT?username=$CH_USER&password=$CH_PASSWORD&database=$CH_DB&x-multi-statement=true" $ACTION
./migrate -path=$PWD/clickhouse/migrations/ -database="clickhouse://$CH_HOST:$CH_PORT?username=$CH_USER&password=$CH_PASSWORD&database=$CH_DB&x-multi-statement=true" version
#docker run -v "$PWD/$DIR_CLICKHOUSE_MGR/":/migrations --rm --network host golang-migrate:latest -path=/migrations/ -database="clickhouse://$CH_HOST:$CH_PORT?username=$CH_USER&password=$CH_PASSWORD&database=$CH_DB&x-multi-statement=true" $ACTION
