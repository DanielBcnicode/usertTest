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
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"usertest.com/broker"
	"usertest.com/user"

	"usertest.com/persistence/memory"
)

func Test_UpdateUserController(t *testing.T) {
	uid := uuid.New()
	data := []user.User{
		{
			ID:        uid,
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
	service := UpdateUserController(&fakeRepo, fakeQueue)

	req, err := http.NewRequest("PUT", "api/v1/user/" + uid.String(), strings.NewReader(`
	{
		"first_name": "NameUpdated",
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

	vars := map[string]string{"id": uid.String()}
	req = mux.SetURLVars(req, vars)

	rr := httptest.NewRecorder()
	handle := http.HandlerFunc(service)
	handle.ServeHTTP(rr, req)

	res := rr.Result()
	resBody, _ := io.ReadAll(res.Body)
	u := user.User{}
	_ = json.Unmarshal(resBody, &u)

	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.Equal(t, "NameUpdated", u.FirstName)
	id := u.ID.String()
	// Check the uuid generated length 
	assert.True(t, len(id) == 36)
	// Check if the element is stored
	assert.Equal(t, 1, len(fakeRepo.Data))
 	// Test if the message was send
	assert.Equal(t, 1, len(fakeQueue.Queue))
}
