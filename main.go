// main.go
package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/dhawalhost/leverflag/config"
	"github.com/dhawalhost/leverflag/database"
	"github.com/dhawalhost/leverflag/logger"
	"github.com/dhawalhost/leverflag/routers"

	"go.uber.org/zap"
)

func main() {
	// Initialize logger
	logger.InitLogger()
	defer logger.Logger.Sync()
	db, err := database.InitDB("memory")
	if err != nil {
		logger.Logger.Fatal("Failed to initialize DB, ", zap.Error(err))
	}
	// Initialize Gin router
	router := config.ServerConfig()
	routers.RegisterAllHandlers(router, db)

	// Start server
	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Logger.Fatal("Server startup failed", zap.Error(err))
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	logger.Logger.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		logger.Logger.Fatal("Server shutdown failed", zap.Error(err))
	}
	logger.Logger.Info("Server exiting")
}
