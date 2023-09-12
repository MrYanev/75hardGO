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

	log.Printf("running on 8080")
	service.Router.Run()

}
