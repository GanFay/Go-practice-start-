package main

import (
	"fmt"
	"strings"
)

type golan struct {
	name    string
	dollars float64
	euro    float64
}

func main() {

	Users := map[string]golan{}

	for {
		fmt.Println("Какую операцию вы хотите зделать, CDE(Convert dollars to euro), CED(Convert euro to dollars), Dep, exit, add, list")
		var nig string
		fmt.Scanln(&nig)
		nig = strings.ToLower(nig)

		switch nig {
		case "add":
			addUser(Users)
		case "cde":
			{
				CDE(Users)
			}
		case "ced":
			{
				CED(Users)
			}
		case "dep":
			{
				DEP(Users)
			}
		case "exit":
			{
				return
			}
		case "list":
			{
				listusers(Users)
			}
		default:
			{
				fmt.Println("Пиши нормально")
			}
		}
	}
}

func listusers(Users map[string]golan) {
	if len(Users) == 0 {
		fmt.Println("Нет пользователей")
		return
	}
	for _, user := range Users {
		fmt.Printf("%s: $%.2f, €%.2f\n", user.name, user.dollars, user.euro)
	}
}

func DEP(Users map[string]golan) {
	name := ""
	fmt.Println("Имя пользователя: ")
	fmt.Scanln(&name)
	user, ok := Users[name]
	if !ok {
		fmt.Println("Пользователь не найден")
		return
	}
	fmt.Println("dollars or euro")
	ds := ""
	_, err := fmt.Scanln(&ds)
	if err != nil {
		return
	}
	if ds == "dollars" {
		fmt.Println("Сколько деп?")
		depd := 0.0
		_, err2 := fmt.Scanln(&depd)
		if err2 != nil {
			return
		}
		user.dollars += depd
		fmt.Println("Депнуто:", depd)
	} else if ds == "euro" {
		fmt.Println("Сколько деп?")
		depe := 0.0
		_, err1 := fmt.Scanln(&depe)
		if err1 != nil {
			return
		}
		user.euro += depe
		fmt.Println("Депнуто:", depe)
	}
	Users[name] = user
	fmt.Printf("Ваш баланс: $%.2f, €%.2f\n", user.dollars, user.euro)
}

func addUser(Users map[string]golan) {
	var name string
	fmt.Println("Введите имя:")
	fmt.Scanln(&name)
	if _, ok := Users[name]; ok {
		fmt.Println("Пользователь уже существует")
		return
	}
	Users[name] = golan{name, 0, 0}
	fmt.Println("Пользователь добавлен:", name)
	fmt.Println(Users[name])
}

func CDE(Users map[string]golan) {
	name := ""
	fmt.Print("Имя пользователя: ")
	fmt.Scanln(&name)
	user, ok := Users[name]
	if !ok {
		fmt.Println("Пользователь не найден")
		return
	}
	fmt.Println("У вас долларов: ", user.dollars, "У вас евро:", user.euro)
	fmt.Println("Сколько долларов переводим в евро")
	dollars := 0.0
	fmt.Scanln(&dollars)
	if dollars > user.dollars {
		fmt.Println("Недостаточно средств")
		return
	}
	rate := 0.854
	euro := rate * dollars
	user.euro += euro
	user.dollars -= dollars
	Users[name] = user
	fmt.Println("У вас евро: ", user.dollars, "У вас долларов:", user.euro)
}

func CED(Users map[string]golan) {
	name := ""
	fmt.Print("Имя пользователя: ")
	_, err := fmt.Scanln(&name)
	if err != nil {
		return
	}
	user, ok := Users[name]
	if !ok {
		fmt.Println("Пользователь не найден")
		return
	}
	fmt.Println("У вас евро: ", user.dollars, "У вас долларов:", user.euro)
	fmt.Println("Сколько евро переводим в доллари")
	euro := 0.0
	_, err1 := fmt.Scanln(&euro)
	if err1 != nil {
		return
	}
	if euro > user.euro {
		fmt.Println("Недостаточно средств")
		return
	}
	rate := 1.17
	dollars := rate * euro
	user.dollars += dollars
	user.euro -= euro
	Users[name] = user
	fmt.Println("У вас евро: ", user.dollars, "У вас долларов:", user.euro)
}
