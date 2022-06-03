package models

import (
	"time"

	"database/sql"
	database "github.com/MiguelCiulog/Go-Gin-Boilerplate/db"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	IdUser    int       `json:"id_user"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	BirthDate time.Time `json:"birth_date"`
}

var db *sql.DB = database.DB

func (u User) GetAllUsers() ([]User, error) {
	sql := "SELECT * FROM users"

	results, err := db.Query(sql)

	if err != nil {
		panic(err.Error())
	}

	var users []User
	for results.Next() {
		var user User
		err = results.Scan(
			&user.IdUser,
			&user.Name,
			&user.Email,
			&user.BirthDate,
		)

		if err != nil {
			panic(err.Error())
		}

		users = append(users, user)
	}

	return users, nil
}
