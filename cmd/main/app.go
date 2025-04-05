package main

import (
	"api_server/internal/user"
	"api_server/pkg/logging"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

func StartServer(router *httprouter.Router) {
	logger := logging.GetLogger()

	logger.Infoln("Starting server")

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal()
		return
	}

	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	logger.Infoln("Server is listening port 8080")
	server.Serve(listener)

}

func main() {
	logger := logging.GetLogger()

	logger.Infoln("Creating router")
	router := httprouter.New()

	logger.Infoln("Register user handler")
	handler := user.NewHandler()
	handler.Register(router)

	StartServer(router)
}
