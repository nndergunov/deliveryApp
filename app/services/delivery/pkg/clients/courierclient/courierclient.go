package courierclient

import (
	"fmt"
	"net/http"
	"strconv"

	"delivery/pkg/domain"
)

type CourierClient struct {
	courierURL string
}

func NewCourierClient(url string) *CourierClient {
	return &CourierClient{courierURL: url}
}
func (a CourierClient) GetCourier(courierID int) (*domain.Courier, error) {
	_, err := http.Get(a.courierURL + "v1/couriers/" + strconv.Itoa(courierID))
	if err != nil {
		return nil, fmt.Errorf("sending request: %w", err)
	}

	//todo: add

	return &domain.Courier{}, nil
}
func (a CourierClient) GetNearestCourier(location *domain.Location, radius int) (*domain.Courier, error) {
	_, err := http.Get(a.courierURL + "v1/courier/nearest")
	if err != nil {
		return nil, fmt.Errorf("sending request: %w", err)
	}

	//todo: add this rout to courier service

	return &domain.Courier{}, nil
}
func (a CourierClient) UpdateCourierAvailable(courierID int, available bool) (*domain.Courier, error) {
	_, err := http.Get(a.courierURL + "v1/couriers/available/" + strconv.Itoa(courierID))
	if err != nil {
		return nil, fmt.Errorf("sending request: %w", err)
	}

	//todo: add

	return &domain.Courier{}, nil
}
