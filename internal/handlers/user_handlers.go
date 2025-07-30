package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"log/slog"
	"net/http"

	"online_shop/internal/middlewares/chec_id"
	"online_shop/internal/services/user_services"

	"github.com/go-chi/chi/v5"
)

var userModel = user_services.User{}

func CreateUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	err := json.NewDecoder(r.Body).Decode(&userModel)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		slog.Info(err.Error())
	}

	newUser, err := userModel.CreateUser()
	if err != nil {
		fmt.Println("Ошибка при создании игрока", err)
	}

	json.NewEncoder(w).Encode(userModel)
	_ = newUser
	slog.Info("Игрок создан ")

}

func GetAllUser(w http.ResponseWriter, r *http.Request) {

	users, err := userModel.GetAllUser()
	if err != nil {
		slog.Error("Не удалось получить пользовотелей")
		log.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(users)

	slog.Info("Cписок пользовотелей получен ")

}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	id := chi.URLParam(r, "id")
	user, err := userModel.GetUserById(id)
	if err != nil {
		slog.Error("Не удается получить парамет id у рользовотеля")
		log.Println(err)

	}
	p := chec_id.Param_chec_id{
		Id:           id,
		ModelId:      user.Id,
		Name:         "User",
		Writer:       w,
		ModelSecsees: user,
	}
	chec_id.ChecId(p)
}

// func GetUserByIdAutorisation(w http.ResponseWriter, r *http.Request) {

// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	id := chi.URLParam(r, "id")
// 	user, err := userModel.GetUserById(id)
// 	if err != nil {
// 		slog.Error(fmt.Sprint(err))
// 	}
// 	p := chec_id.Param_chec_id{
// 		Id:           id,
// 		ModelId:      user.Id,
// 		Name:         "User",
// 		Writer:       w,
// 		ModelSecsees: user,
// 	}
// 	chec_id.ChecId(p)
// }
