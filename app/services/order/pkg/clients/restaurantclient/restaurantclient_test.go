package restaurantclient_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	v1 "github.com/nndergunov/deliveryApp/app/pkg/api/v1"
	"github.com/nndergunov/deliveryApp/app/pkg/api/v1/restaurantapi"
	"github.com/nndergunov/deliveryApp/app/services/order/pkg/clients/restaurantclient"
	"github.com/nndergunov/deliveryApp/app/services/order/pkg/domain"
)

func TestCheckIfAvailable(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		restaurantData restaurantapi.ReturnRestaurant
	}{
		{
			name: "restaurant is unavailable test",
			restaurantData: restaurantapi.ReturnRestaurant{
				ID:              0,
				Name:            "test restaurant 1",
				AcceptingOrders: false,
				City:            "test city 1",
				Address:         "test address 1",
				Longitude:       0,
				Altitude:        0,
			},
		},
		{
			name: "restaurant is available test",
			restaurantData: restaurantapi.ReturnRestaurant{
				ID:              0,
				Name:            "test restaurant 2",
				AcceptingOrders: true,
				City:            "test city 2",
				Address:         "test address 2",
				Longitude:       0,
				Altitude:        0,
			},
		},
	}

	for _, currTest := range tests {
		test := currTest

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			mux := http.NewServeMux()

			mux.HandleFunc("/v1/restaurants/0", func(writer http.ResponseWriter, request *http.Request) {
				restData, err := v1.Encode(test.restaurantData)
				if err != nil {
					t.Fatal(err)
				}

				_, err = writer.Write(restData)
				if err != nil {
					t.Fatal(err)
				}
			})

			srv := httptest.NewServer(mux)

			restClient := restaurantclient.NewRestaurantClient(srv.URL)

			available, err := restClient.CheckIfAvailable(0)
			if err != nil {
				t.Fatal(err)
			}

			if available != test.restaurantData.AcceptingOrders {
				t.Fatalf("Availability Expected: %v; Got: %v", test.restaurantData.AcceptingOrders, available)
			}
		})
	}
}

func TestCalculateOrderPrice(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		menuData  restaurantapi.ReturnMenu
		orderData domain.Order
	}{
		{
			name: "restaurant is available test",
			menuData: restaurantapi.ReturnMenu{
				RestaurantID: 0,
				MenuItems: []restaurantapi.ReturnMenuItem{
					{
						ID:     1,
						Name:   "menu item 1",
						Price:  5,
						Course: "first",
					},
					{
						ID:     2,
						Name:   "menu item 2",
						Price:  2,
						Course: "second",
					},
					{
						ID:     3,
						Name:   "menu item 3",
						Price:  3,
						Course: "dessert",
					},
				},
			},
			orderData: domain.Order{
				OrderID:      0,
				FromUserID:   0,
				RestaurantID: 0,
				OrderItems:   []int{1, 2, 3, 3},
				Status:       domain.OrderStatus{},
			},
		},
	}

	for _, currTest := range tests {
		test := currTest

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			mux := http.NewServeMux()

			mux.HandleFunc("/v1/restaurants/0/menu", func(writer http.ResponseWriter, request *http.Request) {
				restData, err := v1.Encode(test.menuData)
				if err != nil {
					t.Fatal(err)
				}

				_, err = writer.Write(restData)
				if err != nil {
					t.Fatal(err)
				}
			})

			srv := httptest.NewServer(mux)

			restClient := restaurantclient.NewRestaurantClient(srv.URL)

			price, err := restClient.CalculateOrderPrice(test.orderData)
			if err != nil {
				t.Fatal(err)
			}

			var expPrice float64

			for _, orderedItemID := range test.orderData.OrderItems {
				for _, menuItem := range test.menuData.MenuItems {
					if menuItem.ID != orderedItemID {
						continue
					}

					expPrice += menuItem.Price
				}
			}

			if expPrice != price {
				t.Fatalf("Expected price: %f, Got: %f", expPrice, price)
			}
		})
	}
}
