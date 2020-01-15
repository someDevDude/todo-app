package api

import (
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
)

//CreateHomeRoutes handles all requests to /
func CreateHomeRoutes(r *mux.Router) {
	r.HandleFunc("/", homeHandler)
}

func homeHandler(rw http.ResponseWriter, r *http.Request) {
	var dbConnectionString strings.Builder
	dbConnectionString.WriteString(os.Getenv("DB_USER"))
	dbConnectionString.WriteString(":")
	dbConnectionString.WriteString(os.Getenv("DB_PW"))
	dbConnectionString.WriteString("@tcp(")
	dbConnectionString.WriteString(os.Getenv("DB_URL"))
	dbConnectionString.WriteString(":3306)/")
	dbConnectionString.WriteString(os.Getenv("DB_NAME"))
	io.WriteString(rw, dbConnectionString.String())
}
