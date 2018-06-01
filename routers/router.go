package routers

import (
	"github.com/gin-gonic/gin"
)

// InitRoutes returns an instance of gin router with routes attached to it
func InitRoutes() *gin.Engine {
	router := gin.Default()
	router = SetUtilityRoutes(router)
	return router
}
