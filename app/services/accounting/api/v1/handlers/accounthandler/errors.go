package accounthandler

import "errors"

var (
	errNoAccountIDParam   = errors.New("account_id param is not found")
	errIncorrectInputData = errors.New("incorrect input data")
	systemErr             = errors.New("system error")
	errAccountNotFound    = errors.New("account not found")
)
