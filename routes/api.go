package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/uzzalhcse/GoSpire/app/http/controllers"
)

func SetupApiRoutes(router *gin.RouterGroup) {
	var TestController = controllers.NewTestController()

	testRoutes := router.Group("/test")
	{
		testRoutes.GET("/", TestController.Test)
		testRoutes.POST("/upload", TestController.Upload)
	}
	router.GET("/posts", TestController.GetPosts)
	router.GET("/posts/insert", TestController.BulkPostInsert)
}
