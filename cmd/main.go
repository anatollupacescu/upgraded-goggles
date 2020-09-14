package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	exchange "github.com/curve/sample/cmd/internal"
)

func main() {
	client := &exchangeRatesHTTPClient{
		endpointURL: os.Getenv("EXCHANGE_RATES_URL"), //TODO fail if not provided
	}

	provider := &exchange.Rates{
		USDWeeklyRates:        []float64{0.83}, //hardcoded sample data, for now
		GBPWeeklyRates:        []float64{1.14},
		ExternalRatesProvider: client.get,
	}

	app := webApp{
		provider: provider,
	}

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/{symbol}", app.serve)

	var port = ":8080" //TODO make port configurable

	log.Printf("Listening on port: %s", port)
	log.Fatal(http.ListenAndServe(port, router))
}
