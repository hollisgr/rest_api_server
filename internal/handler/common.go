package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SendError(c *gin.Context, code int, err error) {
	statusText := http.StatusText(code)
	c.AbortWithStatusJSON(code, gin.H{
		"success": false,
		"status":  statusText,
		"message": fmt.Sprintf("%v", err),
	})
}

func SendSuccess(c *gin.Context, code int, msg string) {
	statusText := http.StatusText(code)
	c.JSON(code, gin.H{
		"success": true,
		"status":  statusText,
		"message": msg,
	})
}

func GetID(c *gin.Context) (int, error) {
	id := 0
	idStr := c.Params.ByName("id")
	_, err := fmt.Sscanf(idStr, "%d", &id)
	return id, err
}
