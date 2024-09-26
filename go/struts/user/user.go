package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {  	 //only captialized values can be exported
	firstName string
	lastName string
	birthdate string
	createdAt time.Time
}

type Admin struct {
	email string
	password string
	User
}




func (u *User) OutputUserDetails() {
	//...
	fmt.Println(u.firstName, u.lastName, u.birthdate, u.createdAt)
}

func (u *User) ClearUserName() { // * pointer is added here to ensure the original is being edited, and not a copy
	u.firstName = " "
	u.lastName = " "
}

func NewAdmin(email, password string) Admin {
	return Admin{
		email: email,
		password: password,
		User: User{
			firstName: "ADMIN",
			lastName: "ADMIN",
			birthdate: "10.26.1985",
			createdAt: time.Now(),
		},
	}
}

func NewUser(firstName, lastName, birthdate string) (*User, error) {
	if firstName == "" || lastName == "" || birthdate == ""{
		return nil, errors.New("First name, last name, and birthdate are required.")
	}
	return &User{
		firstName: firstName,
		lastName: lastName,
		birthdate: birthdate,
		createdAt: time.Now(),
	}, nil
}

