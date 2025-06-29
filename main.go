// Попробую использовать уже готовый фунционал
package main

import (
	"fmt"

	"github.com/zmey56/gomock/models"
	"github.com/zmey56/gomock/service"
)

var user1 service.UserService

// Главное- я получил данные из стороннего кода, почти не нарушая его.
// !!! Единственное исключение- в структуре service.UserService
// !!! изменил название поля repo на Repo иначе оно не видно здесь.
// Теперь тесты с моками должны быть более понятны!

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

	// Суть: в строке пишу искомую профессию, если она соответствует той, что передана выше
	// то выводится сообщение о ее наличии, а в другом методе структура с этой профессией очищается.
	searchBar := "driver"
	// если заменим на "manager" то все найдется и заменится!

	// обращаюсь к service.GetUser
	user1.GetUser(searchBar)
	// обращаюсь к service.DeleteUser
	user1.DeleteUser(searchBar)
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
	if id == s.profession {
		fmt.Println("Удаляю запись о типе работников:", s.division)
		// Нужно очищать первоначальную структуру, т.е user1
		s.profession = ""
		s.division = ""
		user1.Repo = s
		return nil
	} else {
		fmt.Println("Не найдено соответствие для удаления типа работников", id, "\nтекущий тип работников", s.profession)
		return nil
	}
}
