package auth

type AuthInterface interface {
	CreateAuth(login, pass string) string
	FetchAuth(token string) (string, bool)
	DeleteAuth(token string)
}
