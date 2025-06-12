// Попробую использовать уже готовый фунционал
package main

import (
	"fmt"

	"github.com/zmey56/gomock/models"
	"github.com/zmey56/gomock/service"
)

type UserStorage struct {
	Data map[string]string
}

func NewStorage() *UserStorage {
	return &UserStorage{
		Data: make(map[string]string),
	}
}

func main() {
	person := models.User{
		ID:   "2",
		Name: "Pit",
	}

	storageInstance := NewStorage()

	storageInstance.CreateUser(person.ID, person.Name)

}

func (s *UserStorage) CreateUser(id string, url string) {
	s.Data[id] = url

	//попробовать здесь
	// ctrl := gomock.NewController(t)
	// defer ctrl.Finish()

	// mockRepo := mocks.NewMockUserRepository(ctrl)
	//надо заменить mockRepo

	service.NewUserService(mockRepo).DeleteUser("2")

	fmt.Println("User", s.Data[id], ", id", id)
}

// func (s *UserStorage) GetURL(id string) (string, error) {
// 	e, exists := s.Data[id]
// 	if !exists {
// 		return id, errors.New("URL with such id doesn't exist")
// 	}
// 	return e, nil
// }
