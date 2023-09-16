package api

// define routes for endpoints

// flow of logic/inputs should go:
// main -> api -> models -> storage (txt files)

import (
	"75hardgo/models"
	"log"

	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Service) CreateRouting(c *gin.Context) {

	newUser := c.Query("name")

	nu := &models.User{Name: newUser}

	createdUser, err := s.Create(nu)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error creating user: %s", err.Error())
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "User created succsessfully",
			"user":    createdUser,
		})
	}
}

func (s *Service) CheckProgressOnTasksRouting(c *gin.Context) {
	theUser := c.Query("name")

	if err := c.BindJSON(&theUser); err != nil {
		c.String(http.StatusNotFound, "User doesn't exist!")
	}

	userTasks, err := s.CheckTasks(theUser)
	if err != nil {
		c.String(http.StatusInternalServerError, "The shit just hit the fan!")
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "These are your daily tasks:",
			"tasks":   userTasks,
		})

	}
}

func (s *Service) ReadUserDataRouting(c *gin.Context) {
	theUser := c.Query("name")

	reader, err := s.Get(theUser)
	if err != nil {
		log.Printf("Prin map %#v", s.Users)
		c.String(http.StatusNotFound, "User not found!")
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "This is what we have on the record for that user: ",
			"data":    reader,
		})
	}

}

/*
	func (s *Service) CheckUsersTask(c *gin.Context) {
		theUser := c.Query("name")

		nu := &models.User{Name: theUser}
		checker, err := s.CheckTasks(theUser, nu)
		if err != nil {
			c.String(http.StatusNotFound, "User not found!")
		} else {
			c.JSON(http.StatusOK, gin.H{
				"message": "You have the following task to accomplish",
				"tasks":   checker,
			})
		}
	}
*/
func (s *Service) AdderRouting(c *gin.Context) {
	var TaskRequest struct {
		Task string `json:"task"`
		Name string `json:"name"`
	}

	if err := c.BindJSON(&TaskRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := s.AddTask(TaskRequest.Name, TaskRequest.Task); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task added successguly!"})
}

func (s *Service) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello, World!",
	})
}
