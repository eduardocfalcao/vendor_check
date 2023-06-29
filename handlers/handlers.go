package handlers

import (
	"encoding/json"
	"fmt"
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

func HandlerGetAmazonStatus(client httpClient) apiFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vendorResponse, err := getVendorStatus(client, amazonUrl)

		if err != nil {
			log.Printf("An error occurred when trying to check Amazon vendor. %s", err)
			writeJSON(w, http.StatusInternalServerError, ApiErrorResponse{
				ErrorMessage: "Some error occurred.",
			})
			return
		}

		writeJSON(w, http.StatusOK, vendorResponse)
	}
}

func HandlerGetGoogleStatus(client httpClient) apiFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vendorResponse, err := getVendorStatus(client, googleUrl)

		if err != nil {
			log.Printf("An error occurred when trying to check Google vendor. %s", err)
			writeJSON(w, http.StatusInternalServerError, ApiErrorResponse{
				ErrorMessage: "Some error occurred.",
			})
			return
		}

		writeJSON(w, http.StatusOK, vendorResponse)
	}
}

func HandlerGetAllStatus(client httpClient) apiFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response := make([]CheckVendorResponse, 0, 2)
		var wg sync.WaitGroup
		wg.Add(2)
		responseChan := make(chan CheckVendorResponse, 2)
		errorChan := make(chan error, 2)

		for _, vendorUrl := range []string{amazonUrl, googleUrl} {

			go func(url string) {
				defer wg.Done()
				response, err := getVendorStatus(client, url)
				if err != nil {
					errorChan <- fmt.Errorf("an error occurred when trying to check vendor: %s. %w", url, err)
					return
				}
				responseChan <- response

			}(vendorUrl)
		}

		wg.Wait()
		close(responseChan)
		close(errorChan)

		for r := range responseChan {
			response = append(response, r)
		}

		if err := <-errorChan; err != nil {
			log.Print(err)
			writeJSON(w, http.StatusInternalServerError, ApiErrorResponse{
				ErrorMessage: "Some error occurred.",
			})
			return
		}

		writeJSON(w, http.StatusOK, response)
	}
}

func getVendorStatus(client httpClient, url string) (CheckVendorResponse, error) {
	start := time.Now()
	resp, err := client.Head(url)
	duration := time.Since(start).Milliseconds()

	if err != nil {
		return CheckVendorResponse{}, err
	}

	vendorResponse := CheckVendorResponse{
		Url:        url,
		StatusCode: resp.StatusCode,
		Duration:   duration,
		Date:       time.Now().UTC().Unix(),
	}
	return vendorResponse, nil
}

func writeJSON(w http.ResponseWriter, status int, obj any) error {
	w.WriteHeader(status)
	w.Header().Add("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(obj); err != nil {
		return err
	}
	return nil
}
