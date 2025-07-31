package routers

import (
	"online_shop/internal/handlers"

	"github.com/go-chi/chi/v5"
)

func DeviceRouter(r chi.Router) {

	r.Post("/", handlers.CreateDevice)
	r.Post("/device_infos", handlers.CreateDiviceInfo)

	r.Get("/{id}", handlers.GetDeviceByID)
	r.Get("/type/brand/{id}", handlers.GetDeviceBrandTypeByIdDevice)

}
