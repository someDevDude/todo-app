package database

import (
	"github.com/jmoiron/sqlx"
	"github.com/someDevDude/todo-server/util"
)

//DB holder for app
var DB *sqlx.DB

//Connect to DB
func Connect() {
	db, err := sqlx.Connect("mysql", "root:@/todolist?charset=utf8")

	DB = db

	util.CheckErr(err, func(err error) { panic(err) })
}
