package main

import "fmt"

type golan struct {
	name    string
	dollars float64
	euro    float64
}

func main() {

	Users := map[string]golan{}

	Users["Maks"] = golan{
		"Maks",
		0,
		0,
	}
	fmt.Println(Users)

	suck := ""
	fmt.Scanln(&suck)

	lol(Users, suck)
}

func lol(Users map[string]golan, suck string) {
	fmt.Println("Какую операцию вы хотите зделать, CDE, CED")
	var nig string
	fmt.Scanln(&nig)

	if nig == "CDE" {
		fmt.Println("Сколько доларов хотите перевести?")
		dol := 0.0
		fmt.Scanln(&dol)
		convertdoltoeur(Users, suck, dol)
		fmt.Println(Users[suck])
		lol(Users, suck)
	} else if nig == "CED" {
		fmt.Println("Сколько евро хотите перевести?")
		euro := 0.0
		fmt.Scanln(&euro)
		converteurtodol(Users, suck, euro)
		fmt.Println(Users[suck])
		lol(Users, suck)
	} else if nig == "Dep" {
		fmt.Println("Сколько доларов dep?")
		dodep := 0.0
		fmt.Scanln(&dodep)
		dep(Users, suck, dodep)
		fmt.Println(Users[suck])
		lol(Users, suck)
	} else {
		fmt.Println("Пиши нормально уёбище")
		return
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
