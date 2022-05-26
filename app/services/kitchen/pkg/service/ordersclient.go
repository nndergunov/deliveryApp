package service

import "github.com/nndergunov/deliveryApp/app/pkg/api/v1/orderapi"

type OrdersClient interface {
	GetIncompleteOrders(id int) (*orderapi.ReturnOrderList, error)
}
