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
	/*	err := service.LoadUsersFromTxtFiles(dataFolder)
		if err != nil {
			return nil, err
		}
	*/
	return service, nil
}

func (s *Service) SetUp() {
	//For cycle over all user files in the data dir
	//On each file read and unmarshall json user and
	//add it to the map
	file, err := os.Open(s.Path)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()
	var user models.User
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&user)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	//s.Users[]
}

func (s *Service) LoadUsersFromTxtFiles(folderName string) error {
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
	fileName := fmt.Sprintf("%s\\%s.json", directory, u.Name)
	log.Printf("Dir is %s", directory)
	err = os.WriteFile(fileName, userJSON, 0644)
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
