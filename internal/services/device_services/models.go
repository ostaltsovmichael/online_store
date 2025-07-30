package device_services

type Device struct {
	Id      uint64  `json:"id"`
	Name    string  `json:"name"`
	Price   float32 `json:"price"`
	Rating  int     `json:"rating"`
	Img     string  `json:"img"`
	TypeId  int     `json:"type_id"`
	BrandId int     `json:"brand_id"`
}

type DeviceInfo struct {
	Id          uint64 `json:"id"`
	DeviceId    int    `json:"device_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type DeviceTypeBrand struct {
	Id         uint64  `json:"id"`
	Name       string  `json:"name"`
	Price      float32 `json:"price"`
	Rating     int     `json:"rating"`
	Img        string  `json:"img"`
	Brand      string  `json:"brand"`
	TypeDevice string  `json:"tytype_device"`
}

type DeviceInfoDevice struct {
	Id          uint64  `json:"id"`
	DeviceId    int     `json:"device_id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Name        string  `json:"name"`
	Price       float32 `json:"price"`
	Rating      int     `json:"rating"`
	Img         string  `json:"img"`
	Brand       string  `json:"brand"`
	TypeDevice  string  `json:"tytype_device"`
}
