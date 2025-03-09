package sqlite3interaction_test

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	sqliteinteraction "github.com/av-belyakov/golang_structures_and_algorithms/databaseinteractions/sqlite3interaction"
)

const Path_Sqlite3_DB string = "sqlite3.db"

var (
	ctx    context.Context
	cancel context.CancelFunc

	sqlite3Client *sqliteinteraction.ConnectionSqlite3

	err error
)

func TestMain(m *testing.M) {
	ctx, cancel = context.WithCancel(context.Background())

	sqlite3Client, err = sqliteinteraction.NewConnectionSqlite3(Path_Sqlite3_DB)
	if err != nil {
		log.Fatalln(err)
	}

	os.Exit(m.Run())
}

func TestGetDumpdb(t *testing.T) {
	defer func() {
		//закрываем соединение с Sqlite3 БД
		sqlite3Client.ConnectionClose()
	}()

	t.Run("Тест 1. Получаем эхо ответ от Sqlite3 database", func(t *testing.T) {
		err = sqlite3Client.Ping(ctx)
		assert.NoError(t, err)
	})

	t.Run("Тест 2. Создаем новую таблицу", func(t *testing.T) {
		/*queryRes*/ _, err := sqlite3Client.Client.ExecContext(ctx, "CREATE TABLE IF NOT EXISTS information_table (id INT, name TEXT, event TEXT, age INT)")
		assert.NoError(t, err)
	})

	t.Run("Тест 3. Добавляем записи", func(t *testing.T) {
		record, err := sqlite3Client.Client.PrepareContext(ctx, "INSERT INTO information_table (id, name, event, age) VALUES (?,?,?,?)")
		assert.NoError(t, err)

		container := sqliteinteraction.ListExample{
			sqliteinteraction.Information{1, "Ivan", "Jump", 13},
			sqliteinteraction.Information{2, "Maks", "Jump", 12},
			sqliteinteraction.Information{3, "Olga", "Run", 15},
			sqliteinteraction.Information{4, "Mihail", "Swim", 11},
			sqliteinteraction.Information{5, "Elena", "Run", 12},
			sqliteinteraction.Information{6, "Nikolay", "Go", 13},
			sqliteinteraction.Information{7, "Alex", "Sleep", 14},
		}

		for _, v := range container {
			_, err = record.ExecContext(ctx, v.Id, v.Name, v.Event, v.Age)
			assert.NoError(t, err)
		}
	})

	t.Run("Тест 4. Получаем все записи", func(t *testing.T) {
		rows, err := sqlite3Client.Client.QueryContext(ctx, "SELECT * FROM information_table")
		assert.NoError(t, err)
		//закрываем дескриптор
		defer rows.Close()

		for rows.Next() {
			var (
				id    int
				name  string
				event string
				age   int
			)

			err := rows.Scan(&id, &name, &event, &age)
			assert.NoError(t, err)

			t.Logf("id:%d, name:%s, event:%s, age:%d\n", id, name, event, age)
		}
	})
}
