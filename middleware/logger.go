// middleware/logger.go
package middleware

import (
	"time"

	"github.com/dhawalhost/leverflag/logger"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// LoggerMiddleware is a middleware function to log incoming requests
func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		end := time.Now()
		logger.Logger.Info("Request",
			zap.String("path", c.Request.URL.Path),
			zap.String("method", c.Request.Method),
			zap.Int("status", c.Writer.Status()),
			zap.Duration("latency", end.Sub(start)),
		)
	}
}
