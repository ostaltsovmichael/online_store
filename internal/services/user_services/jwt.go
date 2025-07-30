package user_services

import (
	"online_shop/internal/configs"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var JwtKey = []byte(configs.Cfg.Secret_jwt_key) // В реальном приложении используйте безопасный ключ из конфига

type JWTClaims struct {
	Id    int    `json:"id"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}

// HashPassword хеширует пароль с помощью bcrypt
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 7)
	return string(bytes), err
}

// CheckPasswordHash проверяет, соответствует ли пароль хешу
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// GenerateJWT создает новый JWT токен для пользователя
func (user User) GenerateJWT() (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)

	claims := &JWTClaims{
		Id:    int(user.Id),
		Email: user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "myapp",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JwtKey)
}
