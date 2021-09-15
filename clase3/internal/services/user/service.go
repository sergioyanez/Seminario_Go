package user

type User struct {
	Name  string
	Login string
	Pass  string
}
type UserService interface {
	Login(user, pass string) (*User, error)
}

type userSrv struct {
}

func (us *userSrv) Login(user, pass string) (*User, error) {
	return &User{"Sergio", "sergio", "123456"}, nil
}

func NewUserService() (UserService, error) {
	return &userSrv{}, nil
}
