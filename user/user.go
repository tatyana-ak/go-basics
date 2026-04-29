package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	FirstName string
	LastName  string
	Birthdate string
	CreatedAt time.Time
}

type Admin struct {
	Email    string
	Password string
	User     User
}

func (u User) OutputUserData() {
	fmt.Println(u.FirstName, u.LastName, u.Birthdate, u.CreatedAt)
}

func (u *User) ClearUserName() {
	u.FirstName = ""
	u.LastName = ""
}

func NewAdmin(email, password string) Admin {
	return Admin{
		email,
		password,
		User{
			FirstName: "ADMIN",
			LastName:  "ADMIN",
			Birthdate: "07/07/1990",
		},
	}
}

func New(firstName, lastName, birthdate string) (*User, error) {
	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("The data are required")
	}

	return &User{
		firstName,
		lastName,
		birthdate,
		time.Now(),
	}, nil
}
