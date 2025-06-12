package service

import (
	"github.com/zmey56/gomock/internal"
	"github.com/zmey56/gomock/models"
)

// Как я понимаю логика работы (теста) здесь.

// структура с полем repo равным интерфейсу UserRepository (получается так можно)
// Получается, что она определяет возможные действия ("сервисы") с пользователем ?!
type UserService struct {
	repo internal.UserRepository
}

// при запросе создает и возвращает объект UserService
// со значением поля repo переданного сюда в параметре ro
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
