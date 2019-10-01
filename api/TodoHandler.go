package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"github.com/someDevDude/todo-server/models"
	"github.com/someDevDude/todo-server/repository"
	"github.com/someDevDude/todo-server/util"
)

//CreateTodoRoutes creates a subrouter that handles all todo requests
func CreateTodoRoutes(r *mux.Router) {
	r.HandleFunc("/todo", queryTodosHandler).Methods("GET")
	r.HandleFunc("/todo", createTodoHandler).Methods("POST")
}

//queryTodosHandler returns list of todos
func queryTodosHandler(rw http.ResponseWriter, r *http.Request) {
	var params models.ListParams

	err := schema.NewDecoder().Decode(&params, r.URL.Query())
	util.CheckErr(err, func(err error) { panic(err) })

	results := repository.QueryTodos(params)

	resp, err := json.Marshal(results)
	util.CheckErr(err, func(err error) { panic(err) })

	rw.Write(resp)
}

//CreateTodo creates a todo
func createTodoHandler(rw http.ResponseWriter, r *http.Request) {
	var todo models.Todo

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&todo)
	util.CheckErr(err, func(err error) { panic(err) })

	repository.CreateTodo(todo)
}
