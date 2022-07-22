package deliveryhandler

import (
	"github.com/nndergunov/deliveryApp/app/services/delivery/api/v1/rest/deliveryapi"
	"github.com/nndergunov/deliveryApp/app/services/delivery/pkg/domain"
)

func requestToOrder(req *deliveryapi.AssignOrderRequest) *domain.Order {
	return &domain.Order{FromUserID: req.FromUserID, FromRestaurantID: req.RestaurantID}
}
