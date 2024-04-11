package database

import (
	"github.com/dhawalhost/leverflag/models"
	"github.com/jmoiron/sqlx"
)

const (
	StatusDisabled = iota
	StatusPending
	StatusEnabled
)

// UserRepository represents the repository for user data.
type UserRepository struct {
	db *sqlx.DB
}

// NewUserRepository creates a new UserRepository instance.
func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db: db}
}

// CreateUser inserts a new user record into the database.
func (r *UserRepository) CreateUser(user models.User) (int64, error) {
	res, err := r.db.Exec("INSERT INTO users (username, email, first_name, last_name, display_name, status, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);", user.Username, user.Email, user.FirstName, user.LastName, user.FirstName+" "+user.LastName, StatusEnabled)
	if err != nil {
		return 0, err
	}
	id, _ := res.LastInsertId()
	return id, err
}

// GetUserByID retrieves a user record from the database by ID.
func (r *UserRepository) GetUserByID(userID int) (models.User, error) {
	var user models.User
	err := r.db.Get(&user, "SELECT id, username, email, first_name, last_name, display_name, status, created_at, updated_at FROM users WHERE id = ?", userID)
	return user, err
}

// GetUserLogin retrieves a user login record from the database by username.
func (r *UserRepository) GetUserLogin(username string) (models.UserLogin, error) {
	var user models.UserLogin
	err := r.db.Get(&user, "SELECT id, username, email, password FROM userlogin WHERE id = ?", username)
	return user, err
}
