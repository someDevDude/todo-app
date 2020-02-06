package errorresponse

import (
	"encoding/json"
	"net/http"

	"github.com/someDevDude/todo-server/util"
)

// ErrorParsingQueryParams throws error
func ErrorParsingQueryParams(rw http.ResponseWriter) {
	resp, err := json.Marshal(ErrorResponse{http.StatusBadRequest, InternalErrorCodeParsingQueryParams, "Error", "Error parsing url query params"})
	if err != nil {
		util.Error("error creating ErrorParsingQueryParams response")
	}
	rw.WriteHeader(http.StatusBadRequest)
	rw.Write(resp)
}
