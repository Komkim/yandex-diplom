package request

import "net/http"

type User struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func (u *User) Bind(r *http.Request) error {
	return nil
}
