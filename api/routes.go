package api

import (
	"75hardgo/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// define routes for endpoints

// flow of logic/inputs should go:
// main -> api -> models -> storage (txt files)

func CreateRouting(s *Service, r *gin.Engine) {
	s.Router.POST("/create", func(c *gin.Context) {
		var newUser models.User
		if err := c.BindJSON(&newUser); err != nil {
			c.String(http.StatusBadRequest, "Invalid Input")
			return
		}

		createdUser, err := s.Create(&newUser)
		if err != nil {
			c.String(http.StatusInternalServerError, "Error creating user: %s", err.Error())
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "User created succsessfully",
			"user":    createdUser,
		})
	})

}

func CheckProgressOnTasksRouting(s *Service, r *gin.Engine) {
	s.Router.GET("/check", func(c *gin.Context) {
		var newUser models.User
		if err := c.BindJSON(&newUser); err != nil {
			c.String(http.StatusBadRequest, "User doesn't exist!")
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
