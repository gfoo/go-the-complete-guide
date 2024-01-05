package main

import (
	"fmt"

	"org.gfoo/structs/user"
)

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}

func main() {
	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")
	userType := getUserData("Please enter user type (User/Admin): ")

	var u *user.User
	var a *user.Admin

	var err error

	if userType == "User" {
		u, err = user.NewUser(firstName, lastName, birthdate)
	} else if userType == "Admin" {
		email := getUserData("Please enter your email: ")
		password := getUserData("Please enter your password: ")
		a, err = user.NewAdmin(firstName, lastName, birthdate, email, password)
	} else {
		fmt.Println("Invalid user type")
		return
	}

	if err != nil {
		fmt.Println("Error creating user:", err)
		return
	}

	if userType == "User" {
		fmt.Println(u)
		u.ClearUserName()
		fmt.Println(u)
	} else if userType == "Admin" {
		fmt.Println(a)
		// a.User.ClearUserName()
		a.ClearUserName()
		a.ClearUserName()
		fmt.Println(a)
	}
}

// package main

// import (
// 	"errors"
// 	"fmt"
// 	"time"
// )

// type user struct {
// 	firstName string
// 	lastName  string
// 	birthdate string
// 	created   time.Time
// }

// // func outputUserData(u *user) {
// // 	// fmt.Println((*u).firstName, (*u).lastName, (*u).birthdate)
// // 	fmt.Println(u.firstName, u.lastName, u.birthdate)
// // }

// // with Pointer receivers
// // func (u user) outputUserData() {
// // 	// fmt.Println((*u).firstName, (*u).lastName, (*u).birthdate)
// // 	fmt.Println(u.firstName, u.lastName, u.birthdate)
// // }

// // toString like
// func (u user) String() string {
// 	return fmt.Sprintf("user{firstName: %s, lastName: %s, birthdate: %s, created: %v}", u.firstName, u.lastName, u.birthdate, u.created)
// }

// func (u *user) clearUserName() {
// 	// (*u).firstName = ""
// 	// (*u).lastName = ""
// 	u.firstName = ""
// 	u.lastName = ""
// }

// // constructor like for user
// func newUser(firstName, lastName, birthdate string) (*user, error) {

// 	// check all string are not empty or return an error object with message
// 	if firstName == "" {
// 		return nil, errors.New("first name cannot be empty")
// 	}
// 	if lastName == "" {
// 		return nil, errors.New("last name cannot be empty")
// 	}
// 	if birthdate == "" {
// 		return nil, errors.New("birthdate cannot be empty")
// 	}
// 	return &user{
// 		firstName: firstName,
// 		lastName:  lastName,
// 		birthdate: birthdate,
// 		created:   time.Now(),
// 	}, nil
// }

// func main() {

// 	// firstName := getUserData("Please enter your first name: ")
// 	// lastName := getUserData("Please enter your last name: ")
// 	// birthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

// 	// // create a User from previous data
// 	// var appUser = user{
// 	// 	firstName,
// 	// 	lastName,
// 	// 	birthdate,
// 	// 	created:   time.Now(),
// 	// }

// 	// fmt.Println(firstName, lastName, birthdate)

// 	// appUser := user{
// 	// 	firstName: getUserData("Please enter your first name: "),
// 	// 	lastName:  getUserData("Please enter your last name: "),
// 	// 	birthdate: getUserData("Please enter your birthdate (MM/DD/YYYY): "),
// 	// 	created:   time.Now(),
// 	// }

// 	// appUser := newUser(
// 	// 	getUserData("Please enter your first name: "),
// 	// 	getUserData("Please enter your last name: "),
// 	// 	getUserData("Please enter your birthdate (MM/DD/YYYY): "),
// 	// )

// 	firstName := getUserData("Please enter your first name: ")
// 	lastName := getUserData("Please enter your last name: ")
// 	birthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

// 	appUser, err := newUser(firstName, lastName, birthdate)
// 	if err != nil {
// 		fmt.Println("Error creating user:", err)
// 		return
// 	}

// 	// print out the user data
// 	// outputUserData(&appUser)
// 	// appUser.outputUserData()
// 	fmt.Println(appUser)

// 	appUser.clearUserName()
// 	fmt.Println(appUser)

// }

// func getUserData(promptText string) string {
// 	fmt.Print(promptText)
// 	var value string
// 	fmt.Scanln(&value)
// 	return value
// }
