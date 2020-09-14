package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

type mockProvider struct {
}

func (m mockProvider) Get(string) (float64, string) {
	return 0.0, ""
}

func TestMissingSymbolShouldReturn400(t *testing.T) {
	app := webApp{
		provider: &mockProvider{},
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(app.serve)

	req, err := http.NewRequest("GET", "/", nil)

	// req = mux.SetURLVars(req, map[string]string{"symbol": "USD"})

	if err != nil {
		t.Fatal(err)
	}

	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
}

func TestShouldCallProvider(t *testing.T) {
	t.SkipNow()
}
