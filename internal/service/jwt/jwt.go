package jwt

import (
	"rest_api/internal/cfg"
	"rest_api/internal/service/dto"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CreateToken(u dto.UserLoadDTO) string {
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
	t, _ := token.SignedString([]byte(cfg.GetConfig().JWT.SecretKey))
	return t
}

func ParseToken(tokenString string, key string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})
}
