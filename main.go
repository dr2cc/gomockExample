// Попробую использовать уже готовый фунционал
package main

import (
	"fmt"

	"github.com/zmey56/gomock/models"
)

type UserStorage struct {
	Data map[string]string
}

// UserStorage должна имплементировать эти методы!
// type UserRepository interface {
// 	GetUserByID(id string) (*models.User, error)
// 	DeleteUser(id string) error
// }

func NewStorage() *UserStorage {
	return &UserStorage{
		Data: make(map[string]string),
	}
}

func main() {

	storageInstance := NewStorage()

	fmt.Println(storageInstance.GetUserByID("1"))

}

// Первый метод UserRepository interface есть!!
func (s *UserStorage) GetUserByID(id string) (*models.User, error) {

	person := models.User{
		ID:   id,
		Name: "Pit",
	}

	// var ro internal.UserRepository
	// // // Ошибки нет! Осталось присвоить правильную переменную!
	// //ro = person
	// service.NewUserService(ro).DeleteUser("2")

	return &person, nil
}

// // Вид UserRepository interface уже готов!!
func (s *UserStorage) DeleteUser(id string) error {
	// e, exists := s.Data[id]
	// if !exists {
	// 	return id, errors.New("URL with such id doesn't exist")
	// }
	return nil
}
