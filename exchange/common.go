package exchange

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	// Root URL for the restcountries.eu api
	RestCountriesRoot = "https://restcountries.eu/rest/v2"
	// Root URL for the exchangeratesapi.io api
	ExchangeRatesAPIRoot = "https://api.exchangeratesapi.io"
)

// Country represents a country as given by `restcountries.eu`.
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

// Currency of a Country.
type Currency struct {
	Code   string `json:"code"`
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
}

// Language of a Country.
type Language struct {
	ISO639_1   string `json:"iso639_1"`
	ISO639_2   string `json:"iso639_2"`
	Name       string `json:"name"`
	NativeName string `json:"nativeName"`
}

// RegionalBlocks of a Country.
type RegionalBlock struct {
	Acronym       string   `json:"acronym"`
	Name          string   `json:"name"`
	OtherAcronyms []string `json:"otherAcronyms"`
	OtherNames    []string `json:"otherNames"`
}

// ServerError describes an internal server error and what http status code it should return.
type ServerError struct {
	error string
	// StatusCode is the http status code that should be returned by the server when handling this error.
	StatusCode int
}

func (e *ServerError) Error() string {
	return e.error
}

// GetCountryCode returns the country code given a country name.
func GetCountryCode(name string) (string, *ServerError) {
	country := make([]map[string]string, 1)

	res, err := http.Get(RestCountriesRoot + "/name/" + name + "?fullText=true&fields=alpha3Code")
	if err != nil {
		return "", &ServerError{fmt.Sprintf("Get country failed with: %s", err.Error()), res.StatusCode}
	}

	err = json.NewDecoder(res.Body).Decode(&country)
	if err != nil {
		return "", &ServerError{"Failed to decode json response from restcountries.eu", http.StatusInternalServerError}
	}
	res.Body.Close()

	return country[0]["alpha3Code"], nil
}

// GetCountryByName return country information given a country name.
func GetCountryByName(name string) (Country, *ServerError) {
	var country Country
	// Get country code
	countryCode, err := GetCountryCode(name)
	if err != nil {
		return country, err
	}

	// Get country information from restcountries.eu
	country, err = GetCountryByCode(countryCode)
	if err != nil {
		return country, err
	}

	return country, nil
}

// GetCountryByCode returns country information given country code.
func GetCountryByCode(code string) (Country, *ServerError) {
	var country Country

	res, err := http.Get(RestCountriesRoot + "/alpha/" + code)
	if err != nil {
		return country, &ServerError{fmt.Sprintf("Get country failed with: %s", err.Error()), res.StatusCode}
	}

	err = json.NewDecoder(res.Body).Decode(&country)
	if err != nil {
		return country, &ServerError{"Failed to decode json response from restcountries.eu", http.StatusInternalServerError}
	}
	res.Body.Close()

	return country, nil
}
