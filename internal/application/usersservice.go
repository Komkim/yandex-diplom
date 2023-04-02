package application

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/vzglad-smerti/password_hash"
	"yandex-diplom/internal/domain/aggregate"
	"yandex-diplom/internal/mistake"
)

type UsersService struct {
	UsersRepo aggregate.UsersRepo
}

func NewUsersService(db *pgxpool.Pool) UsersService {
	ur := aggregate.NewUsersRepo(db)
	return UsersService{UsersRepo: ur}
}

func (u *UsersService) Register(login, pass string) error {
	hashPassword, err := password.Hash(pass)
	if err != nil {
		return err
	}

	user, err := u.UsersRepo.GetOne(login)
	if err != nil {
		return err
	}

	if user != nil {
		return mistake.ErrLoginIsTaken
	}

	err = u.UsersRepo.SetOne(login, hashPassword)
	if err != nil {
		return err
	}

	return nil
}
func (u *UsersService) Login(login, pass string) error {
	user, err := u.UsersRepo.GetOne(login)
	if err != nil {
		return err
	}
	if user == nil {
		return mistake.ErrNotAuthenticated
	}

	ok, err := password.Verify(user.HashedPassword, pass)
	if err != nil {
		return err
	}
	if !ok {
		return mistake.ErrNotAuthenticated
	}

	return nil
}
