#!/bin/bash

RESET="\033[0m"

BOLD="\e[1m"
ITALIC="\e[3m"

#ACTION=${1:-"up"}

DIR_POSTGRES="postgres"
DIR_CLICKHOUSE="clickhouse"
DIR_MIGRATION="migrations"

function printInfo() {   
    echo -e "\033[34m[INFO]\033[0m $1"
}

function printSuccess() {
    echo -e "\033[32m[SUCCESS]\033[0m $1"
}

function printWarrning() {
    echo -e "\033[33m[WARRNING]\033[0m $1"
}

function printProcessing() {
    echo -e "\033[35m[PROCESSING]\033[0m $1"
}

function printError() {
    echo -e "\033[31m[ERROR]\033[0m $1"
}

function checkPostgres() {
    printInfo "проверяем подключение к PostgreSQL..."
    printf "хост: %s, порт: %s, пользователь: %s, БД: %s\n" ${POSTGRES_HOST} ${POSTGRES_PORT} ${POSTGRES_USER} ${POSTGRES_DB}

    local pg_timeout=${5:-"5"}
    
    if command -v pg_isready &> /dev/null; then
        if pg_isready -h ${POSTGRES_HOST} -p ${POSTGRES_PORT} -U ${POSTGRES_USER} -d ${POSTGRES_DB} -t ${pg_timeout} &> /dev/null; then
            printSuccess "PostgreSQL доступен"
            return 0
        else
            printError "не удалось подключиться к PostgreSQL"
            return 1
        fi
    fi

    printWarrning "утилита pg_isready не найдена"
}

function checkClickhouse() {
    printInfo "проверяем подключение к Clickhouse..."
    printf "хост: %s, порт: %s, пользователь: %s, БД: %s\n" ${CLICKHOUSE_HOST} ${CLICKHOUSE_HTTP_PORT} ${CLICKHOUSE_USER} ${CLICKHOUSE_DB}

    local ch_timeout=10
    local curl_cmd="curl -s -o /dev/null -w \"%{http_code}\" --connect-timeout $ch_timeout"
    
    if [ -n ${CLICKHOUSE_PASSWORD} ]; then
        curl_cmd="$curl_cmd -u $CLICKHOUSE_USER:$CLICKHOUSE_PASSWORD"
    fi
    
    local response=$(eval "$curl_cmd \"http://$CLICKHOUSE_HOST:$CLICKHOUSE_HTTP_PORT/\" --data \"SELECT 1\"")
    
    if [ ${response} = "200" ]; then
        printSuccess "Clickhouse доступен"
        return 0
    else
        printError "не удалось подключиться к Clickhouse"
        return 1
    fi	
}

function checkDatabaseAvailability() {
	if checkPostgres && checkClickhouse; then
        return 0
    fi

	return 1 
}

function checkPostgresMigrationDirectory() {
    if [ ! -d ${DIR_POSTGRES}/${DIR_MIGRATION} ]; then
        mkdir -p ${DIR_POSTGRES}/${DIR_MIGRATION}
        printInfo "создана директория '${ITALIC}${DIR_POSTGRES}/${DIR_MIGRATION}${RESET}' под миграцию в Postgres"
    fi
    
    printSuccess "директория '${ITALIC}${DIR_POSTGRES}/${DIR_MIGRATION}${RESET}' под миграцию в Postgres доступна"
}

function createMigrationFiles() {
    local tmpFile="tmp.file"
    local dirTemplates="templates"
    local dirMigration=${DIR_MIGRATION}
    local dirClickhouse=${DIR_CLICKHOUSE}
    local fileTemplateStorage=".template-storage"
    local marker="Below are the rollback commands"

    if [ ! -f ${dirClickhouse}/${dirTemplates} ]; then
        mkdir -p ${dirClickhouse}/${dirTemplates}
    fi
    printSuccess "директория '${ITALIC}${dirTemplates}${RESET}' с шаблонами для Clickhouse доступна"

    if [ ! -f ${dirClickhouse}/${dirTemplates}/${fileTemplateStorage} ]; then
        touch ${dirClickhouse}/${dirTemplates}/${fileTemplateStorage}
        printInfo "создан новый файл-хранилище списка обработаных шаблонов '${BOLD}${fileTemplateStorage}${RESET}'"
    fi
    printSuccess "файл-хранилище списка обработаных шаблонов '${BOLD}${fileTemplateStorage}${RESET}' доступен"

    if [ ! -d ${dirClickhouse}/${dirMigration} ]; then
        mkdir -p ${dirClickhouse}/${dirMigration}
        printInfo "создана директория '${ITALIC}${dirClickhouse}/${dirMigration}${RESET}' под миграцию в Clickhouse"
    fi
    printSuccess "директория '${ITALIC}${dirClickhouse}/${dirMigration}${RESET}' под миграцию в Clickhouse доступна"

    newFileIsExist=false
    printInfo "ищем файлы вида '[0-9]{4}_*.tmp.sql' в директории '${ITALIC}${dirClickhouse}/${dirTemplates}${RESET}'"
    for file in ${dirClickhouse}/${dirTemplates}/* 
    do 
        filename=$(basename ${file})
        if [ ${filename} = ${fileTemplateStorage} ]; then
            continue
        fi

        if [[ ! ${filename} =~ ^[0-9]{4}.*tmp\.sql$ ]]; then
            printError "файл '${BOLD}${filename}${RESET}' не соответствует шаблону '[0-9]{4}_*.tmp.sql', пропускаем"
            continue        
        fi

        if grep -qFx ${filename} ${dirClickhouse}/${dirTemplates}/${fileTemplateStorage}; then
            printWarrning "файл '${BOLD}${filename}${RESET}' уже обрабатывался, пропускаем"
            continue
        fi      

        migrationName=
        fileMigrationUp=""
        fileMigrationDown=""
        newFileIsExist=true

        printProcessing "найден новый файл шаблонов '${BOLD}/${filename}/${RESET}', обрабатываем..."  
        
        migrationName="${filename#*_}"
        migrationName="${migrationName%%.*}"

        # для тестов снаружи докер контейнера
        #response=$(docker run --rm -u $(id -u):$(id -g) -v $PWD/${dirClickhouse}/${dirMigration}:/db/migrations:rw golang-migrate:latest create -ext sql -dir /db/migrations -ext sql -tz UTC -seq create_next_migration 2>&1 | awk -F/ '{print $NF}')
        response=$(migrate create -ext sql -dir /db/migrations -ext sql -tz UTC -seq ${migrationName} 2>&1 | awk -F/ '{print $NF}')

        # преобразуем в массив
        IFS=$'\n' read -rd '' -a files <<< "$(echo -e ${response})"
        for file in ${files[@]}; do
            if [[ ${file} == *"up.sql"* ]]; then
                fileMigrationUp=${file}
            fi

            if [[ ${file} == *"down.sql"* ]]; then
                fileMigrationDown=${file}
            fi
        done

        # до маркера разделителя, запросы миграции типа 'UP'
        sed "/$marker/q" ${dirClickhouse}/${dirTemplates}/${filename} | head -n -1 > ${tmpFile}
        envsubst < ${tmpFile} > ${dirClickhouse}/${dirMigration}/${fileMigrationUp}
        rm ${tmpFile}

        # после маркера разделителя, запросы миграции типа 'DOWN' 
        sed -n "/$marker/,\$p" ${dirClickhouse}/${dirTemplates}/${filename} | tail -n +2 > ${tmpFile}
        envsubst < ${tmpFile} > ${dirClickhouse}/${dirMigration}/${fileMigrationDown}
        rm ${tmpFile}

        # проверяем что файл '*.up.sql' существует и не пустой        
        if [ -s ${dirClickhouse}/${dirMigration}/${fileMigrationUp} ]; then
            printSuccess "файл миграции '${BOLD}${fileMigrationUp}${RESET}' создан и не пуст"
        else
            printWarrning "создан пустой файл миграции '${BOLD}${fileMigrationUp}${RESET}'" 
        fi
        
        # проверяем что файл '*.down.sql' существует и не пустой        
        if [ -s ${dirClickhouse}/${dirMigration}/${fileMigrationDown} ]; then
            printSuccess "файл миграции '${BOLD}${fileMigrationDown}${RESET}' создан и не пуст"  
        else
            printWarrning "создан пустой файл миграции '${BOLD}${fileMigrationDown}${RESET}'"
        fi

        # для тестов пока закоментирую
        printInfo "добавляем новый файл шаблонов '${BOLD}${filename}${RESET}' в хранилище"
        echo ${filename} >> ${dirClickhouse}/${dirTemplates}/${fileTemplateStorage}
    done

    if ! ${newFileIsExist}; then
        printError "нет новых файлов соответствующих шаблону, нечего обрабатывать"
    fi
}

# Поиск файла с переменными окружения
if [ ! -f .env ]; then
    printError "файл с переменными окружения не найден, останов"
    exit 1
fi

printInfo "экспортируем переменные окружения"
export $(grep -v '^#' .env | grep -v '^$' | grep '=' | sed 's/"//g' | xargs)

#MIGRATE_ACTION=down
#MIGRATE_STEP=1

STEP=""
ACTION=${MIGRATE_ACTION:-"up"}
if [[ $ACTION != "up" && $ACTION != "down" ]]; then
    printError "ошибка передачи параметра миграции, значение может быть только 'up' или 'down'"
    exit 1
fi
if [[ $ACTION == "down" ]]; then
    STEP=${MIGRATE_STEP}
fi

printInfo "проверяем доступность баз данных"
while true; do
    if checkDatabaseAvailability; then
        break
    fi

    sleep 3
done

checkPostgresMigrationDirectory

# Выполняем миграцию для Postgres #
###################################
printInfo "выполняем миграцию для Postgres..."
if ! migrate -path=$PWD/postgres/migrations/ -database postgresql://$POSTGRES_USER:$POSTGRES_PASSWORD@$POSTGRES_HOST:$POSTGRES_PORT/$POSTGRES_DB?sslmode=disable $ACTION $STEP; then
    printError "ошибка миграции в Postgres"
    exit 1
fi
if ! migrate -path=$PWD/postgres/migrations/ -database postgresql://$POSTGRES_USER:$POSTGRES_PASSWORD@$POSTGRES_HOST:$POSTGRES_PORT/$POSTGRES_DB?sslmode=disable version; then
    printError "ошибка проверки версии миграции в Postgres"
    exit 1
fi

# Создаём файлы миграции для Clickhouse на основе файлов шаблонов #
###################################################################
createMigrationFiles

# Выполняем миграцию для Clickhouse #
#####################################
printInfo "выполняем миграцию для Clickhouse..."
if ! migrate -path=$PWD/clickhouse/migrations/ -database="clickhouse://$CLICKHOUSE_HOST:$CLICKHOUSE_PORT?username=$CLICKHOUSE_USER&password=$CLICKHOUSE_PASSWORD&database=$CLICKHOUSE_DB&x-multi-statement=true" $ACTION $STEP; then
    printError "ошибка миграции в Clickhouse"
    exit 1
fi
