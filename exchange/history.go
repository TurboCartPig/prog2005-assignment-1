package exchange

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/go-chi/chi"
)

// ExchangeRates describes the rates for exchange between a base currency and some other currency in a time period.
type ExchangeRateHistory struct {
	// Base currency
	Base string `json:"base"`
	// Starting date for lookup
	StartAt string `json:"start_at"`
	// End date for lookup
	EndAt string `json:"end_at"`
	// Rate for date for currency
	Rates map[string]map[string]float32 `json:"rates"`
}

// parseParameters from request into country name, start date and end date.
func parseParameters(r *http.Request) (string, string, string, *ServerError) {
	country := chi.URLParam(r, "country")
	startYear := chi.URLParam(r, "start_yyyy")
	startMonth := chi.URLParam(r, "start_mm")
	startDay := chi.URLParam(r, "start_dd")
	endYear := chi.URLParam(r, "end_yyyy")
	endMonth := chi.URLParam(r, "end_mm")
	endDay := chi.URLParam(r, "end_dd")

	// Format parameters as expected by APIs
	startDateStr := fmt.Sprintf("%s-%s-%s", startYear, startMonth, startDay)
	endDateStr := fmt.Sprintf("%s-%s-%s", endYear, endMonth, endDay)

	// Make sure the parameters make sense
	startDate, startErr := time.Parse(time.RFC3339, startDateStr+"T00:00:00Z")
	endDate, endErr := time.Parse(time.RFC3339, endDateStr+"T00:00:00Z")
	if endDate.Before(startDate) || startErr != nil || endErr != nil {
		return "", "", "", &ServerError{"Bad date format", http.StatusBadRequest}
	}

	return country, startDateStr, endDateStr, nil
}

// GetExchangeRateHistory for a given currency in the given time period, relative to given base currency.
func GetExchangeRateHistory(base string, currency Currency, startDate, endDate string) (ExchangeRateHistory, *ServerError) {
	var rates ExchangeRateHistory

	// Construct URL
	url, _ := url.Parse(ExchangeRatesAPIRoot)
	url.Path += "/history"
	queries := url.Query()
	queries.Set("base", base)
	queries.Set("start_at", startDate)
	queries.Set("end_at", endDate)
	queries.Set("symbols", currency.Code)
	url.RawQuery = queries.Encode()

	// Get the rates
	res, err := http.Get(url.String())
	if err != nil {
		return rates, &ServerError{err.Error(), res.StatusCode}
	}

	err = json.NewDecoder(res.Body).Decode(&rates)
	if err != nil {
		return rates, &ServerError{err.Error(), http.StatusInternalServerError}
	}
	res.Body.Close()

	return rates, nil
}

// HistoryHandler responds to request for currency history for a given country.
func HistoryHandler(rw http.ResponseWriter, r *http.Request) {
	countryName, startDate, endDate, err := parseParameters(r)
	if err != nil {
		http.Error(rw, err.Error(), err.StatusCode)
		return
	}

	// Get country information
	country, err := GetCountryByName(countryName)
	if err != nil {
		http.Error(rw, err.Error(), err.StatusCode)
		return
	}

	// Get currency
	rates, err := GetExchangeRateHistory("EUR", country.Currencies[0], startDate, endDate)
	if err != nil {
		http.Error(rw, err.Error(), err.StatusCode)
		return
	}

	// Write json encoded rates to response body
	json.NewEncoder(rw).Encode(rates)
}
