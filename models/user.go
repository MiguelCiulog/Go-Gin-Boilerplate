package models

import (
	"fmt"
	"time"

	database "github.com/MiguelCiulog/Go-Gin-Boilerplate/db"
)

type User struct {
	IdUser    int       `json:"id_user"`
	Name      string    `json:"name"`
	Email     *string   `json:"email"`
	BirthDate time.Time `json:"birth_date"`
}

func (u User) GetAllUsers() ([]User, error) {
	db := database.GetDatabase()
	fmt.Println(db)
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

func (u User) GetUserById(id int) (User, error) {
	db := database.GetDatabase()
	fmt.Println(db)
	sql := fmt.Sprintf("SELECT * FROM users WHERE id_user = %d", id)

	results, err := db.Query(sql)

	if err != nil {
		panic(err.Error())
	}

	var user User
	for results.Next() {
		err = results.Scan(
			&user.IdUser,
			&user.Name,
			&user.Email,
			&user.BirthDate,
		)

		if err != nil {
			panic(err.Error())
		}
	}

	return user, nil
}
