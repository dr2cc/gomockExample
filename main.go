// Попробую использовать уже готовый фунционал
package main

import (
	"fmt"

	"github.com/zmey56/gomock/models"
	"github.com/zmey56/gomock/service"
)

// Цель: создать тип данных реализующий интерфейс internal.UserRepository .
// Сделал. Методы VS Code перечислит ниже:
type resident struct {
	ID   string
	Name string
}

// конструктор
func newResident(id string, name string) resident {
	return resident{
		ID:   id,
		Name: name,
	}
}

func main() {
	// Может потом пойму почему в метод с одним параметром,
	// нужно отправлять два!?
	user, _ := resident.GetUserByID(newResident("2", "Kid"), "2")
	fmt.Println(*user)

	// Здесь главное!
	// Обращаюсь к конструктору service.NewUserService, принимающему параметром
	// интерфейс internal.UserRepository .
	//
	// В Go передача интерфейса параметром в функцию означает,
	// что функция может принимать на вход объект любого типа,
	// который реализует определенный интерфейс.
	//
	// Таковым объектом является мой объект resident .
	// Т.к. он должен в точности реализовывать методы интерфейса (заданные не мной!),
	// то метод GetUserByID возвращает тип *models.User ,
	// а уже его я привожу- resident(*user)
	hz := service.NewUserService(resident(*user))
	fmt.Println(hz)
	// Пока (13.06.25) я получаю
	// &{{2 Kid}}
	// Но это не так важно!
	// Главное- я получил данные из стороннего кода, не нарушая его.
	// Теперь тесты с моками должны быть более понятны!

	// Включение или нет, вызова DeleteUser сути дела не меняет.
	// Главное- мой объект resident здесь реализует второй метод
	// интерфейса internal.UserRepository
	resident.DeleteUser(resident(*user), "2")

}

// Первый метод интерфейса internal.UserRepository
func (s resident) GetUserByID(id string) (*models.User, error) {

	person := models.User{
		ID:   s.ID,
		Name: s.Name,
	}

	return &person, nil
}

// Второй метод интерфейса internal.UserRepository
func (s resident) DeleteUser(id string) error {
	// Стираю данные в структуре
	s.ID = ""
	s.Name = ""
	fmt.Println(s)

	return nil
}
