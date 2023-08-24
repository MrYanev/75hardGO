package api

import (
	"75hardgo/models"

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

// Create API endpoint to reach from frontend or api client
func (s *Service) Create(u *models.User) (*models.User, error) {
	// add user in the map and create a txt file for him
	return u.Create()
}

func (s *Service) Get(name string) (*models.User, error) {
	// print out info of user from his txt file or return err
	return nil, nil
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
