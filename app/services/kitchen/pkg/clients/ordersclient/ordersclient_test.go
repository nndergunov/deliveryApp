package ordersclient_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	v1 "github.com/nndergunov/deliveryApp/app/pkg/api/v1"
	"github.com/nndergunov/deliveryApp/app/pkg/api/v1/orderapi"
	"github.com/nndergunov/deliveryApp/app/services/kitchen/pkg/clients/ordersclient"
)

func TestGetIncompleteOrders(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		orderList orderapi.ReturnOrderList
	}{
		{
			name: "get incomplete orders test",
			orderList: orderapi.ReturnOrderList{
				Orders: []orderapi.ReturnOrder{
					{
						OrderID:      1,
						FromUserID:   1,
						RestaurantID: 0,
						OrderItems:   nil,
						Status:       "incomplete",
					},
					{
						OrderID:      2,
						FromUserID:   2,
						RestaurantID: 0,
						OrderItems:   nil,
						Status:       "incomplete",
					},
					{
						OrderID:      3,
						FromUserID:   3,
						RestaurantID: 0,
						OrderItems:   nil,
						Status:       "incomplete",
					},
				},
			},
		},
	}

	for _, currTest := range tests {
		test := currTest

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			mux := http.NewServeMux()

			mux.HandleFunc("/v1/orders", func(writer http.ResponseWriter, request *http.Request) {
				restData, err := v1.Encode(test.orderList)
				if err != nil {
					t.Fatal(err)
				}

				_, err = writer.Write(restData)
				if err != nil {
					t.Fatal(err)
				}
			})

			srv := httptest.NewServer(mux)

			ordrsClient := ordersclient.NewOrdersClient(srv.URL)

			orders, err := ordrsClient.GetIncompleteOrders(0)
			if err != nil {
				t.Fatal(err)
			}

			if len(orders.Orders) != len(test.orderList.Orders) {
				t.Errorf("wrong number of received errors: expected: %d; got: %d",
					len(test.orderList.Orders), len(orders.Orders))
			}

			for _, rcvdOrder := range orders.Orders {
				found := false

				for _, expOrdr := range test.orderList.Orders {
					if rcvdOrder.OrderID == expOrdr.OrderID {
						found = true

						break
					}
				}

				if !found {
					t.Errorf("got order that is not in the orderlist: %v", rcvdOrder)
				}
			}

			for _, expOrdr := range orders.Orders {
				found := false

				for _, rcvdOrder := range test.orderList.Orders {
					if rcvdOrder.OrderID == expOrdr.OrderID {
						found = true

						break
					}
				}

				if !found {
					t.Errorf("did not get order that is in the orderlist: %v", expOrdr)
				}
			}
		})
	}
}
