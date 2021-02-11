package exchange

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/go-chi/chi"
)

// TODO: Handle errors
// TODO: Validate request parameters

// Gets countries mathing given name
func getCountries(name string) Countries {
	var countries Countries

	// Set fullText to only return full matches, not partial ones
	res, err := http.Get(RestCountriesRoot + "/name/norway?fullText=true")
	if err != nil {
		log.Fatalf("Get country failed with: %s", err.Error())
	}

	json.NewDecoder(res.Body).Decode(&countries)
	res.Body.Close()

	return countries
}

// Gets exchangerates history for given currencies in the given timeperiod
func getExchangeRates(currencies []Currency, startDate, endDate string) Rates {
	var rates Rates

	// Construct URL
	url, _ := url.Parse(ExchangeRatesAPIRoot)
	queries := url.Query()
	queries.Set("base", "EUR")
	queries.Set("start_at", startDate)
	queries.Set("end_at", endDate)
	for _, currency := range currencies {
		queries.Add("symbols", currency.Code)
	}
	url.RawQuery = queries.Encode()
	url.Path += "/history"

	// Get the rates
	res, err := http.Get(url.String())
	if err != nil {
		log.Fatalf("Get exchange rates failed with: %s", err.Error())
	}

	json.NewDecoder(res.Body).Decode(&rates)
	res.Body.Close()

	return rates
}

func HistoryHandler(rw http.ResponseWriter, r *http.Request) {
	country := chi.URLParam(r, "country")
	start_yyyy := chi.URLParam(r, "start_yyyy")
	start_mm := chi.URLParam(r, "start_mm")
	start_dd := chi.URLParam(r, "start_dd")
	end_yyyy := chi.URLParam(r, "end_yyyy")
	end_mm := chi.URLParam(r, "end_mm")
	end_dd := chi.URLParam(r, "end_dd")

	startDate := fmt.Sprintf("%s-%s-%s", start_yyyy, start_mm, start_dd)
	endDate := fmt.Sprintf("%s-%s-%s", end_yyyy, end_mm, end_dd)

	// Get country information from restcountries.eu
	countries := getCountries(country)

	if len(countries) > 1 {
		log.Println("Given parameter matched more the one country. Make sure the country name is spelled correctly.")
	}

	// Get currency
	rates := getExchangeRates(countries[0].Currencies, startDate, endDate)

	// Write json encoded rates to response body
	json.NewEncoder(rw).Encode(rates)
}
