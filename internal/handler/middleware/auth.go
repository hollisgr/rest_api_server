package middleware

import (
	"net/http"
	"rest_api/internal/cfg"
	"rest_api/internal/service/jwt"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		h := c.GetHeader("Authorization")
		hArr := strings.Split(h, "Bearer ")

		_, err := jwt.ParseToken(hArr[1], cfg.GetConfig().JWT.SecretKey)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"success": "false",
				"status":  "Unauthorized access",
				"msg":     "empty or incorrect token",
			})
		}
		c.Next()
	}
}
