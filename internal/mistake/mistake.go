package mistake

import "errors"

var (
	ErrNotAuthenticated = errors.New("user not authenticated")
	ErrLoginIsTaken     = errors.New("login is taken")
	ErrDbId             = errors.New("uuid is not correct")

	ErrUserNullError = errors.New("zero user value returned from database")

	ErrOrderInvalidNumber              = errors.New("invalid order number")
	ErrOrderAlreadyUploadedThisUser    = errors.New("order number has already been uploaded by this user")
	ErrOrderAlreadyUploadedAnotherUser = errors.New("order number has already been uploaded by another user")

	ErrBalanceNotEnouhgFunds = errors.New("tThere are not enough funds on the account")
)
