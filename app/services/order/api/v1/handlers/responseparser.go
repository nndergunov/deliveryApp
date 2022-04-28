package handlers

import (
	"net/http"

	v1 "github.com/nndergunov/deliveryApp/app/pkg/api/v1"
	"github.com/nndergunov/deliveryApp/app/pkg/api/v1/orderapi"
	"github.com/nndergunov/deliveryApp/app/services/order/pkg/domain"
)

func (e endpointHandler) respond(response any, w http.ResponseWriter) {
	data, err := v1.Encode(response)
	if err != nil {
		e.log.Println(err)

		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	w.WriteHeader(http.StatusOK)

	_, err = w.Write(data)
	if err != nil {
		e.log.Println(err)

		w.WriteHeader(http.StatusInternalServerError)

		return
	}
}

func orderToResponse(order domain.Order) orderapi.ReturnOrder {
	return orderapi.ReturnOrder{
		OrderID:      order.OrderID,
		FromUserID:   order.FromUserID,
		RestaurantID: order.RestaurantID,
		OrderItems:   order.OrderItems,
		Status:       order.Status,
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
