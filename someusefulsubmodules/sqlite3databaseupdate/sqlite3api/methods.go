package sqlite3api

import (
	"context"
	"database/sql"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func (m *Module) Start(ctx context.Context) error {
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

// GetAllEventsId получить весь список событий MISP
func (m *Module) GetAllEventsId(ctx context.Context) (map[int]int, error) {
	listEvent := map[int]int{}

	rows, err := m.client.QueryContext(ctx, "SELECT caseId, eventId FROM placeholder_misp")
	if err != nil {
		return listEvent, err
	}
	defer rows.Close()

	result := struct {
		caseId, eventId int
	}{}

	for rows.Next() {
		if err = rows.Scan(&result); err != nil {
			return listEvent, err
		}

		listEvent[result.eventId] = result.caseId
	}

	if err = rows.Err(); err != nil {
		return listEvent, err
	}

	return listEvent, nil
}
