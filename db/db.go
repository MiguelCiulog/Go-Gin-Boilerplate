package db

// TODO: Change to mysql database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Init() {
	var err error
	DB, err = sql.Open("mysql", "root:passwd@tcp(0.0.0.0:3306)/user")
	if err != nil {
		panic("Error when creating database")
	}
	Migrate()
}
