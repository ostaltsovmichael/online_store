package routers

import "online_shop/internal/handlers"

func BrandRouter() {

	router := Router

	router.Post("/brands", handlers.CreateBrand)

}
