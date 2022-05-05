package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/nndergunov/deliveryApp/app/pkg/api/v1/orderapi"
	"github.com/nndergunov/deliveryApp/app/services/order/pkg/domain"
)

func getIDFromEndpoint(code string, r *http.Request) (int, error) {
	vars := mux.Vars(r)

	idVar := vars[code]
	if idVar == "" {
		return 0, fmt.Errorf("getIDFromEndpoint: %w", errNoIDInEndpoint)
	}

	id, err := strconv.Atoi(idVar)
	if err != nil {
		return 0, fmt.Errorf("getIDFromEndpoint: %w", err)
	}

	return id, nil
}

func requestToOrder(orderData orderapi.OrderData) domain.Order {
	return domain.Order{
		OrderID:      0,
		FromUserID:   orderData.FromUserID,
		RestaurantID: orderData.RestaurantID,
		OrderItems:   orderData.OrderItems,
		Status:       "",
	}
}
