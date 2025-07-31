package device_services

import (
	"database/sql"
	"fmt"
	"log"
	"log/slog"
	"online_shop/internal/databases"
)

var db = databases.DB()

// @Summary Создать девайс
// @Description Создать девайс
// @Tags device
// @Accept json
// @Produce json
// @Success 200 {object} Device
// @Router /devices/[post]
func (d *Device) CreateDevice() (*sql.Stmt, error) {

	device, err := db.Prepare(`INSERT INTO devices (name, price, img, typeid,brandid) VALUES ($1, $2, $3, $4, $5)`)
	if err != nil {
		slog.Error("Не коректный запрос к БД таблице devices")
		log.Println(err)
	}

	// img, err := os.OpenFile("./static/img/iphone16wite.jpeg")
	device.Exec(&d.Name, &d.Price, &d.Img, &d.TypeId, &d.BrandId)
	return device, nil

}

func (d *Device) GetDeviceByID(id string) (Device, error) {

	row := db.QueryRow(`SELECT name, price, img, typeid,brandid,id  FROM devices where id=$1`, id)

	row.Scan(&d.Name, &d.Price, &d.Img, &d.TypeId, &d.BrandId, &d.Id)

	if id != fmt.Sprint(d.Id) {
		return Device{}, fmt.Errorf("Девайса не существует")
	} else {
		return *d, nil
	}
}

func (d *DeviceTypeBrand) GetDeviceTypeBrandByID(id string) (DeviceTypeBrand, error) {
	row := db.QueryRow(`SELECT (devices.id),(devices.name), price, ratind, img,(brands.name),(types.name) 
FROM devices  
RIGHT JOIN brands ON (devices.brandid) = (brands.id) 
RIGHT JOIN types ON (devices.typeid) = (types.id)   WHERE (devices.id) = $1;`, id)

	row.Scan(&d.Id, &d.Name, &d.Price, &d.Rating, &d.Img, &d.Brand, &d.TypeDevice)

	return *d, nil
}

// func (d *Device) GetAllDevices() ([]Device, error)  {
// 	db.Prepare(`SELECT * FROM devices`)

// 	return []d, nil
// }
