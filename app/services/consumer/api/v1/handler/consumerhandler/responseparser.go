package consumerhandler

import (
	"consumer/api/v1/consumerapi"
	"consumer/domain"
)

func consumerToResponse(consumer domain.Consumer) consumerapi.ConsumerResponse {
	return consumerapi.ConsumerResponse{
		ID:        consumer.ID,
		Firstname: consumer.Firstname,
		Lastname:  consumer.Lastname,
		Email:     consumer.Email,
		Phone:     consumer.Phone,
		ConsumerLocation: consumerapi.ConsumerLocationResponse{
			LocationAlt: consumer.ConsumerLocation.LocationAlt,
			LocationLat: consumer.ConsumerLocation.LocationLat,
			Country:     consumer.ConsumerLocation.Country,
			City:        consumer.ConsumerLocation.City,
			Region:      consumer.ConsumerLocation.Region,
			Street:      consumer.ConsumerLocation.Street,
			HomeNumber:  consumer.ConsumerLocation.HomeNumber,
			Floor:       consumer.ConsumerLocation.Floor,
			Door:        consumer.ConsumerLocation.Door,
		},
	}
}

func consumerListToResponse(consumerList []domain.Consumer) consumerapi.ReturnConsumerList {
	consumerResponseList := make([]consumerapi.ConsumerResponse, 0, len(consumerList))

	for _, consumer := range consumerList {
		courierResponse := consumerapi.ConsumerResponse{
			ID:        consumer.ID,
			Firstname: consumer.Firstname,
			Lastname:  consumer.Lastname,
			Email:     consumer.Email,
			Phone:     consumer.Phone,
			ConsumerLocation: consumerapi.ConsumerLocationResponse{
				LocationAlt: consumer.ConsumerLocation.LocationAlt,
				LocationLat: consumer.ConsumerLocation.LocationLat,
				Country:     consumer.ConsumerLocation.Country,
				City:        consumer.ConsumerLocation.City,
				Region:      consumer.ConsumerLocation.Region,
				Street:      consumer.ConsumerLocation.Street,
				HomeNumber:  consumer.ConsumerLocation.HomeNumber,
				Floor:       consumer.ConsumerLocation.Floor,
				Door:        consumer.ConsumerLocation.Door,
			},
		}

		consumerResponseList = append(consumerResponseList, courierResponse)
	}
	return consumerapi.ReturnConsumerList{
		ConsumerList: consumerResponseList,
	}
}

func consumerLocationToResponse(consumerLocation domain.ConsumerLocation) consumerapi.ConsumerLocationResponse {
	return consumerapi.ConsumerLocationResponse{
		LocationAlt: consumerLocation.LocationAlt,
		LocationLat: consumerLocation.LocationLat,
		Country:     consumerLocation.Country,
		City:        consumerLocation.City,
		Region:      consumerLocation.Region,
		Street:      consumerLocation.Street,
		HomeNumber:  consumerLocation.HomeNumber,
		Floor:       consumerLocation.Floor,
		Door:        consumerLocation.Door,
	}
}
