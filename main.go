package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Todo struct {
	ID        int
	Task      string
	Completed bool
}

var todos []Todo

func addTask(task string) {
	newTodo := Todo{
		ID:        len(todos) + 1,
		Task:      task,
		Completed: false,
	}

	todos = append(todos, newTodo)
}

func listTasks() {
	for _, todo := range todos {
		status := "Incomplete"

		if todo.Completed {
			status = "Complete"
		}

		fmt.Printf("%d. %s [%s]\n", todo.ID, todo.Task, status)
	}
}

func completeTask(id int) {
	for i := range todos {
		if todos[i].ID == id {
			todos[i].Completed = true
			break
		}
	}
}

func deleteTask(id int) {
	for i := range todos {
		if todos[i].ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			break
		}
	}
}

func saveTasks() {
	file, err := json.Marshal(todos)

	if err != nil {
		fmt.Println("Error encoding todos to JSON:", err)
		return
	}

	err = os.WriteFile("todos.json", file, 0644)

	if err != nil {
		fmt.Println("Error writing to file: ", err)
	}
}

func loadTasks() {
	if _, err := os.Stat("todos.json"); os.IsNotExist(err) {
		todos = []Todo{}
	}

	file, err := os.ReadFile("todos.json")

	if err != nil {
		fmt.Println("Error reading file: ", err)
		return
	}
	err = json.Unmarshal(file, &todos)

	if err != nil {
		fmt.Println("Error decoding JSON: ", err)
	}
}

func main() {

	loadTasks()
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\nTodo App")
		fmt.Println("1. Create new todo: ")
		fmt.Println("2. List of todos: ")
		fmt.Println("3. Complete todo: ")
		fmt.Println("4. Delete todo: ")
		fmt.Println("5. Exit")
		fmt.Print("Enter number: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			fmt.Print("Enter Task: ")
			task, _ := reader.ReadString('\n')
			task = strings.TrimSpace(task)
			addTask(task)
			saveTasks()
		case "2":
			listTasks()
		case "3":
			fmt.Print("Enter Task Id: ")
			taskIdStr, _ := reader.ReadString('\n')
			taskIdStr = strings.TrimSpace(taskIdStr)
			taskId, _ := strconv.Atoi(taskIdStr)
			completeTask(taskId)
			saveTasks()
		case "4":
			fmt.Print("Enter Task Id: ")
			taskIdStr, _ := reader.ReadString('\n')
			taskIdStr = strings.TrimSpace(taskIdStr)
			taskId, _ := strconv.Atoi(taskIdStr)
			deleteTask(taskId)
			saveTasks()
		case "5":
			fmt.Println("exiting ...")
			return
		}
	}
}
