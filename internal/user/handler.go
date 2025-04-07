package user

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"rest_api_server/internal/handlers"
	"rest_api_server/pkg/logging"

	"github.com/julienschmidt/httprouter"
)

const (
	usersURL = "/users"
	userURL  = "/users/:uuid"
)

type handler struct {
	logger  *logging.Logger
	storage Storage
}

func NewHandler(logger *logging.Logger, storage Storage) handlers.Handler {
	return &handler{
		logger:  logger,
		storage: storage,
	}
}

func (h *handler) Register(router *httprouter.Router) {
	router.GET(usersURL, h.GetList)
	router.GET(userURL, h.GetUserByUUID)
	router.POST(userURL, h.CreateUser)
	router.DELETE(userURL, h.DeleteUser)
}

func (h *handler) GetList(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	h.logger.Infoln("incoming request to getlist")
	users, err := h.storage.FindAllUsers(context.Background())
	if err != nil {
		w.WriteHeader(400)
	}
	result, mErr := json.Marshal(users)
	if mErr != nil {
		h.logger.Infoln("Marshall error")
	}

	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func (h *handler) GetUserByUUID(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	h.logger.Infoln("incoming request to user by id")
	idStr := params.ByName("uuid")

	h.logger.Infoln("idstr = ", idStr)

	var id int64

	fmt.Sscanf(idStr, "%d", &id)

	user, err := h.storage.FindUser(context.Background(), id)
	if err != nil {
		w.WriteHeader(400)
		return
	}
	result, mErr := json.Marshal(user)
	if mErr != nil {
		h.logger.Infoln("Marshall error")
	}
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}
func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	h.logger.Infoln("incoming request to create user")
	w.Write([]byte(fmt.Sprintln("THIS IS CREATE USER")))

}
func (h *handler) DeleteUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	h.logger.Infoln("incoming request to user by id")
	idStr := params.ByName("uuid")

	h.logger.Infoln("idstr = ", idStr)

	var id int64

	fmt.Sscanf(idStr, "%d", &id)

	err := h.storage.DeleteUser(context.Background(), id)
	if err != nil {
		w.WriteHeader(400)
		return
	}
	w.WriteHeader(http.StatusOK)
}
