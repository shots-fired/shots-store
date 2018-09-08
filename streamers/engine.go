package streamers

import (
	"encoding/json"

	"github.com/shots-fired/shots-store/store"
)

type (
	// Engine describes all the functions that the streamers engine can perform
	Engine interface {
		SetStreamer(field string, val Streamer) error
		GetStreamer(field string) (Streamer, error)
	}

	engine struct {
		Store store.Store
	}
)

// NewEngine return a new engine implementation
func NewEngine(store store.Store) Engine {
	return engine{
		Store: store,
	}
}

// SetStreamer sets a streamer data in the streamers key under the given field
func (e engine) SetStreamer(field string, val Streamer) error {
	return e.Store.Set("streamers", field, val)
}

// GetStreamer returns the streamer under the streamers key with the given field
func (e engine) GetStreamer(field string) (Streamer, error) {
	res, err := e.Store.Get("streamers", field)
	if err != nil {
		return Streamer{}, err
	}

	var streamer Streamer
	unmarshalErr := json.Unmarshal([]byte(res), &streamer)
	if unmarshalErr != nil {
		return Streamer{}, err
	}

	return streamer, nil
}
