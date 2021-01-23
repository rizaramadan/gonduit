package users

type RegisterService struct{}

func (r RegisterService) Register(i *RegisterInputWrapper) (UserOutputWrapper, error) {
	u := i.GetUser()
	o := GetUserOutput(u)
	return o, nil
}
