package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"github.com/someDevDude/todo-server/database"
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
	var results []models.Todo

	err := schema.NewDecoder().Decode(&params, r.URL.Query())
	util.CheckErr(err, func(err error) { panic(err) })

	queryString := "SELECT * FROM todo"
	separator := " WHERE"
	var args []interface{}

	if params.Q != "" {
		queryString += separator + " LOWER(title) LIKE LOWER('%' +  ?  +'%')"
		separator = " AND"
		args = append(args, params.Q)
	}

	if params.StartIndex != 0 {
		queryString += separator + " OFFSET ?"
		separator = " AND"
		args = append(args, params.StartIndex)
	}

	if params.MaxResults != 0 {
		queryString += separator + " LIMIT ?"
		separator = " AND"
		args = append(args, params.MaxResults)
	}

https: //github.com/jmoiron/sqlx/

	if params.Sort != "" {
		queryString += separator + " ORDER BY ? ?"
		sortParams := strings.SplitAfter(params.Sort, ":")
		fmt.Printf("%#v\n", sortParams)
		fmt.Println(sortParams[0][:len(sortParams[0])-1])
		args = append(args, sortParams[0][:len(sortParams[0])-1])
		args = append(args, "ASC")
	}

	fmt.Println(queryString)

	stmt, err := database.DB.Prepare(queryString)
	util.CheckErr(err, func(err error) { panic(err) })

	fmt.Println(args...)

	rows, err := stmt.Query(args...)
	util.CheckErr(err, func(err error) { panic(err) })

	for rows.Next() {
		var r models.Todo

		err = rows.Scan(&r.ID, &r.Title, &r.Description, &r.Done)

		results = append(results, r)
	}

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

	stmt, err := database.DB.Prepare("INSERT todo SET title = ?, dewscription = ?, done = 0")
	util.CheckErr(err, func(err error) { panic(err) })

	_, err = stmt.Exec(todo.Title, todo.Description)
	util.CheckErr(err, func(err error) { panic(err) })
}
