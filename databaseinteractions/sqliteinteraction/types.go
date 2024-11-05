// Примерный пример взаимодействия с API Sqlite
package sqliteinteraction

import (
	"database/sql"

	"github.com/av-belyakov/golang_structures_and_algorithms/interfaces"
)

type SqliteApiOptions struct {
	path   string
	db     *sql.DB
	logger interfaces.Logger
}

// routeSettings настройки маршрутизатора
type routeSettings struct {
	command      string
	taskId       string
	service      string
	data         []byte
	chanResponse chan<- ChanOutputApiSqlite
}

// ChanApiSqlite канал для взаимодействия с API SQLite
type ChanApiSqlite struct {
	Command      string                     //команда которую должен выполнить API SQLite
	TaskID       string                     //id задачи
	Service      string                     //имя сервиса, за пределами NATS, от имени которого происходит запрос (например MISP, ES)
	Data         []byte                     //данные передаваемые в API SQLite
	ChanResponse chan<- ChanOutputApiSqlite //канал для ответа
}

// ChanOutputApiSqlite
type ChanOutputApiSqlite struct {
	Status bool   //статус выполнения
	TaskID string //id задачи
	Data   []byte //передаваемые данные
}
