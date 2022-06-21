package deliveryservice_test

import (
	"bytes"
	"net/http"
	"strconv"
	"testing"

	v1 "github.com/nndergunov/deliveryApp/app/pkg/api/v1"
	"github.com/nndergunov/deliveryApp/app/pkg/api/v1/courierapi"
	"github.com/nndergunov/deliveryApp/app/pkg/api/v1/deliveryapi"
)

const baseAddr = "http://localhost:8081"

func TestGetEstimateDeliveryEndpoint(t *testing.T) {
	tests := []struct {
		name         string
		consumerID   string
		restaurantID string
	}{
		{
			"GetEstimateDeliveryValues simple test",
			"1",
			"1",
		},
	}

	for _, currentTest := range tests {
		test := currentTest

		t.Run(test.name, func(t *testing.T) {
			estimateDeliveryResp, err := http.Get(baseAddr + "/v1/estimate" + "?" + test.consumerID + "?" + test.restaurantID)
			if err != nil {
				t.Fatal(err)
			}

			if estimateDeliveryResp.StatusCode != http.StatusOK {
				t.Fatalf("Response status: %d", estimateDeliveryResp.StatusCode)
			}

			respData := deliveryapi.EstimateDeliveryResponse{}
			if err = courierapi.DecodeJSON(estimateDeliveryResp.Body, &respData); err != nil {
				t.Fatal(err)
			}

			if err := estimateDeliveryResp.Body.Close(); err != nil {
				t.Error(err)
			}

			if respData.Time == "" {
				t.Errorf("Time: Expected : non empty string, Got: %v", respData.Time)
			}

			if respData.Cost == 0 {
				t.Errorf("Cost: Expected: >0, Got: %v", respData.Cost)
			}
		})
	}
}

func TestAssignOrderEndpoint(t *testing.T) {
	tests := []struct {
		name            string
		assignOrderData deliveryapi.AssignOrderRequest
		orderID         string
	}{
		{
			"AssignOrder simple test",
			deliveryapi.AssignOrderRequest{
				FromUserID:   1,
				RestaurantID: 1,
			},
			"1",
		},
	}

	for _, currentTest := range tests {
		test := currentTest

		t.Run(test.name, func(t *testing.T) {
			reqBody, err := v1.Encode(test.assignOrderData)
			if err != nil {
				t.Fatal(err)
			}

			resp, err := http.Post(baseAddr+"/v1/orders/"+test.orderID, "application/json", bytes.NewBuffer(reqBody))
			if err != nil {
				t.Fatal(err)
			}

			if resp.StatusCode != http.StatusOK {
				t.Fatalf("Response status: %d", resp.StatusCode)
			}

			respData := deliveryapi.AssignOrderResponse{}
			if err = deliveryapi.DecodeJSON(resp.Body, &respData); err != nil {
				t.Fatal(err)
			}

			orderIDInt, err := strconv.Atoi(test.orderID)
			if err != nil {
				t.Fatal(err)
			}

			if respData.OrderID != orderIDInt {
				t.Errorf("OrderID: Expected: %v, Got: %v", orderIDInt, respData.OrderID)
			}

			if respData.CourierID < 0 {
				t.Errorf("CourierID: Expected: >0, Got: %v", respData.CourierID)
			}
		})
	}
}
