package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/eduardocfalcao/vendors_checker/handlers"
)

const (
	port = 8000
)

func main() {
	address := fmt.Sprintf(":%d", port)
	mux := http.NewServeMux()

	registerRoutes(mux)

	server := &http.Server{
		Addr:         address,
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	server.ListenAndServe()
}

func registerRoutes(mux *http.ServeMux) {
	httpClient := &http.Client{Timeout: 1 * time.Second}

	mux.HandleFunc("/v1/amazon-status", handlers.HandlerGetAmazonStatus(httpClient))
	mux.HandleFunc("/v1/google-status", handlers.HandlerGetGoogleStatus(httpClient))
	mux.HandleFunc("/v1/all-status", handlers.HandlerGetAllStatus(httpClient))
}
