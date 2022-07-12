//go:build integration
// +build integration

package db_test

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"testing"

	"github.com/adrianbrad/psqldocker"
	"github.com/adrianbrad/psqltest"
	"github.com/nndergunov/deliveryApp/app/pkg/configreader"
	"github.com/nndergunov/deliveryApp/app/services/order/pkg/db"
	"github.com/nndergunov/deliveryApp/app/services/order/pkg/domain"
)

func numToPointer(num int) *int {
	return &num
}

func TestMain(m *testing.M) {
	err := configreader.SetConfigFile("../../config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	var (
		usr           = configreader.GetString("database.user")
		password      = configreader.GetString("database.password")
		dbName        = configreader.GetString("database.dbName")
		containerName = "order_docker_test"
	)

	sqlFile, err := ioutil.ReadFile("order_test.sql")
	if ioErr != nil {
		log.Fatal(err)
	}

	sql := string(sqlFile)

	c, err := psqldocker.NewContainer(
		usr,
		password,
		dbName,
		psqldocker.WithContainerName(containerName),
		psqldocker.WithSql(sql),
	)
	if err != nil {
		log.Fatalf("err while creating new psql container: %s", err)
	}

	var ret int

	defer func() {
		err = c.Close()
		if err != nil {
			log.Fatalf("err while tearing down db container: %s", err)
		}

		os.Exit(ret)
	}()

	dsn := fmt.Sprintf(
		"user=%s "+
			"password=%s "+
			"dbname=%s "+
			"host=localhost "+
			"port=%s "+
			"sslmode=disable",
		usr,
		password,
		dbName,
		c.Port(),
	)

	psqltest.Register(dsn)

	ret = m.Run()
}

func TestGetAllOrders(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		order domain.Order
	}{
		{
			name: "Insert Order",
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

			sqlDB := psqltest.NewTransactionTestingDB(t)
			database := db.NewDatabaseFromSource(sqlDB)

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
			name:         "Insert Order",
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

			sqlDB := psqltest.NewTransactionTestingDB(t)
			database := db.NewDatabaseFromSource(sqlDB)

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
			name: "Insert Order",
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

			sqlDB := psqltest.NewTransactionTestingDB(t)
			database := db.NewDatabaseFromSource(sqlDB)

			orderID, err := database.InsertOrder(test.order)
			if err != nil {
				t.Fatal(err)
			}
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
			name: "Insert Order",
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

			sqlDB := psqltest.NewTransactionTestingDB(t)
			database := db.NewDatabaseFromSource(sqlDB)

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
			name: "Update Order",
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

			sqlDB := psqltest.NewTransactionTestingDB(t)
			database := db.NewDatabaseFromSource(sqlDB)

			orderID, _ := database.InsertOrder(test.initialOrder)

			test.updatedOrder.OrderID = orderID

			err := database.UpdateOrder(test.updatedOrder)
			if err != nil {
				t.Fatal(err)
			}
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
			name: "Delete Order",
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

			sqlDB := psqltest.NewTransactionTestingDB(t)
			database := db.NewDatabaseFromSource(sqlDB)

			orderID, _ := database.InsertOrder(test.order)

			err := database.DeleteOrder(orderID)
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
			name: "Insert Order",
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

			sqlDB := psqltest.NewTransactionTestingDB(t)
			database := db.NewDatabaseFromSource(sqlDB)

			orderID, _ := database.InsertOrder(test.order)

			err := database.UpdateOrderStatus(orderID, test.newStatus)
			if err != nil {
				t.Fatal(err)
			}
		})
	}
}
