package tests

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"sort"
	"strconv"
	"testing"

	v1 "github.com/nndergunov/deliveryApp/app/pkg/api/v1"
	"github.com/nndergunov/deliveryApp/app/pkg/api/v1/orderapi"
)

const baseAddr = "http://localhost:8083"

// Run this test suite on an empty database.

func TestCreateOrder(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		order orderapi.OrderData
	}{
		{
			name: "mock order 1",
			order: orderapi.OrderData{
				FromUserID:   6513,
				RestaurantID: 5,
				OrderItems:   []int{863866, 632, 821},
			},
		},
	}

	for _, currTest := range tests {
		test := currTest

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			postReq, err := v1.Encode(test.order)
			if err != nil {
				t.Fatal(err)
			}

			resp, err := http.Post(baseAddr+"/v1/orders", "application/json", bytes.NewBuffer(postReq))
			if err != nil {
				t.Fatal(err)
			}

			if resp.StatusCode != http.StatusOK {
				t.Fatalf("Response status: %d", resp.StatusCode)
			}

			respBody, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				t.Fatal(err)
			}

			err = resp.Body.Close()
			if err != nil {
				t.Fatal(err)
			}

			createdOrder, err := orderapi.DecodeReturnOrder(respBody)
			if err != nil {
				t.Fatal(err)
			}

			if createdOrder.FromUserID != test.order.FromUserID {
				t.Errorf("Order From User (id): Expected: %d, Got: %d",
					test.order.FromUserID, createdOrder.FromUserID)
			}

			if createdOrder.RestaurantID != test.order.RestaurantID {
				t.Errorf("Order RestaurantID: Expected: %d, Got: %d",
					test.order.RestaurantID, createdOrder.RestaurantID)
			}

			sort.Ints(createdOrder.OrderItems)
			sort.Ints(test.order.OrderItems)

			for i := 0; i < len(test.order.OrderItems) && i < len(createdOrder.OrderItems); i++ {
				if createdOrder.OrderItems[i] != test.order.OrderItems[i] {
					t.Errorf("Order Element%d: Expected: %d, Got: %d",
						i, test.order.OrderItems[i], createdOrder.OrderItems[i])
				}
			}

			if len(test.order.OrderItems) != len(createdOrder.OrderItems) {
				t.Errorf("Order Element Number: Expected: %d, Got: %d",
					len(createdOrder.OrderItems), len(test.order.OrderItems))
			}
		})
	}
}

func TestGetOrder(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		order orderapi.OrderData
	}{
		{
			name: "mock order 1",
			order: orderapi.OrderData{
				FromUserID:   887,
				RestaurantID: 25154,
				OrderItems:   []int{20, 3},
			},
		},
	}

	for _, currTest := range tests {
		test := currTest

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			postReq, err := v1.Encode(test.order)
			if err != nil {
				t.Fatal(err)
			}

			createResp, err := http.Post(baseAddr+"/v1/orders", "application/json", bytes.NewBuffer(postReq))
			if err != nil {
				t.Fatal(err)
			}

			createRespBody, err := ioutil.ReadAll(createResp.Body)
			if err != nil {
				t.Fatal(err)
			}

			err = createResp.Body.Close()
			if err != nil {
				t.Fatal(err)
			}

			createdOrder, err := orderapi.DecodeReturnOrder(createRespBody)
			if err != nil {
				t.Fatal(err)
			}

			orderID := createdOrder.OrderID

			getResp, err := http.Get(baseAddr + "/v1/orders/" + strconv.Itoa(orderID))
			if err != nil {
				t.Fatal(err)
			}

			if getResp.StatusCode != http.StatusOK {
				t.Fatalf("Response status: %d", getResp.StatusCode)
			}

			getRespBody, err := ioutil.ReadAll(getResp.Body)
			if err != nil {
				t.Fatal(err)
			}

			err = getResp.Body.Close()
			if err != nil {
				t.Fatal(err)
			}

			gotOrder, err := orderapi.DecodeReturnOrder(getRespBody)
			if err != nil {
				t.Fatal(err)
			}

			if gotOrder.FromUserID != test.order.FromUserID {
				t.Errorf("Order From User (id): Expected: %d, Got: %d",
					test.order.FromUserID, gotOrder.FromUserID)
			}

			if gotOrder.RestaurantID != test.order.RestaurantID {
				t.Errorf("Order RestaurantID: Expected: %d, Got: %d",
					test.order.RestaurantID, gotOrder.RestaurantID)
			}

			sort.Ints(gotOrder.OrderItems)
			sort.Ints(test.order.OrderItems)

			for i := 0; i < len(test.order.OrderItems) && i < len(gotOrder.OrderItems); i++ {
				if gotOrder.OrderItems[i] != test.order.OrderItems[i] {
					t.Errorf("Order Element%d: Expected: %d, Got: %d",
						i, test.order.OrderItems[i], gotOrder.OrderItems[i])
				}
			}

			if len(test.order.OrderItems) != len(gotOrder.OrderItems) {
				t.Errorf("Order Element Number: Expected: %d, Got: %d",
					len(gotOrder.OrderItems), len(test.order.OrderItems))
			}
		})
	}
}

func TestUpdateOrder(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		order    orderapi.OrderData
		updOrder orderapi.OrderData
	}{
		{
			name: "mock order 1",
			order: orderapi.OrderData{
				FromUserID:   9878,
				RestaurantID: 7373552,
				OrderItems:   []int{2516054},
			},
			updOrder: orderapi.OrderData{
				FromUserID:   3540,
				RestaurantID: 714163,
				OrderItems:   []int{2259},
			},
		},
	}

	for _, currTest := range tests {
		test := currTest

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			postReq, err := v1.Encode(test.order)
			if err != nil {
				t.Fatal(err)
			}

			createResp, err := http.Post(baseAddr+"/v1/orders", "application/json", bytes.NewBuffer(postReq))
			if err != nil {
				t.Fatal(err)
			}

			createRespBody, err := ioutil.ReadAll(createResp.Body)
			if err != nil {
				t.Fatal(err)
			}

			err = createResp.Body.Close()
			if err != nil {
				t.Fatal(err)
			}

			createdOrder, err := orderapi.DecodeReturnOrder(createRespBody)
			if err != nil {
				t.Fatal(err)
			}

			orderID := createdOrder.OrderID

			mockCLient := http.DefaultClient

			updBody, err := v1.Encode(test.updOrder)
			if err != nil {
				t.Fatal(err)
			}

			updReq, err := http.NewRequest(http.MethodPut, baseAddr+"/v1/orders/"+strconv.Itoa(orderID),
				bytes.NewBuffer(updBody))
			if err != nil {
				t.Fatal(err)
			}

			updResp, err := mockCLient.Do(updReq)
			if err != nil {
				t.Fatal(err)
			}

			if updResp.StatusCode != http.StatusOK {
				t.Fatalf("Response status: %d", updResp.StatusCode)
			}

			updRespBody, err := ioutil.ReadAll(updResp.Body)
			if err != nil {
				t.Fatal(err)
			}

			err = updResp.Body.Close()
			if err != nil {
				t.Fatal(err)
			}

			updOrder, err := orderapi.DecodeReturnOrder(updRespBody)
			if err != nil {
				t.Fatal(err)
			}

			if updOrder.FromUserID != test.updOrder.FromUserID {
				t.Errorf("Order From User (id): Expected: %d, Got: %d",
					test.updOrder.FromUserID, updOrder.FromUserID)
			}

			if updOrder.RestaurantID != test.updOrder.RestaurantID {
				t.Errorf("Order RestaurantID: Expected: %d, Got: %d",
					test.updOrder.RestaurantID, updOrder.RestaurantID)
			}

			sort.Ints(updOrder.OrderItems)
			sort.Ints(test.updOrder.OrderItems)

			for i := 0; i < len(test.updOrder.OrderItems) && i < len(updOrder.OrderItems); i++ {
				if createdOrder.OrderItems[i] != test.order.OrderItems[i] {
					t.Errorf("Order Element%d: Expected: %d, Got: %d",
						i, test.updOrder.OrderItems[i], updOrder.OrderItems[i])
				}
			}

			if len(test.updOrder.OrderItems) != len(updOrder.OrderItems) {
				t.Errorf("Order Element Number: Expected: %d, Got: %d",
					len(updOrder.OrderItems), len(test.updOrder.OrderItems))
			}
		})
	}
}

func TestGetIncompleteOrders(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name            string
		restaurantID    int
		incompleteOrder orderapi.OrderData
		completeOrder   orderapi.OrderData
	}{
		{
			name:         "mock order 1",
			restaurantID: 61,
			incompleteOrder: orderapi.OrderData{
				FromUserID:   739,
				RestaurantID: 61,
				OrderItems:   []int{6, 708},
			},
		},
	}

	for _, currTest := range tests {
		test := currTest

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			incompletePostReq, err := v1.Encode(test.incompleteOrder)
			if err != nil {
				t.Fatal(err)
			}

			_, err = http.Post(baseAddr+"/v1/orders", "application/json", bytes.NewBuffer(incompletePostReq))
			if err != nil {
				t.Fatal(err)
			}

			getResp, err := http.Get(
				baseAddr + "/v1/admin/orders/restaurant=" + strconv.Itoa(test.restaurantID) + "/incomplete")
			if err != nil {
				t.Fatal(err)
			}

			if getResp.StatusCode != http.StatusOK {
				t.Fatalf("Response status: %d", getResp.StatusCode)
			}

			getRespBody, err := ioutil.ReadAll(getResp.Body)
			if err != nil {
				t.Fatal(err)
			}

			err = getResp.Body.Close()
			if err != nil {
				t.Fatal(err)
			}

			gotOrders, err := orderapi.DecodeReturnOrderList(getRespBody)
			if err != nil {
				t.Fatal(err)
			}

			gotOrder := gotOrders.Orders[0]

			if gotOrder.FromUserID != test.incompleteOrder.FromUserID {
				t.Errorf("Order From User (id): Expected: %d, Got: %d",
					test.incompleteOrder.FromUserID, gotOrder.FromUserID)
			}

			if gotOrder.RestaurantID != test.incompleteOrder.RestaurantID {
				t.Errorf("Order RestaurantID: Expected: %d, Got: %d",
					test.incompleteOrder.RestaurantID, gotOrder.RestaurantID)
			}

			sort.Ints(gotOrder.OrderItems)
			sort.Ints(test.incompleteOrder.OrderItems)

			for i := 0; i < len(test.incompleteOrder.OrderItems) && i < len(gotOrder.OrderItems); i++ {
				if gotOrder.OrderItems[i] != test.incompleteOrder.OrderItems[i] {
					t.Errorf("Order Element%d: Expected: %d, Got: %d",
						i, test.incompleteOrder.OrderItems[i], gotOrder.OrderItems[i])
				}
			}

			if len(test.incompleteOrder.OrderItems) != len(gotOrder.OrderItems) {
				t.Errorf("Order Element Number: Expected: %d, Got: %d",
					len(gotOrder.OrderItems), len(test.incompleteOrder.OrderItems))
			}

			if len(gotOrders.Orders) != 1 {
				t.Errorf("Wrong number of received orders: %d", len(gotOrders.Orders))
			}
		})
	}
}
