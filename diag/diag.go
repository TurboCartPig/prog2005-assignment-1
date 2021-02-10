package diag

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

// Response from the diagnostic interface
type Response struct {
	ExchangeRatesAPI int    `json:"excheratesapi"`
	RestCountries    int    `json:"restcountries"`
	Version          string `json:"version"`
	Uptime           int    `json:"uptime"`
}

// Get the status code of a http server of some kind
func getStatusOf(addr string) int {
	res, err := http.Head(addr)
	if err != nil {
		log.Printf("Head request failed with: %s", err.Error())
		// If the request failed, assume it's out fault and not the remote
		// FIXME: What happends if the remote is down? Or has misconfigured DNS settings?
		return http.StatusBadRequest
	}
	res.Body.Close()
	return res.StatusCode
}

// Returns a handler function for the diagnostic endpoint
func NewHandler(startTime time.Time, version string) func(http.ResponseWriter, *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {
		uptime := int(time.Since(startTime).Seconds())

		response := Response{
			getStatusOf("https://api.exchangeratesapi.io"),
			getStatusOf("https://restcountries.eu/rest/v2"),
			version,
			uptime,
		}

		rw.Header().Add("Content-Type", "application/json")
		json.NewEncoder(rw).Encode(response)
	}
}
