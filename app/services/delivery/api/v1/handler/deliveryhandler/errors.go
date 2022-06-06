package deliveryhandler

import "errors"

var errNoOrderIDParam = errors.New("order_id param is not found")
var errIncorrectInputData = errors.New("incorrect input data")
var systemErr = errors.New("system error")
