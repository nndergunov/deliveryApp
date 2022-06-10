package accounthandler

import "errors"

var errNoConsumerIDParam = errors.New("consumer_id param is not found")
var errNoUserIDParam = errors.New("user_id param is not found")
var errNoUserTypeParam = errors.New("user_type param is not found")
var errIncorrectInputData = errors.New("incorrect input data")
var systemErr = errors.New("system error")
var errAccountNotFound = errors.New("account not found")
