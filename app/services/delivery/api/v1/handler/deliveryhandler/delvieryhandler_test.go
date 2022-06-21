package deliveryhandler_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/nndergunov/deliveryApp/app/pkg/api/v1/deliveryapi"

	"github.com/nndergunov/deliveryApp/app/services/delivery/api/v1/handler/deliveryhandler"
	"github.com/nndergunov/deliveryApp/app/services/delivery/pkg/domain"

	"github.com/nndergunov/deliveryApp/app/pkg/api/v1"
	"github.com/nndergunov/deliveryApp/app/pkg/logger"
)

var (
	MockEstimateDeliveryData = &domain.EstimateDeliveryResponse{
		Time: "2006-01-02 15:04:05.999999999 -0700 MST",
		Cost: 100,
	}

	MockAssignOrderData = &domain.AssignOrder{
		OrderID:   1,
		CourierID: 1,
	}
)

type MockService struct{}

func (m MockService) GetEstimateDelivery(consumerID, restaurantID string) (*domain.EstimateDeliveryResponse, error) {
	return MockEstimateDeliveryData, nil
}

func (m MockService) AssignOrder(orderID string, order *domain.Order) (*domain.AssignOrder, error) {
	return MockAssignOrderData, nil
}

func TestGetEstimateDeliveryValuesEndpoint(t *testing.T) {
	t.Parallel()

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
			t.Parallel()

			mockService := new(MockService)

			log := logger.NewLogger(os.Stdout, test.name)
			handler := deliveryhandler.NewDeliveryHandler(deliveryhandler.Params{
				Logger:          log,
				DeliveryService: mockService,
			})

			resp := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, "/v1/estimate"+"?"+test.consumerID+"?"+test.restaurantID, nil)

			handler.ServeHTTP(resp, req)

			if resp.Code != http.StatusOK {
				t.Fatalf("StatusCode: %d", resp.Code)
			}

			respData := deliveryapi.EstimateDeliveryResponse{}
			if err := deliveryapi.DecodeJSON(resp.Body, &respData); err != nil {
				t.Fatal(err)
			}

			if respData.Time != MockEstimateDeliveryData.Time {
				t.Errorf("Time: Expected: %v, Got: %v", MockEstimateDeliveryData.Time, respData.Time)
			}

			if respData.Cost != MockEstimateDeliveryData.Cost {
				t.Errorf("Cost: Expected: %s, Got: %s", MockEstimateDeliveryData.Cost, respData.Cost)
			}
		})
	}
}

func TestAssignOrderEndpoint(t *testing.T) {
	t.Parallel()

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
			t.Parallel()

			mockService := new(MockService)

			log := logger.NewLogger(os.Stdout, test.name)
			handler := deliveryhandler.NewDeliveryHandler(deliveryhandler.Params{
				Logger:          log,
				DeliveryService: mockService,
			})

			reqBody, err := v1.Encode(test.assignOrderData)
			if err != nil {
				t.Fatal(err)
			}

			resp := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPost, "/v1/orders/"+test.orderID+"/assign", bytes.NewBuffer(reqBody))

			handler.ServeHTTP(resp, req)

			if resp.Code != http.StatusOK {
				t.Fatalf("StatusCode: %d", resp.Code)
			}

			respData := deliveryapi.AssignOrderResponse{}
			if err = deliveryapi.DecodeJSON(resp.Body, &respData); err != nil {
				t.Fatal(err)
			}

			if respData.OrderID != MockAssignOrderData.OrderID {
				t.Errorf("OrderID: Expected: %v, Got: %v", MockAssignOrderData.OrderID, respData.OrderID)
			}

			if respData.CourierID != MockAssignOrderData.CourierID {
				t.Errorf("CourierID: Expected: %v, Got: %v", MockAssignOrderData.CourierID, respData.CourierID)
			}
		})
	}
}
