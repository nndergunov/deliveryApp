package consumerclient

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/nndergunov/deliveryApp/app/services/consumer/api/v1/rest/consumerapi"

	"github.com/nndergunov/deliveryApp/app/services/delivery/api/v1/rest/deliveryapi"
)

type ConsumerClient struct {
	consumerURL string
}

func NewConsumerClient(url string) *ConsumerClient {
	return &ConsumerClient{consumerURL: url}
}

func (a ConsumerClient) GetLocation(consumerID int) (*consumerapi.LocationResponse, error) {
	resp, err := http.Get(a.consumerURL + "/v1/locations/" + strconv.Itoa(consumerID))
	if err != nil {
		return nil, fmt.Errorf("sending request: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("not ok status: %v", resp.StatusCode)
	}

	locationData := consumerapi.LocationResponse{}
	if err = deliveryapi.DecodeJSON(resp.Body, &locationData); err != nil {
		return nil, fmt.Errorf("decoding : %w", err)
	}

	if err := resp.Body.Close(); err != nil {
		return nil, err
	}

	return &locationData, nil
}
