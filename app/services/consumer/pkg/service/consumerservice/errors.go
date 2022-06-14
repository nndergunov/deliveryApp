package consumerservice

import "errors"

var (
	errWrongConsumerIDType    = errors.New("wrong consumer_id type")
	errConsumerWithIDNotFound = errors.New("consumer with this id not found")
	systemErr                 = errors.New("system error")
)
