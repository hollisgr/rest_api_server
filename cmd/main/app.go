package main

import (
	"api_server/internal/config"
	"api_server/internal/user"
	"api_server/pkg/logging"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"time"

	"github.com/julienschmidt/httprouter"
)

func StartServer(router *httprouter.Router, cfg *config.Config) {
	logger := logging.GetLogger()

	logger.Infoln("Starting server")

	var listener net.Listener
	var listenErr error

	if cfg.Listen.Type == "sock" {
		appDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			logger.Fatal(err)
		}
		logger.Infoln("Creating socket")
		socketPath := path.Join(appDir, "app.sock")
		logger.Debugln("socket path: ", socketPath)

		logger.Infoln("listen unix socket")
		listener, listenErr = net.Listen("unix", socketPath)
		logger.Infoln("Server is listening unix socket:", socketPath)
	} else {
		logger.Infoln("listen tcp")
		listener, listenErr = net.Listen("tcp", fmt.Sprintf("%s:%s", cfg.Listen.BindIP, cfg.Listen.Port))
		logger.Infof("Server is listening %s:%s", cfg.Listen.BindIP, cfg.Listen.Port)
	}

	if listenErr != nil {
		log.Fatal(listenErr)
		return
	}

	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	server.Serve(listener)
}

func main() {
	logger := logging.GetLogger()

	logger.Infoln("Creating router")
	router := httprouter.New()

	cfg := config.GetConfig()

	logger.Infoln("Register user handler")
	handler := user.NewHandler(logger)
	handler.Register(router)

	StartServer(router, cfg)
}
