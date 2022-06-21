package deliveryservice

import "errors"

var errWrongLocData = errors.New("wrong location data")

var systemErr = errors.New("system error")

var errWrongOrderIDType = errors.New("wrong order id type")
