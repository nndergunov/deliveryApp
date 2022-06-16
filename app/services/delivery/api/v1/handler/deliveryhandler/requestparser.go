package deliveryhandler

import (
	"github.com/nndergunov/deliveryApp/app/pkg/api/v1/deliveryapi"

	"delivery/pkg/domain"
)

func requestToEstimateDelivery(req *deliveryapi.EstimateDeliveryRequest) *domain.EstimateDeliveryRequest {
	return &domain.EstimateDeliveryRequest{
		ConsumerID:   req.ConsumerID,
		RestaurantID: req.RestaurantID,
	}
}

func requestToOrder(req *deliveryapi.AssignOrderRequest) *domain.Order {
	return &domain.Order{FromUserID: req.FromUserID, FromRestaurantID: req.RestaurantID}
}
