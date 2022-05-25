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
