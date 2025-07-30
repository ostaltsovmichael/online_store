package app

import (
	"online_shop/internal/routers"
	"online_shop/internal/server"
)

func StartApp() {

	//godotenv.Load()

	router := routers.GetRouters()

	server.Server(router)
}
