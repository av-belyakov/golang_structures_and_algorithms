package elasticsearchexamples_test

import (
	"testing"

	elasticsearchexamples "github.com/av-belyakov/golang_structures_and_algorithms/databaseinteractions/elasticsearch"
	"github.com/stretchr/testify/assert"
)

const (
	Host = "localhost"
	Port = 9200
	// получаем после создания пользователя
	User   = "development"
	Passwd = "7kY-bQv-HNy-vwx"
	// получаем после создания ApiKey
	ApiKey = "VTkxclRKd0J3WGVtNzJadzlnamk6ZTBMU3BzSTdoWmc0NmpXNUkzM3Mxdw=="
)

func TestElasticsearchIntaraction(t *testing.T) {
	t.Run("Тест 1.1. Подключение к БД с использованием API-ключа.", func(t *testing.T) {
		es, err := elasticsearchexamples.NewClient(
			elasticsearchexamples.WithApiKey(ApiKey),
			elasticsearchexamples.WithHost(Host),
			elasticsearchexamples.WithPort(Port),
			elasticsearchexamples.WithStorage(map[string]string{
				"alert": "alert_indexes",
			}),
		)
		assert.NoError(t, err)
		assert.NoError(t, es.Connect(t.Context()))
	})

	t.Run("Тест 1.2. Подключение к БД с использованием имени пользователя и пароля.", func(t *testing.T) {
		es, err := elasticsearchexamples.NewClient(
			elasticsearchexamples.WithUser(User),
			elasticsearchexamples.WithPasswd(Passwd),
			elasticsearchexamples.WithHost(Host),
			elasticsearchexamples.WithPort(Port),
			elasticsearchexamples.WithStorage(map[string]string{
				"alert": "alert_indexes",
			}),
		)
		assert.NoError(t, err)
		assert.NoError(t, es.Connect(t.Context()))
	})

	t.Run("Тест 2. Выполнение различных запросов к БД.", func(t *testing.T) {
		// ...
	})
}
