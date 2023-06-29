package handlers

type CheckVendorResponse struct {
	Url        string `json:"url"`
	StatusCode int    `json:"statusCode"`
	Duration   int64  `json:"duration"`
	Date       int64  `json:"date"`
}

type ApiErrorResponse struct {
	ErrorMessage string `json:"error_message"`
}
