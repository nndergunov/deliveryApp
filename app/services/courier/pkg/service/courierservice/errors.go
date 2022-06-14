package courierservice

import "errors"

var (
	errWrongCourierIDType    = errors.New("wrong courier_id type")
	errCourierWithIDNotFound = errors.New("courier with this id not found")
	systemErr                = errors.New("system error")
)
