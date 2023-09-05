package main

import (
	"75hardgo/api"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	apiService := api.NewService()

	api.CreateRouting(apiService, r)

	api.CheckProgressOnTasksRouting(apiService, r)

	api.ReadUserDataRouting(apiService, r)

	r.Run("localhost:8080")

}
