package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type ExchangeRateProvider interface {
	Get(string) (float64, string)
}

type webApp struct {
	provider ExchangeRateProvider
}

func (app webApp) serve(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	symbol := vars["symbol"]
	if len(symbol) == 0 {
		http.Error(w, "bad symbol", http.StatusBadRequest)
		return
	}

	exchangeRate, suggestion := app.provider.Get(symbol)

	var res = struct {
		EUR        float64 `json:"EUR"`
		Suggestion string  `json:"suggestion"`
	}{
		EUR:        exchangeRate,
		Suggestion: suggestion,
	}

	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(res)
	if err != nil {
		http.Error(w, "encode response", http.StatusInternalServerError)
	}
}
