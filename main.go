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
	} else {
		fmt.Print("You have no choice lazy bastard!\n")
	}

	user1.Tasks = []string{
		"drink 3 liters of water",
		"two 45 minute workouts",
		"read 10 pages of a non-fiction book",
		"follow a diet",
	}

	fmt.Println("Do you want to add tasks to the challenge or you will stick to the original one? (Yes/No)")
	fmt.Scanln(&input)
	/*
		for strings.ToLower(input) == "yes" || strings.ToLower(input) == "no" {
			fmt.Println("Sorry, I did not quite catch that?")
			fmt.Scanln(&input)
		}
	*/
	if strings.ToLower(input) == "yes" {
		fmt.Printf("I like your spirit, %v!\n", user1.Name)

		fmt.Println("What else would you like me to add to the list?")
		fmt.Scanln(&input)
		user1.Tasks = append(user1.Tasks, input)

		fmt.Println("Is ther anything else you want ot add?")
		fmt.Scanln(&input)

		/*
			for strings.ToLower(input) != "yes" || strings.ToLower(input) != "no" {
				fmt.Println("Sorry, I did not quite catch that?")
				fmt.Scanln(&input)
				if strings.ToLower(input) == "yes" || strings.ToLower(input) == "no" {
					break
				}
			}
		*/

		if strings.ToLower(input) == "no" {
			fmt.Println("Cool! These are your daily goals for the next 75 days!")
			fmt.Println(user1.Tasks)

		} else if strings.ToLower(input) == "yes" {
			for strings.ToLower(input) == "yes" {
				fmt.Println("What else should I add to your challange?")
				fmt.Scanln(&input)
				user1.Tasks = append(user1.Tasks, input)
				fmt.Println("Is ther anything else you want ot add?")
				fmt.Scanln(&input)
				if strings.ToLower(input) == "no" {
					break
				} else {
					input = "yes"
				}
			}
		}
	}

	fmt.Println("In that case these will be the tasks you must accomplish each day: ")
	for _, i := range user1.Tasks {
		fmt.Println(i)
	}
}
