package exchange

import (
	"log"
	"net/http"
	"strconv"
)

func BorderHandler(rw http.ResponseWriter, r *http.Request) {
	limitParam := r.URL.Query().Get("limit")
	limit, err := strconv.Atoi(limitParam)
	if err != nil {
		log.Printf("Invalid limit parameter recived: %s", err.Error())
		http.Error(rw, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	log.Printf("Query parameter 'limit' = %d", limit)
}
