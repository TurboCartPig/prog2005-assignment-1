package exchange

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
)

const (
	// Root URL for the restcountries.eu api
	RestCountriesRoot = "https://restcountries.eu/rest/v2"
	// Root URL for the exchangeratesapi.io api
	ExchangeRatesAPIRoot = "https://api.exchangeratesapi.io"
)

type Countries = []Country

type Country struct {
	Name           string            `json:"name"`
	TopLevelDomain []string          `json:"topLevelDomain"`
	Alpha2Code     string            `json:"alpha2Code"`
	Alpha3Code     string            `json:"alpha3Code"`
	CallingCodes   []string          `json:"callingCodes"`
	Capital        string            `json:"capital"`
	AltSpellings   []string          `json:"altSpellings"`
	Region         string            `json:"region"`
	Subregion      string            `json:"subregion"`
	Population     int               `json:"population"`
	LatLng         []float32         `json:"latlng"`
	Demonym        string            `json:"demonym"`
	Area           float32           `json:"area"`
	Gini           float32           `json:"gini"`
	Timezones      []string          `json:"timezones"`
	Borders        []string          `json:"borders"`
	NativeName     string            `json:"nativeName"`
	NumericCode    string            `json:"numericCode"`
	Currencies     []Currency        `json:"currencies"`
	Languages      []Language        `json:"languages"`
	Translations   map[string]string `json:"translations"`
	Flag           string            `json:"flag"`
	RegionalBlocks []RegionalBlock   `json:"regionalBlocks"`
	Cioc           string            `json:"cioc"`
}

type Currency struct {
	Code   string `json:"code"`
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
}

type Language struct {
	ISO639_1   string `json:"iso639_1"`
	ISO639_2   string `json:"iso639_2"`
	Name       string `json:"name"`
	NativeName string `json:"nativeName"`
}

type RegionalBlock struct {
	Acronym       string   `json:"acronym"`
	Name          string   `json:"name"`
	OtherAcronyms []string `json:"otherAcronyms"`
	OtherNames    []string `json:"otherNames"`
}

// Describes the rates for exchange between a base currency and a list of other currencies in some timeperiod.
type Rates struct {
	// Base currency
	Base string `json:"base"`
	// Starting date for lookup
	StartAt string `json:"start_at"`
	// End date for lookup
	EndAt string `json:"end_at"`
	// Rate for date for currency
	Rates map[string]map[string]float32 `json:"rates"`
}

// Gets countries matching given name
func GetCountries(name string) (Countries, int) {
	var countries Countries

	// Set fullText to only return full matches, not partial ones
	res, err := http.Get(RestCountriesRoot + "/name/norway?fullText=true")
	if err != nil {
		log.Fatalf("Get country failed with: %s", err.Error())
	}

	json.NewDecoder(res.Body).Decode(&countries)
	res.Body.Close()

	return countries, res.StatusCode
}

// Gets exchange-rates history for given currencies in the given time period, relative to given base.
// Where base is a currency code.
func GetExchangeRates(base string, currencies []Currency, startDate, endDate string) (Rates, int) {
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
		log.Fatal(err)
	}
	json.NewDecoder(res.Body).Decode(&rates)
	res.Body.Close()

	return rates, res.StatusCode
}
