package errorresponse

import (
	"net/http"

	"encoding/json"

	"github.com/someDevDude/todo-server/util"
)

// ErrorParsingTodos throws error
func ErrorParsingTodos(rw http.ResponseWriter) {
	rw.WriteHeader(http.StatusBadRequest)

	resp, err := json.Marshal(ErrorResponse{http.StatusBadRequest, InternalErrorCodeParsingResponse, "Error", "Error parsing todos"})
	if err != nil {
		util.Error("Error creating error response")
	}
	rw.Write(resp)
}
