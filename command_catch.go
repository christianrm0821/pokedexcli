package main

import (
	"encoding/json"
	"fmt"
	"math/rand"

	"github.com/christianrm0821/pokedexcli/internal/pokemon"
)

func commandCatch(c *config) error {
	fmt.Printf("Throwing a Pokeball at %v", *(c.Pokemon))
	fmt.Println()
	pokeData, err := pokemon.PokemonData(*c.Pokemon)
	if err != nil {
		fmt.Print("error 1")
		return err

	}
	var pokiStruct pokemon.PokemonStruct
	err = json.Unmarshal(pokeData, &pokiStruct)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		fmt.Println("make sure you spelt the name correctly")
		fmt.Println()
		return err
	}
	options := pokiStruct.BaseExperience / 75
	options++

	randomNum := rand.Intn(options)
	randomNum2 := rand.Intn(options)

	if randomNum == randomNum2 {
		fmt.Printf("%v was caught!\n", *(c.Pokemon))
		(*c).Pokedex[*(c.Pokemon)] = pokeData
	} else {
		fmt.Printf("%v escaped!\n", *(c.Pokemon))
	}
	return nil
}
