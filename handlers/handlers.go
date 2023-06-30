package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"time"
)

const (
	amazonUrl = "https://www.amazon.com"
	googleUrl = "https://www.google.com"
)

type apiFunc func(http.ResponseWriter, *http.Request)

type httpClient interface {
	Head(url string) (*http.Response, error)
}

// HandlerGetAllStatus perforns the check operation for Amazon vendor
func HandlerGetAmazonStatus(client httpClient) apiFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vendorResponse, err := getVendorStatus(client, amazonUrl)

		if err != nil {
			log.Printf("An error occurred when trying to check Amazon vendor. %s", err)
		}

		w.Header().Add("Content-Type", "application/json")

		if err := json.NewEncoder(w).Encode(vendorResponse); err != nil {
			log.Printf("An error occurred when enconding the response. %s", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

// HandlerGetAllStatus perforns the check operation for Google vendor
func HandlerGetGoogleStatus(client httpClient) apiFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vendorResponse, err := getVendorStatus(client, googleUrl)

		if err != nil {
			log.Printf("An error occurred when trying to check Google vendor. %s", err)
		}

		if err := json.NewEncoder(w).Encode(vendorResponse); err != nil {
			log.Printf("An error occurred when enconding the response. %s", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

// HandlerGetAllStatus performs the check operation for the google and amazon vendors
func HandlerGetAllStatus(client httpClient) apiFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var wg sync.WaitGroup
		var mux sync.Mutex
		responses := []CheckVendorResponse{}

		for _, url := range []string{amazonUrl, googleUrl} {
			wg.Add(1)
			go func(u string) {
				defer wg.Done()
				r, err := getVendorStatus(client, u)
				mux.Lock()
				responses = append(responses, r)
				mux.Unlock()
				if err != nil {
					log.Printf("an error occurred when trying to check vendor: %s. %s", googleUrl, err)
				}
			}(url)
		}

		wg.Wait()

		if err := json.NewEncoder(w).Encode(responses); err != nil {
			log.Printf("An error occurred when enconding the response. %s", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

// getVendorStatus makes http HEAD requests to the given URL.
//
// It will always return a filled CheckVendorResponse value,
// but in case of errors on the client, it will assign status
// code 500 in the StatusCode of the struct field.
func getVendorStatus(client httpClient, url string) (CheckVendorResponse, error) {
	start := time.Now()
	resp, err := client.Head(url)
	duration := time.Since(start).Milliseconds()

	if err != nil {
		return CheckVendorResponse{
			Url:        url,
			StatusCode: http.StatusInternalServerError,
			Duration:   duration,
			Date:       time.Now().UTC().Unix(),
		}, err
	}

	vendorResponse := CheckVendorResponse{
		Url:        url,
		StatusCode: resp.StatusCode,
		Duration:   duration,
		Date:       time.Now().UTC().Unix(),
	}
	return vendorResponse, nil
}
