package consumeraccountingservice

import "errors"

var errWrongConsumerIDType = errors.New("wrong consumer_id type")
var errConsumerAccountNotFound = errors.New("consumer account not found")
var systemErr = errors.New("system error")
var errConsumerAccountExist = errors.New("consumer account already exist")
var errWrongAmount = errors.New("wrong amount")
var errNotEnoughBalance = errors.New("not enough balance")
