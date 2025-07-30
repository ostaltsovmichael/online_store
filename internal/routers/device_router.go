package routers

import "online_shop/internal/handlers"

func DeviceRouter() {

	router := Router

	router.Post("/devices", handlers.CreateDevice)
	router.Post("/device_infos", handlers.CreateDiviceInfo)

	router.Get("/devices/{id}", handlers.GetDeviceByID)
	router.Get("/devices/type/brand/{id}", handlers.GetDeviceBrandTypeByIdDevice)

}
