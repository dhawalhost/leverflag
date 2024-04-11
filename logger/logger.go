// logger/logger.go
package logger

import (
	"go.uber.org/zap"
)

var (
	Logger *zap.Logger
)

// InitLogger initializes the logger
func InitLogger() {
	logger, _ := zap.NewProduction()
	Logger = logger
}
