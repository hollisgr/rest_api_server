package jwt

import (
	"rest_api/internal/cfg"
	"rest_api/internal/service/dto"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CreateToken(u dto.JWTTokenCreate) (string, error) {
	expTime := time.Now().Add(time.Hour * time.Duration(cfg.GetConfig().JWT.ExpTime)).Unix()

	payload := jwt.MapClaims{
		"id":          u.Id,
		"login":       u.Login,
		"first_name":  u.FirstName,
		"second_name": u.SecondName,
		"email":       u.Email,
		"exp":         expTime,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	t, err := token.SignedString([]byte(cfg.GetConfig().JWT.SecretKey))

	if err != nil {
		return t, err
	}
	return t, err
}

func ParseToken(tokenString string, key string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})
}
