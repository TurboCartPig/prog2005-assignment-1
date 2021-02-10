package exchange

import (
	"log"
	"net/http"
)

func BorderHandler(rw http.ResponseWriter, r *http.Request) {
	limit := r.URL.Query().Get("limit")
	log.Printf("Query parameter 'limit' = %s", limit)
}
