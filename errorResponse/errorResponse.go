package errorresponse

// ErrorResponse for api errors
type ErrorResponse struct {
	ExternalErrorCode int    `schema:"externalCode"`
	internalErrorCode int    `schema:"internalCode"`
	Title             string `schema:"title"`
	message           string `schema:"message"`
}
