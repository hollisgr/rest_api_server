package handler

import (
	"fmt"
	"net/http"
	"rest_api/docs"
	"rest_api/internal/handler/middleware"
	"rest_api/internal/handler/repository"
	"rest_api/internal/service/dto"
	"rest_api/internal/service/user_repository"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	UserRepository user_repository.UserRepository
}

func NewHandler(u user_repository.UserRepository) repository.Handler {
	//swagger init
	docs.SwaggerInfo.Title = "Swagger Example API"
	docs.SwaggerInfo.Description = "This is a sample server Petstore server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "127.0.0.1:8181"
	docs.SwaggerInfo.BasePath = "/"

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
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

// Authentification godoc
//
//	@Summary		Authentification
//	@Description	auth user by login and password
//	@Accept			json
//	@Produce		json
//	@Param			account	body		dto.WebUserAuth	true	"Account data"
//	@Success		200		{object}	handler.Msg
//	@Failure		400		{object}	handler.Err
//	@Failure		401		{object}	handler.Err
//	@Router			/auth [post]
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

// CreateUser godoc
//
//	@Summary		Create New User
//	@Description	creating user using data
//	@Accept			json
//	@Produce		json
//	@Param			account	body		dto.WebUserCreate	true	"User create data"
//	@Success		200		{object}	handler.Msg
//	@Failure		400		{object}	handler.Err
//	@Failure		401		{object}	handler.Err
//	@Router			/users [post]
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

// UserList godoc
//
//	@Summary		User list
//	@Description	List of all users
//	@Tags			protected
//	@Produce		json
//	@Param			Authorization	header		string	true	"user jwt token"
//	@Success		200				{object}	handler.Msg
//	@Failure		400				{object}	handler.Err
//	@Failure		401				{object}	handler.Err
//	@Router			/users [get]
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

// UserById godoc
//
//	@Summary		User by id
//	@Description	User data by id
//	@Produce		json
//	@Tags			protected
//	@Param			Authorization	header		string	true	"user jwt token"
//	@Param			id				path		int		true	"User ID"
//	@Success		200				{object}	handler.UserLoadResp
//	@Failure		400				{object}	handler.Err
//	@Failure		401				{object}	handler.Err
//	@Router			/users/:id [get]
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

// UpdateUserByID godoc
//
//	@Summary		Update user by id
//	@Description	Update user data by id
//	@Tags			protected
//	@Produce		json
//	@Param			Authorization	header		string	true	"user jwt token"
//	@Param			id				path		int		true	"User ID"
//	@Success		200				{object}	handler.UpdateUserResp
//	@Failure		400				{object}	handler.Err
//	@Failure		401				{object}	handler.Err
//	@Router			/users/:id [patch]
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

// DeleteUserByID godoc
//
//	@Summary		Delete user by id
//	@Description	Delete user by id
//	@Tags			protected
//	@Produce		json
//	@Param			Authorization	header		string	true	"user jwt token"
//	@Param			id				path		int		true	"User ID"
//	@Success		200				{object}	handler.Msg
//	@Failure		400				{object}	handler.Err
//	@Failure		401				{object}	handler.Err
//	@Router			/users/:id [delete]
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
