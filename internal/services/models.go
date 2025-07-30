package services

type Bascet struct {
	Id     uint64 `json:"id"`
	UserId int    `json:"user_id"`
}

type Rating struct {
	Id       uint64 `json:"id"`
	UserId   int    `json:"user_id"`
	DeviceId int    `json:"device_id"`
	Rate     int    `json:"rate"`
}

type BascetDevice struct {
	Id       uint64 `json:"id"`
	BascetId int    `json:"bascet_id"`
	DeviceId int    `json:"device_id"`
}

type Type struct {
	Id   uint64 `json:"id"`
	Name string `json:"name"`
}

type Err_mes struct {
	Err_mes string `json:"err_mes"`
}
