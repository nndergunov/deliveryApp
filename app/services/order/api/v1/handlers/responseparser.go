package handlers

import (
	v1 "github.com/nndergunov/deliveryApp/app/services/order/api/v1/communications"
	"github.com/nndergunov/deliveryApp/app/services/order/pkg/domain"
)

func orderToResponse(order domain.Order) v1.ReturnOrder {
	return v1.ReturnOrder{
		OrderID:      order.OrderID,
		FromUserID:   order.FromUserID,
		RestaurantID: order.RestaurantID,
		OrderItems:   order.OrderItems,
		Status:       order.Status.Status,
	}
}

func orderListToResponse(orderList []domain.Order) v1.ReturnOrderList {
	returnOrderList := make([]v1.ReturnOrder, 0, len(orderList))

	for _, el := range orderList {
		returnOrderList = append(returnOrderList, orderToResponse(el))
	}

	return v1.ReturnOrderList{
		Orders: returnOrderList,
	}
}
