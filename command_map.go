package main

import (
	"encoding/json"
	"fmt"

	"github.com/christianrm0821/pokedexcli/internal/pokeapi"
)

func mapLookUp(c *config) error {
	myurl := "https://pokeapi.co/api/v2/" + "/location-area"

	if c.NextLocation != nil {
		myurl = *((*c).NextLocation)
	}

	data, ok := c.Cache.Get(myurl)
	if !ok {
		tmp, err := c.PokeClient.Location_list(c.NextLocation)
		if err != nil {
			return err
		}
		c.Cache.Add(myurl, *tmp)
		data = *tmp
	}

	var res pokeapi.ResponseStruct
	err := json.Unmarshal(data, &res)
	if err != nil {
		return err
	}
	curr := (*c).NextLocation
	(*c).NextLocation = &res.Next
	(*c).PreviousLocation = curr
	fmt.Println()
	for i := 0; i < 20; i++ {
		fmt.Printf("%v. %v", i+1, res.Results[i].Name)
		fmt.Println()
	}
	return nil
}

func mapbLookup(c *config) error {
	var err error
	if c.PreviousLocation == nil {
		err = fmt.Errorf("this is the First page")
		fmt.Println("This is the First page")
		return err
	}
	data, ok := c.Cache.Get(*c.PreviousLocation)
	var res pokeapi.ResponseStruct
	if !ok {
		tmp, err := c.PokeClient.Location_list(c.PreviousLocation)
		if err != nil {
			return err
		}
		c.Cache.Add(*c.PreviousLocation, *tmp)
		data = *tmp
	}

	err = json.Unmarshal(data, &res)
	if err != nil {
		return err
	}

	fmt.Println()
	for i := 0; i < 20; i++ {
		fmt.Printf("%v. %v", i+1, res.Results[i].Name)
		fmt.Println()
	}
	c.NextLocation = &res.Next
	c.PreviousLocation = res.Previous

	return nil
}
