package models

import (
	"regexp"
	"time"
)

type UserLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
	ID       int    `json:"id"`
}

type User struct {
	Username    string    `json:"username"`
	FirstName   string    `json:"firstname"`
	LastName    string    `json:"lastname"`
	DisplayName string    `json:"displayName"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	ID          int       `json:"id"`
}

func (u *UserLogin) IsValid() bool {
	if len(u.Username) == 0 {
		return false
	}
	if IsValidPassword(u.Password) {
		return false
	}
	return true
}

func IsValidPassword(password string) bool {
	if len(password) < 8 {
		return false
	}
	hasUC := false
	hasLC := false
	hasSC := false

	for _, char := range password {
		if char >= 'A' && char <= 'Z' {
			hasUC = true
			continue
		}
		if char >= 'a' && char <= 'z' {
			hasLC = true
			continue

		}
		if IsSpecialChar(string(char)) {
			hasSC = true
		}
	}

	return hasUC && hasLC && hasSC
}

func IsSpecialChar(str string) bool {
	pattern := `^[&!@#$\(\)]$`
	regex := regexp.MustCompile(pattern)
	return regex.MatchString(str)
}
