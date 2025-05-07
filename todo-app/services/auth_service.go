package services

import (
	"errors"
	"todo-app/models"
)

var users = map[string]models.User{
	"user1": {
		ID:       1,
		Username: "user1",
		Password: "password1",
		Type:     1, // Normal user
	},
	"admin": {
		ID:       2,
		Username: "admin",
		Password: "admin123",
		Type:     2, // Admin user
	},
}

func AuthenticateUser(username, password string) (*models.User, error) {
	user, exists := users[username]
	if !exists || user.Password != password {
		return nil, errors.New("invalid credentials")
	}
	return &user, nil
}
