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
var Version string = "v1"

// Root endpoint path
var RootPath string = "/exchange/" + Version

// Functions
/////////////////////////////////////////////////////////////////////////////////////////////

// Get the port from environment variable $PORT, or use default if the variable is not set
func port() int {
	if port := os.Getenv("PORT"); port != "" {
		port, _ := strconv.Atoi(port)
		return port
	}
	return 3000
}

// Middleware for setting content-type to json
func returnJSON(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(rw, r)
	})
}

// Setup all the top level routes the server serves on
func setupRoutes() *chi.Mux {
	r := chi.NewRouter()

	// Use Middleware
	r.Use(middleware.Logger)
	r.Use(returnJSON)

	// Define endpoints
	r.Get(
		RootPath+"/exchangehistory/{country:[a-z]+}/{start_yyyy}-{start_mm}-{start_dd}-{end_yyyy}-{end_mm}-{end_dd}",
		exchange.HistoryHandler,
	)
	r.Get(
		RootPath+"/exchangeborder/{country:[a-z]+}",
		exchange.BorderHandler,
	)
	r.Get(
		RootPath+"/diag",
		diag.NewHandler(StartTime, Version),
	)

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
