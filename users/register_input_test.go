package users

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetUser(t *testing.T) {
	i := RegisterInput{
		Email:    "a@a.com",
		Password: "123",
		Username: "a",
	}
	w := RegisterInputWrapper{
		User: i,
	}
	u := w.GetUser()
	assert.Equal(t, i.Email, u.Email)
	assert.Equal(t, i.Password, u.Password)
	assert.Equal(t, i.Username, u.Username)
}
