package redisdb

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

// NewConnectionRedis устанавливает новое соединение с БД
func NewConnectionRedis(host string, port int) *ConnectionRedis {
	return &ConnectionRedis{
		Client: redis.NewClient(&redis.Options{
			Addr: fmt.Sprintf("%s:%d", host, port),
		})}
}

// Ping проверка соединения с БД
func (db *ConnectionRedis) Ping(ctx context.Context) (string, error) {
	status := db.Client.Ping(ctx)

	return status.Result()
}

// ConnectionClose закрывает соединение с БД
func (db *ConnectionRedis) ConnectionClose() {
	db.Client.Close()
}
