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

const rateDolToEur = 0.854
const rateEurToDol = 1.17

func main() {

	Users := map[string]*golan{}

	for {
		fmt.Println("Какую операцию вы хотите зделать: convert, dep, exit, add, list")
		var nig string
		_, err := fmt.Scanln(&nig)
		if err != nil {
			return
		}
		nig = strings.ToLower(nig)

		switch nig {
		case "add":
			addUser(Users)
		case "convert":
			{
				CDE(Users)
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

func listusers(Users map[string]*golan) {
	if len(Users) == 0 {
		fmt.Println("Нет пользователей")
		return
	}
	for _, user := range Users {
		fmt.Printf("%s: $%.2f, €%.2f\n", user.name, user.dollars, user.euro)
	}
}

func DEP(Users map[string]*golan) {
	name := ""
	fmt.Println("Имя пользователя: ")
	_, err := fmt.Scanln(&name)
	if err != nil {
		return
	}
	user, ok := Users[name]
	if !ok {
		fmt.Println("Пользователь не найден")
		return
	}
	fmt.Println("dollars or euro")
	ds := ""
	_, err = fmt.Scanln(&ds)
	if err != nil {
		return
	}
	fmt.Println("Сколько деп?")
	dep := 0.0
	_, err = fmt.Scanln(&dep)
	if err != nil {
		return
	}
	if dep <= 0 {
		fmt.Println("Сумма должна быть больше 0")
		return
	}

	switch strings.ToLower(ds) {
	case "dollars":
		{
			user.dollars += dep
		}
	case "euro":
		{
			user.euro += dep
		}
	default:
		{
			fmt.Println("error")
			return
		}
	}

	fmt.Printf("Ваш баланс: $%.2f, €%.2f\n", user.dollars, user.euro)
}

func addUser(Users map[string]*golan) {
	var name string
	fmt.Println("Введите имя:")
	_, err := fmt.Scanln(&name)
	if err != nil {
		return
	}
	if name == " " {
		fmt.Println("Имя не может быть пустым")
		return
	}
	if _, ok := Users[name]; ok {
		fmt.Println("Пользователь уже существует")
		return
	}
	Users[name] = &golan{name, 0, 0}
	fmt.Println("Пользователь добавлен:", name)
	fmt.Println(Users[name])
}

func CDE(Users map[string]*golan) {
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
	fmt.Println("У вас долларов: ", user.dollars, "У вас евро:", user.euro)

	fmt.Println("`dollartoeuro` or `eurotodollar`")
	doloreur := ""
	_, err = fmt.Scanln(&doloreur)
	if err != nil {
		return
	}
	doloreur = strings.ToLower(doloreur)

	summa := 0.0
	fmt.Println("Сколько переводим?")
	_, err = fmt.Scanln(&summa)
	if err != nil {
		return
	}

	if summa <= 0 {
		fmt.Println("Сумма должна быть больше 0")
		return
	}

	switch doloreur {
	case "dollartoeuro":
		{
			if summa > user.dollars {
				fmt.Println("Недостаточно средств.")
				return
			}
			user.euro += summa * rateDolToEur
			user.dollars -= summa
		}
	case "eurotodollar":
		{
			if summa > user.euro {
				fmt.Println("Недостаточно средств.")
				return
			}
			user.dollars += summa * rateEurToDol
			user.euro -= summa
		}
	default:
		{
			fmt.Println("error")
			return
		}
	}

	fmt.Printf("Новый баланс: $%.2f, €%.2f\n", user.dollars, user.euro)
}
