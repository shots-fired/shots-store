package store

import (
	"os"

	"github.com/go-redis/redis"
)

type (
	redisClient struct {
		Client *redis.Client
	}
)

// New returns a new Redis implementation
func New() Store {
	return &redisClient{
		Client: redis.NewClient(&redis.Options{
			Addr:     os.Getenv("REDIS_ADDRESS"),
			Password: os.Getenv("REDIS_PASSWORD"),
			DB:       0,
		}),
	}
}

// Ping pings the redis instance
func (r redisClient) Ping() (string, error) {
	return r.Client.Ping().Result()
}

// Get returns the value for the key and field
func (r redisClient) Get(key, field string) (string, error) {
	return r.Client.HGet(key, field).Result()
}

// GetAll returns all values for the given key
func (r redisClient) GetAll(key string) (map[string]string, error) {
	return r.Client.HGetAll(key).Result()
}

// Set is used to set a value for a specific field under a key
func (r redisClient) Set(key, field string, value interface{}) error {
	return r.Client.HSet(key, field, value).Err()
}

// Delete deletes a field within a key
func (r redisClient) Delete(key, field string) error {
	return r.Client.HDel(key, field).Err()
}
