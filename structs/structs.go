package main

import (
	"fmt"
	"example.com/structs/user"
	// "time"
)

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	/*
	var appUser = user{
		firstName: userFirstName,
		lastName: userLastName,
		birthdate: userBirthdate,
		createdAt: time.Now(),
	}
	*/

	/*
	appUser := user.User{
		FirstName: userFirstName,
		LastName: userLastName,
		Birthdate: userBirthdate,
		CreatedAt: time.Now(),
	}
	*/

	appUser, err := user.New(userFirstName, userLastName, userBirthdate)
	if err != nil {
		fmt.Println(err)
		return
	}

	userAdmin := user.NewAdmin("abc@example.com", "Password123")
	userAdmin.User.OutputUserDetailsMethod()
	userAdmin.User.ClearUserName()
	userAdmin.User.OutputUserDetailsMethod()
	// ... do something awesome with that gathered data!

	// fmt.Println(outputUserDetails(appUser))
	appUser.OutputUserDetailsMethod()
	appUser.ClearUserName()
	appUser.OutputUserDetailsMethod()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}

/*
func outputUserDetails(u user) (string, string, string){
	return u.firstName, u.lastName, u.birthdate
}
*/

/*
// Using pointers
func outputUserDetailsPointer(u *user) (string, string, string){
	return (*u).firstName, (*u).lastName, (*u).birthdate
}
*/
