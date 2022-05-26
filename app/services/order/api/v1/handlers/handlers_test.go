package handlers_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"os"
	"sort"
	"testing"

	v1 "github.com/nndergunov/deliveryApp/app/pkg/api/v1"
	"github.com/nndergunov/deliveryApp/app/pkg/api/v1/orderapi"
	"github.com/nndergunov/deliveryApp/app/pkg/logger"
	"github.com/nndergunov/deliveryApp/app/services/order/api/v1/handlers"
	"github.com/nndergunov/deliveryApp/app/services/order/pkg/domain"
)

var defaultOrder = domain.Order{
	OrderID:      0,
	FromUserID:   0,
	RestaurantID: 0,
	OrderItems:   []int{1, 2, 3},
	Status: domain.OrderStatus{
		OrderID: 0,
		Status:  "incomplete",
	},
}

type mockService struct{}

func (m mockService) CreateOrder(order domain.Order) (*domain.Order, error) {
	return &order, nil
}

func (m mockService) ReturnOrder(_ int) (*domain.Order, error) {
	return &defaultOrder, nil
}

func (m mockService) UpdateOrder(order domain.Order) (*domain.Order, error) {
	return &order, nil
}

func (m mockService) ReturnOrderList(_ domain.SearchParameters) ([]domain.Order, error) {
	return []domain.Order{defaultOrder}, nil
}

func (m mockService) UpdateStatus(status domain.OrderStatus) (*domain.OrderStatus, error) {
	return &status, nil
}

func TestCreateOrderEndpoint(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		orderData orderapi.OrderData
	}{
		{
			"Simple create order test",
			orderapi.OrderData{
				FromUserID:   0,
				RestaurantID: 0,
				OrderItems:   []int{1, 2, 3},
			},
		},
	}

	for _, currentTest := range tests {
		test := currentTest

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			handler := handlers.NewEndpointHandler(mockService{}, logger.NewLogger(os.Stdout, test.name))

			reqBody, err := v1.Encode(test.orderData)
			if err != nil {
				t.Fatal(err)
			}

			resp := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPost, "/v1/orders", bytes.NewBuffer(reqBody))

			handler.ServeHTTP(resp, req)

			order := new(orderapi.ReturnOrder)

			err = v1.Decode(resp.Body.Bytes(), order)
			if err != nil {
				t.Fatal(err)
			}

			sort.Ints(test.orderData.OrderItems)
			sort.Ints(order.OrderItems)

			if len(order.OrderItems) != len(test.orderData.OrderItems) {
				t.Fatal("Wrong number of order items received")
			}

			for i := 0; i < len(test.orderData.OrderItems); i++ {
				if test.orderData.OrderItems[i] != order.OrderItems[i] {
					t.Errorf("Item number %d: Expected ID: %d, Got ID: %d",
						i, test.orderData.OrderItems[i], order.OrderItems[i])
				}
			}
		})
	}
}

func TestReturnOrderEndpoint(t *testing.T) {
	t.Parallel()

	t.Run("Simple return restaurant endpoint test", func(t *testing.T) {
		t.Parallel()

		handler := handlers.NewEndpointHandler(mockService{}, logger.NewLogger(os.Stdout, t.Name()))

		resp := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/v1/orders/0", nil)

		handler.ServeHTTP(resp, req)

		order := new(orderapi.ReturnOrder)

		err := v1.Decode(resp.Body.Bytes(), order)
		if err != nil {
			t.Fatal(err)
		}

		sort.Ints(order.OrderItems)

		if len(order.OrderItems) != len(defaultOrder.OrderItems) {
			t.Fatal("Wrong number of order items received")
		}

		for i := 0; i < len(defaultOrder.OrderItems); i++ {
			if defaultOrder.OrderItems[i] != order.OrderItems[i] {
				t.Errorf("Item number %d: Expected ID: %d, Got ID: %d",
					i, defaultOrder.OrderItems[i], order.OrderItems[i])
			}
		}
	})
}

func TestUpdateOrderEndpoint(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		orderData orderapi.OrderData
	}{
		{
			"Simple update order test",
			orderapi.OrderData{
				FromUserID:   0,
				RestaurantID: 0,
				OrderItems:   []int{1, 2, 3},
			},
		},
	}

	for _, currentTest := range tests {
		test := currentTest

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			handler := handlers.NewEndpointHandler(mockService{}, logger.NewLogger(os.Stdout, test.name))

			reqBody, err := v1.Encode(test.orderData)
			if err != nil {
				t.Fatal(err)
			}

			resp := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPut, "/v1/orders/0", bytes.NewBuffer(reqBody))

			handler.ServeHTTP(resp, req)

			order := new(orderapi.ReturnOrder)

			err = v1.Decode(resp.Body.Bytes(), order)
			if err != nil {
				t.Fatal(err)
			}
			sort.Ints(test.orderData.OrderItems)
			sort.Ints(order.OrderItems)

			if len(order.OrderItems) != len(test.orderData.OrderItems) {
				t.Fatal("Wrong number of order items received")
			}

			for i := 0; i < len(test.orderData.OrderItems); i++ {
				if test.orderData.OrderItems[i] != order.OrderItems[i] {
					t.Errorf("Item number %d: Expected ID: %d, Got ID: %d",
						i, test.orderData.OrderItems[i], order.OrderItems[i])
				}
			}
		})
	}
}

func TestUpdateOrderStatusEndpoint(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		newStatus orderapi.OrderStatusData
	}{
		{
			"Simple update order status test",
			orderapi.OrderStatusData{
				Status: "incomplete",
			},
		},
	}

	for _, currentTest := range tests {
		test := currentTest

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			handler := handlers.NewEndpointHandler(mockService{}, logger.NewLogger(os.Stdout, test.name))

			reqBody, err := v1.Encode(test.newStatus)
			if err != nil {
				t.Fatal(err)
			}

			resp := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPost, "/v1/admin/orders/0/status", bytes.NewBuffer(reqBody))

			handler.ServeHTTP(resp, req)
		})
	}
}

func TestReturnOrderListEndpoint(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		orderData orderapi.OrderData
	}{
		{
			name: "Simple create order test",
			orderData: orderapi.OrderData{
				FromUserID:   0,
				RestaurantID: 0,
				OrderItems:   []int{1, 2, 3},
			},
		},
	}

	for _, currentTest := range tests {
		test := currentTest

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			handler := handlers.NewEndpointHandler(mockService{}, logger.NewLogger(os.Stdout, test.name))

			resp := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, "/v1/orders", nil)

			handler.ServeHTTP(resp, req)

			orders := new(orderapi.ReturnOrderList)

			err := v1.Decode(resp.Body.Bytes(), orders)
			if err != nil {
				t.Fatal(err)
			}

			order := orders.Orders[0]

			sort.Ints(order.OrderItems)

			if len(order.OrderItems) != len(defaultOrder.OrderItems) {
				t.Fatal("Wrong number of order items received")
			}

			for i := 0; i < len(defaultOrder.OrderItems); i++ {
				if defaultOrder.OrderItems[i] != order.OrderItems[i] {
					t.Errorf("Item number %d: Expected ID: %d, Got ID: %d",
						i, defaultOrder.OrderItems[i], order.OrderItems[i])
				}
			}
		})
	}
}
