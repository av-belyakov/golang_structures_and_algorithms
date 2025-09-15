## Шпаргалка с примерами структур и различных алгоритмов на языке Golang

### Структура проекта golang_structures_and_algorithms

##### commonfunctions - Общие вспомогательные функции

##### configurationpatterns - Шаблоны написания конфигурационных опций

- builder - установка конфигурационных опций на основе шаблона "Строитель" (Builder pattern)
- fucnoptions - установка конфигурационных опций на основе "шаблона функциональных опций" (Functional options pattern)

##### errorshandling - Шаблоны группировки произвольного количества ошибок

##### generics - Примеры написания дженериков (Generic)

##### databaseinteractions - Примеры взаимодействий с API различных БД

###### sqlite3interaction - Пример взаимодействия с Sqlite3 DB

###### redisdbinteraction - Пример взаимодействия с Redis DB

##### inmemorycache - Временное хранилище информации в памяти приложения (In-memory cache)

##### lineardatastructures - Линейные структуры данных

- doublylinkedlist - линейная структура данных типа "Двойной связанный список" (Doubly Linked List)
- linkedlist - линейная структура данных типа "Связанный список" (Linked List)
- queues - структуры данных типа "Очереди" (Queues) или типа "Синхронизированная очередь" (Synchronized queue)
- sets - структура данных типа "Множество" или "Набор" (Set)

##### stringspackage - Примеры работы со строками

##### structuraldesignpatterns - Примеры шаблонов структур и алгоритмов

- adapterpattern - шаблон "Адаптер" (Adapter)
- bridgepattern - шаблон "Мост" (Bridge)
- compositepattern - шаблон "Составной" (Composite)
- decoratorpattern - шаблон "Декоратор" (Decorator)
- facadepattern - шаблон "Фасад" (Facade)
- flyweightpattern - шаблон Легчайший (Flyweight)
- privateclassdata - шаблон "Данные частного класса" (Private class data)
- proxypattern - шаблон Прокси (Proxy)
- singletonpattern - шаблон Одиночка (Singleton)

##### syncpackage - Некоторые примеры работы с пакетом sync

##### testingpattern - Примеры шаблонов тестирования

##### timepackage - Примеры работы с пакетом time

##### workingswithfiles - Примеры работы с файлами

- checkfileexist - проверка существования файлов
- checkrwpermissions - проверка прав доступа файла на запись и чтение
- compressuncompressfile - упаковка и распаковка файла спомощью gzip
- copyfile - копирование байтов из исходного файла в целевой
- createfile - создание файлов
- openandclosefile - открытие и закрытие файла с различными атрибутами и правами доступа
- readwithscanner - чтение файла в режиме сканера
- renamefile - переименование файла
- seekpositionsfile - поиск позиции в файле
- usebufferedreader - использование буфера для чтения файла
- usebufferedwriter - использование буфера для записи в файл
- writebyteorstringtofile - запись среза байт или строки в файл

##### traversingdirectories - Примеры обхода директорий

##### filemonitoring - Мониторинг файлов

##### mechanicspipes - Механизмы конвейеров

- anonymouspipes - анонимные конвееры
- navigatingnamedpipes - навигация по именованным каналам

##### unixsocket - Unix сокеты

##### errorspackage - Некоторые примеры работ с ошибками

##### flagpackage - Пример работы с пакетом flag

##### certificategenerator - Пример генерации самоподписных сертификатов

##### miniogopackage - Пример взаимодействия с MinIO с помощью пакета minio-go

##### dbmigration - Примеры настройки пакетов для миграции базы данных

- golangmigrate - Пример работы с пакетом golang-migrate
