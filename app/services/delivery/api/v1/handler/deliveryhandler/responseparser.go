package deliveryhandler

import (
	"delivery/api/v1/deliveryapi"
	"delivery/pkg/domain"
)

func estimateDeliveryToResponse(estimateDelivery *domain.EstimateDeliveryResponse) deliveryapi.EstimateDeliveryResponse {
	return deliveryapi.EstimateDeliveryResponse{Time: estimateDelivery.Time, Cost: estimateDelivery.Cost}
}

func assignOrderResponse(assignOrderToCourier *domain.AssignOrder) deliveryapi.AssignOrderResponse {
	return deliveryapi.AssignOrderResponse{OrderID: assignOrderToCourier.OrderID, CourierID: assignOrderToCourier.CourierID}
}
