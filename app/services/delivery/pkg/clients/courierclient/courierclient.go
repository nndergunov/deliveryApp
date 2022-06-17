package courierclient

import (
	"fmt"
	"github.com/nndergunov/deliveryApp/app/pkg/api/v1/courierapi"
	"github.com/nndergunov/deliveryApp/app/pkg/api/v1/deliveryapi"
	"net/http"
	"strconv"
)

type CourierClient struct {
	courierURL string
}

func NewCourierClient(url string) *CourierClient {
	return &CourierClient{courierURL: url}
}

func (a CourierClient) GetCourier(courierID int) (*courierapi.CourierResponse, error) {
	resp, err := http.Get(a.courierURL + "v1/couriers/" + strconv.Itoa(courierID))
	if err != nil {
		return nil, fmt.Errorf("sending request: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("not ok status: %w", resp.StatusCode)
	}

	courierData := courierapi.CourierResponse{}
	if err = deliveryapi.DecodeJSON(resp.Body, &courierData); err != nil {
		return nil, fmt.Errorf("decoding : %w", err)
	}

	if err := resp.Body.Close(); err != nil {
		return nil, err
	}

	return &courierData, nil
}

func (a CourierClient) GetLocation(city string) (*courierapi.LocationResponseList, error) {
	resp, err := http.Get(a.courierURL + "v1/locations?city=" + city)
	if err != nil {
		return nil, fmt.Errorf("sending request: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("not ok status: %w", resp.StatusCode)
	}

	locationDataList := courierapi.LocationResponseList{}
	if err = deliveryapi.DecodeJSON(resp.Body, &locationDataList); err != nil {
		return nil, fmt.Errorf("decoding : %w", err)
	}

	if err := resp.Body.Close(); err != nil {
		return nil, err
	}

	return &locationDataList, nil
}

func (a CourierClient) UpdateCourierAvailable(courierID int, available string) (*courierapi.CourierResponse, error) {
	resp, err := http.Get(a.courierURL + "v1/couriers-available/" + strconv.Itoa(courierID) + "?available=" + available)
	if err != nil {
		return nil, fmt.Errorf("sending request: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("not ok status: %w", resp.StatusCode)
	}

	courierData := courierapi.CourierResponse{}
	if err = deliveryapi.DecodeJSON(resp.Body, &courierData); err != nil {
		return nil, fmt.Errorf("decoding : %w", err)
	}

	if err := resp.Body.Close(); err != nil {
		return nil, err
	}

	return &courierData, nil
}
