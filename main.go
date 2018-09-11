package main

import (
	"log"

	"github.com/shots-fired/shots-store/apis"
	"github.com/shots-fired/shots-store/store"
	"github.com/shots-fired/shots-store/streamers"
)

func main() {
	log.Fatal(apis.NewServer(apis.NewRouter(streamers.NewEngine(store.New()))).ListenAndServe())
}
