package user

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID
	FirstName string
	LastName  string
	NickName  string
	Password  string
	Email     string
	Country   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// NewUserFromName create a User entity from a firstname and lastname
func NewUserFromName(firstName, lastName string) User {
	return User{
		ID:        uuid.New(),
		FirstName: firstName,
		LastName:  lastName,
		CreatedAt: time.Now(),
	}
}

func NewUser() User {
	return User{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
	}
}
