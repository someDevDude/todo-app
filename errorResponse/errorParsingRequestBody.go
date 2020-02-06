package errorresponse

import (
	"encoding/json"
	"net/http"

	"github.com/someDevDude/todo-server/util"
)

// ErrorParsingRequestBody throws error
func ErrorParsingRequestBody(rw http.ResponseWriter) {
	resp, err := json.Marshal(ErrorResponse{http.StatusBadRequest, InternalErrorCodeParsingRequestBody, "Error", "Error parsing request body"})
	if err != nil {
		util.Error("error creating ErrorParsingRequestBody response")
	}
	rw.WriteHeader(http.StatusBadRequest)
	rw.Write([]byte(resp))
}
