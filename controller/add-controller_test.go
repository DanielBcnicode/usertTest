package controller

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"usertest.com/broker"
	"usertest.com/user"

	"usertest.com/persistence/memory"
)

func Test_AddNewUserController(t *testing.T) {
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
	}
	fakeRepo := memory.NewMemoryUserRepository(data)
	fakeQueue := broker.NewFakeMemoryQueue()
	service := AddNewUserController(&fakeRepo, fakeQueue)

	req, err := http.NewRequest("POST", "/user/add", strings.NewReader(`
	{
		"first_name": "Name",
		"last_name" : "LastName",
		"nickname": "test",
		"password": "password",
		"email": "test@test.com",
		"country": "ES"
	}
	`))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handle := http.HandlerFunc(service)
	handle.ServeHTTP(rr, req)

	res := rr.Result()
	resBody, _ := io.ReadAll(res.Body)
	u := user.User{}
	_ = json.Unmarshal(resBody, &u)

	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.Equal(t, "Name", u.FirstName)
	id := u.ID.String()
	// Check the uuid generated length 
	assert.True(t, len(id) == 36)
	// Check if the element is stored
	assert.Equal(t, 2, len(fakeRepo.Data))
 	// Test if the message was send
	assert.Equal(t, 1, len(fakeQueue.Queue))

}
