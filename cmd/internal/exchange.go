package exchange

import "log"

type Rates struct {
	USDWeeklyRates []float64
	GBPWeeklyRates []float64

	ExternalRatesProvider func(string) (float64, error)
}

func (r *Rates) Get(symbol string) (float64, string) {
	currentRate, err := r.ExternalRatesProvider(symbol)
	if err != nil {
		log.Fatal(err)
	}

	historic := r.USDWeeklyRates

	if symbol == "GBP" {
		historic = r.GBPWeeklyRates
	}

	if len(historic) == 0 {
		return currentRate, "NO DATA"
	}

	lastWeek := historic[len(historic)-1]

	var suggestion = "SELL"

	if lastWeek > currentRate {
		suggestion = "BUY"
	}

	return currentRate, suggestion
}
