package main

import (
	"api_server/internal/user"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

func StartServer(router *httprouter.Router) {

	log.Println("Starting server")

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

	log.Println("Server is listening port 8080")
	server.Serve(listener)

}

func main() {
	log.Println("Creating router")
	router := httprouter.New()

	log.Println("Register user handler")
	handler := user.NewHandler()
	handler.Register(router)

	StartServer(router)
}
