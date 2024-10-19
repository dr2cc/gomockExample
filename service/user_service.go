package service

import (
	"github.com/zmey56/gomock/internal"
	"github.com/zmey56/gomock/models"
)

type UserService struct {
	repo internal.UserRepository
}

func NewUserService(repo internal.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetUser(id string) (*models.User, error) {
	return s.repo.GetUserByID(id)
}

func (s *UserService) DeleteUser(id string) error {
	return s.repo.DeleteUser(id)
}
