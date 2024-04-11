package handlers

import (
	"github.com/gin-gonic/gin"
)

func Response(c *gin.Context, statusCode int, data interface{}, err error) {
	if err != nil {
		c.JSON(statusCode, gin.H{"error": err.Error()})
		return
	}
	c.JSON(statusCode, data)
}
