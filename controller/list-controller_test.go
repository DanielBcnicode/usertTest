package controller

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"usertest.com/user"

	"usertest.com/persistence/memory"
)

func Test_ListUserController(t *testing.T) {
	data := []user.User{
		{
			ID:        uuid.New(),
			FirstName: "a",
			LastName:  "a",
			Nickname:  "a",
			Password:  "a",
			Email:     "a",
			Country:   "a",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			ID:        uuid.New(),
			FirstName: "b",
			LastName:  "b",
			Nickname:  "b",
			Password:  "b",
			Email:     "b",
			Country:   "b",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
	fakeRepo := memory.NewMemoryUserRepository(data)

	service := ListUserController(&fakeRepo)

	req, err := http.NewRequest("GET", "/user/add", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handle := http.HandlerFunc(service)
	handle.ServeHTTP(rr, req)

	res := rr.Result()
	resBody, _ := io.ReadAll(res.Body)
	users := make([]user.User, 0)
	err = json.Unmarshal(resBody, &users)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.Equal(t, 2, len(users))
}
