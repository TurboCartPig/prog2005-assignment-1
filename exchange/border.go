package exchange

import (
	"log"
	"net/http"
	"strconv"
)

// BorderHandler responds to requests for currency information related to a country's bordering countries.
func BorderHandler(rw http.ResponseWriter, r *http.Request) {
	limitParam := r.URL.Query().Get("limit")
	limit, err := strconv.Atoi(limitParam)
	if err != nil {
		log.Printf("Invalid limit parameter received: %s", err.Error())
		http.Error(rw, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	log.Printf("Query parameter 'limit' = %d", limit)
}
