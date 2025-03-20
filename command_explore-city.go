package main

import (
	"fmt"

	"github.com/christianrm0821/pokedexcli/internal/Explore"
)

func commandExploreCity(c *config) error {
	pokemon, err := Explore.GetPokemonCity(*(c.City))
	fmt.Println()
	if err != nil {
		fmt.Println("Make sure you typed the city name correctly")
		return err
	}
	fmt.Println("Exploring " + (*c.City) + "...")
	for _, val := range pokemon {
		fmt.Printf("- %v", val)
		fmt.Println()
	}
	return nil
}
