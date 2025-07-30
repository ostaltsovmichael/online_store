package brand_services

import (
	"database/sql"
	"log"
	"log/slog"
	"online_shop/internal/databases"
)

var db = databases.DB()

func (b *Brand) GetBrandById(id string) (Brand, error) {

	row := db.QueryRow(`SELECT id,name FROM brands where id=$1`, id)

	row.Scan(&b.Id, &b.Name)

	return *b, nil
}

func (t *Brand) CreateBrand() (*sql.Stmt, error) {

	row, err := db.Prepare(`INSERT INTO brands (name) VALUES ($1)`)
	if err != nil {
		slog.Error("Не корекный запрос к бд")
		log.Fatal(err)
	}
	row.Exec(&t.Name)
	slog.Info("Бренд создан")
	return row, nil
}
