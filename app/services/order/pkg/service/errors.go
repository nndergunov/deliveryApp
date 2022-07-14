package service

import "errors"

var (
	// ErrLowBalance returns if the consumer has not enough money to pay for the order.
	ErrLowBalance = errors.New("client has not enough money on their balance")
	// ErrRestaurantOffline returns if the restaurant does not accept orders.
	ErrRestaurantOffline = errors.New("the chosen restaurant is offline and can not accept orders")
)
