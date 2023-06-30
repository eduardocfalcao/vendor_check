package handlers

type CheckVendorResponse struct {
	Url        string `json:"url"`
	StatusCode int    `json:"status_code"`
	Duration   int64  `json:"duration"`
	Date       int64  `json:"date"`
}
