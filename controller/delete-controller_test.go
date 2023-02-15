package controller

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"usertest.com/broker"
	"usertest.com/user"

	"usertest.com/persistence/memory"
)

func Test_DeleteUserController(t *testing.T) {
	suuid := uuid.New()
	data := []user.User{
		{
			ID:        suuid,
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
	service := DeleteUserController(&fakeRepo, fakeQueue)

	req, err := http.NewRequest("DELETE", "api/v1/user/"+suuid.String(), nil)
	if err != nil {
		t.Fatal(err)
	}
	vars := map[string]string{"id":suuid.String()}
	req = mux.SetURLVars(req, vars)

	rr := httptest.NewRecorder()
	handle := http.HandlerFunc(service)
	handle.ServeHTTP(rr, req)

	res := rr.Result()

	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.Equal(t, 0, len(fakeRepo.Data)) // Test if the row is deleted
	assert.Equal(t, 1, len(fakeQueue.Queue)) // Test if the message was send

}
