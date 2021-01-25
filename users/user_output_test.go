package users

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetUserOutput(t *testing.T) {
	i := User{
		Username: "a",
		Token:    "123",
		Email:    "a@a.com",
		Bio:      "bio",
		Image:    "image",
	}
	w := GetUserOutput(i)
	assert.Equal(t, i.Email, w.User.Email)
	assert.Equal(t, i.Username, w.User.Username)
	assert.Equal(t, i.Token, w.User.Token)
	assert.Equal(t, i.Bio, w.User.Bio)
	assert.Equal(t, i.Image, w.User.Image)
}
