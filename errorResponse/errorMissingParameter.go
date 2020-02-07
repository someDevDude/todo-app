package errorresponse

import (
	"encoding/json"
	"net/http"

	"github.com/someDevDude/todo-server/util"
)

// ErrorMissingParameter throws error
func ErrorMissingParameter(rw http.ResponseWriter, parameterError string) {
	resp, err := json.Marshal(ErrorResponse{http.StatusBadRequest, InternalErrorCodeMissingParamater, "Error", parameterError})
	if err != nil {
		util.Error("error creating ErrorMissingParameter response")
	}
	rw.WriteHeader(http.StatusBadRequest)

	rw.Write(resp)
}
