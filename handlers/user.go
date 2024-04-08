package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/dhawalhost/leverflag/database"
	"github.com/dhawalhost/leverflag/models"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userRepository *database.UserRepository
}

// NewUserHandler creates a new UserHandler instance with the given database connection.
func NewUserHandler(userRepository *database.UserRepository) *UserHandler {
	return &UserHandler{userRepository: userRepository}
}

func (h *UserHandler) CreateUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.BindJSON(&user); err != nil {
		Response(ctx, http.StatusBadRequest, nil, err)
		return
	}
	insertedId, err := h.userRepository.CreateUser(user)
	if err != nil {
		Response(ctx, http.StatusInternalServerError, nil, err)
		return
	}
	user, err = h.userRepository.GetUserByID(int(insertedId))
	if err != nil {
		Response(ctx, http.StatusBadRequest, nil, err)

	}
	Response(ctx, http.StatusCreated, user, nil)

}

// GetUser handles GET requests to retrieve a user by ID.
func (h *UserHandler) GetUser(ctx *gin.Context) {

	userIdStr := ctx.Param("id")
	userID, err := strconv.Atoi(userIdStr) // Example user ID for demonstration
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"errors": ""})
	}
	user, err := h.userRepository.GetUserByID(userID)
	if err != nil {
		Response(ctx, http.StatusNotFound, nil, err)
		return
	}
	if user.Username == "" {
		Response(ctx, http.StatusNotFound, nil, errors.New("no user found"))

	}
	Response(ctx, http.StatusOK, user, nil)

}
