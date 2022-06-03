package deliveryhandler

import (
	"delivery/api/v1/deliveryapi"
	"delivery/pkg/domain"
)

func deliveryTimeToResponse(deliveryTime *domain.DeliveryTime) deliveryapi.DeliveryTimeResponse {
	return deliveryapi.DeliveryTimeResponse{Time: deliveryTime.Time}
}

func deliveryCostToResponse(deliveryCost *domain.DeliveryCost) deliveryapi.DeliveryCostResponse {
	return deliveryapi.DeliveryCostResponse{Cost: deliveryCost.Cost}
}

func deliveryAssignedCourierResponse(assignOrderToCourier *domain.AssignedCourier) deliveryapi.DeliveryAssignedCourierResponse {
	return deliveryapi.DeliveryAssignedCourierResponse{OrderID: assignOrderToCourier.OrderID, CourierID: assignOrderToCourier.CourierID}
}
