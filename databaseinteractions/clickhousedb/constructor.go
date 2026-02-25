package clickhousedb

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"net"
	"time"

	"github.com/ClickHouse/clickhouse-go/v2"
)

func New(opts ...Options) (*ClickHouseClient, error) {
	chc := &ClickHouseClient{
		parameters: ClickhouseParameters{
			database: "clickhouse",
			host:     "127.0.0.1",
			port:     9000,
		},
	}

	for _, opt := range opts {
		if err := opt(chc); err != nil {
			return chc, err
		}
	}

	chc.options = &clickhouse.Options{
		Addr: []string{
			net.JoinHostPort(chc.parameters.host, fmt.Sprintf("%d", chc.parameters.port)),
		},
		Auth: clickhouse.Auth{
			Database: chc.parameters.database,
			Username: chc.parameters.user,
			Password: chc.parameters.password,
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
	}

	return chc, nil
}

// WithHost имя или ip адрес
func WithHost(v string) Options {
	return func(th *ClickHouseClient) error {
		if v == "" {
			return errors.New("the value of 'host' cannot be empty")
		}

		th.parameters.host = v

		return nil
	}
}

// WithPort сетевой порт
func WithPort(v int) Options {
	return func(th *ClickHouseClient) error {
		if v <= 0 || v > 65535 {
			return errors.New("an incorrect network port value was received")
		}

		th.parameters.port = v

		return nil
	}
}

// WithDatabase имя базы данных
func WithDatabase(v string) Options {
	return func(th *ClickHouseClient) error {
		if v == "" {
			return errors.New("the value of 'database' cannot be empty")
		}

		th.parameters.database = v

		return nil
	}
}

// WithUser имя пользователя для авторизации в БД
func WithUser(v string) Options {
	return func(th *ClickHouseClient) error {
		if v == "" {
			return errors.New("the value of 'user' cannot be empty")
		}

		th.parameters.user = v

		return nil
	}
}

// WithPassword пароль для авторизации в БД
func WithPassword(v string) Options {
	return func(th *ClickHouseClient) error {
		if v == "" {
			return errors.New("the value of 'passwd' cannot be empty")
		}

		th.parameters.password = v

		return nil
	}
}
