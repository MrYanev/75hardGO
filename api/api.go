package api

import (
	"75hardgo/models"
	"fmt"

	"github.com/gin-gonic/gin"
)

type Service struct {
	Router *gin.Engine
	Users  map[string]*models.User
}

func NewService() *Service {
	return &Service{
		Router: gin.Default(),
		Users:  make(map[string]*models.User),
	}
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
	s.Users[u.Name] = u

	fileName := fmt.Sprintf("%s_user_data.txt", u.Name)
	createdUser, err := u.Create(fileName)
	if err != nil {
		return nil, err
	}

	return createdUser, nil
}

func (s *Service) Get(name string) (*models.User, error) {
	// print out info of user from his txt file or return err
	user, ok := s.Users[name]
	if !ok {
		return nil, fmt.Errorf("User not found")
	}

	fileName := fmt.Sprintf("%s_user_data.txt", user.Name)
	userFromFile, err := models.ReadUserFromFile(fileName)
	if err != nil {
		return nil, err
	}
	return userFromFile, nil
}

func (s *Service) AddTask(task string, userName string) error {
	// get user by name from service map of users
	// add task to user (might be a func of the User struct)
	// update the User's txt file with the new task/progress
	// check if errors or return nil
	return nil
}

func (s *Service) CheckTasks() ([]string, error) {
	// just display tasks (u.CheckTasks())
	return nil, nil
}
