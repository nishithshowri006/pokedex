package main

import (
	"time"

	"github.com/nishithshowri006/pokedex/internal/pokecache"
)

func main() {
	c := pokecache.NewCache(time.Second * 5)
	startRepl(c)
}
