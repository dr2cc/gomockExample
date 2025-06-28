// Попробую использовать уже готовый фунционал
package main

import (
	"fmt"

	"github.com/zmey56/gomock/models"
	"github.com/zmey56/gomock/service"
)

var user1 service.UserService

func main() {
	user1 = *service.NewUserService(newEmployee("manager", "sales"))

	// // Что за тип данных в user1 - service.UserService !
	// fmt.Printf("%T\n", user1)

	// Создается объект типа service.UserService
	// Теперь я обращаюсь к его (service.UserService) методам.
	// А каждый из них уже обращается
	// к методам типа данных (employee) реализующего интерфейс internal.UserRepository
	// Что я в них буду делать не регламентируется,
	// регламентируются только параметры и возвращаемые значения.

	// обращаюсь к service.GetUser
	user1.GetUser("manager")
	// обращаюсь к service.DeleteUser
	user1.DeleteUser("manager")
	// Что теперь в user1 :
	fmt.Println(user1)
}

// Тип данных реализующий интерфейс internal.UserRepository .
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

// реализую этот интерфейс
// type UserRepository interface {
// 	GetUserByID(id string) (*models.User, error)
// 	DeleteUser(id string) error
// }

// Первый метод типа employee , реализующий интерфейс internal.UserRepository , нахожу запись
func (s employee) GetUserByID(id string) (*models.User, error) {
	if id == s.profession {
		fmt.Println("Эта профессия относится к отделу:", s.division)
		person := models.User{
			ID:   s.profession,
			Name: s.division,
		}
		return &person, nil
	} else {
		fmt.Println("Не найдена профессия", id)
		person := models.User{
			ID:   "",
			Name: "",
		}
		return &person, nil
	}
}

// Второй метод типа employee , реализующий интерфейс internal.UserRepository , стираю запись
func (s employee) DeleteUser(id string) error {
	//var v service.UserService = user1
	// // запрос такой "go как обратится к полю структуры являющейся интерфейсом"
	if id == s.profession {
		fmt.Println("Удаляю запись о типе работников:", s.division)
		// Вот здесь написать правильную работу с указателем.
		// Цель- очищать первоначальную структуру, т.е user1
		//user1
		v := &s
		v.profession = ""
		v.division = ""
		fmt.Println(*v)
		return nil
	} else {
		fmt.Println("Не найдено соответствие для удаления типа работников", id, "\nтекущий тип работников", s.profession)
		return nil
	}
}

// // main старый, может комменты тут полезные?
// func main() {
// 	// Может потом пойму почему в метод с одним параметром,
// 	// нужно отправлять два!?
// 	//user, _ := resident.GetUserByID(newResident("2", "Kid"), "2")
// 	//fmt.Println(*user)

// 	// Здесь главное!
// 	// Обращаюсь к конструктору service.NewUserService, принимающему параметром
// 	// интерфейс internal.UserRepository .
// 	//
// 	// В Go передача интерфейса параметром в функцию означает,
// 	// что функция может принимать на вход объект любого типа,
// 	// который реализует этот интерфейс.
// 	//
// 	// Таковым объектом является мой объект resident
// 	//
// 	// Т.к. он должен в точности реализовывать методы интерфейса (заданные не мной!),
// 	// то метод GetUserByID возвращает тип *models.User
// 	//
// 	// а уже его я привожу- resident(*user)
// 	//
// 	//service.NewUserService(resident(*user))
// 	//
// 	// Получаю тип main.employee
// 	fmt.Printf("%T\n", newEmployee("clerk", "accounting"))
// 	user := *service.NewUserService(newEmployee("manager", "sales"))
// 	// Получаю тип service.UserService
// 	fmt.Printf("%T\n", user)
// 	//
// 	// обращаюсь к service.GetUser
// 	fmt.Println(user.GetUser("3"))
// 	// обращаюсь к service.DeleteUser
// 	fmt.Println(user.DeleteUser("3"))

// 	// Пока (13.06.25) я получаю
// 	// &{{2 Kid}}
// 	// Но это не так важно!
// 	//
// 	// Главное- я получил данные из стороннего кода, не нарушая его.
// 	// Теперь тесты с моками должны быть более понятны!
// 	//
// 	// get, _ := resident.GetUserByID(resident(*&user.DeleteUser), "2")
// 	// fmt.Println("найдена запись:", *get)
// 	// // Включение или нет, вызова DeleteUser сути дела не меняет.
// 	// // Главное- мой объект resident здесь реализует второй метод
// 	// // интерфейса internal.UserRepository
// 	// resident.DeleteUser(resident(*user), "2")

// }
