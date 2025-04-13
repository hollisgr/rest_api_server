package user

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"rest_api_server/internal/handlers"
	"rest_api_server/pkg/logging"

	"github.com/julienschmidt/httprouter"
)

const (
	usersURL = "/users"
	userURL  = "/users/:id"
)

type handler struct {
	logger  *logging.Logger
	storage Storage
	respMsg RestMsg
}

func NewHandler(logger *logging.Logger, storage Storage) handlers.Handler {
	return &handler{
		logger:  logger,
		storage: storage,
	}
}

func (h *handler) Register(router *httprouter.Router) {
	router.GET(usersURL, h.GetList)
	router.GET(userURL, h.GetUserByID)
	router.POST(usersURL, h.CreateUser)
	router.DELETE(userURL, h.DeleteUser)
}

func (h *handler) GetList(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	h.logger.Infoln("incoming request to getlist")
	users, err := h.storage.FindAllUsers(context.Background())
	if err != nil {
		msg := h.respMsg.CreateMsgJson(404, "Not found", "Userlist is empty")
		w.WriteHeader(http.StatusNotFound)
		w.Write(msg)
		return
	}
	result, mErr := json.Marshal(users)
	if mErr != nil {
		h.logger.Infoln("Marshall error")
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func (h *handler) GetUserByID(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	h.logger.Infoln("incoming request to user by id")
	idStr := params.ByName("id")

	var id int64

	n, _ := fmt.Sscanf(idStr, "%d", &id)
	if n < 1 {
		h.logger.Infoln("Uncorrect id")
		msg := h.respMsg.CreateMsgJson(400, "Bad Request", "Uncorrect id")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(msg)
		return
	}

	user, err := h.storage.FindUser(context.Background(), id)
	if err != nil {
		h.logger.Infoln("User not found")
		msg := h.respMsg.CreateMsgJson(404, "Not found", "User not found")
		w.WriteHeader(http.StatusNotFound)
		w.Write(msg)
		return
	}
	result, mErr := json.Marshal(user)
	if mErr != nil {
		h.logger.Infoln("Marshall error")
		msg := h.respMsg.CreateMsgJson(400, "Bad Request", "Marshal err")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(msg)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}
func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	h.logger.Infoln("incoming request to create user")
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	newuser := User{}
	var uid string

	err = json.Unmarshal(body, &newuser)
	if err != nil {
		h.logger.Infoln("cant unmarshal")
		msg := h.respMsg.CreateMsgJson(400, "Bad Request", "Unmarshal err")
		w.Write(msg)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	uid, err = h.storage.CreateUser(context.Background(), newuser)
	if err != nil {
		h.logger.Infoln("cant create")
		msg := h.respMsg.CreateMsgJson(400, "Bad Request", "Create err")
		w.Write(msg)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	msg := h.respMsg.CreateMsgJson(201, "Created", fmt.Sprintf("Successful created user uid: %s", uid))
	w.WriteHeader(http.StatusOK)
	w.Write(msg)
}
func (h *handler) DeleteUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	h.logger.Infoln("incoming request to delete user by id")
	idStr := params.ByName("id")

	var id int64

	n, _ := fmt.Sscanf(idStr, "%d", &id)
	if n < 1 {
		h.logger.Infoln("Uncorrect id")
		msg := h.respMsg.CreateMsgJson(400, "Bad Request", "Uncorrect id")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(msg)
		return
	}

	err := h.storage.DeleteUser(context.Background(), id)
	if err != nil {
		h.logger.Infoln("Marshall error")
		msg := h.respMsg.CreateMsgJson(404, "Not found", "User not found")
		w.WriteHeader(http.StatusNotFound)
		w.Write(msg)
		return
	}
	msg := h.respMsg.CreateMsgJson(200, "OK", fmt.Sprintf("Successful deleted id: %d", id))
	w.WriteHeader(http.StatusOK)
	w.Write(msg)
}
