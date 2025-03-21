package main

import (
	"encoding/json"
	"fmt"

	"github.com/christianrm0821/pokedexcli/internal/pokemon"
)

func commandInspect(c *config) error {
	fmt.Println()
	fmt.Println("---------------------------------------")
	val, ok := (*c).Pokedex[*(c.Pokemon)]
	if !ok {
		fmt.Println("you have not caught that pokemon")
		return nil
	}
	var myStruct pokemon.PokemonStruct
	err := json.Unmarshal(val, &myStruct)
	if err != nil {
		return err
	}
	fmt.Printf("Name: %v\n", myStruct.Name)
	fmt.Printf("Height: %v\n", myStruct.Height)
	fmt.Printf("Weight: %v\n", myStruct.Weight)
	fmt.Println()
	for _, val := range myStruct.Stats {
		fmt.Printf("%v: %v\n", val.Stat.Name, val.BaseStat)
	}
	fmt.Println()
	fmt.Println("Type: ")
	for _, v := range myStruct.Types {
		fmt.Printf("- %v\n", v.Type.Name)
	}
	fmt.Println("---------------------------------------")
	return nil
}

/*

type PokemonStruct struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
	Height         int    `json:"height"`
	Weight         int    `json:"weight"`
	Stats          []struct {
		BaseStat int `json:"base_stat"`
		Effort   int `json:"effort"`
		Stat     struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Type struct {
			Name string `json:"name"`
		} `json:"type"`
	} `json:"types"`
}
*/
