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

func requestToNewConsumerLocation(req *consumerapi.NewConsumerLocationRequest) domain.ConsumerLocation {
	return domain.ConsumerLocation{
		Altitude:   req.Altitude,
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

func requestToUpdateConsumerLocation(req *consumerapi.UpdateConsumerLocationRequest) domain.ConsumerLocation {
	return domain.ConsumerLocation{
		Altitude:   req.Altitude,
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
