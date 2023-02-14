package user

import (
	"context"
	"time"

	"github.com/google/uuid"
)

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

func NewUserFromModel(
	id uuid.UUID,
	firstName string,
	lastName string,
	nickName string,
	password string,
	email string,
	country string,
	createdAt time.Time,
	updateAt time.Time,
) (User, error) {
	return User{
		ID:        id,
		FirstName: firstName,
		LastName:  lastName,
		Nickname:  nickName,
		Password:  password,
		Email:     email,
		CreatedAt: createdAt,
		UpdatedAt: updateAt,
	}, nil
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

func (u *User) UpdateDate() {
	u.UpdatedAt = time.Now()
}

type RepositoryFilter struct {
	Filters map[string]string
}

type Paginator struct {
	PagSize     int
	CurrentPage int
	NextPage    int
}

type UserRepo interface {
	Save(ctx context.Context, user *User) error
	Update(ctx context.Context, user *User) error
	FindByFilter(ctx context.Context, filter RepositoryFilter, paginator *Paginator) ([]User, error)
	Delete(ctx context.Context, userId uuid.UUID) error
}
