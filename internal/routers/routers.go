package routers

import (
	"online_shop/internal/configs"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

var Router = chi.NewRouter()

func GetRouters() *chi.Mux {
	Router.Use(middleware.Logger)
	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{configs.Cfg.Host + "/*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	})
	Router.Use(cors.Handler)
	DeviceRouter()
	UserRouter()
	BrandRouter()
	TypeRouter()

	return Router
}
