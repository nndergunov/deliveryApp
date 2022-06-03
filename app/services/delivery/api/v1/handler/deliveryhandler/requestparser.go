package deliveryhandler

import (
	"delivery/api/v1/deliveryapi"
	"delivery/pkg/domain"
)

func requestToDeliveryLocation(req *deliveryapi.DeliveryTimeRequest) *domain.DeliveryDistanceLocation {
	return &domain.DeliveryDistanceLocation{
		FromLocation: &domain.Location{
			Latitude:   req.FromLocation.Latitude,
			Longitude:  req.FromLocation.Longitude,
			Country:    req.FromLocation.Country,
			City:       req.FromLocation.City,
			Region:     req.FromLocation.Region,
			Street:     req.FromLocation.Street,
			HomeNumber: req.FromLocation.HomeNumber,
			Floor:      req.FromLocation.Floor,
			Door:       req.FromLocation.Door,
		},
		ToLocation: &domain.Location{
			Latitude:   req.FromLocation.Latitude,
			Longitude:  req.FromLocation.Longitude,
			Country:    req.FromLocation.Country,
			City:       req.FromLocation.City,
			Region:     req.FromLocation.Region,
			Street:     req.FromLocation.Street,
			HomeNumber: req.FromLocation.HomeNumber,
			Floor:      req.FromLocation.Floor,
			Door:       req.FromLocation.Door,
		},
	}
}

func requestToDeliveryCostLocation(req *deliveryapi.DeliveryCostRequest) *domain.DeliveryDistanceLocation {
	return &domain.DeliveryDistanceLocation{
		FromLocation: &domain.Location{
			Latitude:   req.FromLocation.Latitude,
			Longitude:  req.FromLocation.Longitude,
			Country:    req.FromLocation.Country,
			City:       req.FromLocation.City,
			Region:     req.FromLocation.Region,
			Street:     req.FromLocation.Street,
			HomeNumber: req.FromLocation.HomeNumber,
			Floor:      req.FromLocation.Floor,
			Door:       req.FromLocation.Door,
		},
		ToLocation: &domain.Location{
			Latitude:   req.FromLocation.Latitude,
			Longitude:  req.FromLocation.Longitude,
			Country:    req.FromLocation.Country,
			City:       req.FromLocation.City,
			Region:     req.FromLocation.Region,
			Street:     req.FromLocation.Street,
			HomeNumber: req.FromLocation.HomeNumber,
			Floor:      req.FromLocation.Floor,
			Door:       req.FromLocation.Door,
		},
	}
}

func requestToOrder(req *deliveryapi.OrderRequest) *domain.Order {
	return &domain.Order{FromUserID: req.FromUserID, FromRestaurantID: req.RestaurantID}
}
