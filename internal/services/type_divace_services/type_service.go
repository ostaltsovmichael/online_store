package type_divace_services

import (
	"database/sql"
	"log"
	"log/slog"
	"online_shop/internal/databases"
)

var db = databases.DB()

func (t *Type) CreateType() (*sql.Stmt, error) {
	row, err := db.Prepare(`INSERT INTO types (name) VALUES ($1)`)
	if err != nil {
		slog.Error("Не корекный запрос к бд")
		log.Fatal(err)
	}
	row.Exec(&t.Name)
	slog.Info("Тип создан")
	return row, nil
}

func (t *Type) GetTypeById(id string) (Type, error) {

	row := db.QueryRow(`SELECT id, name FROM types where id=$1`, id)

	row.Scan(&t.Id, &t.Name)

	return *t, nil
}
