package consumerhandler

import (
	"consumer/api/v1/consumerapi"
	"consumer/pkg/domain"
)

func requestToNewConsumer(req *consumerapi.NewConsumerRequest) domain.Consumer {
	return domain.Consumer{
		Firstname: req.Firstname,
		Lastname:  req.Lastname,
		Email:     req.Email,
		Phone:     req.Phone,
	}
}

func requestToUpdateConsumer(req *consumerapi.UpdateConsumerRequest) domain.Consumer {
	return domain.Consumer{
		Firstname: req.Firstname,
		Lastname:  req.Lastname,
		Email:     req.Email,
		Phone:     req.Phone,
	}
}

func requestToNewLocation(req *consumerapi.NewLocationRequest) domain.Location {
	return domain.Location{
		Latitude:   req.Latitude,
		Longitude:  req.Longitude,
		Country:    req.Country,
		City:       req.City,
		Region:     req.Region,
		Street:     req.Street,
		HomeNumber: req.HomeNumber,
		Floor:      req.Floor,
		Door:       req.Door,
	}
}

func requestToUpdateLocation(req *consumerapi.UpdateLocationRequest) domain.Location {
	return domain.Location{
		Latitude:   req.Latitude,
		Longitude:  req.Longitude,
		Country:    req.Country,
		City:       req.City,
		Region:     req.Region,
		Street:     req.Street,
		HomeNumber: req.HomeNumber,
		Floor:      req.Floor,
		Door:       req.Door,
	}
}
