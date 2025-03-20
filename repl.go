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
	City             *string
	Area             *string
}
type Pokemon struct {
	Name string
	URL  string
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
		"explore": {
			name:        "explore",
			description: "displays the names of pokemon found in the area",
			callback:    commandExplore,
		},
		"explore-city": {
			name:        "explore-city",
			description: "displays the names of pokemon found in the city(must be a city name)",
			callback:    commandExploreCity,
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
				if input[0] == "explore" {
					if len(input) == 1 {
						fmt.Println("need to input a area")
						fmt.Print("Pokedex > ")
						continue
					}
					c.Area = &input[1]
				}

				if input[0] == "explore-city" {
					if len(input) == 1 {
						fmt.Println("need to input a city")
						fmt.Print("Pokedex > ")
						continue
					}
					c.City = &input[1]
				}

				val.callback(c)
				fmt.Print("Pokedex > ")
			} else {
				fmt.Println("Unknown command")
				fmt.Print("Pokedex > ")
			}
		}

	}
}
