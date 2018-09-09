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
		GetAllStreamers() (Streamers, error)
		DeleteStreamer(field string) error
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

// GetAllStreamers returns a slice of all streamers in the streamers key
func (e engine) GetAllStreamers() (Streamers, error) {
	res, err := e.Store.GetAll("streamers")
	if err != nil {
		return Streamers{}, err
	}

	var streamers Streamers
	for _, v := range res {
		var streamer Streamer
		unmarshalErr := json.Unmarshal([]byte(v), &streamer)
		if unmarshalErr != nil {
			return Streamers{}, unmarshalErr
		}
		streamers = append(streamers, streamer)
	}

	return streamers, nil
}

// DeleteStreamer deletes the streamer with the key of field
func (e engine) DeleteStreamer(field string) error {
	return e.Store.Delete("streamers", field)
}
