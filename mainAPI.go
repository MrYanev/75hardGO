package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type User struct {
	Name     string   `json:"name"`
	Progress int      `json:"progress"`
	Tasks    []string `json:"tasks"`
}

var input string

func takeName() {
	fmt.Println("How may I call you?")
	fmt.Scanln(&input)
	user1 := User{Name: input}
	fmt.Println("Hello", user1.Name)
}

func main() {
	router := gin.Default()

	router.Run("localhost:8080")
}
