package database

import (
	"os"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/someDevDude/todo-server/util"
)

//DB holder for app
var DB *sqlx.DB

//Connect to DB
func Connect() {

	var dbConnectionString strings.Builder
	dbConnectionString.WriteString(os.Getenv("DB_USER"))
	dbConnectionString.WriteString(":")
	dbConnectionString.WriteString(os.Getenv("DB_PW"))
	dbConnectionString.WriteString("@tcp(")
	dbConnectionString.WriteString(os.Getenv("DB_URL"))
	dbConnectionString.WriteString(":3306)/")
	dbConnectionString.WriteString(os.Getenv("DB_NAME"))

	db, err := sqlx.Connect("mysql", dbConnectionString.String())

	DB = db

	util.CheckErr(err, func(err error) { panic(err) })
}
