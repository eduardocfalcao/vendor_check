package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type httpClientMock struct {
	mock.Mock
}

func (h *httpClientMock) Head(url string) (*http.Response, error) {
	args := h.Called(url)
	return args.Get(0).(*http.Response), args.Error(1)
}

func TestHandlerGetAmazonStatus(t *testing.T) {
	testCases := []struct {
		Name               string
		HeadStatusCode     int
		ExpectedStatusCode int
		ErrToReturn        error
	}{
		{
			Name:               "Should Return status code 200",
			HeadStatusCode:     http.StatusOK,
			ExpectedStatusCode: http.StatusOK,
			ErrToReturn:        nil,
		},
		{
			Name:               "Should Return status code 200",
			HeadStatusCode:     http.StatusServiceUnavailable,
			ExpectedStatusCode: http.StatusOK,
			ErrToReturn:        nil,
		},
		{
			Name:               "Should Return an error",
			HeadStatusCode:     http.StatusOK,
			ExpectedStatusCode: http.StatusInternalServerError,
			ErrToReturn:        errors.New("some error"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {

			mock := &httpClientMock{}
			mock.On("Head", amazonUrl).Return(&http.Response{StatusCode: tc.HeadStatusCode}, tc.ErrToReturn)

			req := httptest.NewRequest(http.MethodGet, "/", nil)
			w := httptest.NewRecorder()

			HandlerGetAmazonStatus(mock)(w, req)
			res := w.Result()
			defer res.Body.Close()

			assert.Equal(t, tc.ExpectedStatusCode, w.Code)
			if w.Code == http.StatusOK {
				var checkVendor CheckVendorResponse

				if err := json.NewDecoder(res.Body).Decode(&checkVendor); err != nil {
					assert.FailNow(t, "Not possible to parse the response body")
				}
				assert.Equal(t, tc.HeadStatusCode, checkVendor.StatusCode)
				assert.Equal(t, amazonUrl, checkVendor.Url)
			}
		})
	}
}

func TestHandlerGetGoogleStatus(t *testing.T) {
	testCases := []struct {
		Name               string
		HeadStatusCode     int
		ExpectedStatusCode int
		ErrToReturn        error
	}{
		{
			Name:               "Should Return status code 200",
			HeadStatusCode:     http.StatusOK,
			ExpectedStatusCode: http.StatusOK,
			ErrToReturn:        nil,
		},
		{
			Name:               "Should Return status code 200",
			HeadStatusCode:     http.StatusServiceUnavailable,
			ExpectedStatusCode: http.StatusOK,
			ErrToReturn:        nil,
		},
		{
			Name:               "Should Return an error",
			HeadStatusCode:     http.StatusOK,
			ExpectedStatusCode: http.StatusInternalServerError,
			ErrToReturn:        errors.New("some error"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {

			mock := &httpClientMock{}
			mock.On("Head", googleUrl).Return(&http.Response{StatusCode: tc.HeadStatusCode}, tc.ErrToReturn)

			req := httptest.NewRequest(http.MethodGet, "/", nil)
			w := httptest.NewRecorder()

			HandlerGetGoogleStatus(mock)(w, req)
			res := w.Result()
			defer res.Body.Close()

			assert.Equal(t, tc.ExpectedStatusCode, w.Code)
			if w.Code == http.StatusOK {
				var checkVendor CheckVendorResponse

				if err := json.NewDecoder(res.Body).Decode(&checkVendor); err != nil {
					assert.FailNow(t, "Not possible to parse the response body")
				}
				assert.Equal(t, tc.HeadStatusCode, checkVendor.StatusCode)
				assert.Equal(t, googleUrl, checkVendor.Url)
			}
		})
	}
}

func TestHandlerGetAllStatus(t *testing.T) {
	type VendorCall struct {
		Url            string
		HeadStatusCode int
		ErrToReturn    error
	}

	testCases := []struct {
		Name               string
		VendorCalls        []VendorCall
		ExpectedStatusCode int
	}{
		{
			Name: "Should Return status code 200",
			VendorCalls: []VendorCall{
				{
					Url:            amazonUrl,
					HeadStatusCode: http.StatusOK,
					ErrToReturn:    nil,
				},
				{
					Url:            googleUrl,
					HeadStatusCode: http.StatusBadGateway,
					ErrToReturn:    nil,
				},
			},
			ExpectedStatusCode: http.StatusOK,
		},
		{
			Name: "Should return an error if google vendor returns an error",
			VendorCalls: []VendorCall{
				{
					Url:            amazonUrl,
					HeadStatusCode: http.StatusOK,
					ErrToReturn:    nil,
				},
				{
					Url:            googleUrl,
					HeadStatusCode: http.StatusOK,
					ErrToReturn:    errors.New("some error"),
				},
			},
			ExpectedStatusCode: http.StatusInternalServerError,
		},
		{
			Name: "Should return an error if amazon vendor returns an error",
			VendorCalls: []VendorCall{
				{
					Url:            amazonUrl,
					HeadStatusCode: http.StatusOK,
					ErrToReturn:    errors.New("some error"),
				},
				{
					Url:            googleUrl,
					HeadStatusCode: http.StatusOK,
					ErrToReturn:    nil,
				},
			},
			ExpectedStatusCode: http.StatusInternalServerError,
		},
		{
			Name: "Should return error if both vendors returns an error",
			VendorCalls: []VendorCall{
				{
					Url:            amazonUrl,
					HeadStatusCode: http.StatusOK,
					ErrToReturn:    errors.New("some error"),
				},
				{
					Url:            googleUrl,
					HeadStatusCode: http.StatusOK,
					ErrToReturn:    errors.New("some error"),
				},
			},
			ExpectedStatusCode: http.StatusInternalServerError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {

			mock := &httpClientMock{}
			for _, vendorCall := range tc.VendorCalls {
				mock.On("Head", vendorCall.Url).Return(&http.Response{StatusCode: vendorCall.HeadStatusCode}, vendorCall.ErrToReturn)
			}

			req := httptest.NewRequest(http.MethodGet, "/", nil)
			w := httptest.NewRecorder()

			HandlerGetAllStatus(mock)(w, req)
			res := w.Result()
			defer res.Body.Close()

			assert.Equal(t, tc.ExpectedStatusCode, w.Code)

			if w.Code == http.StatusOK {
				var checkVendorsResult []CheckVendorResponse

				if err := json.NewDecoder(res.Body).Decode(&checkVendorsResult); err != nil {
					assert.FailNow(t, "Not possible to parse the response body")
				}

				callMap := map[string]VendorCall{}
				for _, call := range tc.VendorCalls {
					callMap[call.Url] = call
				}

				for _, vendorResult := range checkVendorsResult {
					expectedVendor := callMap[vendorResult.Url]

					assert.Equal(t, expectedVendor.HeadStatusCode, vendorResult.StatusCode)
					assert.Equal(t, expectedVendor.Url, vendorResult.Url)
				}

			}
		})
	}
}
