package db_test

import (
	"fmt"
	"sort"
	"testing"

	"github.com/nndergunov/deliveryApp/app/pkg/configreader"
	"github.com/nndergunov/deliveryApp/app/services/order/pkg/db"
	"github.com/nndergunov/deliveryApp/app/services/order/pkg/domain"
)

func numToPointer(num int) *int {
	return &num
}

func TestGetAllOrders(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		order domain.Order
	}{
		{
			name: "Test Insert Order",
			order: domain.Order{
				OrderID:      0,
				FromUserID:   2,
				RestaurantID: 64851,
				OrderItems:   []int{2311},
				Status: domain.OrderStatus{
					OrderID: 0,
					Status:  "",
				},
			},
		},
	}

	for _, currentTest := range tests {
		test := currentTest

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			err := configreader.SetConfigFile("config.yaml")
			if err != nil {
				t.Fatal(err)
			}

			dbURL := fmt.Sprintf("host=" + configreader.GetString("database.host") +
				" port=" + configreader.GetString("database.port") +
				" user=" + configreader.GetString("database.user") +
				" password=" + configreader.GetString("database.password") +
				" dbname=" + configreader.GetString("database.dbName") +
				" sslmode=" + configreader.GetString("database.sslmode"))

			database, err := db.NewDatabase(dbURL)
			if err != nil {
				t.Fatal(err)
			}

			orderID, _ := database.InsertOrder(test.order)

			orders, err := database.GetAllOrders(nil)
			if err != nil {
				t.Fatal(err)
			}

			var (
				order domain.Order
				found bool
			)

			for _, ordr := range orders {
				if ordr.OrderID == orderID {
					order = ordr
					found = true

					break
				}
			}

			if !found {
				t.Fatal("Did not find inserted incompleteOrder")
			}

			if test.order.RestaurantID != order.RestaurantID {
				t.Errorf("RestaurantID: Expected: %d, Got: %d", test.order.RestaurantID, order.RestaurantID)
			}

			if test.order.FromUserID != order.FromUserID {
				t.Errorf("FromUserID: Expected: %d, Got: %d", test.order.FromUserID, order.FromUserID)
			}

			sort.Ints(test.order.OrderItems)
			sort.Ints(order.OrderItems)

			for i := 0; i < len(test.order.OrderItems); i++ {
				if test.order.OrderItems[i] != order.OrderItems[i] {
					t.Errorf("OrderItem %d: Expected: %d, Got: %d",
						i, test.order.OrderItems[i], order.OrderItems[i])
				}
			}

			_ = database.DeleteOrder(orderID)
		})
	}
}

func TestGetAllIncompleteOrdersFromRestaurant(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name            string
		restaurantID    int
		incompleteOrder domain.Order
		completeOrder   domain.Order
		params          domain.SearchParameters
	}{
		{
			name:         "Test Insert Order",
			restaurantID: 673473346,
			incompleteOrder: domain.Order{
				OrderID:      0,
				FromUserID:   2,
				RestaurantID: 673473346,
				OrderItems:   []int{2311},
				Status: domain.OrderStatus{
					OrderID: 0,
					Status:  "",
				},
			},
			completeOrder: domain.Order{
				OrderID:      0,
				FromUserID:   182,
				RestaurantID: 673473346,
				OrderItems:   []int{},
				Status: domain.OrderStatus{
					OrderID: 0,
					Status:  "",
				},
			},
			params: domain.SearchParameters{
				FromRestaurantID: numToPointer(673473346),
				Statuses:         nil,
				ExcludeStatuses:  []string{"complete"},
			},
		},
	}

	for _, currentTest := range tests {
		test := currentTest

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			err := configreader.SetConfigFile("config.yaml")
			if err != nil {
				t.Fatal(err)
			}

			dbURL := fmt.Sprintf("host=" + configreader.GetString("database.host") +
				" port=" + configreader.GetString("database.port") +
				" user=" + configreader.GetString("database.user") +
				" password=" + configreader.GetString("database.password") +
				" dbname=" + configreader.GetString("database.dbName") +
				" sslmode=" + configreader.GetString("database.sslmode"))

			database, err := db.NewDatabase(dbURL)
			if err != nil {
				t.Fatal(err)
			}

			orderID, _ := database.InsertOrder(test.incompleteOrder)
			complOrderID, _ := database.InsertOrder(test.completeOrder)

			_ = database.UpdateOrderStatus(complOrderID, "complete")

			orders, err := database.GetAllOrders(&test.params)
			if err != nil {
				t.Fatal(err)
			}

			var (
				order domain.Order
				found bool
			)

			for _, ordr := range orders {
				if !found {
					if ordr.OrderID == orderID {
						order = ordr
						found = true
					}
				}

				if ordr.Status.Status == "complete" {
					t.Errorf("found complete order")
				}
			}

			if !found {
				t.Fatal("Did not find inserted incompleteOrder")
			}

			if test.incompleteOrder.RestaurantID != order.RestaurantID {
				t.Errorf("RestaurantID: Expected: %d, Got: %d", test.incompleteOrder.RestaurantID, order.RestaurantID)
			}

			if test.incompleteOrder.FromUserID != order.FromUserID {
				t.Errorf("FromUserID: Expected: %d, Got: %d", test.incompleteOrder.FromUserID, order.FromUserID)
			}

			sort.Ints(test.incompleteOrder.OrderItems)
			sort.Ints(order.OrderItems)

			for i := 0; i < len(test.incompleteOrder.OrderItems); i++ {
				if test.incompleteOrder.OrderItems[i] != order.OrderItems[i] {
					t.Errorf("OrderItem %d: Expected: %d, Got: %d",
						i, test.incompleteOrder.OrderItems[i], order.OrderItems[i])
				}
			}

			_ = database.DeleteOrder(orderID)
		})
	}
}

func TestInsertOrder(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		order domain.Order
	}{
		{
			name: "Test Insert Order",
			order: domain.Order{
				OrderID:      0,
				FromUserID:   9,
				RestaurantID: 8,
				OrderItems:   []int{1691131},
				Status: domain.OrderStatus{
					OrderID: 0,
					Status:  "",
				},
			},
		},
	}

	for _, currentTest := range tests {
		test := currentTest

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			err := configreader.SetConfigFile("config.yaml")
			if err != nil {
				t.Fatal(err)
			}

			dbURL := fmt.Sprintf("host=" + configreader.GetString("database.host") +
				" port=" + configreader.GetString("database.port") +
				" user=" + configreader.GetString("database.user") +
				" password=" + configreader.GetString("database.password") +
				" dbname=" + configreader.GetString("database.dbName") +
				" sslmode=" + configreader.GetString("database.sslmode"))

			database, err := db.NewDatabase(dbURL)
			if err != nil {
				t.Fatal(err)
			}

			orderID, err := database.InsertOrder(test.order)
			if err != nil {
				t.Fatal(err)
			}

			_ = database.DeleteOrder(orderID)
		})
	}
}

func TestGetOrder(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		order domain.Order
	}{
		{
			name: "Test Insert Order",
			order: domain.Order{
				OrderID:      0,
				FromUserID:   7368911,
				RestaurantID: 900501,
				OrderItems:   []int{5643},
				Status: domain.OrderStatus{
					OrderID: 0,
					Status:  "",
				},
			},
		},
	}

	for _, currentTest := range tests {
		test := currentTest

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			err := configreader.SetConfigFile("config.yaml")
			if err != nil {
				t.Fatal(err)
			}

			dbURL := fmt.Sprintf("host=" + configreader.GetString("database.host") +
				" port=" + configreader.GetString("database.port") +
				" user=" + configreader.GetString("database.user") +
				" password=" + configreader.GetString("database.password") +
				" dbname=" + configreader.GetString("database.dbName") +
				" sslmode=" + configreader.GetString("database.sslmode"))

			database, err := db.NewDatabase(dbURL)
			if err != nil {
				t.Fatal(err)
			}

			orderID, _ := database.InsertOrder(test.order)

			order, err := database.GetOrder(orderID)
			if err != nil {
				t.Fatal(err)
			}

			if test.order.RestaurantID != order.RestaurantID {
				t.Errorf("RestaurantID: Expected: %d, Got: %d", test.order.RestaurantID, order.RestaurantID)
			}

			if test.order.FromUserID != order.FromUserID {
				t.Errorf("FromUserID: Expected: %d, Got: %d", test.order.FromUserID, order.FromUserID)
			}

			sort.Ints(test.order.OrderItems)
			sort.Ints(order.OrderItems)

			for i := 0; i < len(test.order.OrderItems); i++ {
				if test.order.OrderItems[i] != order.OrderItems[i] {
					t.Errorf("OrderItem %d: Expected: %d, Got: %d",
						i, test.order.OrderItems[i], order.OrderItems[i])
				}
			}

			_ = database.DeleteOrder(orderID)
		})
	}
}

func TestUpdateOrder(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name         string
		initialOrder domain.Order
		updatedOrder domain.Order
	}{
		{
			name: "Test Update Order",
			initialOrder: domain.Order{
				OrderID:      0,
				FromUserID:   210,
				RestaurantID: 2126847,
				OrderItems:   []int{85},
				Status: domain.OrderStatus{
					OrderID: 0,
					Status:  "",
				},
			},
			updatedOrder: domain.Order{
				OrderID:      0,
				FromUserID:   210,
				RestaurantID: 2126847,
				OrderItems:   []int{716410},
				Status: domain.OrderStatus{
					OrderID: 0,
					Status:  "",
				},
			},
		},
	}

	for _, currentTest := range tests {
		test := currentTest

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			err := configreader.SetConfigFile("config.yaml")
			if err != nil {
				t.Fatal(err)
			}

			dbURL := fmt.Sprintf("host=" + configreader.GetString("database.host") +
				" port=" + configreader.GetString("database.port") +
				" user=" + configreader.GetString("database.user") +
				" password=" + configreader.GetString("database.password") +
				" dbname=" + configreader.GetString("database.dbName") +
				" sslmode=" + configreader.GetString("database.sslmode"))

			database, err := db.NewDatabase(dbURL)
			if err != nil {
				t.Fatal(err)
			}

			orderID, _ := database.InsertOrder(test.initialOrder)

			test.updatedOrder.OrderID = orderID

			err = database.UpdateOrder(test.updatedOrder)
			if err != nil {
				t.Fatal(err)
			}

			_ = database.DeleteOrder(orderID)
		})
	}
}

func TestDeleteOrder(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		order domain.Order
	}{
		{
			name: "Test Delete Order",
			order: domain.Order{
				OrderID:      0,
				FromUserID:   1635,
				RestaurantID: 76,
				OrderItems:   []int{597590},
				Status: domain.OrderStatus{
					OrderID: 0,
					Status:  "",
				},
			},
		},
	}

	for _, currentTest := range tests {
		test := currentTest

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			err := configreader.SetConfigFile("config.yaml")
			if err != nil {
				t.Fatal(err)
			}

			dbURL := fmt.Sprintf("host=" + configreader.GetString("database.host") +
				" port=" + configreader.GetString("database.port") +
				" user=" + configreader.GetString("database.user") +
				" password=" + configreader.GetString("database.password") +
				" dbname=" + configreader.GetString("database.dbName") +
				" sslmode=" + configreader.GetString("database.sslmode"))

			database, err := db.NewDatabase(dbURL)
			if err != nil {
				t.Fatal(err)
			}

			orderID, _ := database.InsertOrder(test.order)

			err = database.DeleteOrder(orderID)
			if err != nil {
				t.Fatal(err)
			}
		})
	}
}

func TestUpdateOrderStatus(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		order     domain.Order
		newStatus string
	}{
		{
			name: "Test Insert Order",
			order: domain.Order{
				OrderID:      0,
				FromUserID:   36605,
				RestaurantID: 51553,
				OrderItems:   []int{1226842977},
				Status: domain.OrderStatus{
					OrderID: 0,
					Status:  "",
				},
			},
			newStatus: "cooking",
		},
	}

	for _, currentTest := range tests {
		test := currentTest

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			err := configreader.SetConfigFile("config.yaml")
			if err != nil {
				t.Fatal(err)
			}

			dbURL := fmt.Sprintf("host=" + configreader.GetString("database.host") +
				" port=" + configreader.GetString("database.port") +
				" user=" + configreader.GetString("database.user") +
				" password=" + configreader.GetString("database.password") +
				" dbname=" + configreader.GetString("database.dbName") +
				" sslmode=" + configreader.GetString("database.sslmode"))

			database, err := db.NewDatabase(dbURL)
			if err != nil {
				t.Fatal(err)
			}

			orderID, _ := database.InsertOrder(test.order)

			err = database.UpdateOrderStatus(orderID, test.newStatus)
			if err != nil {
				t.Fatal(err)
			}

			_ = database.DeleteOrder(orderID)
		})
	}
}
