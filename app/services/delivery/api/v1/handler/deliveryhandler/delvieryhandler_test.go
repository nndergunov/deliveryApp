package deliveryhandler_test

import (
	"bytes"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/nndergunov/deliveryApp/app/pkg/api/v1/deliveryapi"

	"github.com/nndergunov/deliveryApp/app/services/delivery/api/v1/handler/deliveryhandler"
	"github.com/nndergunov/deliveryApp/app/services/delivery/pkg/domain"

	"github.com/nndergunov/deliveryApp/app/pkg/api/v1"
	"github.com/nndergunov/deliveryApp/app/pkg/logger"

	mockservice "github.com/nndergunov/deliveryApp/app/services/delivery/pkg/mocks"
)

func TestGetEstimateDeliveryValuesEndpoint(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		in   struct {
			consumerID   string
			restaurantID string
		}
		out deliveryapi.EstimateDeliveryResponse
	}{
		{
			"get_estimate_delivery_values_test",
			struct {
				consumerID   string
				restaurantID string
			}{consumerID: "1", restaurantID: "1"},
			deliveryapi.EstimateDeliveryResponse{
				Time: "2006-01-02 15:04:05.999999999 -0700 MST",
				Cost: 100,
			},
		},
	}

	for _, currentTest := range tests {
		test := currentTest

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			ctl := gomock.NewController(t)
			service := mockservice.NewMockDeliveryService(ctl)

			mockOutData := &domain.EstimateDeliveryResponse{
				Time: test.out.Time,
				Cost: test.out.Cost,
			}
			service.EXPECT().GetEstimateDelivery(test.in.consumerID, test.in.restaurantID).Return(mockOutData, nil)

			handler := deliveryhandler.NewDeliveryHandler(deliveryhandler.Params{
				Logger:          logger.NewLogger(os.Stdout, test.name),
				DeliveryService: service,
			})

			resp := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, "/v1/estimate"+"?consumer_id="+test.in.consumerID+"&restaurant_id="+test.in.restaurantID, nil)

			handler.ServeHTTP(resp, req)

			if resp.Code != http.StatusOK {
				t.Fatalf("StatusCode: %d", resp.Code)
			}

			respData := deliveryapi.EstimateDeliveryResponse{}
			err := deliveryapi.DecodeJSON(resp.Body, &respData)
			require.NoError(t, err)

			assert.Equal(t, respData, test.out)
		})
	}
}

func TestAssignOrderEndpoint(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		InID   string
		inBody deliveryapi.AssignOrderRequest
		out    deliveryapi.AssignOrderResponse
	}{
		{
			"assign_order_test",
			"1",
			deliveryapi.AssignOrderRequest{
				FromUserID:   1,
				RestaurantID: 1,
			},
			deliveryapi.AssignOrderResponse{
				OrderID:   1,
				CourierID: 2,
			},
		},
	}

	for _, currentTest := range tests {
		test := currentTest

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			ctl := gomock.NewController(t)
			service := mockservice.NewMockDeliveryService(ctl)

			mockInData := &domain.Order{
				FromUserID:       test.inBody.FromUserID,
				FromRestaurantID: test.inBody.RestaurantID,
			}

			mockOutData := &domain.AssignOrder{
				OrderID:   test.out.OrderID,
				CourierID: test.out.CourierID,
			}
			service.EXPECT().AssignOrder(test.InID, mockInData).Return(mockOutData, nil)

			handler := deliveryhandler.NewDeliveryHandler(deliveryhandler.Params{
				Logger:          logger.NewLogger(os.Stdout, test.name),
				DeliveryService: service,
			})

			reqBody, err := v1.Encode(test.inBody)
			require.NoError(t, err)

			resp := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPost, "/v1/orders/"+test.InID+"/assign", bytes.NewBuffer(reqBody))

			handler.ServeHTTP(resp, req)

			if resp.Code != http.StatusOK {
				t.Fatalf("StatusCode: %d", resp.Code)
			}

			respData := deliveryapi.AssignOrderResponse{}
			err = deliveryapi.DecodeJSON(resp.Body, &respData)
			require.NoError(t, err)

			assert.Equal(t, respData, test.out)
		})
	}
}
