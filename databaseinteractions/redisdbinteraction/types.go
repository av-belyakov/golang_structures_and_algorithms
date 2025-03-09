// Пример взаимодействия с Redis DB
package redisdbinteraction

import "github.com/redis/go-redis/v9"

// ConnectionRedis
type ConnectionRedis struct {
	Client *redis.Client
}
