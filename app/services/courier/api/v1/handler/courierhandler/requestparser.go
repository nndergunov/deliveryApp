package courierhandler

import (
	"github.com/nndergunov/deliveryApp/app/pkg/api/v1/courierapi"

	"github.com/nndergunov/deliveryApp/app/services/courier/pkg/domain"
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

func requestToNewLocation(req *courierapi.NewLocationRequest) domain.Location {
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

func requestToUpdateLocation(req *courierapi.UpdateLocationRequest) domain.Location {
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
