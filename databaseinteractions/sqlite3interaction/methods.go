package sqlite3interaction

import (
	"context"
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

// NewConnectionSqlite3 устанавливает новое соединение с БД
func NewConnectionSqlite3(dbPath string) (*ConnectionSqlite3, error) {
	sqlite3Client, err := sql.Open("sqlite3", dbPath)

	return &ConnectionSqlite3{
		Client: sqlite3Client,
	}, err
}

// Ping проверка соединения с БД
func (db *ConnectionSqlite3) Ping(ctx context.Context) error {
	return db.Client.PingContext(ctx)
}

// ConnectionClose закрывает соединение с БД
func (db *ConnectionSqlite3) ConnectionClose() {
	db.Client.Close()
}
