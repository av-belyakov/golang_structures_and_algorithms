package clickhousedb

import (
	"context"
	"database/sql/driver"

	"github.com/ClickHouse/clickhouse-go/v2"
)

type ClickHouseClient struct {
	connect    driver.Conn
	options    *clickhouse.Options
	ctx        context.Context
	parameters ClickhouseParameters
}

type ClickhouseParameters struct {
	database string
	user     string
	passwd   string
	host     string
	port     int
}

type Options func(*ClickHouseClient) error
