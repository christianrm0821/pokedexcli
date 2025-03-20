package Explore

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetLocationData(city string) (*LocationStruct, error) {
	myurl := baseURL + "/location/" + city

	res, err := http.Get(myurl)
	if err != nil {
		return nil, err

	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: status code %d", res.StatusCode)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var responseLocation LocationStruct
	err = json.Unmarshal(data, &responseLocation)
	if err != nil {
		return nil, err
	}
	return &responseLocation, nil

}

func GetPokemonCity(city string) ([]string, error) {
	locationData, err := GetLocationData(city)
	if err != nil {
		return nil, err
	}
	var pokemon []string

	for _, val := range (*locationData).Areas {
		areaName := val.Name
		myurl := baseURL + "location-area/" + areaName

		res, err := http.Get(myurl)
		if err != nil {
			return nil, err
		}

		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("error: status code %d", res.StatusCode)
		}

		data, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}

		var area LocationArea
		err = json.Unmarshal(data, &area)
		if err != nil {
			return nil, err
		}
		for _, val := range area.PokemonEncounters {
			pokemon = append(pokemon, val.Pokemon.Name)
		}
	}
	return pokemon, nil
}

func GetPokemonArea(area string) ([]string, error) {
	myurl := baseURL + "location-area/" + area
	res, err := http.Get(myurl)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: status code %d", res.StatusCode)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var myArea LocationArea
	var myPokemon []string
	err = json.Unmarshal(data, &myArea)
	if err != nil {
		return nil, err
	}
	for _, val := range myArea.PokemonEncounters {
		myPokemon = append(myPokemon, val.Pokemon.Name)
	}
	return myPokemon, nil
}
