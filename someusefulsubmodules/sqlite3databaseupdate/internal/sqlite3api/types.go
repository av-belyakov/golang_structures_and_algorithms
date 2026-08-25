package sqlite3api

import "database/sql"

type Module struct {
	client     *sql.DB
	pathFileDb string
}

type ColumnInfo struct {
	CID        int
	Name       string
	Type       string
	NotNull    int
	DefaultVal sql.NullString
	PK         int
}
