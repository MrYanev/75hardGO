package models

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

func (u *User) Create() (*User, error) {
	// create a txt file for the user and save it
	return nil, nil
}
