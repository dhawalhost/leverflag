package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

// HealthHandler handles HTTP requests related to health checks.
type HealthHandler struct {
	db *sqlx.DB
}

// NewHealthHandler creates a new HealthHandler instance.
func NewHealthHandler(db *sqlx.DB) *HealthHandler {
	return &HealthHandler{db: db}
}

// HealthCheck handles GET requests to check the health of the application.
func (h *HealthHandler) HealthCheck(c *gin.Context) {
	// Check database connectivity
	if err := h.db.Ping(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connectivity issue"})
		return
	}

	// Return success response
	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}
