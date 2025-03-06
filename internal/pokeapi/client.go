package pokeapi

import (
	"net/http"
	"time"
)

type Client struct {
	httpClient http.Client
}

// NewClient -
func NewClient(timeout time.Duration) Client {
	myClient := Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
	return myClient
}
