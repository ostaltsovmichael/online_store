package middlewares

// import (
// 	"fmt"
// 	"net/http"

// 	"online_shop/internal/services/user_services"

// 	"github.com/golang-jwt/jwt/v5"
// )

// func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		authHeader := r.Header.Get("Authorization")
// 		if authHeader == "" {
// 			w.WriteHeader(http.StatusUnauthorized)
// 			fmt.Fprint(w, "Missing Authorization Header")
// 			return
// 		}
// 		tokenString := authHeader[len("Bearer "):]
// 		claims := &user_services.JWTClaims{}

// 		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
// 			return user_services.JwtKey, nil
// 		})
// 		if err != nil || !token.Valid {
// 			w.WriteHeader(http.StatusUnauthorized)
// 			fmt.Fprint(w, "Invalid token")
// 			return
// 		}

// 		next(w, r)

// 	}

// }
