package main

import (
	"encoding/json"
	"fmt"

	"github.com/christianrm0821/pokedexcli/internal/Explore"
)

func commandExplore(c *config) error {
	myurl := "https://pokeapi.co/api/v2/location-area/" + *(c.Area)
	var err error
	data, tmp := c.Cache.Get(myurl)
	if !tmp {
		data, err = Explore.GetPokemonArea(*(c.Area))
		if err != nil {
			return err
		}
		c.Cache.Add(myurl, data)
	}
	var myArea Explore.LocationArea
	var myPokemon []string
	err = json.Unmarshal(data, &myArea)

	if err != nil {
		fmt.Println("Make sure you typed the Area name correctly")
		return err
	}

	for _, val := range myArea.PokemonEncounters {
		myPokemon = append(myPokemon, val.Pokemon.Name)
	}

	fmt.Println()
	fmt.Println("Exploring " + *(c.Area) + "...")
	for _, val := range myPokemon {
		fmt.Printf("- %v", val)
		fmt.Println()
	}
	return nil
}
