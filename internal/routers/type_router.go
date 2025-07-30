package routers

import "online_shop/internal/handlers"

func TypeRouter() {

	router := Router

	router.Post("/types", handlers.CreateType)
	router.Get("/types/{id}", handlers.GetTypeById)

}
