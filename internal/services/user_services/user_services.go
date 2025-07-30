package user_services

import (
	"database/sql"
	"fmt"
	"log"
	"log/slog"
	"online_shop/internal/databases"
)

var db = databases.DB()

func (u *User) CreateUser() (*sql.Stmt, error) {

	user, err := db.Prepare(`INSERT INTO users (email, password) VALUES ($1, $2)`)
	if err != nil {
		slog.Error("Не корекный запрос к бд")
		log.Fatal(err)
	}

	user.Exec(u.Email, u.Password)

	return user, nil
}

func (u *User) GetAllUser() ([]User, error) {

	rows, err := db.Query(`SELECT id,email,password,role FROM users`)
	if err != nil {
		slog.Error("Не корекный запрос к бд")
		log.Println(err)

	}

	users := []User{}

	for rows.Next() {
		err := rows.Scan(&u.Id, &u.Email, &u.Password, &u.Role)
		if err != nil {
			slog.Error("Не удалось получить пользовотелей")
			log.Println(err)
		}

		users = append(users, *u)

	}

	return users, nil
}

func (u *User) GetUserById(id string) (User, error) {

	row := db.QueryRow(`SELECT id,email,password,role FROM users where id=$1`, id)

	row.Scan(&u.Id, &u.Email, &u.Password, &u.Role)

	chec_id := fmt.Sprint(u.Id)

	if id != chec_id {

		return User{
			Email: "Пользователя не существует",
		}, fmt.Errorf("пользовотель не существует err")
	} else if id == chec_id {

		log.Println("Пользовотель по id: ", u.Id, " получен")

	}

	return *u, nil
}
