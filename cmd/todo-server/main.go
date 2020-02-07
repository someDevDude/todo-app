package main

import (
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gorilla/mux"

	"github.com/someDevDude/todo-server/api"
	"github.com/someDevDude/todo-server/database"
	"github.com/someDevDude/todo-server/util"
)

func main() {
	util.InitialiseLogger()

	database.Connect()
	defer database.DB.Close()

	r := mux.NewRouter()
	r.Use(middleware)

	api.CreateHomeRoutes(r)
	api.CreateHealthCheckRoutes(r)
	api.CreateTodoRoutes(r)

	http.ListenAndServe(":8080", r)
}

func middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Content-Type", "application/json; charset=utf-8")
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(rw, r)
	})
}
