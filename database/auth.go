package database

import (
	"errors"

	"github.com/dhawalhost/leverflag/models"
)

var (
	ErrIncorrectCreds = errors.New("username/password is incorrect")
)

type AuthService interface {
	Authenticate(username, password string) (models.UserLogin, error)
}

// AuthServiceImpl represents the concrete implementation of AuthService.
type AuthServiceImpl struct {
	userRepository UserRepository
}

// NewAuthServiceImpl creates a new AuthServiceImpl instance.
func NewAuthServiceImpl(userRepository UserRepository) AuthService {
	return &AuthServiceImpl{userRepository: userRepository}
}

// Authenticate authenticates a user with the given username and password.
func (s *AuthServiceImpl) Authenticate(username, password string) (models.UserLogin, error) {
	// Fetch user by username from the repository
	user, err := s.userRepository.GetUserLogin(username)
	if err != nil {
		return models.UserLogin{}, err
	}
	if user.Username != username {
		return models.UserLogin{}, ErrIncorrectCreds
	}
	if user.Password != password {
		return models.UserLogin{}, ErrIncorrectCreds
	}

	return user, nil
}
