package user

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id,omitempty"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	NickName  string    `json:"nickname"`
	Password  string    `json:"password"`
	Email     string    `json:"email"`
	Country   string    `json:"country"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
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
	t := time.Now()
	return User{
		ID:        uuid.New(),
		CreatedAt: t,
		UpdatedAt: t,
	}
}
