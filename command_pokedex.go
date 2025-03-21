package main

import (
	"fmt"
)

func commandPokedex(c *config) error {
	fmt.Println()
	fmt.Println("Pokedex")
	fmt.Println("-------------------------------------")
	if len(c.Pokedex) == 0 {
		fmt.Println("You have no Pokemon in your pokedex")
		return nil
	}
	for k, _ := range c.Pokedex {
		fmt.Printf("- %v\n", k)
	}
	fmt.Println("-------------------------------------")
	return nil

}
