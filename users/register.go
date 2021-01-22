package users

//RegisterInput is type represent register input
type RegisterInput struct {
	User struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	} `json:"user"`
}

//GetUser it a method to return user from register input
func (r RegisterInput) GetUser() User {
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
