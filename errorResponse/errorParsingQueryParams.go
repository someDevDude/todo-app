package errorresponse

import (
	"encoding/json"
	"net/http"

	"github.com/someDevDude/todo-server/util"
)

// ErrorParsingQueryParams throws error
func ErrorParsingQueryParams(rw http.ResponseWriter) {
	rw.WriteHeader(http.StatusBadRequest)
	resp, err := json.Marshal(ErrorResponse{http.StatusBadRequest, InternalErrorCodeParsingQueryParams, "Error", "Error parsing url query params"})
	util.CheckErr(err, func(err error) {
		util.Error("Error creating error response")

	})
	rw.Write(resp)
}
