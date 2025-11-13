package main

import (
	"os"
	"fmt"
	"time"
	"strings"
	"bufio"
)

type Task struct {
	Title string
	Status bool
	CreatedAt time.Time
	CompletedAt *time.Time
}

func listTasks(tasks []Task) {
	fmt.Println("+-----+----------------------+-----------+---------------------+---------------------+")
  fmt.Println("| No. | Task                 | Status    | Created At          | Completed At        |")
  fmt.Println("+-----+----------------------+-----------+---------------------+---------------------+")

	if len(tasks) == 0 {
  	fmt.Println("|                                 NO TASKS FOUND                                     |")
  	fmt.Println("+------------------------------------------------------------------------------------+")
	} else {
		for i, t := range tasks {
		completed := "Pending"
		completedAt := "N/A"
		if t.Status {
			completed = "Completed"
			completedAt = t.CompletedAt.Format("2006-01-02 15:04:05")
		}
			fmt.Printf("| %-3d | %-20s | %-9s | %-19s | %-19s |\n", i, t.Title, completed, t.CreatedAt.Format("2006-01-02 15:04:05"), completedAt)

  		fmt.Println("+-----+----------------------+-----------+---------------------+---------------------+")
		}
	}
}

func main () {

	tasks := []Task{}

	var input string

	for {
		// List the tasks
		listTasks(tasks)

		// Main Menu
		fmt.Println("1. Add a new task")
		fmt.Println("2. Complete a task")
		fmt.Println("3. Delete a task")
		fmt.Println("Enter q for exit")
		fmt.Print(": ")

		fmt.Scan(&input)

		switch input {
		case "1":
			fmt.Println("Enter the new task")
			fmt.Print(": ")
	
			// scan until a newline, i.e enter
			reader := bufio.NewReader(os.Stdin)
			title, _ := reader.ReadString('\n')
			// remove \n
			title = strings.TrimSpace(title)

			newTask := Task{
				Title: title,
				Status: false,
				CreatedAt: time.Now(),
				CompletedAt: nil,
			}

			tasks = append(tasks, newTask)
			fmt.Println("Task Added.")
		
		case "2":
			var index int
			fmt.Println("Enter task index to complete")
			fmt.Print(": ")
			fmt.Scan(&index)

			if index >= 0 && index < len(tasks) {
				if !tasks[index].Status {
					tasks[index].Status = true
					t := time.Now()
					tasks[index].CompletedAt = &t
					fmt.Println("Task marked as completed.")
				} else {
					fmt.Println("Task is already completed.")
				}
			} else {
				fmt.Println("Invalid index.")
			}
		case "3":
			var index int
			fmt.Println("Enter task index to delete")
			fmt.Print(": ")
			fmt.Scan(&index)

			if index >= 0 && index < len(tasks) {
				tasks = append(tasks[:index], tasks[index+1:]...)
				fmt.Println("Task deleted.")
			} else {
				fmt.Println("Invalid index.")
			}
		case "q":
			fmt.Println("Existing program...")
			return

		default:
			fmt.Println("Invalid choice, try again.")
		}
	}


}
