// user.go
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
	Created   time.Time
}

type Admin struct {
	User
	Email    string
	Password string
}

func (u User) String() string {
	return fmt.Sprintf("user{firstName: %s, lastName: %s, birthdate: %s, created: %v}",
		u.FirstName, u.LastName, u.Birthdate, u.Created)
}

func (a *Admin) String() string {
	return fmt.Sprintf("admin{user: %s, email: %s, password: %s}",
		a.User.String(), a.Email, a.Password)
}

func NewUser(firstName, lastName, birthdate string) (*User, error) {
	if firstName == "" {
		return nil, errors.New("first name cannot be empty")
	}
	if lastName == "" {
		return nil, errors.New("last name cannot be empty")
	}
	if birthdate == "" {
		return nil, errors.New("birthdate cannot be empty")
	}
	return &User{
		FirstName: firstName,
		LastName:  lastName,
		Birthdate: birthdate,
		Created:   time.Now(),
	}, nil
}

func NewAdmin(firstName, lastName, birthdate, email, password string) (*Admin, error) {
	user, err := NewUser(firstName, lastName, birthdate)
	if err != nil {
		return nil, err
	}
	if email == "" {
		return nil, errors.New("email cannot be empty")
	}
	if password == "" {
		return nil, errors.New("password cannot be empty")
	}
	return &Admin{
		User:     *user,
		Email:    email,
		Password: password,
	}, nil
}

func (u *User) ClearUserName() {
	u.FirstName = ""
	u.LastName = ""
}
