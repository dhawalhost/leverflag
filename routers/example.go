package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterHealthRoutes(route *gin.RouterGroup) {
	route.GET("/health", func(ctx *gin.Context) { ctx.JSON(http.StatusOK, gin.H{"live": "ok"}) })
	// route.GET("/v1/example/", controllers.GetData)
	// route.POST("/v1/example/", controllers.Create)

	//Add All route
	//TestRoutes(route)
}
