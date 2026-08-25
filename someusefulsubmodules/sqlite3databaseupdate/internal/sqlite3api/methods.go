package sqlite3api

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func (m *Module) Start(ctx context.Context) error {
	if _, err := os.Stat(m.pathFileDb); err != nil {
		return err
	}

	client, err := sql.Open("sqlite3", m.pathFileDb)
	if err != nil {
		return err
	}
	m.client = client

	go func(ctx context.Context) {
		<-ctx.Done()

		client.Close()
	}(ctx)

	ctx, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()

	if err = m.client.PingContext(ctx); err != nil {
		return err
	}

	return nil
}

// GetTableColumns информация о таблице
func (m *Module) GetTableColumns(ctx context.Context, tableName string) ([]ColumnInfo, error) {
	rows, err := m.client.QueryContext(ctx, fmt.Sprintf("PRAGMA table_info(%s)", tableName))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var columns []ColumnInfo
	for rows.Next() {
		var col ColumnInfo
		if err := rows.Scan(
			&col.CID,
			&col.Name,
			&col.Type,
			&col.NotNull,
			&col.DefaultVal,
			&col.PK,
		); err != nil {
			return nil, err
		}

		columns = append(columns, col)
	}

	return columns, rows.Err()
}

// GetAllEventsId получить весь список событий MISP
func (m *Module) GetAllEventsId(ctx context.Context) (map[int64]int64, error) {
	listEvent := map[int64]int64{}

	rows, err := m.client.QueryContext(ctx, "SELECT caseId, eventId FROM placeholder_misp")
	if err != nil {
		return listEvent, err
	}
	defer rows.Close()

	var caseId, eventId int64
	for rows.Next() {
		if err = rows.Scan(&caseId, &eventId); err != nil {
			return listEvent, err
		}

		listEvent[eventId] = caseId
	}

	return listEvent, rows.Err()
}

// UpdateColumnSource обновить
func (m *Module) UpdateColumnSource(ctx context.Context, tableName string, eventId int64, source string) error {
	query := fmt.Sprintf("UPDATE %s SET source='%s', datetime='%s' WHERE eventId=%d", tableName, source, time.Now().Format(time.RFC3339), eventId)

	fmt.Println("Query:", query)

	if _, err := m.client.ExecContext(ctx, query); err != nil {
		return err
	}

	return nil
}

// AlterTableAddColumn добавляет новые колонки в таблицу
func (m *Module) AlterTableAddColumn(ctx context.Context, tableName, columnName, columnType string) error {
	query := fmt.Sprintf("ALTER TABLE %s ADD COLUMN %s %s", tableName, columnName, columnType)

	fmt.Println("Query:", query)

	if _, err := m.client.ExecContext(ctx, query); err != nil {
		return err
	}

	return nil
}
