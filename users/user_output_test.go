package users

import "testing"

func TestGetUserOutput(t *testing.T) {
	i := User{
		Username: "a",
		Token:    "123",
		Email:    "a@a.com",
		Bio:      "bio",
		Image:    "image",
	}
	w := GetUserOutput(i)
	if i.Username != w.User.Username || i.Email != w.User.Email || i.Token != w.User.Token || i.Bio != w.User.Bio || i.Image != w.User.Image {
		t.Errorf("get user from register input wrapper fail")
	}
}
