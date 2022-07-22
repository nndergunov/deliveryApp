package deliveryhandler

import "errors"

var (
	errNoOrderIDParam     = errors.New("order_id param is not found")
	errIncorrectInputData = errors.New("incorrect input data")
	systemErr             = errors.New("system error")
)
