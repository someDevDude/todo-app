package database

import (
	"fmt"
	"strings"

	"github.com/someDevDude/todo-server/models"
	"github.com/someDevDude/todo-server/util"
)

//QueryTodos queires todos
func QueryTodos(params models.ListParams) ([]models.TodoFull, error) {
	var results []models.TodoFull

	queryString := "SELECT * FROM todo"
	separator := " WHERE"
	args := make(map[string]interface{})

	if params.Q != "" {
		queryString += separator + " LOWER(title) LIKE LOWER(CONCAT('%', :query, '%'))"
		separator = " AND"
		args["query"] = params.Q
	}

	if params.StartIndex != 0 {
		queryString += separator + " OFFSET :startIndex"
		separator = " AND"
		args["startIndex"] = params.StartIndex
	}

	if params.Sort != "" {
		queryString += " ORDER BY :sortColumn "
		sortParams := strings.SplitAfter(params.Sort, ":")
		args["sortColumn"] = sortParams[0][:len(sortParams[0])-1] + "  " + sortParams[1]
	}

	if params.MaxResults != 0 {
		queryString += " LIMIT :maxResults"
		args["maxResults"] = params.MaxResults
	}

	rows, err := DB.Query(queryString)
	if err != nil {
		util.Errorf("Error querying todos in database\n%s", err.Error())
		return nil, err
	}

	for rows.Next() {
		var r models.TodoFull

		err = rows.Scan(&r.ID, &r.Title, &r.Description, &r.Done)

		results = append(results, r)
	}

	return results, nil
}

//CreateTodo creates a todo
func CreateTodo(todo *models.TodoFull) (*models.TodoFull, error) {
	stmt, err := DB.Prepare("INSERT todo SET title = ?, description = ?, done = 0")
	if err != nil {
		util.Errorf("Error preparing insert statement for todos\n%s", err.Error())
		return nil, err
	}

	result, err := stmt.Exec(todo.Title, todo.Description)
	if err != nil {
		util.Errorf("Error inseting todo into database\n%s", err.Error())
		return nil, err
	}

	lastInsertedRowID, err := result.LastInsertId()
	if err != nil {
		util.Errorf("Error inserting todo into database\n%s", err.Error())
		return nil, err
	}

	todo, err = GetTodo(lastInsertedRowID)
	if err != nil {
		return nil, err
	}

	fmt.Printf("%v", result)

	return todo, nil
}

// GetTodo returns todo by id
func GetTodo(id int64) (*models.TodoFull, error) {
	row := DB.QueryRowx("SELECT * FROM todo WHERE id = ?", id)

	var todo models.TodoFull
	err := row.StructScan(&todo)
	if err != nil {
		util.Errorf("Error parsing row to todo\n%s", err.Error())
		return nil, err
	}

	return &todo, nil
}
