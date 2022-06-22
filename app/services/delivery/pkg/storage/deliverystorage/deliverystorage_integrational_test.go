package deliverystorage_test

import (
	"os"
	"strings"
	"testing"

	"github.com/nndergunov/deliveryApp/app/pkg/configreader"

	"github.com/nndergunov/deliveryApp/app/services/delivery/pkg/db"
	"github.com/nndergunov/deliveryApp/app/services/delivery/pkg/domain"
	"github.com/nndergunov/deliveryApp/app/services/delivery/pkg/storage/deliverystorage"
)

const configFile = "/config.yaml"

func TestAssignOrder(t *testing.T) {
	tests := []struct {
		name        string
		assignOrder domain.AssignOrder
		response    domain.AssignOrder
	}{
		{
			name: "TestAssignOrder",
			assignOrder: domain.AssignOrder{
				OrderID:   1,
				CourierID: 1,
			},

			response: domain.AssignOrder{
				OrderID:   1,
				CourierID: 1,
			},
		},
	}

	for _, currentTest := range tests {
		test := currentTest

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			line, err := os.Getwd()
			if err != nil {
				t.Fatal(err)
			}
			confPath := strings.TrimSuffix(line, "\\pkg\\storage\\deliverystorage")

			err = configreader.SetConfigFile(confPath + configFile)
			if err != nil {
				t.Fatal(err)
			}

			database, err := db.OpenDB("postgres", configreader.GetString("DB.test"))
			if err != nil {
				t.Fatal(err)
			}

			storage := deliverystorage.NewDeliveryStorage(deliverystorage.Params{DB: database})
			if err != nil {
				t.Fatal(err)
			}

			assignedOrder, err := storage.AssignOrder(test.assignOrder)
			if err != nil {
				t.Fatal(err)
			}

			if assignedOrder == nil {
				t.Errorf("assignedOrder: Expected: %s, Got: %s", "not nil", "nil")
			}

			if assignedOrder.OrderID != test.response.OrderID {
				t.Errorf("OrderID: Expected: %v, Got: %v", test.assignOrder.OrderID, assignedOrder.OrderID)
			}

			if assignedOrder.CourierID != test.response.CourierID {
				t.Errorf("CourierID: Expected: %v, Got: %v", test.assignOrder.CourierID, assignedOrder.CourierID)
			}

			if err = storage.DeleteAssignedOrder(assignedOrder.OrderID); err != nil {
				t.Error(err)
			}

			if err := database.Close(); err != nil {
				t.Error(err)
			}
		})
	}
}
