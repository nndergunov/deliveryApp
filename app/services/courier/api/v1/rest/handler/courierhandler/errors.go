package courierhandler

import "errors"

var (
	errNoCourierIDParam   = errors.New("courier_id param is not found")
	errNoAvailableParam   = errors.New("available param is not found")
	errIncorrectInputData = errors.New("incorrect input data")
	systemErr             = errors.New("system error")
)
