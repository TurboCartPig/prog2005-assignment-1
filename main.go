package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"assignment-1/diag"
	"assignment-1/exchange"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

// Globals
/////////////////////////////////////////////////////////////////////////////////////////////

// The instant the server was started
var StartTime time.Time = time.Now()

// Version of the program
var Version string = "0.1.0"

// Functions
/////////////////////////////////////////////////////////////////////////////////////////////

// Get the port from enviroment variable $PORT, or use default if the variable is not set
func port() int {
	if port := os.Getenv("PORT"); port != "" {
		port, _ := strconv.Atoi(port)
		return port
	}
	return 3000
}

// Setup all the top level routes the server serves on
func setupRoutes() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/exchange/v1/exchangehistory/{country:[a-z]+}/{start:[0-9-]+}-{end:[0-9-]+}", exchange.HistoryHandler)
	r.Get("/exchange/v1/exchangeborder/{country:[a-z]+}", exchange.BorderHandler)
	r.Get("/exchange/v1/diag", diag.NewHandler(StartTime, Version))

	return r
}

// Serve the resources as defined by routes in `r`
func serve(r *chi.Mux) {
	port := port()
	addr := fmt.Sprintf(":%d", port)
	http.ListenAndServe(addr, r)
}

func main() {
	r := setupRoutes()
	serve(r)
}
