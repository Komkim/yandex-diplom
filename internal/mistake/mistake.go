package mistake

import "errors"

var (
	NotAuthenticated = errors.New("User not authenticated.")
	LoginIsTaken     = errors.New("Login is taken")
	DbIdError        = errors.New("Uuid is not correct")

	UserNullError = errors.New("zero user value returned from database")

	OrderInvalidNumber              = errors.New("Invalid order number.")
	OrderAlreadyUploadedThisUser    = errors.New("Order number has already been uploaded by this user")
	OrderAlreadyUploadedAnotherUser = errors.New("Order number has already been uploaded by another user")

	BalanceNotEnouhgFunds = errors.New("There are not enough funds on the account")
)
