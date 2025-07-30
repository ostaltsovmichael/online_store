package auth_services

import (
	"database/sql"
	"log"
	"log/slog"
	"online_shop/internal/databases"
	"online_shop/internal/models"
)

type User struct {
	models.User
}

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
