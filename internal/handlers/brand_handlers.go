package handlers

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"online_shop/internal/middlewares/chec_id"
	"online_shop/internal/services/brand_services"

	"github.com/go-chi/chi/v5"
)

var brandServive = brand_services.Brand{}

func CreateBrand(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	err := json.NewDecoder(r.Body).Decode(&brandServive)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		slog.Info(err.Error())
	}

	typeNew, err := brandServive.CreateBrand()
	if err != nil {
		fmt.Println("Ошибка при создании типа", err)
	}

	json.NewEncoder(w).Encode(brandServive)
	_ = typeNew

}

func GetBrandById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	id := chi.URLParam(r, "id")
	brand, err := brandServive.GetBrandById(id)
	if err != nil {
		slog.Error("Не удалось получить бренд")
	}

	p := chec_id.Param_chec_id{
		Id:           id,
		ModelId:      brand.Id,
		Name:         "brand",
		Writer:       w,
		ModelSecsees: brand,
	}
	chec_id.ChecId(p)

}
