package main

import (
	"bufio"
	//	"encoding/json"
	"fmt"
	//	"io"
	//	"net/http"
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
		}

	}

}
