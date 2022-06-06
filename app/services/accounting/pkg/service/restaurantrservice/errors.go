package restaurantservice

import "errors"

var errWrongRestaurantIDType = errors.New("wrong restaurant_id type")
var errWrongRestaurantID = errors.New("wrong restaurant id")
var errRestaurantAccountNotFound = errors.New("restaurant account not found")
var systemErr = errors.New("system error")
var errRestaurantAccountExist = errors.New("restaurant account already exist")
var errWrongAmount = errors.New("wrong amount")
var errNotEnoughBalance = errors.New("not enough balance")
