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

func getUrl(config *config) string {
	return ""
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

/*
type Location struct {
	Name string
	URL  string
}
*/

type config struct {
	nextLocation     *string
	previousLocation *string
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
		"mapb": {
			name:        "mapb",
			description: "displays the names of the previous page's 20 location areas in the Pokemon world",
			callback:    mapbLookup,
		},
	}
}

func repl() {
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
		}

	}
}
