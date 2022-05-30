package courierservice

import "errors"

var errWrongCourierIDType = errors.New("wrong courier_id type")
var errCourierAccountNotFound = errors.New("courier account not found")
var systemErr = errors.New("system error")
var errCourierAccountExist = errors.New("courier account already exist")
var errWrongAmount = errors.New("wrong amount")
var errNotEnoughBalance = errors.New("not enough balance")
