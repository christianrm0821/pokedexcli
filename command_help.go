package main

import "fmt"

func describeCommands() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for key, val := range commandMap {
		fmt.Printf("%v: %v\n", key, val.description)
	}
	return fmt.Errorf("")
}
