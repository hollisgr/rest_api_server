package handler

import (
	"fmt"
	"net/http"
	"rest_api/internal/service/dto"

	"github.com/gin-gonic/gin"
)

type WebUserAuth struct {
	Login    string `json:"login" validate:"required,min=5,max=20"`
	Password string `json:"password" validate:"required"`
}

type UserLoadResp struct {
	Success bool            `json:"success" example:"true"`
	User    dto.WebUserLoad `json:"user"`
}

type UpdateUserResp struct {
	Success bool            `json:"success" example:"true"`
	User    dto.WebUserLoad `json:"user"`
}

type Msg struct {
	Success bool   `json:"success" example:"true"`
	Status  string `json:"status" example:"status text"`
	Message string `json:"message" example:"message text"`
}

type Err struct {
	Success bool   `json:"success" example:"false"`
	Status  string `json:"status" example:"status text"`
	Message string `json:"message" example:"message text"`
}

func SendError(c *gin.Context, code int, err error) {
	statusText := http.StatusText(code)
	respErr := Err{
		Success: false,
		Status:  statusText,
		Message: fmt.Sprintf("%v", err),
	}
	c.AbortWithStatusJSON(code, respErr)
}

func SendSuccess(c *gin.Context, code int, msg string) {
	statusText := http.StatusText(code)
	respMsg := Msg{
		Success: true,
		Status:  statusText,
		Message: msg,
	}
	c.JSON(code, respMsg)
}

func GetID(c *gin.Context) (int, error) {
	id := 0
	idStr := c.Params.ByName("id")
	_, err := fmt.Sscanf(idStr, "%d", &id)
	return id, err
}
