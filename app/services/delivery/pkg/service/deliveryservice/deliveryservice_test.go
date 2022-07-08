package deliveryservice_test

import (
	"os"
	"strconv"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/nndergunov/deliveryApp/app/pkg/api/v1/consumerapi"
	"github.com/nndergunov/deliveryApp/app/pkg/api/v1/courierapi"
	"github.com/nndergunov/deliveryApp/app/pkg/api/v1/restaurantapi"
	"github.com/nndergunov/deliveryApp/app/pkg/logger"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/nndergunov/deliveryApp/app/services/delivery/pkg/domain"
	mock "github.com/nndergunov/deliveryApp/app/services/delivery/pkg/mocks"
	"github.com/nndergunov/deliveryApp/app/services/delivery/pkg/service/deliveryservice"
)

func TestGetEstimateDelivery(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		in   struct {
			consumerID   string
			restaurantID string
		}
		out *domain.EstimateDeliveryResponse
	}{
		{
			"get_estimate_delivery_test",
			struct {
				consumerID   string
				restaurantID string
			}{consumerID: "1", restaurantID: "1"},
			&domain.EstimateDeliveryResponse{
				Time: "4m20s",
				Cost: 0.48,
			},
		},
	}

	for _, currentTest := range tests {
		test := currentTest

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			ctl := gomock.NewController(t)
			storage := mock.NewMockDeliveryStorage(ctl)
			courierClient := mock.NewMockCourierClient(ctl)
			consumerClient := mock.NewMockConsumerClient(ctl)
			restaurantClient := mock.NewMockRestaurantClient(ctl)

			mockConsumerClientOutData := &consumerapi.LocationResponse{
				UserID:     1,
				Latitude:   "41.03641945369733",
				Longitude:  "28.919665086287385",
				Country:    "TestCountry",
				City:       "TestCity",
				Region:     "TestRegion",
				Street:     "TestStreet",
				HomeNumber: "TestHomeNumber",
				Floor:      "TestFloor",
				Door:       "TestDoor",
			}

			mockRestaurantClientOutData := &restaurantapi.ReturnRestaurant{
				ID:              1,
				Name:            "testRestaurant",
				AcceptingOrders: true,
				City:            "TestCity",
				Address:         "TestAddress",
				Longitude:       28.868111948612256,
				Latitude:        41.03630071727614,
			}

			mockConsumerClientInData, err := strconv.Atoi(test.in.consumerID)
			require.NoError(t, err)

			mockRestaurantClientInData, err := strconv.Atoi(test.in.restaurantID)
			require.NoError(t, err)

			consumerClient.EXPECT().GetLocation(mockConsumerClientInData).Return(mockConsumerClientOutData, nil)
			restaurantClient.EXPECT().GetRestaurant(mockRestaurantClientInData).Return(mockRestaurantClientOutData, nil)

			service := deliveryservice.NewDeliveryService(deliveryservice.Params{
				DeliveryStorage:  storage,
				Logger:           logger.NewLogger(os.Stdout, "service: "),
				RestaurantClient: restaurantClient,
				CourierClient:    courierClient,
				ConsumerClient:   consumerClient,
			})

			resp, err := service.GetEstimateDelivery(test.in.consumerID, test.in.restaurantID)
			require.NoError(t, err)

			assert.Equal(t, test.out, resp)
		})
	}
}

func TestAssignOrder(t *testing.T) {
	tests := []struct {
		name   string
		inID   string
		inBody *domain.Order
		out    *domain.AssignOrder
	}{
		{
			"assign_order_test",
			"1",
			&domain.Order{
				OrderID:          1,
				FromUserID:       1,
				FromRestaurantID: 1,
			},
			&domain.AssignOrder{
				OrderID:   1,
				CourierID: 1,
			},
		},
	}

	for _, currentTest := range tests {
		test := currentTest

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			ctl := gomock.NewController(t)
			storage := mock.NewMockDeliveryStorage(ctl)
			courierClient := mock.NewMockCourierClient(ctl)
			consumerClient := mock.NewMockConsumerClient(ctl)
			restaurantClient := mock.NewMockRestaurantClient(ctl)

			mockRestaurantClientOutData := &restaurantapi.ReturnRestaurant{
				ID:              1,
				Name:            "testRestaurant",
				AcceptingOrders: true,
				City:            "TestCity",
				Address:         "TestAddress",
				Longitude:       28.868111948612256,
				Latitude:        41.03630071727614,
			}

			restaurantClient.EXPECT().GetRestaurant(test.inBody.FromRestaurantID).Return(mockRestaurantClientOutData, nil)

			mockCourierClientOutData := &courierapi.LocationResponseList{
				LocationResponseList: []courierapi.LocationResponse{
					{
						UserID:     1,
						Latitude:   "41.03641945369733",
						Longitude:  "28.919665086287385",
						Country:    "TestCountry",
						City:       "TestCity",
						Region:     "TestRegion",
						Street:     "TestStreet",
						HomeNumber: "TestHomeNumber",
						Floor:      "TestFloor",
						Door:       "TestDoor",
					},
				},
			}

			courierClient.EXPECT().GetLocation(mockRestaurantClientOutData.City).Return(mockCourierClientOutData, nil)

			mockStorageInData := domain.AssignOrder{
				OrderID:   test.out.OrderID,
				CourierID: test.out.CourierID,
			}

			mockStorageOutData := &domain.AssignOrder{
				OrderID:   test.out.OrderID,
				CourierID: test.out.CourierID,
			}

			storage.EXPECT().AssignOrder(mockStorageInData).Return(mockStorageOutData, nil)

			courierClient.EXPECT().UpdateCourierAvailable(mockCourierClientOutData.LocationResponseList[0].UserID, "false").Return(&courierapi.CourierResponse{}, nil)

			service := deliveryservice.NewDeliveryService(deliveryservice.Params{
				DeliveryStorage:  storage,
				Logger:           logger.NewLogger(os.Stdout, "service: "),
				RestaurantClient: restaurantClient,
				CourierClient:    courierClient,
				ConsumerClient:   consumerClient,
			})

			resp, err := service.AssignOrder(test.inID, test.inBody)
			require.NoError(t, err)

			assert.Equal(t, resp, test.out)
		})
	}
}
