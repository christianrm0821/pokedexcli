package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (myClient *Client) Location_list(url *string) (ResponseStruct, error) {
	myurl := baseURL + "/location-area"
	if url != nil {
		myurl = *url
	}
	var myResponse ResponseStruct

	req, err := http.NewRequest("GET", myurl, nil)
	if err != nil {
		return myResponse, err
	}

	res, err := (*myClient).httpClient.Do(req)
	if err != nil {
		return myResponse, err
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)

	if err != nil {
		return myResponse, err
	}

	err = json.Unmarshal(data, &myResponse)
	if err != nil {
		return myResponse, nil
	}
	return myResponse, nil
}
