package main

import (
	"net/http"
	"paper/api/v1/router"
	"paper/config"
)

func main() {
	startWebServer()
}

func startWebServer() {
	r := router.Router()
	newConfig := config.ViperConfig()
	server := &http.Server{
		Addr:    newConfig.Server.Host + ":" + newConfig.Server.Port,
		Handler: r,
	}
	server.ListenAndServe()
}
