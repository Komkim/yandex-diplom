package mistake

import "errors"

var (
	InvalidOrderNumber = errors.New("Invalid order number.")
	NotAuthenticated   = errors.New("User not authenticated.")
	LoginIsTaken       = errors.New("Login is taken")
	DbIdError          = errors.New("Uuid is not correct")
)
