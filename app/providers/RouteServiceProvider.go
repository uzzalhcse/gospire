package providers

import (
	"github.com/gin-gonic/gin"
	"github.com/uzzalhcse/GoSpire/app/http/middleware"
	"github.com/uzzalhcse/GoSpire/routes"
	"github.com/uzzalhcse/go-platform-detection"
)

type RouteServiceProvider struct{}

func (r RouteServiceProvider) Register(router *gin.Engine) *gin.Engine {
	router.Use(middleware.Cors)
	router.Use(platform.ResolveDevice())
	v1 := router.Group("/v1")
	{
		routes.SetupApiRoutes(v1)
		routes.SetupAuthRoutes(v1)
	}

	routes.SetupWebRoutes(router)

	return router
}
