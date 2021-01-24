package users

import "errors"

//RegisterInput is type for input
type RegisterInput struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

//RegisterInputWrapper is type represent register input
type RegisterInputWrapper struct {
	User RegisterInput `json:"user"`
}

//GetUser it a method to return user from register input
func (r RegisterInputWrapper) GetUser() User {
	user := User{
		Username: r.User.Username,
		Email:    r.User.Email,
		Password: r.User.Password,
		Token:    "",
		Bio:      "",
		Image:    "",
	}
	return user
}

//Validate is to validate RegisterInputWrapper
func (r RegisterInputWrapper) Validate() error {
	if r.User.Email == "" || r.User.Password == "" || r.User.Username == "" {
		return errors.New("one of the mandatory register input is missing")
	}
	return nil
}
