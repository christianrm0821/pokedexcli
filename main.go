package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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
	res, err := http.Get("https://pokeapi.co/api/v2/location-area")
	if err != nil {
		return fmt.Errorf("error with get request: %w", err)
	}
	defer res.Body.Close()

	bytes, err2 := io.ReadAll(res.Body)
	if err2 != nil {
		return fmt.Errorf("error reading res.body: %w", err2)
	}

	mapLocations := MapStruct{}

	err = json.Unmarshal(bytes, &mapLocations)
	if err != nil {
		return fmt.Errorf("error with unmarshal/creating request: %w", err)
	}

	return fmt.Errorf("")
}

func getUrl(config *config) string {
	return ""
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
	config      *config
}

type Location struct {
	Name string
	URL  string
}

type Named struct {
	Count    int
	Next     string
	Previous *string
	Results  []struct {
		Name string
		URL  string
	}
}

type config struct {
	Next     string
	Previous *string
}

type MapStruct struct {
	Location Location
	Named    Named
}

var commandMap map[string]cliCommand

func init() {
	commandMap = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the pokedex",
			callback:    commandExit,
			config:      *config,
		},
		"help": {
			name:        "help",
			description: "Gives description of all commands",
			callback:    describeCommands,
			config:      *config,
		},
		"map": {
			name:        "map",
			description: "displays the names of 20 location areas in the Pokemon world",
			callback:    mapLookUp,
			config:      *config,
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
		}

	}

}
