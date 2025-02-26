package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cleanInput(text string) []string {
	mySlice := strings.Fields(strings.ToLower(text))
	return mySlice
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return fmt.Errorf("Closing the Pokedex... Goodbye!")
}

func describeCommands() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for key, val := range commandMap {
		fmt.Printf("%v: %v\n", key, val.description)
	}
	return fmt.Errorf("")
}

func mapLookUp() error {
	return fmt.Errorf("")

}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

var commandMap map[string]cliCommand

func init() {
	commandMap = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Gives description of all commands",
			callback:    describeCommands,
		},
		"map": {
			name:        "map",
			description: "displays the names of 20 location areas in the Pokemon world",
			callback:    mapLookUp,
		},
	}
}

func main() {
	myScanner := bufio.NewScanner(os.Stdin)
	for {
		if myScanner.Scan() {
			fmt.Print("Pokedex > ")
			input := cleanInput(myScanner.Text())
			val, ok := commandMap[input[0]]
			if ok {
				fmt.Printf("%v", val.callback())
			} else {
				fmt.Println("Unknown command")
			}
			//firstWord := input[0]
			//fmt.Printf("Your command was: %v", firstWord)
		}

	}

}
