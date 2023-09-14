package main

import (
	"75hardgo/api"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	service := api.NewService()

	// Add all routes here like this
	service.Router.GET("/ping", service.Ping)

	//Route for adding a new task
	service.Router.POST("/adder", service.AdderRouting)

	//Route for creating an user
	service.Router.POST("/creates", service.CreateRouting)

	//Route for ReadingUserData
	service.Router.GET("/reader", service.ReadUserDataRouting)

	//Route for CheckProgressOnTask
	service.Router.GET("/chekcer", service.CheckProgressOnTasksRouting)

	log.Printf("Running on 8080")
	service.Router.Run()

}
