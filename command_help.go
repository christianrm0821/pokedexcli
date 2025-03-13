package main

import "fmt"

func describeCommands(c *config) error {
	fmt.Println()
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("--------------------------------")
	i := 1
	for key, val := range GetCommands() {
		fmt.Printf("%v. %v: %v\n", i, key, val.description)
		i++
		//fmt.Println()
	}
	fmt.Println("--------------------------------")
	return nil
}
