package store

type (
	// Store describes our store implementation
	Store interface {
		Ping() (string, error)
		Get(id string) (string, error)
	}
)
