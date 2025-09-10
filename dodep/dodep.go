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

	suck := login()

	Users[suck] = golan{
		suck,
		0,
		0,
	}
	for {
		lol(Users, suck)
	}
}

func login() string {
	fmt.Println("Введите имя")
	suck := ""
	_, err := fmt.Scanln(&suck)
	if err != nil {
		return ""
	}
	println(suck)
	return suck
}

func lol(Users map[string]golan, suck string) {
	fmt.Println("Какую операцию вы хотите зделать, CDE(Convert dollars to euro), CED(Convert euro to dollars), Dep, exit")
	var nig string
	fmt.Scanln(&nig)
	nig = strings.ToLower(nig)

	switch nig {
	case "cde":
		{
			fmt.Println("Сколько доларов хотите перевести?")
			dol1 := 0.0
			fmt.Scanln(&dol1)
			convertdoltoeur(Users, suck, dol1)
			fmt.Println(Users[suck])
		}
	case "ced":
		{
			fmt.Println("Сколько евро хотите перевести?")
			euro1 := 0.0
			fmt.Scanln(&euro1)
			converteurtodol(Users, suck, euro1)
			fmt.Println(Users[suck])
		}
	case "generate":
		{
			fmt.Println("Сколько доларов dep?")
			dodep := 0.0
			fmt.Scanln(&dodep)
			dep(Users, suck, dodep)
			fmt.Println(Users[suck])
		}
	case "dep":
		{
			fmt.Println("Сколько доларов dep?")
			dodep := 0.0
			fmt.Scanln(&dodep)
			dep(Users, suck, dodep)
			fmt.Println(Users[suck])
		}
	case "exit":
		{
			login()
		}
	default:
		{
			fmt.Println("Пиши нормально")
		}
	}
}

func convertdoltoeur(Users map[string]golan, name string, dollars float64) {
	user := Users[name]
	rate := 0.854
	euro := rate * dollars
	user.euro += euro
	user.dollars -= dollars
	Users[name] = user
}

func converteurtodol(Users map[string]golan, name string, euro float64) {
	user := Users[name]
	rate := 1.17
	dollars := rate * euro
	user.dollars += dollars
	user.euro -= euro
	Users[name] = user
}

func dep(Users map[string]golan, name string, x float64) {
	user := Users[name]
	user.dollars += x
	Users[name] = user
}
