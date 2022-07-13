package service_test

import (
	"testing"

	"github.com/nndergunov/deliveryApp/app/services/kitchen/pkg/service"
	"github.com/nndergunov/deliveryApp/app/services/order/api/v1/communication"
)

var mockOrderList = communication.ReturnOrderList{
	Orders: []communication.ReturnOrder{
		{
			OrderID:      0,
			FromUserID:   0,
			RestaurantID: 0,
			OrderItems:   []int{1, 2, 3, 4, 5},
			Status:       "",
		},
	},
}

type mockCommunicator struct{}

func (m mockCommunicator) GetIncompleteOrders(_ int) (*communication.ReturnOrderList, error) {
	return &mockOrderList, nil
}

func TestGetTasks(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
	}{
		{
			name: "Get tasks simple",
		},
	}

	for _, currentTest := range tests {
		test := currentTest

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			mockComm := mockCommunicator{}

			serviceInstance := service.NewService(mockComm)

			tasks, err := serviceInstance.GetTasks(0)
			if err != nil {
				return
			}

			for i := 1; i <= 5; i++ {
				if tasks[i] != 1 {
					t.Errorf("Wrong number of ordered items, expected 1, got: %d", tasks[i])
				}
			}
		})
	}
}
