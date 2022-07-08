package courierhandler

import (
	"github.com/nndergunov/deliveryApp/app/services/courier/api/v1/courierapi"
	"github.com/nndergunov/deliveryApp/app/services/courier/pkg/domain"
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

func courierListToResponse(courierList []domain.Courier) courierapi.CourierResponseList {
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
	return courierapi.CourierResponseList{
		CourierResponseList: courierResponseList,
	}
}

func locationToResponse(location domain.Location) courierapi.LocationResponse {
	return courierapi.LocationResponse{
		UserID:     location.UserID,
		Latitude:   location.Latitude,
		Longitude:  location.Longitude,
		Country:    location.Country,
		City:       location.City,
		Region:     location.Region,
		Street:     location.Street,
		HomeNumber: location.HomeNumber,
		Floor:      location.Floor,
		Door:       location.Door,
	}
}

func locationListToResponse(locationList []domain.Location) courierapi.LocationResponseList {
	locationResponseList := make([]courierapi.LocationResponse, 0, len(locationList))

	for _, location := range locationList {
		locationResponse := courierapi.LocationResponse{
			UserID:     location.UserID,
			Latitude:   location.Latitude,
			Longitude:  location.Longitude,
			Country:    location.Country,
			City:       location.City,
			Region:     location.Region,
			Street:     location.Street,
			HomeNumber: location.HomeNumber,
			Floor:      location.Floor,
			Door:       location.Door,
		}

		locationResponseList = append(locationResponseList, locationResponse)
	}
	return courierapi.LocationResponseList{
		LocationResponseList: locationResponseList,
	}
}
