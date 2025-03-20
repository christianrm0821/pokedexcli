package pokeapi

import (
	"io"
	"net/http"
)

func (myClient *Client) Location_list(url *string) (*[]byte, error) {
	myurl := baseURL + "/location-area"
	if url != nil {
		myurl = *url
	}

	req, err := http.NewRequest("GET", myurl, nil)
	if err != nil {
		return nil, err
	}

	res, err := (*myClient).httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}
	return &data, nil
}
