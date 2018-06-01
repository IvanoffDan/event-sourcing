package routers

import (
	"github.com/IvanoffDan/event-sourcing/controllers"
	"github.com/gin-gonic/gin"
)

// SetUtilityRoutes attaches auth routes to an instance of gin router
func SetUtilityRoutes(router *gin.Engine) *gin.Engine {
	user := router.Group("/utility")
	{
		user.GET("/health", controllers.HealthCheckHandler)
	}

	return router
}
