package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Task struct {
	Description string
	Completed   bool
}

var tasks []Task

func addTask(description string) {
	task := Task{Description: description, Completed: false}
	tasks = append(tasks, task)
	fmt.Println("Added task:", description)
}

func listTasks() {
	if len(tasks) == 0 {
		fmt.Println("No tasks available")
		return
	}
	for i, task := range tasks {
		status := "✗"
		if task.Completed {
			status = "✓"
		}
		fmt.Printf("%d. [%s] %s\n", i+1, status, task.Description)
	}
}

func completeTask(index int) {
	if index < 0 || index >= len(tasks) {
		fmt.Println("Invalid task number.")
		return
	}
	tasks[index].Completed = true
	fmt.Printf("Completed task: %s\n", tasks[index].Description)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("\nChoose an option.")
		fmt.Println("1. Add task")
		fmt.Println("2. List Tasks")
		fmt.Println("3. Complete Task")
		fmt.Println("4. Exit")

		scanner.Scan()
		choice := strings.TrimSpace(scanner.Text())

		switch choice {
		case "1":
			fmt.Print("Enter task description: ")
			scanner.Scan()
			description := strings.TrimSpace(scanner.Text())
			addTask(description)
		case "2":
			listTasks()
		case "3":
			fmt.Println("Enter task number to complete:")
			scanner.Scan()
			var index int
			fmt.Sscan(scanner.Text(), &index)
			completeTask(index - 1)
		case "4":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println(fmt.Println("Invalid option, please try again"))
		}
	}
}
