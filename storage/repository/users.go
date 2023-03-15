package storage

type Users interface {
	Register(login, password string) error
	Login(login, password string) error
}
