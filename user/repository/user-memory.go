package repository

import (
	"context"
	"errors"
	"log"
	"sync"

	"github.com/google/uuid"
	"usertest.com/user"
)

type MemoryUserRepository struct {
	Data []user.User
	mu   sync.Mutex
}

func NewMemoryUserRepository(data []user.User) MemoryUserRepository {
	return MemoryUserRepository{Data: data}
}

func (u *MemoryUserRepository) Save(ctx context.Context, user *user.User) error {
	u.mu.Lock()
	u.Data = append(u.Data, *user)
	u.mu.Unlock()
	
	return nil
}

func (u *MemoryUserRepository) Update(ctx context.Context, user *user.User) error {
	u.mu.Lock()
	defer u.mu.Unlock()
	po := -1

	for i := 0; i < len(u.Data); i++ {
		if u.Data[i].ID == user.ID {
			po = i
			break
		}
	}

	if po == -1 {
		log.Printf("ERROR: user with id = %s not found\n", user.ID)
		return errors.New("user to update not found in database")
	}

	u.Data[po] = *user

	return nil
}

func (u *MemoryUserRepository) Delete(ctx context.Context, userId uuid.UUID) error {
	u.mu.Lock()
	defer u.mu.Unlock()
	po := -1

	for i := 0; i < len(u.Data); i++ {
		if u.Data[i].ID == userId {
			po = i
			break
		}
	}

	if po == -1 {
		log.Printf("ERROR: user with id = %s not found\n", userId)
		return errors.New("user to delete not found in database")
	}

	ret := make([]user.User, 0)
    ret = append(ret, u.Data[:po]...)
    u.Data = append(ret, u.Data[po+1:]...)

	return nil
}

// FindByFilter returns the users using pagination, in this memory repository implementation the filter
// is avoided. 
func (u *MemoryUserRepository) FindByFilter(ctx context.Context, filter user.RepositoryFilter, paginator *user.Paginator) ([]user.User, error) {
	u.mu.Lock()
	defer u.mu.Unlock()
	
	d := make([]user.User, 0)
	limit := 10
	offset := 0
	if paginator != nil {
		if paginator.PagSize > 0 {
			limit = paginator.PagSize
		}
		offset = limit * paginator.CurrentPage
	}
	
	m := len(u.Data)
	if offset > m {
		return d, nil
	}
	
	t := limit
	if offset + limit > m {
		t = m - offset
	}

	for i := 0; i < t; i++ {
		d = append(d, u.Data[offset + i])
	}
	
	return d, nil
}
