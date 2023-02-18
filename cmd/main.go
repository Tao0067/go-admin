package main

import (
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"net/http"
	"paper/config"
	"paper/internal/router"
	"paper/pkg/html"
	"paper/pkg/logger"
	"paper/pkg/models"
)

func init() {
	config.ViperConfig()
	config.Logger = logger.NewLogger(&lumberjack.Logger{
		Filename:  config.ServiceConfig.Logging.FileName,
		MaxSize:   config.ServiceConfig.Logging.MaxSize,
		MaxAge:    config.ServiceConfig.Logging.MaxAge,
		LocalTime: true,
		Compress:  config.ServiceConfig.Logging.Compress,
	}, "", log.LstdFlags).WithCaller(2)

	config.DB, _ = models.NewDb()
}

func main() {
	startWebServer()
}

func startWebServer() {

	r := router.Router()
	r.StaticFS("/static", http.Dir("./web/public/"))
	// 设置模板解析函数
	render := html.LoadTemplates()
	r.HTMLRender = render
	server := &http.Server{
		Addr:    config.ServiceConfig.Server.Host + ":" + config.ServiceConfig.Server.Port,
		Handler: r,
	}
	server.ListenAndServe()
}
