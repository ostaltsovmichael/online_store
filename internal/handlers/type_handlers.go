package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"online_shop/internal/services/type_divace_services"

	"github.com/go-chi/chi/v5"
)

var typeModel = type_divace_services.Type{}

func CreateType(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	err := json.NewDecoder(r.Body).Decode(&typeModel)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		slog.Info(err.Error())
	}

	typeNew, err := typeModel.CreateType()
	if err != nil {
		fmt.Println("Ошибка при создании типа", err)
	}

	json.NewEncoder(w).Encode(typeModel)
	_ = typeNew

}

func GetTypeById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	newType, err := typeModel.GetTypeById(id)
	if err != nil {
		slog.Error("Не удается получить тип по id")
		log.Println(err)
	}
	err = json.NewEncoder(w).Encode(newType)
	if err != nil {
		slog.Error("Не удается получить данные по типу")
	}
	slog.Info("Тип получен по id: " + id)
}

