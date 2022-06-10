package accounthandler

import "errors"

var errNoAccountIDParam = errors.New("account_id param is not found")
var errIncorrectInputData = errors.New("incorrect input data")
var systemErr = errors.New("system error")
var errAccountNotFound = errors.New("account not found")
