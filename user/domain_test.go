package user

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	t.Run("Constructor and UpdateData test", func(t *testing.T) {
		got := NewUser()
		assert.True(t, uuid.Nil != got.ID)
		assert.Greater(t, got.CreatedAt.Unix(), int64(0))
		assert.Equal(t, got.UpdatedAt.Unix(), got.CreatedAt.Unix())
		tm := time.Now().Add(-time.Hour)

		// Put CreatedAt and UpdateAt in the past
		got.CreatedAt = tm  
		got.UpdatedAt = tm
		got.UpdateDate() 
		assert.NotEqual(t, got.UpdatedAt.Unix(), got.CreatedAt.Unix())
	})

}
