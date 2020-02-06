package api

import (
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

//CreateHomeRoutes handles all requests to /
func CreateHomeRoutes(r *mux.Router) {
	r.HandleFunc("/", homeHandler)
}

func homeHandler(rw http.ResponseWriter, r *http.Request) {
	io.WriteString(rw, "pong")
}
