package service

import (
	"github.com/zmey56/gomock/internal"
	"github.com/zmey56/gomock/models"
)

// Как я понимаю логика работы (теста) здесь.

// Структура с полем repo равным интерфейсу UserRepository (получается так можно!)
// Она определяет возможные действия ("сервисы") с пользователем ?!
// VS Code любезно покажет методы:
type UserService struct {
	repo internal.UserRepository
}

// Конструктор объектов UserService (при запросе создает и возвращает объект UserService,
// со значением поля repo переданного сюда в параметре ro)
//
// В Go передача интерфейса параметром в функцию означает,
// что функция может принимать на вход объект любого типа,
// который реализует определенный интерфейс.
func NewUserService(ro internal.UserRepository) *UserService {
	// // было записано так (это простая запись структуры, если у нее одно поле):
	// return &UserService{repo: ro}

	// напишу как привык:
	return &UserService{
		repo: ro,
	}
}

// метод GetUser объекта UserService
// т.к. NewUserService принимает параметром интерфейс NewUserService
// то и метод NewUserService (!!!!)
func (s *UserService) GetUser(id string) (*models.User, error) {
	return s.repo.GetUserByID(id)
}

// метод DeleteUser объекта UserService
// т.к. NewUserService принимает параметром интерфейс NewUserService
// то и метод NewUserService (!!!!)
func (s *UserService) DeleteUser(id string) error {
	return s.repo.DeleteUser(id)
}
