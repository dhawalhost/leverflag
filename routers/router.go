package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func RegisterAllHandlers(router *gin.Engine, db *sqlx.DB) {
	v1 := router.Group("/api/v1")

	RegisterAuthRoutes(v1, db)
	RegisterHealthRoutes(v1)
}
