package main

///Need to fix the explore might need to start over but we will see

import (
	"time"

	"github.com/christianrm0821/pokedexcli/internal/pokeapi"
	"github.com/christianrm0821/pokedexcli/internal/pokecache"
)

func main() {
	c := config{
		PokeClient:       pokeapi.NewClient(10 * time.Second),
		NextLocation:     nil,
		PreviousLocation: nil,
		Cache:            pokecache.NewCache(10 * time.Second),
		City:             nil,
		Area:             nil,
		Pokedex:          make(map[string][]byte),
		Pokemon:          nil,
	}
	repl(&c)
}
