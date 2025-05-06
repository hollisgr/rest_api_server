package user

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"rest_api_server/internal/config"
	"rest_api_server/internal/handlers"
	"rest_api_server/pkg/logging"

	"github.com/julienschmidt/httprouter"
)

const (
	usersURL   = "/users"
	userURL    = "/users/:id"
	adminURL   = "/admin"
	tgUsersURL = "/tg_users"
)

type handler struct {
	logger  *logging.Logger
	storage Storage
	respMsg RestMsg
	cfg     config.Config
}

func NewHandler(logger *logging.Logger, storage Storage) handlers.Handler {
	return &handler{
		logger:  logger,
		storage: storage,
		cfg:     *config.GetConfig(),
	}
}

func (h *handler) Register(router *httprouter.Router) {
	router.GET(usersURL, h.GetList)
	router.GET(userURL, h.GetUserByID)
	router.GET(tgUsersURL, h.GetListTgUsers)
	router.POST(usersURL, h.CreateUser)
	router.POST(tgUsersURL, h.CreateTgUser)
	router.POST(adminURL, h.SetAdmin)
	router.DELETE(userURL, h.DeleteUser)
}

func (h *handler) GetList(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	h.logger.Infoln("incoming request to getlist")
	users, _ := h.storage.FindAllUsers(context.Background())
	if len(users) < 1 {
		h.respMsg.SendMsgJson(w, http.StatusNotFound, "Not found", "Userlist is empty")
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

func (h *handler) GetListTgUsers(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	h.logger.Infoln("incoming request to getlist")
	users, _ := h.storage.FindAllUsers(context.Background())
	if len(users) < 1 {
		h.respMsg.SendMsgJson(w, http.StatusNotFound, "Not found", "Userlist is empty")
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
		h.respMsg.SendMsgJson(w, http.StatusBadRequest, "Bad Request", "Uncorrect id")
		return
	}

	user, err := h.storage.FindUser(context.Background(), id)
	if err != nil {
		h.logger.Infoln("User not found")
		h.respMsg.SendMsgJson(w, http.StatusNotFound, "Not found", "User not found")
		return
	}
	result, mErr := json.Marshal(user)
	if mErr != nil {
		h.logger.Infoln("Marshall error")
		h.respMsg.SendMsgJson(w, http.StatusBadRequest, "Bad Request", "Marshal err")
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func (h *handler) CreateTgUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	h.logger.Infoln("incoming request to create tg user")
	body, err := io.ReadAll(r.Body)
	if err != nil {
		h.respMsg.SendMsgJson(w, http.StatusBadRequest, "Bad Request", "Incorrect userdata")
		return
	}

	newuser := TGUser{}

	err = json.Unmarshal(body, &newuser)
	if err != nil {
		h.logger.Infoln("cant unmarshal")
		h.respMsg.SendMsgJson(w, http.StatusBadRequest, "Bad Request", "Unmarshal err")
		return
	}

	err = h.storage.CreateTgUser(context.Background(), newuser)

	if err != nil {
		h.logger.Error(err)
		h.respMsg.SendMsgJson(w, http.StatusBadRequest, "Bad Request", "Create err")
		return
	}

	h.respMsg.SendMsgJson(w, http.StatusCreated, "Created", "Successful created tg user")
}

func (h *handler) SetAdmin(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	h.logger.Infoln("incoming request to set admin")
	body, err := io.ReadAll(r.Body)
	if err != nil {
		h.respMsg.SendMsgJson(w, http.StatusBadRequest, "Bad Request", "Incorrect userdata")
		return
	}

	newuser := TGAdmin{}

	err = json.Unmarshal(body, &newuser)
	if err != nil {
		h.logger.Infoln("cant unmarshal")
		h.respMsg.SendMsgJson(w, http.StatusBadRequest, "Bad Request", "Unmarshal err")
		return
	}

	if newuser.ADMIN_PWD != h.cfg.AdminPassword {
		h.logger.Infoln("incorrect password")
		h.respMsg.SendMsgJson(w, http.StatusBadRequest, "Bad Request", "Incorrect Password")
		return
	}

	err = h.storage.SetAdmin(context.Background(), newuser.TG_ID)

	if err != nil {
		h.logger.Error(err)
		h.respMsg.SendMsgJson(w, http.StatusBadRequest, "Bad Request", "user not found")
		return
	}

	h.respMsg.SendMsgJson(w, http.StatusAccepted, "Access granted", "Successful admin access")
}

func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	h.logger.Infoln("incoming request to create user")
	body, err := io.ReadAll(r.Body)
	if err != nil {
		h.respMsg.SendMsgJson(w, http.StatusBadRequest, "Bad Request", "Incorrect userdata")
		return
	}

	newuser := User{}

	err = json.Unmarshal(body, &newuser)
	if err != nil {
		h.logger.Infoln("cant unmarshal")
		h.respMsg.SendMsgJson(w, http.StatusBadRequest, "Bad Request", "Unmarshal err")
		return
	}

	err = isAdmin(h.storage, &newuser)

	if err != nil {
		h.logger.Infoln("access denied")
		h.respMsg.SendMsgJson(w, http.StatusUnauthorized, "Unauthorized", "Access denied")
		return
	}

	valErr := pwdValidation(w, h, newuser.Password)
	if valErr != nil {
		return
	}

	err = h.storage.CreateUser(context.Background(), newuser)

	if err != nil {
		h.logger.Error(err)
		h.respMsg.SendMsgJson(w, http.StatusBadRequest, "Bad Request", "Create err")
		return
	}

	h.respMsg.SendMsgJson(w, http.StatusCreated, "Created", "Successful created user")
}

func (h *handler) DeleteUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	h.logger.Infoln("incoming request to delete user by id")
	idStr := params.ByName("id")

	body, err := io.ReadAll(r.Body)
	if err != nil {
		h.respMsg.SendMsgJson(w, http.StatusBadRequest, "Bad Request", "Incorrect userdata")
		return
	}

	newuser := User{}

	err = json.Unmarshal(body, &newuser)
	if err != nil {
		h.logger.Infoln("cant unmarshal")
		h.respMsg.SendMsgJson(w, http.StatusBadRequest, "Bad Request", "Unmarshal err")
		return
	}

	fmt.Println(newuser)

	err = isAdmin(h.storage, &newuser)

	if err != nil {
		h.logger.Infoln("access denied")
		h.respMsg.SendMsgJson(w, http.StatusUnauthorized, "Unauthorized", "Access denied")
		return
	}

	var id int64

	n, _ := fmt.Sscanf(idStr, "%d", &id)
	if n < 1 {
		h.logger.Infoln("Uncorrect id")
		h.respMsg.SendMsgJson(w, http.StatusBadRequest, "Bad Request", "Uncorrect id")
		return
	}

	err = h.storage.DeleteUser(context.Background(), id)
	if err != nil {
		h.logger.Infoln("Marshall error")
		h.respMsg.SendMsgJson(w, http.StatusNotFound, "Not found", "User not found")
		return
	}
	h.respMsg.SendMsgJson(w, http.StatusOK, "OK", fmt.Sprintf("Successful deleted id: %d", id))
}
