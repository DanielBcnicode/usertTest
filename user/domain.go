package user

import (
	"context"
	"time"

	"github.com/google/uuid"
)

// User main entity for a user
type User struct {
	ID        uuid.UUID `json:"id,omitempty"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Nickname  string    `json:"nickname"`
	Password  string    `json:"password"`
	Email     string    `json:"email"`
	Country   string    `json:"country"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// User constructor
func NewUser() User {
	t := time.Now()
	return User{
		ID:        uuid.New(),
		CreatedAt: t,
		UpdatedAt: t,
	}
}

// UpdateDate update the UpdatedAt field with now
func (u *User) UpdateDate() {
	u.UpdatedAt = time.Now()
}

// RepositoryFilter holds the url filters
type RepositoryFilter struct {
	Filters map[string]string
}

// Paginator data structure to handle the pages in queries
type Paginator struct {
	PagSize     int
	CurrentPage int
}

// UserRepo is the contract to the user repository
type UserRepo interface {
	Save(ctx context.Context, user *User) error
	Update(ctx context.Context, user *User) error
	FindByFilter(ctx context.Context, filter RepositoryFilter, paginator *Paginator) ([]User, error)
	Delete(ctx context.Context, userId uuid.UUID) error
}
