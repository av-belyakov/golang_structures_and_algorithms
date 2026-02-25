package clickhousedb

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/ClickHouse/clickhouse-go/v2"
)

// Connect выполняет подключение к БД и реализует интерфейс clickhouse.Conn
func (chc *ClickHouseClient) Connect(ctx context.Context) (clickhouse.Conn, error) {
	conn, err := clickhouse.Open(chc.options)
	if err != nil {
		return nil, err
	}

	chc.ctx = ctx
	chc.connect = conn

	if err := chc.connect.Ping(ctx); err != nil {
		if exception, ok := err.(*clickhouse.Exception); ok {
			fmt.Printf("Exception [%d] %s \n%s\n", exception.Code, exception.Message, exception.StackTrace)
		}

		return nil, err
	}

	return chc.connect, nil
}

// ConnectDB использует общую структуру драйвера database/sql
func (chc *ClickHouseClient) ConnectDB(ctx context.Context) (*sql.DB, error) {
	chc.ctx = ctx
	chc.connectDb = clickhouse.OpenDB(chc.options)

	if err := chc.connectDb.Ping(); err != nil {
		if exception, ok := err.(*clickhouse.Exception); ok {
			fmt.Printf("Exception [%d] %s \n%s\n", exception.Code, exception.Message, exception.StackTrace)
		}

		return nil, err
	}

	return chc.connectDb, nil
}

/*
func NewClickhouseConnect(ctx context.Context) (driver.Conn, error) {
	conn, err := clickhouse.Open(&clickhouse.Options{
		Addr: []string{net.JoinHostPort("127.0.0.1", os.Getenv("CLICKHOUSE_SERVER_YOUR_OWN_PORT"))},
		Auth: clickhouse.Auth{
			Database: os.Getenv("CLICKHOUSE_SERVER_DB"),
			Username: os.Getenv("CLICKHOUSE_SERVER_USER"),
			Password: os.Getenv("CLICKHOUSE_SERVER_PASSWORD"),
		},
		DialContext: func(ctx context.Context, addr string) (net.Conn, error) {
			var d net.Dialer
			return d.DialContext(ctx, "tcp", addr)
		},
		Debug: true,
		Debugf: func(format string, v ...any) {
			fmt.Printf(format+"\n", v...)
		},
		TLS: &tls.Config{
			InsecureSkipVerify: false,
		},
		Settings: clickhouse.Settings{
			"max_execution_time": 60,
		},
		Compression: &clickhouse.Compression{
			Method: clickhouse.CompressionLZ4,
		},
		DialTimeout:          time.Second * 30,
		MaxOpenConns:         5,
		MaxIdleConns:         5,
		ConnMaxLifetime:      time.Duration(10) * time.Minute,
		ConnOpenStrategy:     clickhouse.ConnOpenInOrder,
		BlockBufferSize:      10,
		MaxCompressionBuffer: 10240,
		// опционально, информация о клиенте
		ClientInfo: clickhouse.ClientInfo{
			Products: []struct {
				Name, Version string
			}{
				{Name: "clickhouse-go-package", Version: "2.43.0"},
			},
		},
	})
	if err != nil {
		return nil, err
	}

	if err := conn.Ping(ctx); err != nil {
		if exception, ok := err.(*clickhouse.Exception); ok {
			fmt.Printf("Exception [%d] %s \n%s\n", exception.Code, exception.Message, exception.StackTrace)
		}
		return nil, err
	}
	return conn, nil
}
*/
