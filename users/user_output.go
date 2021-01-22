package users

//UserOutput is inner type to represent user output
type UserOutput struct {
	Username string      `json:"username"`
	Email    string      `json:"email"`
	Token    string      `json:"token"`
	Bio      string      `json:"bio"`
	Image    interface{} `json:"image"`
}

//UserOutputWrapper is type to represent user output
type UserOutputWrapper struct {
	User UserOutput `json:"user"`
}

//GetUserOutput create UserOutputWrapper from user
func GetUserOutput(u User) UserOutputWrapper {
	user := UserOutput{
		Username: u.Username,
		Email:    u.Email,
		Token:    u.Token,
		Bio:      u.Bio,
		Image:    u.Image,
	}
	wrapper := UserOutputWrapper{
		User: user,
	}
	return wrapper
}
