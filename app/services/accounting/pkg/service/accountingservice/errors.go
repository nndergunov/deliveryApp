package accountingservice

import "errors"

var (
	errWrongUserID        = errors.New("wrong user_id")
	errWrongUserType      = errors.New("wrong user_type")
	errAccountNotFound    = errors.New("account not found")
	systemErr             = errors.New("system error")
	errMaxNumberOfAccount = errors.New("user has max number of accounts")
	errWrongAmount        = errors.New("wrong amount")
)
