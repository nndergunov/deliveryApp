package accountservice

import "errors"

var errWrongUserID = errors.New("wrong user_id")
var errWrongUserType = errors.New("wrong user_type")
var errAccountNotFound = errors.New("account not found")
var systemErr = errors.New("system error")
var errMaxNumberOfAccount = errors.New("user has max number of accounts")
var errWrongAmount = errors.New("wrong amount")
