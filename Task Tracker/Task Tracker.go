package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Status string

const (
	Todo   Status = "todo"
	InProg Status = "inprogress"
	Done   Status = "done"
)

type list struct {
	Name        string
	Description string
	Id          int
	Status      Status
}

func main() {
	var tasks []*list
	loadJSON(&tasks)
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("\n\\/Task tracker|:=>\n|\n┖\\/Input command|:=>\n |\n ┖= addtask, deltask, update, list <todo || inprogress || done. Or nothing>, exit//.\n\n")
		input, _ := reader.ReadString('\n')
		choice := strings.TrimSpace(strings.ToLower(input))

		parts := strings.Fields(choice)

		if len(parts) == 0 {
			continue
		}

		switch parts[0] {
		case "addtask":
			addTask(&tasks)
		case "deltask":
			deltask(&tasks)
		case "update":
			updateField(&tasks)
		case "list":
			if len(parts) == 1 {
				printTasks(tasks)
			} else if len(parts) == 2 {
				st := parts[1]
				sort(tasks, Status(st))
			}
		default:
			{
				fmt.Println("error")
				return
			}
		}
	}
}

func sort(tasks []*list, status Status) {
	var curenttask []*list
	for i := range tasks {
		if tasks[i].Status == status {
			curenttask = append(curenttask, tasks[i])
			i++
		} else {
			i++
		}
	}
	printTasks(curenttask)
}

func deltask(tasks *[]*list) {
	fmt.Println("Input Id or Name task")
	deltask1 := ""
	_, err := fmt.Scanln(&deltask1)
	if err != nil {
		return
	}
	var found bool
	var task *list
	deltask1 = strings.ToLower(deltask1)
	if deltask1 == "id" {
		fmt.Println("Input Id:")
		id := 0
		_, err = fmt.Scanln(&id)
		if err != nil {
			return
		}
		if task, found = findById(*tasks, id); found {
			fmt.Printf("Нашёл! \nID: %d | Name: %s | Description: %s | Status: %s\n", task.Id, task.Name, task.Description, task.Status)
		} else {
			return
		}
	} else if deltask1 == "name" {
		fmt.Println("Input Id:")
		name := ""
		_, err = fmt.Scanln(&name)
		if err != nil {
			return
		}
		if task, found = findByName(*tasks, name); found {

			fmt.Printf("Нашёл! \nID: %d | Name: %s | Description: %s | Status: %s\n", task.Id, task.Name, task.Description, task.Status)

		} else {
			return
		}
	} else {
		fmt.Println("Invaild")
	}

	var newTasks []*list
	for _, task1 := range *tasks {

		if task1.Id != task.Id {
			newTasks = append(newTasks, task1)
		}
	}
	*tasks = newTasks
	saveJSON(tasks)
}

func findById(tasks []*list, id int) (*list, bool) {
	for i := range tasks {
		if tasks[i].Id == id {
			return tasks[i], true
		}
	}
	return nil, false
}

func findByName(tasks []*list, name string) (*list, bool) {
	for i := range tasks {
		if tasks[i].Name == name {
			return tasks[i], true
		}
	}
	return nil, false
}

func updateField(tasks *[]*list) {
	fmt.Println("ID or Name")

	choice := ""

	_, err := fmt.Scanln(&choice)
	if err != nil {
		return
	}
	choice = strings.ToLower(choice)
	var task *list
	var found bool
	if choice == "id" {
		fmt.Println("Input Id")
		id := 0
		_, err = fmt.Scanln(&id)
		if err != nil {
			return
		}
		if task, found = findById(*tasks, id); found {
			fmt.Printf("Нашёл! \nID: %d | Name: %s | Description: %s | Status: %s\n", task.Id, task.Name, task.Description, task.Status)
		} else {
			fmt.Println("Not found")
		}
	} else if choice == "name" {
		fmt.Println("Input name")
		name := ""
		_, err = fmt.Scanln(&name)
		if err != nil {
			return
		}
		name = strings.ToLower(name)
		if task, found = findByName(*tasks, name); found {
			fmt.Printf("Нашёл! \nID: %d | Name: %s | Description: %s | Status: %s\n", task.Id, task.Name, task.Description, task.Status)
		} else {
			fmt.Println("Не найдено")
		}
	} else {
		fmt.Println("error")
		return
	}
	fmt.Println("Введите на какой вы статус хотите поменять(todo, inprogress, done):")
	status := ""
	_, err = fmt.Scanln(&status)
	if err != nil {
		return
	}
	switch strings.ToLower(status) {
	case "todo":
		{
			task.Status = Todo
		}
	case "inprogress":
		{
			task.Status = InProg
		}
	case "done":
		{
			task.Status = Done
		}
	case "exit":
		{
			fmt.Println("BB")
			return
		}
	default:
		{
			fmt.Println("Unknown command. Error")
			return
		}
	}
	saveJSON(tasks)
	fmt.Println("Задача обновлена и сохранена в файл.")
}

func addTask(tasks *[]*list) {
	fmt.Println("Input name task(very short, for find):")
	var scrp string
	_, err := fmt.Scanln(&scrp)
	if err != nil {
		return
	}
	scrp = strings.ToLower(scrp)
	fmt.Println("Input description task:")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	choice := strings.TrimSpace(strings.ToLower(input))

	newID := 0
	for {
		exists := false
		for _, t := range *tasks {
			if t.Id == newID {
				exists = true
				newID++
				break
			}
		}
		if !exists {
			break
		}
	}

	*tasks = append(*tasks, &list{Name: scrp, Description: choice, Id: newID, Status: Todo})
	saveJSON(tasks)
}

func printTasks(tasks []*list) {
	for _, t := range tasks {
		fmt.Printf("ID: %d | Name: %s | Description: %s | Status: %s\n", t.Id, t.Name, t.Description, t.Status)
	}
}

func loadJSON(tasks *[]*list) {
	fileName := "tasks.json"
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

		err = json.NewDecoder(file).Decode(tasks)
		if err != nil {
			fmt.Println("Error decoding JSON:", err)
			return
		}
		fmt.Println("Tasks loaded from JSON.")
	} else if os.IsNotExist(err) {
		// файла нет → создаём пустой список
		*tasks = []*list{}
		fmt.Println("No tasks.json found. Starting fresh.")
	} else {
		fmt.Println("Error checking file:", err)
	}
}

func saveJSON(tasks *[]*list) {
	fileName := "tasks.json"
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(tasks)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
	}
}
