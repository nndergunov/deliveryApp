package consumerhandler

import (
	"consumer/api/v1/consumerapi"
	"consumer/pkg/domain"
)

func consumerToResponse(consumer domain.Consumer) consumerapi.ConsumerResponse {
	return consumerapi.ConsumerResponse{
		ID:        consumer.ID,
		Firstname: consumer.Firstname,
		Lastname:  consumer.Lastname,
		Email:     consumer.Email,
		Phone:     consumer.Phone,
	}
}

func consumerListToResponse(consumerList []domain.Consumer) consumerapi.ReturnConsumerResponseList {
	consumerResponseList := make([]consumerapi.ConsumerResponse, 0, len(consumerList))

	for _, consumer := range consumerList {
		consumerResponse := consumerapi.ConsumerResponse{
			ID:        consumer.ID,
			Firstname: consumer.Firstname,
			Lastname:  consumer.Lastname,
			Email:     consumer.Email,
			Phone:     consumer.Phone,
		}

		consumerResponseList = append(consumerResponseList, consumerResponse)
	}
	return consumerapi.ReturnConsumerResponseList{
		ConsumerResponseList: consumerResponseList,
	}
}

func locationToResponse(location domain.Location) consumerapi.LocationResponse {
	return consumerapi.LocationResponse{
		UserID:     location.UserID,
		Altitude:   location.Latitude,
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
