package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type exchangeRatesHTTPClient struct {
	endpointURL string
}

func (r *exchangeRatesHTTPClient) get(symbol string) (float64, error) {
	response, err := http.Get(r.endpointURL)

	if err != nil {
		return 0.0, err
	}

	d := json.NewDecoder(response.Body)

	var payload struct {
		Rates map[string]float64 `json:"rates"`
	}

	if err = d.Decode(&payload); err != nil {
		return 0.0, fmt.Errorf("parse response body: %w", err)
	}

	rate, found := payload.Rates[symbol]

	if !found {
		return 0.0, errors.New("rate not found")
	}

	return rate, nil
}
