package clickhousedb

import (
	"context"
	"database/sql"

	"github.com/ClickHouse/clickhouse-go/v2"
)

//
// Для подключения к БД выбираем один из двух вариантов
// 1. Используем нативный интерфейс драйвера Clickhouse
// 2. Используем общую структуру драйвера database/sql
//

type ClickHouseClient struct {
	connect    clickhouse.Conn // нативный интерфейс драйвера Clickhouse
	connectDb  *sql.DB         // общую структуру драйвера database/sql
	options    *clickhouse.Options
	ctx        context.Context
	parameters ClickhouseParameters
}

type ClickhouseParameters struct {
	database string
	user     string
	password string
	host     string
	port     int
}

type Options func(*ClickHouseClient) error
