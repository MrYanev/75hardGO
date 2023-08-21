package main

import (
	"fmt"
	"strings"
)

type User struct {
	Name     string
	Progress int
	Tasks    []string
}

func main() {
	//The program starts by asking the user for a name
	//And then greeting him with that name
	var input string
	fmt.Println("How may I call you?")
	fmt.Scanln(&input)
	user1 := User{Name: input}

	fmt.Println("Hello", user1.Name)
	fmt.Println("Are you ready to become a better version of yourself?")
	fmt.Scanln(&input)
	if strings.ToLower(input) == "yes" {
		fmt.Println("Let's get it on!")
	}

	fmt.Println("Do you want to add tasks to the challenge or you will stick to the original one? (Yes/No)")
	fmt.Scanln(&input)
	if strings.ToLower(input) == "yes" {
		fmt.Printf("I like your spirit, %v!", user1.Name)

		user1.Tasks = []string{
			"drink 3 liters of water",
			"two 45 minute workouts",
			"read 10 pages of a non-fiction book",
			"follow a diet",
		}

		fmt.Println("What else would you like me to add to the list?")
		fmt.Scan(&input)
		user1.Tasks = append(user1.Tasks, input)

		fmt.Println("Is ther anything else you want ot add?")
		fmt.Scan(&input)
		if strings.ToLower(input) == "no" {
			fmt.Println(user1.Tasks)
		}
	}
}
