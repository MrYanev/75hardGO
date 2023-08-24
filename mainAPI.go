package main

import (
	"75hardgo/api"
)

func main() {
	service := api.NewService()

	service.Router.Run("localhost:8080")
}
