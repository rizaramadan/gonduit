package users

//RegisterInput is type for input
type RegisterInput struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

//RegisterInputwrapper is type represent register input
type RegisterInputwrapper struct {
	User RegisterInput `json:"user"`
}

//GetUser it a method to return user from register input
func (r RegisterInputwrapper) GetUser() User {
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
