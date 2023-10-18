package server

import (
	"github.com/gin-gonic/gin"
	"github.com/uzzalhcse/GoSpire/app/exceptions"
	"github.com/uzzalhcse/GoSpire/app/providers"
	"github.com/uzzalhcse/GoSpire/config"
	"log"
)

func RunServer() {
	if !config.App.Debug {
		gin.SetMode(gin.ReleaseMode)
	}
	server := gin.Default()
	router := providers.RouteServiceProvider{}.Register(server)
	err := router.Run(":" + config.App.Port)
	exceptions.PanicIfNeeded(err)
	log.Println("Server is Running on port " + config.App.Port)
}
