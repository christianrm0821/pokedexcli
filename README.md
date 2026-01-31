# pokedexcli
A command-line Pokédex built in Go that interacts with the PokéAPI to explore Pokémon locations, catch and inspect Pokémon, and manage your personal Pokédex.

## 📦 Features
Makes HTTP requests to the PokéAPI to fetch live Pokémon data

Parses JSON responses in Go for meaningful display

Interactive command-line interface for exploring and managing your Pokédex

Implements in-memory caching to improve performance and reduce redundant API calls

## 👥 Who It's For
Fans of Pokémon who enjoy the command line and want a fun way to explore the Pokémon world interactively.

## ⚙️ Installation
Make sure you have Go installed. Then, clone the repository and run the program:

bash
Copy
Edit
go run main.go
No external dependencies are required—only Go's standard library is used.

## 📚 Libraries Used
net/http – for making API requests

encoding/json – to parse JSON responses

fmt, os, bufio, strings, time, math/rand – for CLI interaction and caching

## 💡 Available Commands

#Command	Description
help	Lists and describes all available commands

exit	Exits the application

map	Displays the next 20 Pokémon location areas

mapb	Displays the previous 20 Pokémon location areas

explore	Lists the Pokémon found in a selected area

explore-city	Lists Pokémon found in a given city (requires city name)

catch <name>	Attempts to catch the specified Pokémon

inspect <name>	Shows detailed information about a caught Pokémon

pokedex	Lists all Pokémon you have successfully caught

## ⚠️ Limitations & Notes
Requires valid area/city names as input for some commands

Data depends on the PokéAPI, so availability and performance may vary

Caught Pokémon are stored only in-memory and reset on restart

## 🌱 Future Improvements
Save Pokédex data between sessions

Add more game-like elements (XP, battles, evolutions)

Support fuzzy search for area and Pokémon names
