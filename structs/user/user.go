package user

import (
	"fmt"
	"errors"
	"time"
)

type User struct{
	firstName string
	lastName string
	birthdate string
	createdAt time.Time
}

// Making use of struct embedding

type Admin struct {
	email string
	password string
	User
}

// This is now the Struct user method
func (u User) OutputUserDetailsMethod() {
	fmt.Println(u.firstName, u.lastName, u.birthdate)
}

func (u *User) ClearUserName() {
	u.firstName = ""
	u.lastName = ""
}

func New(userFirstName, userLastName, userBirthdate string) (*User, error) {
	if (userFirstName == "" || userLastName == "" || userBirthdate == ""){
		return nil, errors.New("firstName, lastName and birthDate all the fields are required")
	}
	return & User{
		firstName: userFirstName,
		lastName: userLastName,
		birthdate: userBirthdate,
		createdAt: time.Now(),
	}, nil
}


// Creating Admin constructor
func NewAdmin(userEmail, userPassword string) Admin {
	return Admin{
		email: userEmail,
		password: userPassword,
		User: User{
			firstName: "Admin",
			lastName: "Data",
			birthdate: "----",
			createdAt: time.Now(),
		},
	}
}