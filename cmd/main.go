package main

import (
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"net/http"
	"paper/api/v1/router"
	"paper/config"
	"paper/pkg/logger"
	"paper/pkg/models"
)

func init() {
	config.ViperConfig()
	config.DB, _ = models.NewDb()
	config.Logger = logger.NewLogger(&lumberjack.Logger{
		Filename:  config.ServiceConfig.Logging.FileName,
		MaxSize:   config.ServiceConfig.Logging.MaxSize,
		MaxAge:    config.ServiceConfig.Logging.MaxAge,
		LocalTime: true,
		Compress:  config.ServiceConfig.Logging.Compress,
	}, "", log.LstdFlags).WithCaller(2)
}

func main() {
	startWebServer()
}

func startWebServer() {
	config.Logger.Infof("%s: go-programming-tour-book/%s", "eddycjy", "blog-service")
	r := router.Router()
	server := &http.Server{
		Addr:    config.ServiceConfig.Server.Host + ":" + config.ServiceConfig.Server.Port,
		Handler: r,
	}
	server.ListenAndServe()
}
