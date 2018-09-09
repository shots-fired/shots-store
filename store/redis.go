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

// Get returns the value for the key and field
func (r Redis) Get(key, field string) (string, error) {
	return r.Client.HGet(key, field).Result()
}

// GetAll returns all values for the given key
func (r Redis) GetAll(key string) (map[string]string, error) {
	return r.Client.HGetAll(key).Result()
}

// Set is used to set a value for a specific field under a key
func (r Redis) Set(key, field string, value interface{}) error {
	return r.Client.HSet(key, field, value).Err()
}

// Delete deletes a field within a key
func (r Redis) Delete(key, field string) error {
	return r.Client.HDel(key, field).Err()
}
