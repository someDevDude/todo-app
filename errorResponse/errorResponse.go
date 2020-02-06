package errorresponse

// ErrorResponse for api errors
type ErrorResponse struct {
	ExternalErrorCode int    `schema:"externalCode"`
	InternalErrorCode int    `schema:"internalCode"`
	Title             string `schema:"title"`
	Message           string `schema:"message"`
}
