package main

import (
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gorilla/mux"

	"bitbucket.org/glxdevs/todo-server/api"
	"bitbucket.org/glxdevs/todo-server/database"
)

func main() {
	database.Connect()
	defer database.DB.Close()

	r := mux.NewRouter()

	api.CreateHomeRoutes(r)
	api.CreateTodoRoutes(r)

	http.ListenAndServe(":8080", r)
}
