package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

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
