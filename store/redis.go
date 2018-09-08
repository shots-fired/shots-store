package store

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
func New() Store {
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

// Get returns the value under the given id
func (r Redis) Get(id string) (string, error) {
	return r.Client.Get(id).Result()
}
