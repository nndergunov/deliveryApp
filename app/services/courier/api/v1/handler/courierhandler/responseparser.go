package courierhandler

import (
	"courier/api/v1/courierapi"
	"courier/pkg/domain"
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

func courierListToResponse(courierList []domain.Courier) courierapi.ReturnCourierResponseList {
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
	return courierapi.ReturnCourierResponseList{
		CourierResponseList: courierResponseList,
	}
}

func courierLocationToResponse(courierLocation domain.CourierLocation) courierapi.CourierLocationResponse {
	return courierapi.CourierLocationResponse{
		CourierID:  courierLocation.CourierID,
		Altitude:   courierLocation.Altitude,
		Longitude:  courierLocation.Longitude,
		Country:    courierLocation.Country,
		City:       courierLocation.City,
		Region:     courierLocation.Region,
		Street:     courierLocation.Street,
		HomeNumber: courierLocation.HomeNumber,
		Floor:      courierLocation.Floor,
		Door:       courierLocation.Door,
	}
}
