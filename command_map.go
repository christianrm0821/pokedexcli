package main

import "fmt"

func mapLookUp(c *config) error {
	res, err := c.PokeClient.Location_list(c.NextLocation)
	if err != nil {
		return err
	}
	curr := (*c).NextLocation
	(*c).NextLocation = &res.Next
	(*c).PreviousLocation = curr
	for i := 0; i < 20; i++ {
		fmt.Printf("%v", res.Results[i].Name)
		fmt.Println()
	}
	return nil
}

func mapbLookup(c *config) error {
	res, err := c.PokeClient.Location_list(c.PreviousLocation)
	if err != nil {
		return err
	}
	if res.Previous == nil {
		fmt.Println("This is the First page")
		return err
	}
	for i := 0; i < 20; i++ {
		fmt.Printf("%v", res.Results[i].Name)
		fmt.Println()
	}
	c.NextLocation = &res.Next
	c.PreviousLocation = res.Previous

	return nil
}
