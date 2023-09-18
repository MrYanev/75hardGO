package main

import (
	"75hardgo/api"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	service, err := api.NewService()
	if err != nil {
		fmt.Printf("Error creating a service: %s\n", err)
	}

	// Add all routes here like this
	service.Router.GET("/ping", service.Ping)

	//Route for adding a new task
	service.Router.POST("/adder", service.AdderRouting)

	//Route for creating an user
	service.Router.POST("/creates", service.CreateRouting)

	//Route for ReadingUserData
	service.Router.GET("/reader", service.ReadUserDataRouting)

	//Route for CheckProgressOnTask
	service.Router.GET("/getter", service.CheckProgressOnTasksRouting)

	//Update user task routing
	service.Router.POST("/update", service.UpdateTasks)

	//Route for responses
	service.Router.POST("/responder", service.ResponseRoute)

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-signalCh
		service.WriteUpdatesToFile()
		fmt.Println("Received termination signal. Shutting down...")
		os.Exit(0)
	}()

	log.Printf("Running on 8080")
	service.Router.Run()

}
