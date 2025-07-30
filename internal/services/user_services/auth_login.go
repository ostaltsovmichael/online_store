package user_services

import (
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
)

// LoginHandler обрабатывает аутентификацию пользователя
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	users := []User{}

	var user User
	found := false
	for _, u := range users {
		if u.Email == username && CheckPasswordHash(password, u.Password) {
			user = u
			found = true
			break
		}
	}

	if !found {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Invalid credentials")
		return
	}

	token, err := user.GenerateJWT()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Could not generate token")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"token": "%s"}`, token)
}

// ProtectedHandler пример защищенного маршрута
func ProtectedHandler(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	tokenString := authHeader[len("Bearer "):]

	claims := &JWTClaims{}

	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Invalid token")
		return
	}

	fmt.Fprintf(w, "Hello, %s! Your user ID is %d", claims.Email, claims.Id)
}
