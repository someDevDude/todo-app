package database

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/someDevDude/todo-server/util"
)

//DB holder for app
var DB *sqlx.DB

//Connect to DB
func Connect() {
	fmt.Printf(os.Getenv("DB_URL"))
	db, err := sqlx.Connect("mysql", os.Getenv("DB_URL"))

	DB = db

	util.CheckErr(err, func(err error) { panic(err) })
}
