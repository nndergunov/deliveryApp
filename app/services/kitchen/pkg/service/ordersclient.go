package service

import (
	"github.com/nndergunov/deliveryApp/app/services/order/api/v1/orderapi"
)

// OrdersClient interface shows needed signature for the Order Client.
type OrdersClient interface {
	GetIncompleteOrders(id int) (*orderapi.ReturnOrderList, error)
}
