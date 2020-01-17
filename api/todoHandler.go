package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"github.com/someDevDude/todo-server/database"
	"github.com/someDevDude/todo-server/errorresponse"
	"github.com/someDevDude/todo-server/models"
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
	util.CheckErr(err, func(err error) {
		util.Errorf("Error decoding todo query list params %s\n%s", r.URL, err.Error())
		errorresponse.ErrorParsingQueryParams(rw)
		return
	})

	results := database.QueryTodos(params)

	resp, err := json.Marshal(results)
	util.CheckErr(err, func(err error) {
		util.Errorf("Error decoding todos\n%s", err.Error())
		errorresponse.ErrorParsingTodos(rw)
		return
	})

	rw.Write(resp)
}

//CreateTodo creates a todo
func createTodoHandler(rw http.ResponseWriter, r *http.Request) {
	var todo models.TodoFull

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&todo)
	util.CheckErr(err, func(err error) {
		util.Errorf("Error parsing todos\n%s", err.Error())
		errorresponse.ErrorParsingRequestBody(rw)
		return
	})

	database.CreateTodo(todo)
}
