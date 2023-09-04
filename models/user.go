package models

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type User struct {
	Name     string   `json:"name"`
	Progress int      `json:"progress"`
	Tasks    []string `json:"tasks"`
}

func (u *User) CheckTasks() []string {
	return u.Tasks
}

func (u *User) CheckProgressOnTasks(taskName string) (bool, error) {
	// check if tasks are compleated and reset or print shame message!
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Please comfirm the compleation of each task")
	for i, task := range u.Tasks {
		fmt.Printf("Task %d: %s (yes/no): ", i+1, task)
		answer, _ := reader.ReadString('\n')
		answer = strings.TrimSpace(answer)

		if strings.ToLower(answer) == "no" {
			u.Progress = 0
			fmt.Printf("You have failed to accomplish all tasks today\n")
			fmt.Println("Your progress has been restarted!")
		}
	}
	return true, nil
}

func (u *User) Create(fileName string) (*User, error) {
	file, err := os.Create(fileName)
	if err != nil {
		return nil, fmt.Errorf("Failed to create a file!")
	}
	defer file.Close()

	_, err = fmt.Fprintf(file, "Name: %s\nProgress: %d\n", u.Name, u.Progress)
	if err != nil {
		return nil, fmt.Errorf("Failed to create a file!")
	}
	return u, nil
}

func ReadUserDataFromFile(fileName string) {
	//Function for fetching data drom the text file
}
