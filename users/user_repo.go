package users

type UserRepo interface {
	Create(u *User) error
}
