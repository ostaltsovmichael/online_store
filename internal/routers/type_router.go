package routers

import (
	"online_shop/internal/handlers"

	"github.com/go-chi/chi/v5"
)

func TypeRouter(r chi.Router) {

	r.Post("/", handlers.CreateType)
	r.Get("/{id}", handlers.GetTypeById)

}
