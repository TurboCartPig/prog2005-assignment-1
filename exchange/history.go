package exchange

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
)

// Parse parameters from request into country, start date and end date.
func parseParameters(r *http.Request) (string, string, string, error) {
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
		log.Printf("Bad parameter format")
		return "", "", "", errors.New("Bad Format")
	}

	return country, startDateStr, endDateStr, nil
}

func HistoryHandler(rw http.ResponseWriter, r *http.Request) {
	country, startDate, endDate, err := parseParameters(r)
	if err != nil {
		http.Error(rw, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	// Get country information from restcountries.eu
	countries, status := GetCountries(country)
	if status != http.StatusOK {
		http.Error(rw, http.StatusText(status), status)
		return
	}

	if len(countries) > 1 {
		log.Println("Given parameter matched more the one country. Make sure the country name is spelled correctly.")
	}

	// Get currency
	rates, status := GetExchangeRates("EUR", countries[0].Currencies, startDate, endDate)
	if status != http.StatusOK {
		http.Error(rw, http.StatusText(status), status)
		return
	}

	// Write json encoded rates to response body
	json.NewEncoder(rw).Encode(rates)
}
