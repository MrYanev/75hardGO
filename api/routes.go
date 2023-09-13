package api

// define routes for endpoints

// flow of logic/inputs should go:
// main -> api -> models -> storage (txt files)

import (
	"75hardgo/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Service) CreateRouting(c *gin.Context) {

	newUser, ok := c.GetQuery("name")
	if !ok { //If not OK error message
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing Name query parameter!"})
		return
	}
	if err := c.BindJSON(&newUser); err != nil {
		c.String(http.StatusBadRequest, "Invalid Input")
		return
	}

	nu := &models.User{Name: newUser}
	createdUser, err := s.Create(nu) //To be checked
	if err != nil {
		c.String(http.StatusInternalServerError, "Error creating user: %s", err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User created succsessfully",
		"user":    createdUser,
	})

}

func (s *Service) CheckProgressOnTasksRouting() {
	s.Router.GET("/checker", func(c *gin.Context) {
		var newUser models.User
		if err := c.BindJSON(&newUser); err != nil {
			c.String(http.StatusNotFound, "User doesn't exist!")
		}

		userTasks, err := s.CheckTasks(newUser.Name, &newUser)
		if err != nil {
			c.String(http.StatusInternalServerError, "The shit just hit the fan!")
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "These are your daily tasks:",
			"tasks":   userTasks,
		})
	})
}

func (s *Service) ReadUserDataRouting() {
	s.Router.GET("/reader", func(c *gin.Context) {

	})
}

func (s *Service) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello, World!",
	})
}
