package main

import "fmt"

func mapLookUp(c *config) error {
	//res, err := c.PokeClient.Location_list(strPtr)
	//client_pointer := c.PokeClient

	/*(var strPtr *string
	if (*c).PreviousLocation == nil {
		strPtr = nil
	} else {
		strPtr = c.NextLocation
	}
	*/
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
	return nil
}
