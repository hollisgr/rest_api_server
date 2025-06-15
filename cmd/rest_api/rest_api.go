package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"rest_api/internal/cfg"
	"rest_api/internal/db"
	"rest_api/internal/db/postgres"
	"rest_api/internal/handler"
	"rest_api/internal/service/user_service"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func main() {

	cfg := cfg.GetConfig()

	pool, err := postgres.NewClient(context.Background(), 3, *cfg)

	validate := validator.New(validator.WithRequiredStructEnabled())

	if err != nil {
		log.Fatalln(err)
	}
	log.Println("database connected on host:", cfg.Postgresql.Host, "port:", cfg.Postgresql.Port)

	storage := db.NewStorage(pool)
	err = pool.Ping(context.Background())
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("database pinged OK")

	userRepo := user_service.NewUserService(storage, validate)

	r := gin.New()
	handler := handler.NewHandler(userRepo)
	handler.Register(r)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigChan
		log.Println("Interrupt signal received. Exiting...")
		pool.Close()
		os.Exit(0)
	}()

	ListenAddr := cfg.Server.BindIP + ":" + cfg.Server.ListenPort
	r.Run(ListenAddr)
}
