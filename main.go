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
		}
	}
}
