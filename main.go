package main

import (
	"time"

	"github.com/christianrm0821/pokedexcli/internal/pokeapi"
)

func main() {
	c := config{
		PokeClient:       pokeapi.NewClient(10 * time.Second),
		NextLocation:     nil,
		PreviousLocation: nil,
	}
	repl(&c)
}
