package exchange

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestDiagEndpoint(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/exchange/v1/diag", nil)
	if err != nil {
		t.Fatal(err.Error())
	}

	rr := httptest.NewRecorder()
	handler := NewDiagHandler(time.Now(), "v1")
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Handler returned incorrect status code. Wanted: %d, got: %d", http.StatusOK, rr.Code)
	}

	var body diag
	err = json.NewDecoder(rr.Body).Decode(&body)
	if err != nil {
		t.Fatal(err.Error())
	}

	if body.ExchangeRatesAPI != http.StatusOK {
		t.Errorf("Handler returned incorrect status code. Wanted: %d, got: %d", http.StatusOK, body.ExchangeRatesAPI)
	}
	if body.RestCountries != http.StatusOK {
		t.Errorf("Handler returned incorrect status code. Wanted: %d, got: %d", http.StatusOK, body.RestCountries)
	}
}
