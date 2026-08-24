package sqlite3api

import "database/sql"

type Module struct {
	client     *sql.DB
	pathFileDb string
}
