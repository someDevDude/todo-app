package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"github.com/someDevDude/todo-server/database"
	"github.com/someDevDude/todo-server/errorresponse"
	"github.com/someDevDude/todo-server/models"
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
	if err != nil {
		errorresponse.ErrorParsingQueryParams(rw)
		return
	}

	results, err := database.QueryTodos(params)

	resp, err := json.Marshal(results)
	if err != nil {
		errorresponse.ErrorParsingTodos(rw)
		return
	}

	rw.Write(resp)
}

//CreateTodo creates a todo
func createTodoHandler(rw http.ResponseWriter, r *http.Request) {
	var todo *models.TodoFull

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&todo)
	if err != nil {
		errorresponse.ErrorParsingRequestBody(rw)
		return
	}

	validationError := ValidateTodo(todo)
	if validationError != "" {
		errorresponse.ErrorMissingParameter(rw, validationError)
		return
	}

	todo, err = database.CreateTodo(todo)

	response, err := json.Marshal(todo)
	if err != nil {
		errorresponse.ErrorParsingTodos(rw)
		return
	}

	rw.Write(response)

}

// ValidateTodo validates todo
func ValidateTodo(todo *models.TodoFull) string {
	if todo.Title == "" {
		return "Missing title for todo"
	}

	return ""
}
