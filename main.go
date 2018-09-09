package main

import (
	"github.com/shots-fired/shots-store/apis"
	"github.com/shots-fired/shots-store/store"
	"github.com/shots-fired/shots-store/streamers"
)

func main() {
	s := store.New()
	e := streamers.NewEngine(s)
	r := apis.New(e)
	apis.Host(r)
}
