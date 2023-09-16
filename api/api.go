package api

import (
	"75hardgo/models"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"path/filepath"

	"github.com/gin-gonic/gin"
)

type Service struct {
	Router *gin.Engine
	Users  map[string]*models.User
	Path   string
}

func NewService() (*Service, error) {
	service := &Service{
		Router: gin.Default(),
		Users:  make(map[string]*models.User),
	}
	dir, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("Couldn't take current dirrectory")
	}
	service.Path = filepath.Join(dir, "data")
	service.SetUp()
	return service, nil
}

func (s *Service) SetUp() {
	//For cycle over all user files in the data dir
	//On each file read and unmarshall json user and
	//add it to the map
	log.Printf(s.Path)
	files, err := os.ReadDir(s.Path)
	if err != nil {
		fmt.Println("Error reading the directory:", err)
		return
	}

	for _, file := range files {
		if !file.IsDir() && filepath.Ext(file.Name()) == ".json" {
			data, err := os.ReadFile(filepath.Join(s.Path, file.Name()))
			if err != nil {
				fmt.Printf("Error reading the file %s: %s\n", file.Name(), err)
				continue
			}

			var user *models.User
			err = json.Unmarshal(data, &user)
			if err != nil {
				fmt.Printf("Error unmarshalling the JSON file %s: %s\n", file.Name(), err)
				continue
			}
			s.Users[user.Name] = user
		}
	}
}

func (s *Service) LoadUsersFromTxtFiles() error {
	// on each app restarts read from the folder for user txt files
	// and load those in the service map of users in mem so we have
	// persistent usage between app restarts
	return nil
}

// Create is the api endpoint for creating users by given
// name, progress and tasks.
func (s *Service) Create(u *models.User) (*models.User, error) {
	// add user in the map and create a txt file for him

	createdUser, err := u.Create()
	if err != nil {
		return nil, fmt.Errorf("Failed to create a file!")
	}
	userJSON, err := json.Marshal(createdUser)
	if err != nil {
		fmt.Println("Error", err)
		return nil, err
	}

	directory, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	subFolderPath := filepath.Join(directory, "data")
	fileName := fmt.Sprintf("%s.json", u.Name)
	filePath := filepath.Join(subFolderPath, fileName)

	log.Printf("Dir is %s", directory)
	err = os.WriteFile(filePath, userJSON, 0644)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}
	s.Users[u.Name] = u

	return createdUser, nil
}

func (s *Service) Get(u *models.User) (*models.User, error) {
	// print out info of user from his txt file or return err
	user, ok := s.Users[u.Name]

	if !ok {
		return nil, fmt.Errorf("User not found")
	}
	fileName := fmt.Sprintf("%s_user_data.txt", user.Name)
	theUser, err := u.ReadUserDataFromFile(fileName)
	if err != nil {
		return nil, err
	}
	return theUser, nil
}

func (s *Service) AddTask(userName string, task string) error {
	// get user by name from service map of users
	// add task to user (might be a func of the User struct)
	// update the User's txt file with the new task/progress
	// check if errors or return nil

	user, ok := s.Users[userName]
	if !ok {
		return fmt.Errorf("User not found")
	}

	user.Tasks = append(user.Tasks, task)

	fileName := filepath.Join(s.Path, fmt.Sprintf("%s_user_data.txt", user.Name))
	if err := user.UpdateUserFile(fileName); err != nil {
		return err
	}

	return nil
}

func (s *Service) CheckTasks(name string, u *models.User) (bool, error) {
	// just display tasks (u.CheckTasks())
	user, ok := s.Users[name]
	if !ok {
		return false, fmt.Errorf("User not found")
	}

	tasksUser, err := u.CheckProgressOnTasks(user.Name)
	if err != nil {
		return false, fmt.Errorf("Shit just hit the fan!")
	}

	return tasksUser, nil
}
