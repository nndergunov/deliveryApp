package handlers_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"os"
	"sort"
	"testing"

	v1 "github.com/nndergunov/deliveryApp/app/pkg/api/v1"
	"github.com/nndergunov/deliveryApp/app/pkg/logger"
	"github.com/nndergunov/deliveryApp/app/services/order/api/v1/communication"
	"github.com/nndergunov/deliveryApp/app/services/order/api/v1/handlers"
	"github.com/nndergunov/deliveryApp/app/services/order/pkg/domain"
	"github.com/nndergunov/deliveryApp/app/services/order/pkg/service/mockservice"
	"github.com/stretchr/testify/mock"
)

func TestReturnAllOrders(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name          string
		defaultOrders []domain.Order
	}{
		{
			name: "Simple return orders endpoint",
			defaultOrders: []domain.Order{{
				OrderID:      0,
				FromUserID:   0,
				RestaurantID: 0,
				OrderItems:   []int{1, 2, 3},
				Status: domain.OrderStatus{
					OrderID: 0,
					Status:  "incomplete",
				},
			}},
		},
	}

	for _, currTest := range tests {
		test := currTest

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			repo := &mockservice.App{}

			repo.On("ReturnOrderList", mock.AnythingOfType("domain.SearchParameters")).
				Return(test.defaultOrders, nil).
				Once()

			handler := handlers.NewEndpointHandler(
				repo,
				logger.NewLogger(os.Stdout, t.Name()),
			)

			resp := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, "/v1/orders", nil)

			handler.ServeHTTP(resp, req)

			orders := new(communication.ReturnOrderList)

			err := v1.Decode(resp.Body.Bytes(), orders)
			if err != nil {
				t.Fatal(err)
			}

			if len(orders.Orders) != len(test.defaultOrders) {
				t.Fatalf("Wrong number of orders, expected: %d, got: %d", len(test.defaultOrders), len(orders.Orders))
			}
		})
	}
}

func TestCreateOrderEndpoint(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		orderData domain.Order
	}{
		{
			"Simple create order",
			domain.Order{
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

			repo := &mockservice.App{}

			repo.On("CreateOrder", mock.AnythingOfType("domain.Order"), mock.AnythingOfType("int")).
				Return(&test.orderData, nil).
				Once()

			handler := handlers.NewEndpointHandler(repo, logger.NewLogger(os.Stdout, test.name))

			reqBody, err := v1.Encode(test.orderData)
			if err != nil {
				t.Fatal(err)
			}

			resp := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPost, "/v1/orders", bytes.NewBuffer(reqBody))

			handler.ServeHTTP(resp, req)

			order := new(communication.ReturnOrder)

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

	tests := []struct {
		name      string
		orderData domain.Order
	}{
		{
			name: "Simple return order",
			orderData: domain.Order{
				FromUserID:   0,
				RestaurantID: 0,
				OrderItems:   []int{1, 2, 3},
			},
		},
	}

	for _, currTest := range tests {
		test := currTest

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			repo := &mockservice.App{}

			repo.On("ReturnOrder", mock.AnythingOfType("int")).
				Return(&test.orderData, nil).
				Once()

			handler := handlers.NewEndpointHandler(repo, logger.NewLogger(os.Stdout, t.Name()))

			resp := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, "/v1/orders/0", nil)

			handler.ServeHTTP(resp, req)

			order := new(communication.ReturnOrder)

			err := v1.Decode(resp.Body.Bytes(), order)
			if err != nil {
				t.Fatal(err)
			}

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

func TestUpdateOrderEndpoint(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		orderData domain.Order
	}{
		{
			"Simple update order",
			domain.Order{
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

			repo := &mockservice.App{}

			repo.On("UpdateOrder", mock.AnythingOfType("domain.Order")).
				Return(&test.orderData, nil).
				Once()

			handler := handlers.NewEndpointHandler(repo, logger.NewLogger(os.Stdout, test.name))

			reqBody, err := v1.Encode(test.orderData)
			if err != nil {
				t.Fatal(err)
			}

			resp := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPut, "/v1/orders/0", bytes.NewBuffer(reqBody))

			handler.ServeHTTP(resp, req)

			order := new(communication.ReturnOrder)

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
		newStatus domain.OrderStatus
	}{
		{
			"Simple update order status",
			domain.OrderStatus{
				Status: "incomplete",
			},
		},
	}

	for _, currentTest := range tests {
		test := currentTest

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			repo := &mockservice.App{}

			repo.On("UpdateStatus", mock.AnythingOfType("domain.OrderStatus")).
				Return(&test.newStatus, nil).
				Once()

			handler := handlers.NewEndpointHandler(repo, logger.NewLogger(os.Stdout, test.name))

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
