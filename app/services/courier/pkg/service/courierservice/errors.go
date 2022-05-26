package courierservice

import "errors"

var errWrongCourierIDType = errors.New("wrong courier_id type")
var errCourierWithIDNotFound = errors.New("courier with this id not found")
var systemErr = errors.New("system error")
