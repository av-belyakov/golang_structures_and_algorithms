// Пример взаимодействия с Sqlite3 DB
package sqlite3interaction

import (
	"database/sql"
)

// ConnectionSqlite3
type ConnectionSqlite3 struct {
	Client *sql.DB
}

type ListExample []Information

type Information struct {
	Id    int
	Name  string
	Event string
	Age   int
}
