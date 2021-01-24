package users

import "fmt"

//RegisterService is to serve register
type RegisterService struct {
	UserRepo UserRepo
}

func NewRegisterService(u UserRepo) *RegisterService {
	fmt.Println("NewRegisterService created")
	r := new(RegisterService)
	r.UserRepo = u
	return r
}

//Register is to register a user
func (r *RegisterService) Register(i *RegisterInputWrapper) (UserOutputWrapper, error) {
	u := i.GetUser()
	err := r.UserRepo.Create(&u)
	if err != nil {
		return UserOutputWrapper{}, err
	}
	o := GetUserOutput(u)
	return o, nil
}
