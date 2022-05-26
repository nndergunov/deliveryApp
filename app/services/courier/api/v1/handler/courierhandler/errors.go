package courierhandler

import "errors"

var errNoCourierIDParam = errors.New("courier_id param is not found")
var errNoAvailableParam = errors.New("available param is not found")
var errIncorrectInputData = errors.New("incorrect input data")
var systemErr = errors.New("system error")
