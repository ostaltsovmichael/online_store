package device_services

import (
	"database/sql"
	"log"
	"log/slog"
)

func (d *DeviceInfo) CreateDiviceInfo() (*sql.Stmt, error) {
	row, err := db.Prepare(`INSERT INTO device_infos (title, description, deviceid) VALUES ($1,$2,$3)`)
	if err != nil {
		slog.Error("Не коректный запрос к таблице device_infos")
		log.Println(err)
	}

	row.Exec(&d.Title, &d.Description, &d.DeviceId)
	return row, nil
}
