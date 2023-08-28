package models

import (
	"fmt"
	"os"
)

type User struct {
	Name     string   `json:"name"`
	Progress int      `json:"progress"`
	Tasks    []string `json:"tasks"`
}

func (u *User) CheckTasks() []string {
	return u.Tasks
}

func (u *User) CheckProgressOnTasks(taskName string) bool {
	// check if tasks are compleated and reset or print shame message!
	return false
}

func (u *User) Create(fileName string) (*User, error) {
	file, err := os.Create(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	_, err = fmt.Fprintf(file, "Name: %s\nProgress: %d\n", u.Name, u.Progress)
	if err != nil {
		return nil, err
	}
	return u, nil
}
