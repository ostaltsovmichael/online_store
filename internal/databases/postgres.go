package databases

import (
	"database/sql"
	"log/slog"
	"online_shop/internal/configs"

	"log"

	_ "github.com/lib/pq"
)

func (p Postgres) ConectDB() (*sql.DB, error) {
	p.Conect = configs.Cfg.Postgres_url

	db, err := sql.Open("postgres", p.Conect)
	if err != nil {
		slog.Error("Bad conection postgres")
		log.Fatal(err)
	}
	//defer db.Close()

	return db, nil

}
