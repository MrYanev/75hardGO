package models

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var basicTasks = []string{
	"Read 10 pages of non-fictional book\n",
	"Drink a gallon of water\n",
	"Complete two 45 minutes workouts\n",
	"Follow a food regimen\n",
	"Take a progress photo\n",
}

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

	for _, i := range basicTasks {
		u.Tasks = append(u.Tasks, i)
	}

	_, err = fmt.Fprintf(file, "Name: %s\nProgress: %d\nTasks: %s", u.Name, u.Progress, u.Tasks)
	if err != nil {
		return nil, fmt.Errorf("Failed to create a file!")
	}
	return u, nil
}

func (u *User) ReadUserDataFromFile(fileName string) error {
	//Function for fetching data drom the text file
	file, err := os.Open(fileName)
	if err != nil {
		return fmt.Errorf("Failed to read the file!")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("Error while reading the file!")
	}
	return nil
}
