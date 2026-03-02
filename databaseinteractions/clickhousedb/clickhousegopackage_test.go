package clickhousedb_test

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"

	"github.com/av-belyakov/golang_structures_and_algorithms/databaseinteractions/clickhousedb"
)

// для подробной информации https://clickhouse.com/docs/ru/integrations/go

func TestClickhouseGoPackage(t *testing.T) {
	if err := godotenv.Load("./clickhouseimage/.env"); err != nil {
		log.Fatal(err)
	}

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
					email String,
					geographical_coordinate_latitude Float64,
					geographical_coordinate_longitude Float64,
					date_of_construction Date,
					date_of_commissioning Date,
					total_area INT,
					living_area INT,
					price Float64
					)
				ENGINE = MergeTree()
				PRIMARY KEY (object_name, country)`,
			))
		})

		t.Run("Тест 2.2. Заполняем тестовую таблицу данными", func(t *testing.T) {
			for range 1_000 {
				assert.NoError(t, conn.Exec(t.Context(),
					`INSERT INTO real_estate_objects_example(
					uuid,
					object_name,
					country,
					address,
					email,
					geographical_coordinate_latitude,
					geographical_coordinate_longitude,
					date_of_construction,
					date_of_commissioning,
					total_area,
					living_area,
					price
				) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12);`,
					gofakeit.UUID(),
					gofakeit.Name(),
					gofakeit.Country(),
					gofakeit.Address().Address,
					gofakeit.Email(),
					gofakeit.Latitude(),
					gofakeit.Longitude(),
					gofakeit.Date(),
					gofakeit.Date(),
					gofakeit.IntRange(50, 3_000),
					gofakeit.IntRange(30, 2_500),
					gofakeit.Price(1_000_000, 1_000_000_000),
				))
			}
		})
	})

	t.Run("Тест 3. Выполнить запросы к БД", func(t *testing.T) {
		t.Run("Тест 3.1. Поиск по стране", func(t *testing.T) {
			rows, err := conn.Query(
				t.Context(),
				`SELECT uuid, country, address, email, date_of_construction, price
				FROM real_estate_objects_example 
				WHERE country IN ('Russia', 'Franc', 'USA', 'China', 'Spain', 'Niger');`,
			)
			assert.NoError(t, err)

			num := 1
			for rows.Next() {
				var (
					uuid                 string
					country              string
					address              string
					email                string
					date_of_construction time.Time
					price                float64
				)

				assert.NoError(t, rows.Scan(&uuid, &country, &address, &email, &date_of_construction, &price))

				fmt.Printf(
					"%d. Found country:'%s', uuid:'%s', address:'%s', email:'%s', date_of_construction:'%v', price:'%v'\n",
					num, country, uuid, address, email, date_of_construction, price,
				)

				num++
			}

			rows.Close()
		})

		t.Run("Тест 3.2. Обновление данных", func(t *testing.T) {
			construction := "2002-02-17"
			commissioning := "2021-01-01"

			assert.NoError(t, conn.Exec(
				t.Context(),
				`ALTER TABLE clickhouse_database_example.real_estate_objects_example
				UPDATE date_of_construction = '2002-02-17', date_of_commissioning = '2021-01-01'
				WHERE date_of_construction = '1970-01-01' OR date_of_commissioning = '1970-01-01';`,
			))

			rows, err := conn.Query(
				t.Context(),
				`SELECT uuid, object_name 
				FROM real_estate_objects_example 
				WHERE date_of_construction = $1 AND date_of_commissioning = $2;`,
				construction,
				commissioning,
			)
			assert.NoError(t, err)

			num := 1
			fmt.Println("Found next items:")
			for rows.Next() {
				var uuid, objectName string
				assert.NoError(t, rows.Scan(&uuid, &objectName))

				fmt.Printf("%d. %s %s\n", num, uuid, objectName)

				num++
			}

			rows.Close()
		})
	})

	t.Run("Тест 4. Генерация случайных данных средствами Clickhouse", func(t *testing.T) {
		assert.NoError(t, conn.Exec(
			t.Context(),
			`CREATE TABLE IF NOT EXISTS my_test_table_example
			ENGINE = MergeTree
			ORDER BY tuple()
			AS SELECT * 
			FROM generateRandom(
				'col1 UInt32, col2 String, col3 Float64, col4 DateTime',
				1,  -- seed for data generation
				10  -- number of different random values
			)
			LIMIT 100;`,
		))
	})

	//t.Run("Тест 5. Удаление тестовой таблицы", func(t *testing.T) {
	//	conn.Exec(t.Context(), `DROP TABLE IF EXISTS real_estate_objects_example`)
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
