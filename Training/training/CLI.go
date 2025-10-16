package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type list struct {
	Name  string  `json:"name"`
	Money float64 `json:"money"`
	Id    int     `json:"id"`
}

func main() {
	var lists []*list
	loadJSON(&lists)
	printLists(lists)
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run cli.go <add> <name>")
		return
	}

	arg2 := os.Args[2]

	addUser(&lists, arg2)
}

func loadJSON(lists *[]*list) {
	fileName := "data.json"
	if _, err := os.Stat(fileName); err == nil {
		// файл существует → читаем
		file, err := os.Open(fileName)
		if err != nil {
			fmt.Println("Error opening file:", err)
			return
		}
		defer func(file *os.File) {
			err := file.Close()
			if err != nil {

			}
		}(file)

		err = json.NewDecoder(le).Decode(lists)
		if err != nil {
			fmt.Println("Error decoding JSON:", err)
			return
		}
		fmt.Println("Tasks loaded from JSON.")
	} else if os.IsNotExist(err) {
		// файла нет → создаём пустой список
		*lists = []*list{}
		fmt.Println("No lists.json found. Starting fresh.")
	} else {
		fmt.Println("Error checking file:", err)
	}
}

func saveJSON(lists *[]*list) {
	fileName := "data.json"
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer func(file *os.File) {
		err = file.Close()
		if err != nil {

		}
	}(file)

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(lists)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
	}
}

func printLists(lists []*list) {
	for _, l := range lists {
		fmt.Println("Name:", l.Name, "| Money:", l.Money, "| Id:", l.Id)
	}
}

func addUser(lists *[]*list, name string) {
	newId := 0
	for {
		flag := false
		for _, l := range *lists {
			if l.Id == newId {
				flag = true
				newId++
				break
			}
		}
		if !flag {
			break
		}
	}
	*lists = append(*lists, &list{Name: name, Money: 0, Id: newId})
	fmt.Println("User added!")
	printLists(*lists)
	saveJSON(lists)
}
