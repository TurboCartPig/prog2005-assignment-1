package exchange

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"strconv"

	"github.com/go-chi/chi"
)

// rate is a exchange rate used in ExchangeRates.
type rate struct {
	Currency string  `json:"currency"`
	Rate     float32 `json:"rate"`
}

// ExchangeRates Describes the exchange rates between some base currency and the currencies of some countries.
type ExchangeRates struct {
	Base  string          `json:"base"`
	Rates map[string]rate `json:"rates"`
}

// GetExchangeRates returns exchange rates between multiple currencies.
func GetExchangeRates(base string, countries []Country) (ExchangeRates, *ServerError) {
	ret := ExchangeRates{base, make(map[string]rate)}
	var rates struct {
		Base  string             `json:"base"`
		Rates map[string]float32 `json:"rates"`
	}

	// Construct URL
	url, _ := url.Parse(ExchangeRatesAPIRoot)
	url.Path += "/latest"
	queries := url.Query()
	queries.Set("base", base)
	for _, country := range countries {
		queries.Add("symbols", country.Currencies[0].Code)
	}
	url.RawQuery = queries.Encode()

	// Get the rates
	res, err := http.Get(url.String())
	if err != nil {
		return ret, &ServerError{err.Error(), res.StatusCode}
	}

	err = json.NewDecoder(res.Body).Decode(&rates)
	if err != nil {
		return ret, &ServerError{err.Error(), http.StatusInternalServerError}
	}
	res.Body.Close()

	for _, country := range countries {
		currency := country.Currencies[0].Code
		ret.Rates[country.Name] = rate{currency, rates.Rates[currency]}
	}

	return ret, nil
}

// BorderHandler responds to requests for currency information related to a country's bordering countries.
func BorderHandler(rw http.ResponseWriter, r *http.Request) {
	countryName := chi.URLParam(r, "country")
	country, err := GetCountryByName(countryName)
	if err != nil {
		http.Error(rw, err.Error(), err.StatusCode)
		return
	}

	// Only return min(limit parameter, number of border country)
	limit := len(country.Borders)
	if limitParamS := r.URL.Query().Get("limit"); limitParamS != "" {
		limitParam, convErr := strconv.Atoi(limitParamS)
		if convErr != nil {
			log.Printf("Invalid limit parameter received: %s", convErr.Error())
			http.Error(rw, convErr.Error(), http.StatusBadRequest)
			return
		}
		if limit > limitParam {
			limit = limitParam
		}
	}

	// Get country information from country code
	borderCountries := make([]Country, limit)
	for i := 0; i < limit; i++ {
		borderCountry, serverErr := GetCountryByCode(country.Borders[i])
		if serverErr != nil {
			http.Error(rw, serverErr.Error(), serverErr.StatusCode)
			return
		}
		borderCountries[i] = borderCountry
	}

	// Get exchange rates between countries
	rates, err := GetExchangeRates(country.Currencies[0].Code, borderCountries)
	if err != nil {
		http.Error(rw, err.Error(), err.StatusCode)
		return
	}

	// Write json encoded rates to response body
	json.NewEncoder(rw).Encode(rates)
}
