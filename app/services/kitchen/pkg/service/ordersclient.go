package service

import (
	"github.com/nndergunov/deliveryApp/app/services/order/api/v1/communication"
)

type OrdersClient interface {
	GetIncompleteOrders(id int) (*communication.ReturnOrderList, error)
}
