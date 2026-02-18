package postgresinteraction_test

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"database/sql"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func TestPostgresInteraction(t *testing.T) {
	if err := godotenv.Load("./postgresimage/.env"); err != nil {
		log.Fatal(err)
	}

	dsn := fmt.Sprintf(
		"postgres://%s:%s@localhost/%s?sslmode=require",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	t.Run("Тест 1. Создаём тестовые таблицы и заполняем их данными", func(t *testing.T) {
		row, err := db.QueryContext(t.Context(),
			`CREATE TABLE IF NOT EXISTS public.attack_protocols (
				id SERIAL PRIMARY KEY,
				name varchar(255) NOT NULL,
				attack_level_id integer NOT NULL,
				default_t_protocol varchar(3) NOT NULL,
				default_port integer NOT NULL,
				updated_at timestamp with time zone,
				created_at timestamp with time zone,
				updated_by varchar(100),
				created_by varchar(100)
		);`)
		assert.NoError(t, err)

		row.Close()

		for range 30 {
			row, err := db.QueryContext(t.Context(),
				`INSERT INTO public.attack_protocols(
					name,
					attack_level_id,
					default_t_protocol,
					default_port,
					updated_at,
					created_at,
					updated_by,
					created_by
				) VALUES ($1, $2, $3, $4, $5, $6, $7, $8);`,
				gofakeit.State(),
				gofakeit.Number(1, 10),
				"TCP",
				gofakeit.Number(1, 65535),
				gofakeit.Date(),
				gofakeit.Date(),
				gofakeit.Name(),
				gofakeit.Name(),
			)
			assert.NoError(t, err)
			assert.NotNil(t, row)

			row.Close()
		}
	})

	t.Run("Тест 2. Выполняем запросы к БД", func(t *testing.T) {
		t.Run("2.1. Поиск записей по интервалу портов", func(t *testing.T) {
			rows, err := db.QueryContext(t.Context(),
				`SELECT * 
				FROM public.attack_protocols 
				WHERE default_port BETWEEN $1 AND $2;`,
				1_000,
				30_000,
			)
			assert.NoError(t, err)

			list := []Info(nil)
			for rows.Next() {
				var (
					id            int
					name          string
					attackLevelID int
					defaultTProto string
					defaultPort   int
					updatedAt     time.Time
					createdAt     time.Time
					updatedBy     string
					createdBy     string
				)

				err := rows.Scan(
					&id,
					&name,
					&attackLevelID,
					&defaultTProto,
					&defaultPort,
					&updatedAt,
					&createdAt,
					&updatedBy,
					&createdBy,
				)
				assert.NoError(t, err)

				list = append(list, Info{
					Id:               id,
					Name:             name,
					AttackLevelId:    attackLevelID,
					DefaultTProtocol: defaultTProto,
					DefaultPort:      defaultPort,
					UpdatedAt:        updatedAt,
					CreatedAt:        createdAt,
					UpdatedBy:        updatedBy,
					CreatedBy:        createdBy,
				})
			}
			rows.Close()
			assert.Greater(t, len(list), 0)

			for _, info := range list {
				fmt.Println(info)
			}
		})

		t.Run("2.2. Поиск записей по дате и по идентификатору уровня атаки", func(t *testing.T) {
			rows, err := db.QueryContext(t.Context(),
				`SELECT  name, created_at, attack_level_id
				FROM public.attack_protocols 
				WHERE CAST(created_at AS DATE) BETWEEN '1995-01-01' AND '2025-12-31'
				AND attack_level_id = 2;`,
			)
			assert.NoError(t, err)

			var (
				createdAt     time.Time
				name          string
				attackLevelID int
			)

			list := []Info(nil)
			for rows.Next() {
				err := rows.Scan(
					&name,
					&createdAt,
					&attackLevelID,
				)
				assert.NoError(t, err)

				list = append(list, Info{
					Name:          name,
					CreatedAt:     createdAt,
					AttackLevelId: attackLevelID,
				})
			}

			rows.Close()
			assert.Greater(t, len(list), 0)

			for _, info := range list {
				fmt.Println(info)
			}
		})
	})
	t.Run("Тест 3. Выполняем поиск и обновление записей", func(t *testing.T) {
		t.Run("3.1. Обновление протокола для некоторых типов атак", func(t *testing.T) {
			row, err := db.QueryContext(t.Context(),
				`UPDATE public.attack_protocols
				SET default_t_protocol = 'UDP' 
				WHERE attack_level_id = 2;`,
			)
			assert.NoError(t, err)
			row.Close()
		})
	})

	t.Cleanup(func() {
		_, err := db.QueryContext(context.Background(), `DROP TABLE IF EXISTS public.attack_protocols;`)
		if err != nil {
			fmt.Println("error drop table:", err)
		}

		db.Close()

		os.Unsetenv("POSTGRES_DB")
		os.Unsetenv("POSTGRES_USER")
		os.Unsetenv("POSTGRES_PASSWORD")
	})
}

type Info struct {
	UpdatedAt        time.Time
	CreatedAt        time.Time
	Name             string
	DefaultTProtocol string
	UpdatedBy        string
	CreatedBy        string
	Id               int
	DefaultPort      int
	AttackLevelId    int
}
