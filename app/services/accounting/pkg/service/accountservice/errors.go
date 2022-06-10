package accountservice

import "errors"

var errWrongConsumerIDType = errors.New("wrong consumer_id type")
var errWrongUserID = errors.New("wrong user_id")
var errWrongUserType = errors.New("wrong user_type")
var errAccountNotFound = errors.New("account not found")
var systemErr = errors.New("system error")
var errAccountExist = errors.New("account already exist")
var errMaxNumberOfAccount = errors.New("user has max number of accounts")
var errWrongAmount = errors.New("wrong amount")
var errNotEnoughBalance = errors.New("not enough balance")
