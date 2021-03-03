package exchange

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

// diag is the response from the diagnostic interface.
type diag struct {
	ExchangeRatesAPI int    `json:"excheratesapi"`
	RestCountries    int    `json:"restcountries"`
	Version          string `json:"version"`
	Uptime           int    `json:"uptime"`
}

// getStatusOf returns the status code of a head request to the root path of a remote.
func getStatusOf(addr string) int {
	res, err := http.Head(addr)
	if err != nil {
		log.Printf("Head request failed with: %s", err.Error())
		return http.StatusBadRequest // Assume I did something wrong, all other errors should be "successful"
	}
	res.Body.Close()
	return res.StatusCode
}

// NewDiagHandler returns a handler function for the diagnostic endpoint
func NewDiagHandler(startTime time.Time, version string) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		uptime := int(time.Since(startTime).Seconds())

		response := diag{
			getStatusOf(ExchangeRatesAPIRoot),
			getStatusOf(RestCountriesRoot),
			version,
			uptime,
		}

		json.NewEncoder(rw).Encode(response)
	}
}
