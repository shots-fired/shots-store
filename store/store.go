package store

type (
	// Store describes our store implementation
	Store interface {
		Ping() (string, error)
		Get(key, field string) (string, error)
		GetAll(key string) (map[string]string, error)
		Set(key, field string, value interface{}) error
		Delete(key, field string) error
	}
)
