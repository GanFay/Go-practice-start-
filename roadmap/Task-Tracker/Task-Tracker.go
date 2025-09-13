package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strconv"
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
		fmt.Printf("\n\\/Task tracker|:=>\n|\n┖\\/Input command|:=>\n |\n ┖= addtask, deltask, update, sort\n |\n ┖= list <todo || inprogress || done. Or nothing>, exit.\n\n")
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
				changeStatus(tasks, Status(st))
			}
		case "exit":
			{
				fmt.Println("BB")
				return
			}
		case "sort":
			sortTasks(&tasks)
		default:
			{
				fmt.Println("error")
			}
		}
	}
}

func sortTasks(tasks *[]*list) {
	sort.Slice(*tasks, func(i, j int) bool {
		return (*tasks)[i].Id < (*tasks)[j].Id
	})
	fmt.Println("Sorted Tasks")
	saveJSON(tasks)
	fmt.Println("SavedJSON Tasks")
}

func changeStatus(tasks []*list, status Status) {
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
	if task, found = findById(*tasks, deltask1); found {
		fmt.Printf("Found! \nID: %d | Name: %s | Description: %s | Status: %s\n", task.Id, task.Name, task.Description, task.Status)
	} else if task, found = findByName(*tasks, deltask1); found {
		fmt.Printf("Found! \nID: %d | Name: %s | Description: %s | Status: %s\n", task.Id, task.Name, task.Description, task.Status)
	} else {
		fmt.Println("Not Found")
		return
	}

	nig := ""
	fmt.Println("U sure want to delete this task")
	_, err = fmt.Scanln(&nig)
	if err != nil {
		return
	}
	switch nig {
	case "no":
		{
			deltask(tasks)
		}
	case "-":
		{
			deltask(tasks)
		}
	case "false":
		{
			deltask(tasks)
		}
	}

	var newTasks []*list
	for _, task1 := range *tasks {

		if task1.Id != task.Id {
			newTasks = append(newTasks, task1)
		}
	}
	*tasks = newTasks
	saveJSON(tasks)
	fmt.Println("Delete Task Success")
	sortTasks(tasks)
}

func findById(tasks []*list, id string) (*list, bool) {
	for i := range tasks {
		if strconv.Itoa(tasks[i].Id) == id {
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
	fmt.Println("Input ID or Name")

	choice := ""

	_, err := fmt.Scanln(&choice)
	if err != nil {
		return
	}
	choice = strings.ToLower(choice)
	var task *list
	var found bool

	if task, found = findById(*tasks, choice); found {
		fmt.Printf("Found! \nID: %d | Name: %s | Description: %s | Status: %s\n", task.Id, task.Name, task.Description, task.Status)
	} else if task, found = findByName(*tasks, choice); found {
		fmt.Printf("Found! \nID: %d | Name: %s | Description: %s | Status: %s\n", task.Id, task.Name, task.Description, task.Status)
	} else {
		fmt.Println("Not Found")
		return
	}
	fmt.Println("Input task status (todo, inprogress, done):")
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
	default:
		{
			fmt.Println("Unknown command. Error")
			return
		}
	}
	saveJSON(tasks)
	fmt.Println("Task updated and save successfully.")
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

	for _, pid := range *tasks {
		if pid.Name == scrp {
			fmt.Println("Input name error (name already exists)!")
			return
		}
	}

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

	fmt.Println("Task added!")
	*tasks = append(*tasks, &list{Name: scrp, Description: choice, Id: newID, Status: Todo})
	saveJSON(tasks)
	sortTasks(tasks)
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
