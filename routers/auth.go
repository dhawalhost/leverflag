package routers

import (
	"github.com/dhawalhost/leverflag/database"
	"github.com/dhawalhost/leverflag/handlers"
	"github.com/jmoiron/sqlx"

	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(route *gin.RouterGroup, db *sqlx.DB) {
	auth := route.Group("/auth")
	userRepository := database.NewUserRepository(db)
	authService := database.NewAuthServiceImpl(*userRepository)
	authHandler := handlers.NewAuthHandler(authService)
	{
		auth.POST("/login", authHandler.AuthenticateUser)
		// auth.GET("/me", handlers.AuthorizeJWT(), authHandler.)
	}
}
