package routers

import (
	"online_shop/internal/handlers"
	"online_shop/internal/services/user_services"

	"github.com/go-chi/chi/v5"
)

func UserRouter(r chi.Router) {

	r.Post("/", handlers.CreateUser1)

	r.Get("/", handlers.GetAllUser)
	r.Get("/{id}", handlers.GetUserById)
	//router.Get("/users_auth/{id}", handlers.GetUserByIdAutorisation)
	r.Get("/login", user_services.LoginHandler)
	//router.Get("/protected", middlewares.AuthMiddleware(user_services.ProtectedHandler))

}
