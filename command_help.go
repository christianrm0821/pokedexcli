package main

import "fmt"

func describeCommands(c *config) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	//fmt.Printf("%v", len(commandMap))
	//fmt.Println()
	for key, val := range GetCommands() {
		fmt.Printf("%v: %v\n", key, val.description)
	}
	return nil
}
