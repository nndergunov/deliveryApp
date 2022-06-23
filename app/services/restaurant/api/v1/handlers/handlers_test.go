package handlers_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	v1 "github.com/nndergunov/deliveryApp/app/pkg/api/v1"
	"github.com/nndergunov/deliveryApp/app/pkg/api/v1/restaurantapi"
	"github.com/nndergunov/deliveryApp/app/pkg/logger"
	"github.com/nndergunov/deliveryApp/app/services/restaurant/api/v1/handlers"
	"github.com/nndergunov/deliveryApp/app/services/restaurant/pkg/domain"
	"github.com/nndergunov/deliveryApp/app/services/restaurant/pkg/service/mockservice"
	"github.com/stretchr/testify/mock"
)

/*
var (
	MockReturnRestaurantData = domain.Restaurant{
		ID:      0,
		Name:    "Name",
		City:    "City",
		Address: "Address",
	}

	MockReturnMenuData = domain.Menu{
		RestaurantID: 0,
		Items:        []domain.MenuItem{MockReturnMenuItem},
	}

	MockReturnMenuItem = domain.MenuItem{
		ID:     0,
		MenuID: 0,
		Name:   "Name",
		Course: "Course",
	}
)

type MockService struct{}

func (m MockService) ReturnAllRestaurants() ([]domain.Restaurant, error) {
	return []domain.Restaurant{MockReturnRestaurantData}, nil
}

func (m MockService) ReturnRestaurant(_ int) (*domain.Restaurant, error) {
	return &MockReturnRestaurantData, nil
}

func (m MockService) CreateNewRestaurant(restaurantData domain.Restaurant) (*domain.Restaurant, error) {
	return &restaurantData, nil
}

func (m MockService) UpdateRestaurant(restaurantData domain.Restaurant) (*domain.Restaurant, error) {
	return &restaurantData, nil
}

func (m MockService) DeleteRestaurant(_ int) error {
	return nil
}

func (m MockService) ReturnMenu(_ int) (*domain.Menu, error) {
	return &MockReturnMenuData, nil
}

func (m MockService) CreateMenu(menuData domain.Menu) (*domain.Menu, error) {
	return &menuData, nil
}

func (m MockService) AddMenuItem(_ int, menuItem domain.MenuItem) (*domain.MenuItem, error) {
	return &menuItem, nil
}

func (m MockService) UpdateMenuItem(_ int, menuItem domain.MenuItem) (*domain.MenuItem, error) {
	return &menuItem, nil
}

func (m MockService) DeleteMenuItem(_ int, _ int) error {
	return nil
}
*/

func TestCreateRestaurantEndpoint(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		restaurantData domain.Restaurant
	}{
		{
			name: "Create restaurant simple",
			restaurantData: domain.Restaurant{
				ID:      0,
				Name:    "Name",
				City:    "City",
				Address: "Address",
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
	}{{name: "Get restaurants simple"}}

	for _, currTest := range tests {
		test := currTest

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			repo := mockservice.AppService{}

			repo.On("ReturnAllRestaurants").
				Return(&test.restaurantData, nil).
				Once()

			mockService := new(MockService)
			log := logger.NewLogger(os.Stdout, "Get restaurants simple")
			handler := handlers.NewEndpointHandler(mockService, log)

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
			"Update restaurant simple",
			domain.Restaurant{
				ID:      0,
				Name:    "Name",
				City:    "City",
				Address: "Address",
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

			mockService := new(MockService)
			log := logger.NewLogger(os.Stdout, test.name)
			handler := handlers.NewEndpointHandler(mockService, log)

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

	t.Run("Delete restaurant simple", func(t *testing.T) {
		t.Parallel()

		repo := mockservice.AppService{}

		repo.On("CreateNewRestaurant", mock.AnythingOfType("domain.Restaurant")).
			Return(&test.restaurantData, nil).
			Once()

		mockService := new(MockService)
		log := logger.NewLogger(os.Stdout, "Delete restaurant simple")
		handler := handlers.NewEndpointHandler(mockService, log)

		resp := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodDelete, "/v1/admin/restaurants/0", nil)

		handler.ServeHTTP(resp, req)

		if resp.Code != http.StatusOK {
			t.Fatalf("StatusCode: %d", resp.Code)
		}
	})
}

func TestCreateMenuEndpoint(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name         string
		menuItemData domain.MenuItem
	}{
		{
			"Create menu simple",
			domain.MenuItem{
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

			repo.On("CreateNewRestaurant", mock.AnythingOfType("domain.Restaurant")).
				Return(&test.restaurantData, nil).
				Once()

			mockService := new(MockService)
			log := logger.NewLogger(os.Stdout, test.name)
			handler := handlers.NewEndpointHandler(mockService, log)

			reqData, _ := v1.Encode(restaurantapi.MenuData{
				MenuItems: []restaurantapi.MenuItemData{{
					ID:     0,
					Name:   test.menuItemData.Name,
					Course: test.menuItemData.Course,
				}},
			})

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

			if respData.MenuItems[0].Name != test.menuItemData.Name {
				t.Errorf("Name: Expected: %s, Got: %s", test.menuItemData.Name, respData.MenuItems[0].Name)
			}

			if respData.MenuItems[0].Course != test.menuItemData.Course {
				t.Errorf("Course: Expected: %s, Got: %s", test.menuItemData.Course, respData.MenuItems[0].Course)
			}
		})
	}
}

func TestGetMenuEndpoint(t *testing.T) {
	t.Parallel()

	t.Run("Get menu simple", func(t *testing.T) {
		t.Parallel()

		repo := mockservice.AppService{}

		repo.On("CreateNewRestaurant", mock.AnythingOfType("domain.Restaurant")).
			Return(&test.restaurantData, nil).
			Once()

		mockService := new(MockService)
		log := logger.NewLogger(os.Stdout, "Get menu simple")
		handler := handlers.NewEndpointHandler(mockService, log)

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

		if respData.MenuItems[0].Name != MockReturnMenuItem.Name {
			t.Errorf("Name: Expected: %s, Got: %s", MockReturnMenuItem.Name, respData.MenuItems[0].Name)
		}

		if respData.MenuItems[0].Course != MockReturnMenuItem.Course {
			t.Errorf("Course: Expected: %s, Got: %s", MockReturnMenuItem.Course, respData.MenuItems[0].Course)
		}
	})
}

func TestAddMenuItemEndpoint(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name         string
		menuItemData domain.MenuItem
	}{
		{
			"Add menu item simple",
			domain.MenuItem{
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

			repo.On("CreateNewRestaurant", mock.AnythingOfType("domain.Restaurant")).
				Return(&test.restaurantData, nil).
				Once()

			mockService := new(MockService)
			log := logger.NewLogger(os.Stdout, test.name)
			handler := handlers.NewEndpointHandler(mockService, log)

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
			"Update menu item simple",
			domain.MenuItem{
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

			repo.On("CreateNewRestaurant", mock.AnythingOfType("domain.Restaurant")).
				Return(&test.restaurantData, nil).
				Once()

			mockService := new(MockService)
			log := logger.NewLogger(os.Stdout, test.name)
			handler := handlers.NewEndpointHandler(mockService, log)

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

	t.Run("Delete menu item simple", func(t *testing.T) {
		t.Parallel()

		repo := mockservice.AppService{}

		repo.On("CreateNewRestaurant", mock.AnythingOfType("domain.Restaurant")).
			Return(&test.restaurantData, nil).
			Once()

		mockService := new(MockService)
		log := logger.NewLogger(os.Stdout, "Delete menu item simple")
		handler := handlers.NewEndpointHandler(mockService, log)

		resp := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodDelete, "/v1/admin/restaurants/0/menu/0", nil)

		handler.ServeHTTP(resp, req)

		if resp.Code != http.StatusOK {
			t.Fatalf("StatusCode: %d", resp.Code)
		}
	})
}
