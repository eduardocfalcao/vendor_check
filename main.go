package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
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

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		log.Printf("Starting http server. Listening port %s", address)
		if err := server.ListenAndServe(); err != nil {
			log.Print(err)
			cancel()
		}
	}()

	select {
	case <-c:
		log.Print("Stopping server...")
		server.Shutdown(ctx)
		os.Exit(0)
	case <-ctx.Done():
	}
}

func registerRoutes(mux *http.ServeMux) {
	httpClient := &http.Client{Timeout: 1 * time.Second}

	mux.HandleFunc("/v1/amazon-status", cors(http.HandlerFunc(handlers.HandlerGetAmazonStatus(httpClient))))
	mux.HandleFunc("/v1/google-status", cors(http.HandlerFunc(handlers.HandlerGetGoogleStatus(httpClient))))
	mux.HandleFunc("/v1/all-status", cors(http.HandlerFunc(handlers.HandlerGetAllStatus(httpClient))))
}

func cors(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		next(w, r)
	}
}
