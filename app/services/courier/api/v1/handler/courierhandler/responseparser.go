package courierhandler

import (
	"courier/api/v1/courierapi"
	"courier/domain"
)

func courierToResponse(courier domain.Courier) courierapi.CourierResponse {
	return courierapi.CourierResponse{
		ID:        courier.ID,
		Username:  courier.Username,
		Firstname: courier.Firstname,
		Lastname:  courier.Lastname,
		Email:     courier.Email,
		Phone:     courier.Phone,
		Available: courier.Available,
	}
}

func courierListToResponse(courierList []domain.Courier) courierapi.ReturnCourierList {
	courierResponseList := make([]courierapi.CourierResponse, 0, len(courierList))

	for _, courier := range courierList {
		courierResponse := courierapi.CourierResponse{
			ID:        courier.ID,
			Username:  courier.Username,
			Firstname: courier.Firstname,
			Lastname:  courier.Lastname,
			Email:     courier.Email,
			Phone:     courier.Phone,
			Available: courier.Available,
		}

		courierResponseList = append(courierResponseList, courierResponse)
	}
	return courierapi.ReturnCourierList{
		CourierList: courierResponseList,
	}
}
