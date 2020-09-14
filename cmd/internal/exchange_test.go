package exchange

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNoSuggestion(t *testing.T) {
	r := Rates{
		ExternalRatesProvider: func(s string) (float64, error) {
			assert.Equal(t, s, "GBP")
			return 4.0, nil
		},
	}

	res, suggestion := r.Get("GBP")

	assert.Equal(t, "NO DATA", suggestion)
	assert.Equal(t, 4.0, res)
}

func TestSuggestBuy(t *testing.T) {
	t.SkipNow()
}

func TestSuggestSell(t *testing.T) {
	t.SkipNow()
}
