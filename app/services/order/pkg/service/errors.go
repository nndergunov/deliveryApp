package service

import "errors"

var (
	ErrLowBalance        = errors.New("client has not enough money on their balance")
	ErrRestaurantOffline = errors.New("the chosen restaurant is offline and can not accept orders")
)
