package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/christianrm0821/pokedexcli/internal/pokeapi"
)

func cleanInput(text string) []string {
	mySlice := strings.Fields(strings.ToLower(text))
	return mySlice
}

type cliCommand struct {
	name        string
	description string
	callback    func(c *config) error
}

type config struct {
	PokeClient       pokeapi.Client
	NextLocation     *string
	PreviousLocation *string
}

var commandMap map[string]cliCommand

func GetCommands() map[string]cliCommand {
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
	return commandMap
}

func repl(c *config) {
	fmt.Print("Pokedex > ")
	myScanner := bufio.NewScanner(os.Stdin)
	for {
		if myScanner.Scan() {
			fmt.Print("Pokedex > ")
			input := cleanInput(myScanner.Text())
			mymap := GetCommands()
			val, ok := mymap[input[0]]
			if ok {
				fmt.Printf("%v", val.callback(c))
			} else {
				fmt.Println("Unknown command")
			}
		}

	}
}
