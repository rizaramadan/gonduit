package users

import "testing"

func TestGetUser(t *testing.T) {
	i := RegisterInput{
		Email:    "a@a.com",
		Password: "123",
		Username: "a",
	}
	w := RegisterInputwrapper{
		User: i,
	}
	u := w.GetUser()
	if i.Email != u.Email || i.Password != u.Password || i.Username != u.Username {
		t.Errorf("get user from register input wrapper fail")
	}
}
