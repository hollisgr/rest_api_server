package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"rest_api_server/internal/config"
	"rest_api_server/internal/user"
	"rest_api_server/internal/user/db"
	"rest_api_server/pkg/client/mongodb"
	"rest_api_server/pkg/logging"
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

	cfgMongo := cfg.MongoDB

	mongoDBClient, err := mongodb.NewClient(context.Background(), cfgMongo.Host, cfgMongo.Port,
		cfgMongo.Username, cfgMongo.Password, cfgMongo.Database, cfgMongo.Auth_db)

	if err != nil {
		panic(err)
	}

	storage := db.NewStorage(mongoDBClient, cfgMongo.Collection, logger)

	logger.Infoln("Register user handler")
	handler := user.NewHandler(logger, storage)
	handler.Register(router)

	StartServer(router, cfg)
}
