package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/christianrm0821/pokedexcli/internal/pokeapi"
	"github.com/christianrm0821/pokedexcli/internal/pokecache"
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
	Cache            pokecache.CacheStruct
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
	myScanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Pokedex > ")
	for {
		if myScanner.Scan() {
			fmt.Print("Pokedex > ")
			input := cleanInput(myScanner.Text())
			mymap := GetCommands()
			if len(input) == 0 {
				fmt.Println("You must enter a command")
				fmt.Print("Pokedex > ")
				continue
			}
			val, ok := mymap[input[0]]
			if ok {
				val.callback(c)
				fmt.Print("Pokedex > ")
			} else {
				fmt.Println("Unknown command")
				fmt.Print("Pokedex > ")
			}
		}

	}
}
