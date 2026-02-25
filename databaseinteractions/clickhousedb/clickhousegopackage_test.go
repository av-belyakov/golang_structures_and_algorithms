package clickhousedb_test

import (
	"log"
	"os"
	"strconv"
	"testing"

	"github.com/av-belyakov/golang_structures_and_algorithms/databaseinteractions/clickhousedb"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

// для подробной информации https://clickhouse.com/docs/integrations/go

func TestClickhouseGoPackage(t *testing.T) {
	if err := godotenv.Load("./clickhouseimage/.env"); err != nil {
		log.Fatal(err)
	}

	/*
		Database: os.Getenv("CLICKHOUSE_SERVER_DB"),
		Username: os.Getenv("CLICKHOUSE_SERVER_USER"),
		Password: os.Getenv("CLICKHOUSE_SERVER_PASSWORD"),
	*/
	port, err := strconv.ParseInt(os.Getenv("CLICKHOUSE_SERVER_YOUR_OWN_PORT"), 10, 32)
	if err != nil {
		log.Fatal(err)
	}

	client, err := clickhousedb.New(
		clickhousedb.WithHost(os.Getenv("CLICKHOUSE_SERVER_HOST")),
		clickhousedb.WithPort(int(port)),
		clickhousedb.WithDatabase(os.Getenv("CLICKHOUSE_SERVER_DB")),
		clickhousedb.WithUser(os.Getenv("CLICKHOUSE_SERVER_USER")),
		clickhousedb.WithPassword(os.Getenv("CLICKHOUSE_SERVER_PASSWORD")),
	)
	if err != nil {
		log.Fatal(err)
	}

	conn, err := client.Connect(t.Context())
	if err != nil {
		log.Fatal(err)
	}

	t.Run("Тест 1. Получаем информацию о сервере", func(t *testing.T) {
		drServerVersion, err := conn.ServerVersion()
		assert.NoError(t, err)
		assert.Greater(t, len(drServerVersion.String()), 0)

		//fmt.Println("___ Server version:", drServerVersion.String())
	})

	t.Run("Тест 2. Добавляем новые записи", func(t *testing.T) {
		t.Run("Тест 2.1. Создаём тестовую таблицу", func(t *testing.T) {
			assert.NoError(t, conn.Exec(
				t.Context(),
				`CREATE TABLE IF NOT EXISTS real_estate_objects_example (
					uuid UUID,
					object_name String,
					country String,
					address String,
					geographical_coordinate_latitude Float64,
					geographical_coordinate_longitude Float64,
					date_of_construction Date,
					date_of_commissioning Date,
					total_area Float64,
					living_area Float64,
					price Float64
					)`,
			))
		})

		t.Run("Тест 2.2. Заполняем тестовую таблицу данными", func(t *testing.T) {
			for range 350 {
				assert.NoError(t, conn.Exec(t.Context(),
					`INSERT INTO real_estate_objects_example(
					uuid,
					object_name,
					country,
					address,
					geographical_coordinate_latitude,
					geographical_coordinate_longitude,
					date_of_construction,
					date_of_commissioning,
					total_area,
					living_area,
					price
				) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11);`,
					gofakeit.UUID(),
					gofakeit.Name(),
					gofakeit.Country(),
					gofakeit.Address().Address,
					gofakeit.Latitude(),
					gofakeit.Longitude(),
					gofakeit.Date(),
					gofakeit.Date(),
					gofakeit.Float64(),
					gofakeit.Float64(),
					gofakeit.Price(1_000_000, 1_000_000_000),
				))
			}
		})
	})

	//t.Run("Тест 3. Удаление тестовой таблицы", func(t *testing.T) {
	//	conn.Exec(context.Background(), `DROP TABLE IF EXISTS real_estate_objects_example`)
	//})

	t.Cleanup(func() {
		os.Unsetenv("CLICKHOUSE_SERVER_HOST")
		os.Unsetenv("CLICKHOUSE_SERVER_YOUR_OWN_PORT")

		os.Unsetenv("CLICKHOUSE_SERVER_DB")
		os.Unsetenv("CLICKHOUSE_SERVER_USER")
		os.Unsetenv("CLICKHOUSE_SERVER_PASSWORD")

		conn.Close()
	})
}
