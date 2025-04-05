package user

import (
	"api_server/internal/handlers"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

const (
	usersURL = "/users"
	userURL  = "/users/:uuid"
)

type handler struct {
}

func NewHandler() handlers.Handler {
	return &handler{}
}

func (h *handler) Register(router *httprouter.Router) {
	router.GET(usersURL, h.GetList)
	router.GET(userURL, h.GetUserByUUID)
	router.POST(userURL, h.CreateUser)
	router.DELETE(userURL, h.DeleteUser)
}

func (h *handler) GetList(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte(fmt.Sprintln("THIS IS LIST OF USERS")))
}

func (h *handler) GetUserByUUID(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte(fmt.Sprintln("THIS IS GET USER BY ID")))

}
func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte(fmt.Sprintln("THIS IS CREATE USER")))

}
func (h *handler) DeleteUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte(fmt.Sprintln("THIS IS DELETE USER BY ID")))

}
