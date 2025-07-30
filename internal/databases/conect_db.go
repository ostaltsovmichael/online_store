package databases

import (
	"database/sql"
	"log"
	"log/slog"
)

type Postgres struct {
	Conect string
}
type SQLite struct {
	Conect string
}

type DataBases interface {
	ConectDB(dbPath string) (*sql.DB, error)
}

func DB() *sql.DB {
	var conect = SQLite{}

	db, err := conect.ConectDB()
	if err != nil {
		slog.Error("Bad conection to  db")
		log.Println(err)
	}
	return db
}
