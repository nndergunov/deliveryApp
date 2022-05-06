package handlers

import (
	"github.com/nndergunov/deliveryApp/app/pkg/api/v1/orderapi"
	"github.com/nndergunov/deliveryApp/app/services/order/pkg/domain"
)

func orderToResponse(order domain.Order) orderapi.ReturnOrder {
	return orderapi.ReturnOrder{
		OrderID:      order.OrderID,
		FromUserID:   order.FromUserID,
		RestaurantID: order.RestaurantID,
		OrderItems:   order.OrderItems,
		Status:       order.Status.Status,
	}
}

func orderListToResponse(orderList []domain.Order) orderapi.ReturnOrderList {
	returnOrderList := make([]orderapi.ReturnOrder, 0, len(orderList))

	for _, el := range orderList {
		returnOrderList = append(returnOrderList, orderToResponse(el))
	}

	return orderapi.ReturnOrderList{
		Orders: returnOrderList,
	}
}
