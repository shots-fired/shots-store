package streamers

import "encoding/json"

type (
	// Streamer is the data model for all streamer info
	Streamer struct {
		Name    string `json:"name"`
		Status  string `json:"status"`
		Viewers int    `json:"viewers"`
	}

	// Streamers is a collection of Streamer structs
	Streamers []Streamer
)

// MarshalBinary implements encoding.BinaryMarshaler so Streamer can be saved to the redis cache without any conversion
func (s Streamer) MarshalBinary() (data []byte, err error) {
	return json.Marshal(s)
}
