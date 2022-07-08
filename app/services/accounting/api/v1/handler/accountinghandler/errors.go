package accountinghandler

import "errors"

var (
	errNoIDParam          = errors.New("id param not found")
	errIncorrectInputData = errors.New("incorrect input data")
	systemErr             = errors.New("system error")
	errAccountNotFound    = errors.New("account not found")
)
