package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)
	cfg := NewConfig()
	startRepl(cfg)
	err := commandCatch(cfg, r, "pikachu")
	if err != nil {
		fmt.Println("Error:", err)
	}
}
