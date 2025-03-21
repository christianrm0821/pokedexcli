package pokemon

import (
	"io"
	"net/http"
)

func PokemonData(url string) ([]byte, error) {
	myurl := "https://pokeapi.co/api/v2/pokemon/" + url

	res, err := http.Get(myurl)
	if err != nil {
		return nil, nil
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return data, nil
}
