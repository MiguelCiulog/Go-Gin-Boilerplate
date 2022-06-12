package db

// TODO: Change to mysql database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	config "github.com/MiguelCiulog/Go-Gin-Boilerplate/config"
)

var db *sql.DB

func Init() {
	var err error
	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		config.Env.DB.DatabaseUser, config.Env.DB.DatabasePassword,
		config.Env.DB.DatabaseUrl, config.Env.DB.DatabasePort,
		config.Env.DB.DatabaseName)

	db, err = sql.Open("mysql", connection)
	if err != nil {
		panic("Error when creating database")
	}
	// Optional migration
	Migrate()
}

func GetDatabase() *sql.DB {
	return db
}
