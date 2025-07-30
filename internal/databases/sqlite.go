package databases

import (
	"database/sql"
	"log/slog"
	"online_shop/internal/configs"

	"log"

	_ "github.com/mattn/go-sqlite3"
)

func (p SQLite) ConectDB() (*sql.DB, error) {
	p.Conect = configs.Cfg.Sqlite_path

	db, err := sql.Open("sqlite3", p.Conect)
	if err != nil {
		slog.Error("Bad conection sqlite")
		log.Fatal(err)
	}
	//defer db.Close()

	return db, nil

}
