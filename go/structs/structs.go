package main

import (
	"fmt"

	"example.com/structs/user"
)

// type user struct {
// 	firstName string
// 	lastName string
// 	birthdate string
// 	createdAt time.Time
// }

// func (u *user) outputUserDetails() {
// 	//...
// 	fmt.Println(u.firstName, u.lastName, u.birthdate, u.createdAt)
// }

// func (u *user) clearUserName() { // * pointer is added here to ensure the original is being edited, and not a copy
// 	u.firstName = " "
// 	u.lastName = " "
// }

// func newUser(firstName, lastName, birthdate string) (*user, error) {
// 	if firstName == "" || lastName == "" || birthdate == ""{
// 		return nil, errors.New("First name, last name, and birthdate are required.")
// 	}
// 	return &user{
// 		firstName: firstName,
// 		lastName: lastName,
// 		birthdate: birthdate,
// 		createdAt: time.Now(),
// 	}, nil
// }





func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var appUser *user.User

	appUser, err := user.NewUser(userFirstName,userLastName,userBirthdate)

	if err != nil {
		fmt.Println(err)
		return
	}

	admin := user.NewAdmin("test@admin.com", "769222")

	admin.OutputUserDetails()
	admin.ClearUserName()

	// ... do something awesome with that gathered data!

	appUser.OutputUserDetails()
	appUser.ClearUserName()
	appUser.OutputUserDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}