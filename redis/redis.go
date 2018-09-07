package redis

import (
	"github.com/go-redis/redis"
)

type (
	// Redis is the wrapper around the redis client
	Redis struct {
		Client *redis.Client
	}
)

// New returns a new Redis implementation
func New() *Redis {
	return &Redis{
		Client: redis.NewClient(&redis.Options{
			Addr:     "redis:6379",
			Password: "", // no password set
			DB:       0,  // use default DB
		}),
	}
}

// Ping pings the redis instance
func (r Redis) Ping() (string, error) {
	return r.Client.Ping().Result()
}
