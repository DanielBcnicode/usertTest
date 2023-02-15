package user

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"usertest.com/broker"
	"serest.com/user"
	"usertest.com/user/repository"
)

func Test_AddNewUserController(t *testing.T) {
	data := []user.User{
		{
			ID:        "",
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
	fakeRepo := repository.NewMemoryUserRepository(data)
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

	assert.Equal(t, http.StatusOK, res)
	assert.NotNil(t, resBody)

}
