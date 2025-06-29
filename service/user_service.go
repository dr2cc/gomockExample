package service

import (
	"github.com/zmey56/gomock/internal"
	"github.com/zmey56/gomock/models"
)

// Как я понимаю логика работы (теста) здесь.

// Структура с полем repo равным интерфейсу UserRepository
// !!! Единственное изменение в чужом коде!!!
// !!! Меняю repo на Repo . Иначе это поле не видно из других модулей!
// !!!
// В Go, если интерфейс используется как тип значения в структуре, это означает,
// что поле структуры может хранить значения любого типа, реализующего данный интерфейс.
// VS Code любезно покажет методы:
type UserService struct {
	Repo internal.UserRepository
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
		Repo: ro,
	}
}

// метод GetUser объекта UserService
// т.к. NewUserService принимает параметром интерфейс NewUserService
// то и метод NewUserService (!!!!)
func (s *UserService) GetUser(id string) (*models.User, error) {
	return s.Repo.GetUserByID(id)
}

// метод DeleteUser объекта UserService
// т.к. NewUserService принимает параметром интерфейс NewUserService
// то и метод NewUserService (!!!!)
func (s *UserService) DeleteUser(id string) error {
	return s.Repo.DeleteUser(id)
}
