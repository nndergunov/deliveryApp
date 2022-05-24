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

func consumerLocationToResponse(consumerLocation domain.ConsumerLocation) consumerapi.ConsumerLocationResponse {
	return consumerapi.ConsumerLocationResponse{
		ConsumerID: consumerLocation.ConsumerID,
		Altitude:   consumerLocation.Altitude,
		Longitude:  consumerLocation.Longitude,
		Country:    consumerLocation.Country,
		City:       consumerLocation.City,
		Region:     consumerLocation.Region,
		Street:     consumerLocation.Street,
		HomeNumber: consumerLocation.HomeNumber,
		Floor:      consumerLocation.Floor,
		Door:       consumerLocation.Door,
	}
}
