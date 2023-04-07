package main

import (
	"app/api"
	"app/config"
	"app/pkg/logger"
	"app/storage/postgres"
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()

	var loggerLevel = new(string)

	*loggerLevel = logger.LevelDebug

	switch cfg.Environment {
	case config.DebugMode:
		*loggerLevel = logger.LevelDebug
		gin.SetMode(gin.DebugMode)
	case config.TestMode:
		*loggerLevel = logger.LevelDebug
		gin.SetMode(gin.TestMode)
	default:
		*loggerLevel = logger.LevelInfo
		gin.SetMode(gin.ReleaseMode)
	}

	log := logger.NewLogger("app", *loggerLevel)
	defer func() {
		err := logger.Cleanup(log)
		if err != nil {
			return
		}
	}()

	store, err := postgres.NewPostgres(context.Background(), cfg)
	if err != nil {
		log.Panic("Error connect to postgresql: ", logger.Error(err))
		return
	}
	defer store.CloseDB()

	r := gin.New()

	r.Use(gin.Recovery(), gin.Logger())

	api.SetUpApi(r, &cfg, store, log)

	fmt.Println("Listening Server", cfg.ServerHost+cfg.ServerPort)
	err = r.Run(cfg.ServerHost + cfg.ServerPort)
	if err != nil {
		log.Panic("Error listening server:", logger.Error(err))
		return
	}
}
