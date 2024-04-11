package config

import (
	"github.com/dhawalhost/leverflag/middleware"
	"github.com/gin-gonic/gin"
)

func ServerConfig() *gin.Engine {
	// gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	// Use LoggerMiddleware
	router.Use(middleware.LoggerMiddleware())
	return router
}
