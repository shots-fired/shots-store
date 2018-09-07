package main

import (
	"fmt"

	"github.com/shots-fired/shots-store/redis"
)

func main() {
	r := redis.New()
	pong, err := r.Ping()
	fmt.Println(pong, err)
}
