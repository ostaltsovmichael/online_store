package routers

import (
	"online_shop/internal/handlers"

	"github.com/go-chi/chi/v5"
)

func BrandRouter(r chi.Router) {
	r.Post("/", handlers.CreateBrand)
}
