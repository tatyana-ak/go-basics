package main

import (
	"fmt"

	"examle.com/structs/user"
)

func main() {
	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthdate := getUserData("Please enter your birthday (MM/DD/YYYY): ")

	var appUser *user.User
	appUser, err := user.New(firstName, lastName, birthdate)
	appAdmin := user.NewAdmin("test@example.com", "123456")

	appAdmin.User.OutputUserData()

	if err != nil {
		fmt.Print(err)
		return
	}

	appUser.OutputUserData()
	appUser.ClearUserName()
	appUser.OutputUserData()
}

func getUserData(promtText string) string {
	fmt.Print(promtText)
	var value string
	fmt.Scanln(&value)
	return value
}
