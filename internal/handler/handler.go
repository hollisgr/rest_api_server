package handler

import (
	"fmt"
	"net/http"
	"rest_api/internal/handler/middleware"
	"rest_api/internal/handler/repository"
	"rest_api/internal/service/dto"
	"rest_api/internal/service/user_repository"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	UserRepository user_repository.UserRepository
}

func NewHandler(u user_repository.UserRepository) repository.Handler {
	return &Handler{
		UserRepository: u,
	}
}

func (h *Handler) Register(r *gin.Engine) {
	r.Use(gin.Logger())
	authRequired := r.Group("/")
	authRequired.Use(middleware.AuthMiddleware())
	{
		authRequired.GET("/users/:id", h.UserById)
		authRequired.GET("/users", h.UserList)
		authRequired.PATCH("/users/:id", h.UpdateUserByID)
		authRequired.DELETE("/users/:id", h.DeleteUserByID)
	}

	r.POST("/users", h.CreateNewUser)
	r.POST("/auth", h.Authentification)
}

func (h *Handler) Authentification(c *gin.Context) {
	userData := dto.WebUserAuth{}

	err := c.BindJSON(&userData)
	if err != nil {
		SendError(c, http.StatusUnauthorized, err)
		return
	}

	token, err := h.UserRepository.AuthUser(userData)

	if err != nil {
		SendError(c, http.StatusUnauthorized, err)
		return
	}

	SendSuccess(c, http.StatusOK, fmt.Sprintf("Bearer %s", token))
}

func (h *Handler) CreateNewUser(c *gin.Context) {
	userData := dto.WebUserCreate{}

	err := c.BindJSON(&userData)
	if err != nil {
		SendError(c, http.StatusUnauthorized, err)
		return
	}

	id, err := h.UserRepository.CreateUser(userData)

	if err != nil {
		SendError(c, http.StatusBadRequest, err)
		return
	}

	SendSuccess(c, http.StatusOK, fmt.Sprintf("created user with id: %d", id))
}

func (h *Handler) UserList(c *gin.Context) {
	userList, err := h.UserRepository.LoadUserList()
	if err != nil {
		// SendError(c, http.StatusNotFound, fmt.Errorf("userlist is empty"))
		SendError(c, http.StatusNotFound, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success":    "true",
		"users list": userList,
	})
}

func (h *Handler) UserById(c *gin.Context) {
	idStr := c.Params.ByName("id")
	id := 0

	_, err := fmt.Sscanf(idStr, "%d", &id)
	if err != nil {
		SendError(c, http.StatusBadRequest, fmt.Errorf("incorrect id"))
		return
	}

	u, err := h.UserRepository.LoadUserByID(id)
	if err != nil {
		SendError(c, http.StatusBadRequest, fmt.Errorf("user not found"))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": "true",
		"user":    u,
	})
}

func (h *Handler) UpdateUserByID(c *gin.Context) {
	id, err := GetID(c)
	if err != nil {
		SendError(c, http.StatusBadRequest, fmt.Errorf("incorrect id"))
		return
	}

	u := dto.WebUserUpdate{}
	err = c.BindJSON(&u)
	if err != nil {
		SendError(c, http.StatusBadRequest, err)
		return
	}

	u.Id = id

	err = h.UserRepository.UpdateUser(u)
	if err != nil {
		SendError(c, http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": "true",
		"user":    u,
	})
}

func (h *Handler) DeleteUserByID(c *gin.Context) {
	id, err := GetID(c)
	if err != nil {
		SendError(c, http.StatusBadRequest, fmt.Errorf("incorrect id"))
		return
	}
	err = h.UserRepository.DeleteUser(id)
	if err != nil {
		SendError(c, http.StatusBadRequest, err)
		return
	}

	SendSuccess(c, http.StatusOK, fmt.Sprintf("deleted user_id: %d", id))
}
