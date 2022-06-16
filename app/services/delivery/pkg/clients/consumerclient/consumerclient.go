package restaurantclient

import (
	"fmt"
	"net/http"
	"strconv"

	"delivery/pkg/domain"
)

type ConsumerClient struct {
	consumerURL string
}

func NewConsumerClient(url string) *ConsumerClient {
	return &ConsumerClient{consumerURL: url}
}

func (a ConsumerClient) GetRestaurant(consumerID int) (*domain.Location, error) {
	_, err := http.Get(a.consumerURL + "v1/consumers/location/" + strconv.Itoa(consumerID))
	if err != nil {
		return nil, fmt.Errorf("sending request: %w", err)
	}
	// todo when restaurant add this rout

	return &domain.Location{}, nil
}
