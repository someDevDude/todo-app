package database

import (
	"database/sql"

	"github.com/someDevDude/todo-server/util"
)

//DB holder for app
var DB *sql.DB

//Connect to DB
func Connect() {
	db, err := sql.Open("mysql", "root:@/todolist?charset=utf8")

	DB = db

	util.CheckErr(err, func(err error) { panic(err) })
}
