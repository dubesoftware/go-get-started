package models

import (
	"fmt"

	"github.com/pkg/errors"
)

// User is a struct containing user information.
type User struct {
	ID        int
	FirstName string
	LastName  string
}

var (
	users  []*User
	nextID = 1
)

// GetUsers returns a slice of User object pointers.
func GetUsers() []*User {
	return users
}

// AddUser adds a user object pointer to the users slice.
func AddUser(user User) (User, error) {
	if user.ID != 0 {
		return User{}, errors.New("new user must not include ID or ID must be set to 0")
	}
	user.ID = nextID
	nextID++
	users = append(users, &user)
	return user, nil
}

// GetUserByID returns a user object by ID
func GetUserByID(id int) (User, error) {
	for _, user := range users {
		if user.ID == id {
			return *user, nil
		}
	}
	return User{}, fmt.Errorf("user with ID '%v' not found", id)
}

// UpdateUserByID updates a user object by ID
func UpdateUserByID(user User) (User, error) {
	for i, u := range users {
		if u.ID == user.ID {
			users[i] = &user
			return user, nil
		}
	}
	return User{}, fmt.Errorf("user with ID '%v' not found", user.ID)
}

// RemoveUserByID returns a user object by ID
func RemoveUserByID(id int) error {
	for i, user := range users {
		if user.ID == id {
			users = append(users[:i], users[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("user with ID '%v' not found", id)
}
