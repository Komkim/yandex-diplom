package application

type UsersService struct {
}

func NewUsersService() UsersService {
	return UsersService{}
}

func (u *UsersService) Register(login, password string) error {
	return nil
}
func (u *UsersService) Login(login, password string) error {
	return nil
}
