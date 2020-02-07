package errorresponse

// ErrorResponse for api errors
type ErrorResponse struct {
	ExternalErrorCode int    `json:"externalCode"`
	InternalErrorCode int    `json:"internalCode"`
	Title             string `json:"title"`
	Message           string `json:"message"`
}
