// Попробую использовать уже готовый фунционал
package main

import (
	"fmt"

	"github.com/zmey56/gomock/models"
	"github.com/zmey56/gomock/service"
)

// Цель: создать тип данных реализующий интерфейс internal.UserRepository .
// Сделал. Методы VS Code перечислит ниже:
type employee struct {
	profession string
	division   string
}

// конструктор
func newEmployee(p string, d string) employee {
	return employee{
		profession: p,
		division:   d,
	}
}

func main() {
	// Может потом пойму почему в метод с одним параметром,
	// нужно отправлять два!?
	//user, _ := resident.GetUserByID(newResident("2", "Kid"), "2")
	//fmt.Println(*user)

	// Здесь главное!
	// Обращаюсь к конструктору service.NewUserService, принимающему параметром
	// интерфейс internal.UserRepository .
	//
	// В Go передача интерфейса параметром в функцию означает,
	// что функция может принимать на вход объект любого типа,
	// который реализует этот интерфейс.
	//
	// Таковым объектом является мой объект resident
	//
	// Т.к. он должен в точности реализовывать методы интерфейса (заданные не мной!),
	// то метод GetUserByID возвращает тип *models.User
	//
	// а уже его я привожу- resident(*user)
	//
	//service.NewUserService(resident(*user))
	//
	// Получаю тип main.employee
	fmt.Printf("%T\n", newEmployee("clerk", "accounting"))
	user := *service.NewUserService(newEmployee("manager", "sales"))
	// Получаю тип service.UserService
	fmt.Printf("%T\n", user)
	//
	// обращаюсь к service.GetUser
	fmt.Println(user.GetUser("3"))
	// обращаюсь к service.DeleteUser
	fmt.Println(user.DeleteUser("3"))

	// Пока (13.06.25) я получаю
	// &{{2 Kid}}
	// Но это не так важно!
	//
	// Главное- я получил данные из стороннего кода, не нарушая его.
	// Теперь тесты с моками должны быть более понятны!
	//
	// get, _ := resident.GetUserByID(resident(*&user.DeleteUser), "2")
	// fmt.Println("найдена запись:", *get)
	// // Включение или нет, вызова DeleteUser сути дела не меняет.
	// // Главное- мой объект resident здесь реализует второй метод
	// // интерфейса internal.UserRepository
	// resident.DeleteUser(resident(*user), "2")

}

// Первый метод интерфейса internal.UserRepository , нахожу запись
func (s employee) GetUserByID(id string) (*models.User, error) {
	fmt.Println("Я GetUserByID , вызываюсь из service.GetUser и не понимаю зачем тут ID:", id)

	fmt.Println(s)

	person := models.User{
		ID:   s.profession,
		Name: s.division,
	}

	return &person, nil
	//return models.User(*s),nil
}

// Второй метод интерфейса internal.UserRepository , стираю запись
func (s employee) DeleteUser(id string) error {
	// Стираю данные в структуре
	fmt.Println("Я DeleteUser, вызываюсь из service.DeleteUser и не понимаю зачем тут ID", id, "и удаляю", s)
	s.profession = ""
	s.division = ""
	fmt.Println(s)

	return nil
}
