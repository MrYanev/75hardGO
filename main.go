package main

import (
	"75hardgo/api"
	//"75hardgo/models"
	//"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	apiService := api.NewService()

	api.CreateRouting(apiService, r)

	api.CheckProgressOnTasksRouting(apiService, r)

	r.Run("localhost:8080")

	//Name to be added
	//userFileName := ""

}
