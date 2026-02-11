package redisdbinteraction_test

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/av-belyakov/golang_structures_and_algorithms/databaseinteractions/redisdbinteraction"
)

//Перед началом тестирования необходимо выполнить
//docker run -d -p 6379:6379 --name redisdb redis:latest
//если нужно подключить уже существую БД --volume /home/artemij/go/src/placeholder_misp/test/redis_file/dump.rdb:/data/dump.rdb

const (
	HostRDb string = "127.0.0.1"
	PortRDb int    = 6379
)

var (
	ctx    context.Context
	cancel context.CancelFunc

	redisClient *redisdbinteraction.ConnectionRedis
)

func TestMain(m *testing.M) {
	ctx, cancel = context.WithCancel(context.Background())

	redisClient = redisdbinteraction.NewConnectionRedis(HostRDb, PortRDb)

	os.Exit(m.Run())
}

func TestGetDumpdb(t *testing.T) {
	defer func() {
		//закрываем соединение с Redis БД
		redisClient.ConnectionClose()

		cancel()
	}()

	t.Run("Тест 1. Получаем эхо ответ от Redis database", func(t *testing.T) {
		response, err := redisClient.Ping(ctx)
		assert.NoError(t, err)
		assert.NotEmpty(t, response)
		assert.Equal(t, response, "PONG")
	})

	t.Run("Тест 2. Запись ключ:значение", func(t *testing.T) {
		// .Set(контекст, ключ, значение, время жизни в базе данных)
		err := redisClient.Client.Set(context.Background(), "key1", "test value", 0).Err()
		assert.NoError(t, err)

		err = redisClient.Client.Set(context.Background(), "key2", 333, 30*time.Second).Err()
		assert.NoError(t, err)
	})

	t.Run("Тест 3. Получить значение по ключу", func(t *testing.T) {
		result := redisClient.Client.Get(context.Background(), "key1")
		res, err := result.Result()
		assert.NoError(t, err)
		assert.Equal(t, res, "test value")
	})

	t.Run("Тест 4. Работа со списками", func(t *testing.T) {
		//.RPush(контекст, ключ, список значений)
		err := redisClient.Client.RPush(ctx, "colors", "orange", "green", "red", "blue", "white", "black", "grey").Err()
		assert.NoError(t, err)

		//0, -1 получить всё
		list, err := redisClient.Client.LRange(ctx, "colors", 0, -1).Result()
		assert.NoError(t, err)
		assert.Equal(t, len(list), 7)
	})

	t.Run("Тест 5. Работа с хешами", func(t *testing.T) {
		err := redisClient.Client.HSet(ctx, "user:1001", "name", "Olga", "country", "Russia", "age", 31).Err()
		assert.NoError(t, err)

		res, err := redisClient.Client.HGet(ctx, "user:1001", "name").Result()
		assert.NoError(t, err)
		t.Log("result:", res)

		resAll, err := redisClient.Client.HGetAll(ctx, "user:1001").Result()
		assert.NoError(t, err)
		assert.NotEqual(t, len(resAll), 0)

		for k, v := range resAll {
			t.Logf("%s:%s\n", k, v)
		}
	})

	t.Run("Тест 6. Работа со множествами", func(t *testing.T) {
		err := redisClient.Client.SAdd(ctx, "tags", "golang", "redis", "backend").Err()
		assert.NoError(t, err)

		tags, err := redisClient.Client.SMembers(ctx, "tags").Result()
		assert.NoError(t, err)
		assert.Equal(t, len(tags), 3)
	})
}
