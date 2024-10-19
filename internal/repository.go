package internal

import "github.com/zmey56/gomock/models"

type UserRepository interface {
	GetUserByID(id string) (*models.User, error)
	DeleteUser(id string) error
}
