package main

import (
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gorilla/mux"
	"github.com/someDevDude/todo-server/api"
	"github.com/someDevDude/todo-server/repository"
)

func main() {
	repository.Connect()
	defer repository.DB.Close()

	r := mux.NewRouter()

	api.CreateHomeRoutes(r)
	api.CreateTodoRoutes(r)

	http.ListenAndServe(":8080", r)
}
