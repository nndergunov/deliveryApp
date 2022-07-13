package handlers_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"

	v1 "github.com/nndergunov/deliveryApp/app/pkg/api/v1"
	"github.com/nndergunov/deliveryApp/app/pkg/logger"
	"github.com/nndergunov/deliveryApp/app/services/restaurant/api/v1/handlers"
	"github.com/nndergunov/deliveryApp/app/services/restaurant/api/v1/restaurantapi"
	"github.com/nndergunov/deliveryApp/app/services/restaurant/pkg/domain"
	"github.com/nndergunov/deliveryApp/app/services/restaurant/pkg/service/mockservice"
	"github.com/stretchr/testify/mock"
)

func TestCreateRestaurantEndpoint(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		restaurantData domain.Restaurant
	}{
		{
			name: "Create restaurant simple",
			restaurantData: domain.Restaurant{
				ID:              0,
				Name:            "Name",
				AcceptingOrders: true,
				City:            "City",
				Address:         "Address",
				Longitude:       0,
				Latitude:        0,
			},
		},
	}

	for _, currentTest := range tests {
		test := currentTest

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			repo := mockservice.AppService{}

			repo.On("CreateNewRestaurant", mock.AnythingOfType("domain.Restaurant")).
				Return(&test.restaurantData, nil).
				Once()

			log := logger.NewLogger(os.Stdout, test.name)
			handler := handlers.NewEndpointHandler(&repo, log)

			reqData, _ := v1.Encode(restaurantapi.RestaurantData{
				Name:    test.restaurantData.Name,
				City:    test.restaurantData.City,
				Address: test.restaurantData.Address,
			})

			resp := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPost, "/v1/admin/restaurants", bytes.NewBuffer(reqData))

			handler.ServeHTTP(resp, req)

			if resp.Code != http.StatusOK {
				t.Fatalf("StatusCode: %d", resp.Code)
			}

			respData := new(restaurantapi.ReturnRestaurant)

			err := v1.Decode(resp.Body.Bytes(), respData)
			if err != nil {
				t.Fatal(err)
			}

			if respData.Name != test.restaurantData.Name {
				t.Errorf("Name: Expected: %s, Got: %s", test.restaurantData.Name, respData.Name)
			}

			if respData.City != test.restaurantData.City {
				t.Errorf("City: Expected: %s, Got: %s", test.restaurantData.City, respData.City)
			}

			if respData.Address != test.restaurantData.Address {
				t.Errorf("Address: Expected: %s, Got: %s", test.restaurantData.Address, respData.Address)
			}
		})
	}
}

func TestGetRestaurantsEndpoint(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		restaurantData []domain.Restaurant
	}{
		{
			name: "Get restaurants simple",
			restaurantData: []domain.Restaurant{
				{
					ID:              0,
					Name:            "Name",
					AcceptingOrders: true,
					City:            "City",
					Address:         "Address",
					Longitude:       0,
					Latitude:        0,
				},
			},
		},
	}

	for _, currTest := range tests {
		test := currTest

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			repo := mockservice.AppService{}

			repo.On("ReturnAllRestaurants").
				Return(test.restaurantData, nil).
				Once()

			log := logger.NewLogger(os.Stdout, test.name)
			handler := handlers.NewEndpointHandler(&repo, log)

			resp := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, "/v1/restaurants", nil)

			handler.ServeHTTP(resp, req)

			if resp.Code != http.StatusOK {
				t.Fatalf("StatusCode: %d", resp.Code)
			}

			respData := new(restaurantapi.ReturnRestaurantList)

			err := v1.Decode(resp.Body.Bytes(), respData)
			if err != nil {
				t.Fatal(err)
			}

			if len(respData.List) != len(test.restaurantData) {
				t.Errorf(
					"wrong number of restraurants received: expected: %d, got: %d",
					len(respData.List),
					len(test.restaurantData),
				)
			}
		})
	}
}

func TestUpdateRestaurantEndpoint(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		restaurantData domain.Restaurant
	}{
		{
			name: "Update restaurant simple",
			restaurantData: domain.Restaurant{
				ID:              0,
				Name:            "Name",
				AcceptingOrders: true,
				City:            "City",
				Address:         "Address",
				Longitude:       0,
				Latitude:        0,
			},
		},
	}

	for _, currentTest := range tests {
		test := currentTest

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			repo := mockservice.AppService{}

			repo.On("UpdateRestaurant", mock.AnythingOfType("domain.Restaurant")).
				Return(&test.restaurantData, nil).
				Once()

			log := logger.NewLogger(os.Stdout, test.name)
			handler := handlers.NewEndpointHandler(&repo, log)

			reqData, _ := v1.Encode(restaurantapi.RestaurantData{
				Name:    test.restaurantData.Name,
				City:    test.restaurantData.City,
				Address: test.restaurantData.Address,
			})

			resp := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPut, "/v1/admin/restaurants/0", bytes.NewBuffer(reqData))

			handler.ServeHTTP(resp, req)

			if resp.Code != http.StatusOK {
				t.Fatalf("StatusCode: %d", resp.Code)
			}

			respData := new(restaurantapi.ReturnRestaurant)

			err := v1.Decode(resp.Body.Bytes(), respData)
			if err != nil {
				t.Fatal(err)
			}

			if respData.Name != test.restaurantData.Name {
				t.Errorf("Name: Expected: %s, Got: %s", test.restaurantData.Name, respData.Name)
			}

			if respData.City != test.restaurantData.City {
				t.Errorf("City: Expected: %s, Got: %s", test.restaurantData.City, respData.City)
			}

			if respData.Address != test.restaurantData.Address {
				t.Errorf("Address: Expected: %s, Got: %s", test.restaurantData.Address, respData.Address)
			}
		})
	}
}

func TestDeleteRestaurantEndpoint(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		restaurantData domain.Restaurant
	}{
		{
			name: "Update restaurant simple",
			restaurantData: domain.Restaurant{
				ID:              0,
				Name:            "Name",
				AcceptingOrders: true,
				City:            "City",
				Address:         "Address",
				Longitude:       0,
				Latitude:        0,
			},
		},
	}

	for _, currTest := range tests {
		test := currTest

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			repo := mockservice.AppService{}

			repo.On("DeleteRestaurant", mock.AnythingOfType("int")).
				Return(nil).
				Once()

			log := logger.NewLogger(os.Stdout, test.name)
			handler := handlers.NewEndpointHandler(&repo, log)

			resp := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodDelete, "/v1/admin/restaurants/0", nil)

			handler.ServeHTTP(resp, req)

			if resp.Code != http.StatusOK {
				t.Fatalf("StatusCode: %d", resp.Code)
			}
		})
	}
}

func TestCreateMenuEndpoint(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		menuData domain.Menu
	}{
		{
			name: "Create menu simple",
			menuData: domain.Menu{
				RestaurantID: 0,
				Items: []domain.MenuItem{
					{
						ID:     0,
						MenuID: 0,
						Name:   "Name",
						Price:  0,
						Course: "Course",
					},
				},
			},
		},
	}

	for _, currentTest := range tests {
		test := currentTest

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			repo := mockservice.AppService{}

			repo.On("CreateMenu", mock.AnythingOfType("domain.Menu")).
				Return(&test.menuData, nil).
				Once()

			log := logger.NewLogger(os.Stdout, test.name)
			handler := handlers.NewEndpointHandler(&repo, log)

			reqData, _ := v1.Encode(test.menuData)

			resp := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPost, "/v1/admin/restaurants/0/menu", bytes.NewBuffer(reqData))

			handler.ServeHTTP(resp, req)

			if resp.Code != http.StatusOK {
				t.Fatalf("StatusCode: %d", resp.Code)
			}

			respData := new(restaurantapi.ReturnMenu)

			err := v1.Decode(resp.Body.Bytes(), respData)
			if err != nil {
				t.Fatal(err)
			}

			if len(respData.MenuItems) != len(test.menuData.Items) {
				t.Errorf(
					"Wrong number of menu items: exprcted: %d, got: %d",
					len(test.menuData.Items),
					len(respData.MenuItems),
				)
			}
		})
	}
}

func TestGetMenuEndpoint(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		menuData domain.Menu
	}{
		{
			name: "Get menu simple",
			menuData: domain.Menu{
				RestaurantID: 0,
				Items: []domain.MenuItem{
					{
						ID:     0,
						MenuID: 0,
						Name:   "Name",
						Price:  0,
						Course: "Course",
					},
				},
			},
		},
	}

	for _, currTest := range tests {
		test := currTest

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			repo := mockservice.AppService{}

			repo.On("ReturnMenu", mock.AnythingOfType("int")).
				Return(&test.menuData, nil).
				Once()

			log := logger.NewLogger(os.Stdout, test.name)
			handler := handlers.NewEndpointHandler(&repo, log)

			resp := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, "/v1/restaurants/0/menu", nil)

			handler.ServeHTTP(resp, req)

			if resp.Code != http.StatusOK {
				t.Fatalf("StatusCode: %d", resp.Code)
			}

			respData := new(restaurantapi.ReturnMenu)

			err := v1.Decode(resp.Body.Bytes(), respData)
			if err != nil {
				t.Fatal(err)
			}

			if len(respData.MenuItems) != len(test.menuData.Items) {
				t.Errorf(
					"Wrong number of menu items: exprcted: %d, got: %d",
					len(test.menuData.Items),
					len(respData.MenuItems),
				)
			}
		})
	}
}

func TestAddMenuItemEndpoint(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name         string
		menuItemData domain.MenuItem
	}{
		{
			name: "Add menu item simple",
			menuItemData: domain.MenuItem{
				ID:     0,
				MenuID: 0,
				Name:   "Name",
				Price:  0,
				Course: "Course",
			},
		},
	}

	for _, currentTest := range tests {
		test := currentTest

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			repo := mockservice.AppService{}

			repo.On("AddMenuItem", mock.AnythingOfType("int"), mock.AnythingOfType("domain.MenuItem")).
				Return(&test.menuItemData, nil).
				Once()

			log := logger.NewLogger(os.Stdout, test.name)
			handler := handlers.NewEndpointHandler(&repo, log)

			reqData, _ := v1.Encode(restaurantapi.MenuItemData{
				ID:     0,
				Name:   test.menuItemData.Name,
				Course: test.menuItemData.Course,
			})

			resp := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPut, "/v1/admin/restaurants/0/menu", bytes.NewBuffer(reqData))

			handler.ServeHTTP(resp, req)

			if resp.Code != http.StatusOK {
				t.Fatalf("StatusCode: %d", resp.Code)
			}

			respData := new(restaurantapi.ReturnMenuItem)

			err := v1.Decode(resp.Body.Bytes(), respData)
			if err != nil {
				t.Fatal(err)
			}

			if respData.Name != test.menuItemData.Name {
				t.Errorf("Name: Expected: %s, Got: %s", test.menuItemData.Name, respData.Name)
			}

			if respData.Course != test.menuItemData.Course {
				t.Errorf("Course: Expected: %s, Got: %s", test.menuItemData.Course, respData.Course)
			}
		})
	}
}

func TestUpdateMenuItemEndpoint(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name         string
		menuItemData domain.MenuItem
	}{
		{
			name: "Update menu item simple",
			menuItemData: domain.MenuItem{
				ID:     0,
				MenuID: 0,
				Name:   "Name",
				Course: "Course",
			},
		},
	}

	for _, currentTest := range tests {
		test := currentTest

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			repo := mockservice.AppService{}

			repo.On("UpdateMenuItem", mock.AnythingOfType("int"), mock.AnythingOfType("domain.MenuItem")).
				Return(&test.menuItemData, nil).
				Once()

			log := logger.NewLogger(os.Stdout, test.name)
			handler := handlers.NewEndpointHandler(&repo, log)

			reqData, _ := v1.Encode(restaurantapi.MenuItemData{
				ID:     0,
				Name:   test.menuItemData.Name,
				Course: test.menuItemData.Course,
			})

			resp := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPatch, "/v1/admin/restaurants/0/menu/0", bytes.NewBuffer(reqData))

			handler.ServeHTTP(resp, req)

			if resp.Code != http.StatusOK {
				t.Fatalf("StatusCode: %d", resp.Code)
			}

			respData := new(restaurantapi.ReturnMenuItem)

			err := v1.Decode(resp.Body.Bytes(), respData)
			if err != nil {
				t.Fatal(err)
			}

			if respData.Name != test.menuItemData.Name {
				t.Errorf("Name: Expected: %s, Got: %s", test.menuItemData.Name, respData.Name)
			}

			if respData.Course != test.menuItemData.Course {
				t.Errorf("Course: Expected: %s, Got: %s", test.menuItemData.Course, respData.Course)
			}
		})
	}
}

func TestDeleteMenuItemEndpoint(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name         string
		restaurantID int
		menuItemID   int
	}{
		{
			name:         "Update restaurant simple",
			restaurantID: 0,
			menuItemID:   0,
		},
	}

	for _, currTest := range tests {
		test := currTest

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			repo := mockservice.AppService{}

			repo.On("DeleteMenuItem", mock.AnythingOfType("int"), mock.AnythingOfType("int")).
				Return(nil).
				Once()

			log := logger.NewLogger(os.Stdout, test.name)
			handler := handlers.NewEndpointHandler(&repo, log)

			resp := httptest.NewRecorder()
			req := httptest.NewRequest(
				http.MethodDelete,
				"/v1/admin/restaurants/"+strconv.Itoa(test.restaurantID)+"/menu/"+strconv.Itoa(test.menuItemID),
				nil,
			)

			handler.ServeHTTP(resp, req)

			if resp.Code != http.StatusOK {
				t.Fatalf("StatusCode: %d", resp.Code)
			}
		})
	}
}
