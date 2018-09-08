package main

import (
	"github.com/shots-fired/shots-store/apis"
	"github.com/shots-fired/shots-store/store"
	"github.com/shots-fired/shots-store/streamers"
)

func main() {
	r := store.New()
	e := streamers.NewEngine(r)
	apis.HostAPIs(e)
}
