package main

import (
	"fmt"

	"github.com/shots-fired/shots-store/store"
	"github.com/shots-fired/shots-store/streamers"
)

func main() {
	r := store.New()
	e := streamers.NewEngine(r)
	err := e.SetStreamer("1", streamers.Streamer{
		Name:    "Jake",
		Status:  "online",
		Viewers: 5,
	})
	fmt.Printf("Set error: %v\n", err)
	res, err2 := e.GetStreamer("1")
	fmt.Printf("Get result: %+v, %v", res, err2)
}
