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

	//Route for creating an user
	service.Router.POST("/creates", service.CreateRouting)

	//Route for CheckProgressOnTask
	service.Router.GET("/chekcer", service.CheckProgressOnTasksRouting)

	log.Printf("running on 8080")
	service.Router.Run()

}
