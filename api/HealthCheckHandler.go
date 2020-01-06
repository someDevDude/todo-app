package api

import (
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

//CreateHealthCheckRoutes handles all requests regarding health checks
func CreateHealthCheckRoutes(r *mux.Router) {
	r.HandleFunc("/healthcheck", healthCheckHandler).Methods("GET")
}

func healthCheckHandler(rw http.ResponseWriter, r *http.Request) {
	io.WriteString(rw, `status: upa`)
}
