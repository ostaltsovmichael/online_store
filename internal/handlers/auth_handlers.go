package handlers

import (
	"fmt"
	"log/slog"
	"net/http"
	"online_shop/internal/services/auth_services"
)

func CreateUser1(w http.ResponseWriter, r *http.Request) {

	user := auth_services.User{}
	//newToken := user_services.User{}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	u, err := user.CreateUser()
	if err != nil {
		slog.Error(err.Error())
	}

	fmt.Println(u)

}
