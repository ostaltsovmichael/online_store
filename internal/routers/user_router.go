package routers

import (
	"online_shop/internal/handlers"
	"online_shop/internal/services/user_services"
)

func UserRouter() {

	router := Router

	router.Post("/users", handlers.CreateUser1)

	// router.Get("/users", handlers.GetAllUser)
	// router.Get("/users/{id}", handlers.GetUserById)
	//router.Get("/users_auth/{id}", handlers.GetUserByIdAutorisation)
	router.Get("/login", user_services.LoginHandler)
	//router.Get("/protected", middlewares.AuthMiddleware(user_services.ProtectedHandler))

}
