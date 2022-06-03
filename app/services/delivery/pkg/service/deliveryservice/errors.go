package deliveryservice

import "errors"

var errWrongFromLocLatType = errors.New("wrong from location latitude type")
var errWrongFromLonLatType = errors.New("wrong from location longitude type")

var errWrongToLocLatType = errors.New("wrong to location latitude type")
var errWrongToLonLatType = errors.New("wrong to location longitude type")

var errWrongLocData = errors.New("wrong location data")

var systemErr = errors.New("system error")

var errWrongOrderIDType = errors.New("wrong order id type")
