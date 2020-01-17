package errorresponse

import (
	"encoding/json"
	"github.com/someDevDude/todo-server/util"
	"net/http"
)

// ErrorParsingRequestBody throws error
func ErrorParsingRequestBody(rw http.ResponseWriter) {
	rw.WriteHeader(http.StatusBadRequest)
	resp, err := json.Marshal(ErrorResponse{http.StatusBadRequest, InternalErrorCodeParsingRequestBody, "Error", "Error parsing request body"})
	util.CheckErr(err, func(err error) {
		util.Error("Error creating error response")

	})
	rw.Write(resp)
}
