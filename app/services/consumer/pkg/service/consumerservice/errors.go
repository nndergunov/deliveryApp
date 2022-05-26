package consumerservice

import "errors"

var errWrongConsumerIDType = errors.New("wrong consumer_id type")
var errConsumerWithIDNotFound = errors.New("consumer with this id not found")
var systemErr = errors.New("system error")
