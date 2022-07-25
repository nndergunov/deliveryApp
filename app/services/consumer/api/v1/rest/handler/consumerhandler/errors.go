package consumerhandler

import "errors"

var (
	errNoConsumerIDParam  = errors.New("consumer_id param is not found")
	errIncorrectInputData = errors.New("incorrect input data")
	systemErr             = errors.New("system error")
)
