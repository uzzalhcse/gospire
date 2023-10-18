package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/uzzalhcse/GoSpire/app/http/responses"
	"net/http"
	"runtime"
)

func SetupWebRoutes(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"ginBoilerplateVersion": "v0.01",
			"goVersion":             runtime.Version(),
		})
	})
	router.NoRoute(func(c *gin.Context) {
		responses.Error(c, "Route Not Found", nil)
	})

}
