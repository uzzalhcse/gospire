package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/uzzalhcse/GoSpire/app/http/controllers"
)

func SetupAuthRoutes(router *gin.RouterGroup) {
	var AuthController = controllers.NewAuthController()

	authRoutes := router.Group("/auth")
	{
		authRoutes.POST("/login", AuthController.Login)
		authRoutes.POST("/register", AuthController.Register)
	}
}
