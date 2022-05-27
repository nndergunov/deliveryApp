package courierhandler

import (
	"courier/api/v1/courierapi"
	"courier/pkg/domain"
)

func requestToNewCourier(req *courierapi.NewCourierRequest) domain.Courier {
	return domain.Courier{
		Username:  req.Username,
		Password:  req.Password,
		Firstname: req.Firstname,
		Lastname:  req.Lastname,
		Email:     req.Email,
		Phone:     req.Phone,
	}
}

func requestToUpdateCourier(req *courierapi.UpdateCourierRequest) domain.Courier {
	return domain.Courier{
		Username:  req.Username,
		Firstname: req.Firstname,
		Lastname:  req.Lastname,
		Email:     req.Email,
		Phone:     req.Phone,
	}
}

func requestToNewCourierLocation(req *courierapi.NewCourierLocationRequest) domain.CourierLocation {
	return domain.CourierLocation{
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

func requestToUpdateConsumerLocation(req *courierapi.UpdateCourierLocationRequest) domain.CourierLocation {
	return domain.CourierLocation{
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