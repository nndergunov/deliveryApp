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

const baseAddr = "http://localhost:8082"

func TestGetEstimateDeliveryEndpoint(t *testing.T) {
	tests := []struct {
		name         string
		consumerID   string
		restaurantID string
		response     deliveryapi.EstimateDeliveryResponse
	}{
		{
			"GetEstimateDeliveryValues simple test",
			"1",
			"1",
			deliveryapi.EstimateDeliveryResponse{
				Time: "5h24m50s",
				Cost: 35.73,
			},
		},
	}

	for _, currentTest := range tests {
		test := currentTest

		t.Run(test.name, func(t *testing.T) {
			estimateDeliveryResp, err := http.Get(baseAddr + "/v1/estimate" + "?consumer_id=" + test.consumerID + "&restaurant_id=" + test.restaurantID)
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

			if respData.Time != test.response.Time {
				t.Errorf("Time: Expected : %s, Got: %s", test.response.Time, respData.Time)
			}

			if respData.Cost != test.response.Cost {
				t.Errorf("Cost: Expected: %v, Got: %v", test.response.Cost, respData.Cost)
			}
		})
	}
}

func TestAssignOrderEndpoint(t *testing.T) {
	tests := []struct {
		name            string
		assignOrderData deliveryapi.AssignOrderRequest
		orderID         string
		response        deliveryapi.AssignOrderResponse
	}{
		{
			"AssignOrder simple test",
			deliveryapi.AssignOrderRequest{
				FromUserID:   1,
				RestaurantID: 1,
			},
			"1",
			deliveryapi.AssignOrderResponse{
				OrderID:   1,
				CourierID: 1,
			},
		},
	}

	for _, currentTest := range tests {
		test := currentTest

		t.Run(test.name, func(t *testing.T) {
			reqBody, err := v1.Encode(test.assignOrderData)
			if err != nil {
				t.Fatal(err)
			}

			resp, err := http.Post(baseAddr+"/v1/orders/"+test.orderID+"/assign", "application/json", bytes.NewBuffer(reqBody))
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

			if respData.CourierID != test.response.CourierID {
				t.Errorf("CourierID: Expected: %v, Got: %v", test.response.CourierID, respData.CourierID)
			}
		})
	}
}
