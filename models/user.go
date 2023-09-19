package models

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
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

func (u *User) CheckProgressOnTasks(userName string) error {
	// check if tasks are compleated and reset or print shame message!
	/*
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
				break
			}
		}
	*/
	return nil
}

func (u *User) Create() (*User, error) {
	for _, i := range basicTasks {
		u.Tasks = append(u.Tasks, i)
	}

	return u, nil
}

func (u *User) ReadUserDataFromFile(filePath string) (*User, error) {
	//Function for fetching data drom the text file
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("Failed to read the file!")
	}
	defer file.Close()

	//theUser = &User{}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			continue
		}
		if _, err := fmt.Scanf(line, "Name: %s", &u.Name); err == nil {
			continue
		}
		if _, err := fmt.Scanf(line, "Process: %d", &u.Progress); err == nil {
			continue
		}

		u.Tasks = append(u.Tasks, line)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("Error while reading the file!")
	}

	return u, nil
}

func (u *User) UpdateUserFile(fileName string) error {

	file, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("Error working with the file")
	}

	defer file.Close()

	_, err = fmt.Fprintf(file, "Name: %s\nProgress: %d\nTasks: %v\n", u.Name, u.Progress, u.Tasks)
	if err != nil {
		return err
	}
	for _, task := range u.Tasks {
		_, err = fmt.Fprintf(file, task)
		if err != nil {
			return err
		}
	}

	return nil
}

func (u *User) WriteUpdatesToFile() error {

	directory, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("Failed to get the folder path")
	}
	subFolderPath := filepath.Join(directory, "data")
	fileName := fmt.Sprintf("%s.json", u.Name)
	filePath := filepath.Join(subFolderPath, fileName)

	userJSON, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		return fmt.Errorf("Failed to marshal user data: %v", err)
	}

	err = os.WriteFile(filePath, userJSON, 0644)
	if err != nil {
		return fmt.Errorf("Failed to write user data to file: %v", err)
	}

	fmt.Printf("User data written to %s\n", fileName)
	return nil
}
