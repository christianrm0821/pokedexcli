package main

import (
	"fmt"

	"github.com/christianrm0821/pokedexcli/internal/Explore"
)

func commandExplore(c *config) error {
	pokemon, err := Explore.GetPokemonArea(*(c.Area))
	fmt.Println()
	if err != nil {
		fmt.Println("Make sure you typed the Area name correctly")
		return err
	}
	fmt.Println("Exploring " + *(c.Area) + "...")
	for _, val := range pokemon {
		fmt.Printf("- %v", val)
		fmt.Println()
	}
	return nil
}
