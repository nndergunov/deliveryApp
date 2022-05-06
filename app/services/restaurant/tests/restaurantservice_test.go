package tests

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"strconv"
	"testing"

	v1 "github.com/nndergunov/deliveryApp/app/pkg/api/v1"
	"github.com/nndergunov/deliveryApp/app/pkg/api/v1/restaurantapi"
)

const baseAddr = "http://localhost:8084"

func TestCreateRestaurant(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		restaurant restaurantapi.RestaurantData
	}{
		{
			name: "mock test 1",
			restaurant: restaurantapi.RestaurantData{
				Name:    "Betting",
				City:    "Shaftsbury",
				Address: "Pontiac Street 903",
			},
		},
	}

	for _, currTest := range tests {
		test := currTest

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			req, err := v1.Encode(test.restaurant)
			if err != nil {
				t.Fatal(err)
			}

			resp, err := http.Post(baseAddr+"/v1/admin/restaurants",
				"application/json", bytes.NewBuffer(req))
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
				t.Error(err)
			}

			createdRest, err := restaurantapi.DecodeReturnRestaurant(respBody)
			if err != nil {
				t.Fatal(err)
			}

			if createdRest.Name != test.restaurant.Name {
				t.Errorf("Restaurant Name: expected: %s ; got: %s", createdRest.Name, test.restaurant.Name)
			}

			if createdRest.City != test.restaurant.City {
				t.Errorf("Restaurant City: expected: %s ; got: %s", createdRest.City, test.restaurant.City)
			}

			if createdRest.Address != test.restaurant.Address {
				t.Errorf("Restaurant Address: expected: %s ; got: %s",
					createdRest.Address, test.restaurant.Address)
			}

			deleter := http.DefaultClient

			delReq, err := http.NewRequest(http.MethodDelete,
				baseAddr+"/v1/admin/restaurants/"+strconv.Itoa(createdRest.ID), nil)
			if err != nil {
				t.Error(err)
			}

			_, err = deleter.Do(delReq)
			if err != nil {
				t.Errorf("Could not delete created restaurant: %v", err)
			}
		})
	}
}

func TestGetRestaurants(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		restaurant restaurantapi.RestaurantData
	}{
		{
			name: "mock test 1",
			restaurant: restaurantapi.RestaurantData{
				Name:    "Biggest",
				City:    "Lookout Mountain",
				Address: "Linda Road 2990",
			},
		},
	}

	for _, currTest := range tests {
		test := currTest

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			// Creating restaurant instance.
			createReq, err := v1.Encode(test.restaurant)
			if err != nil {
				t.Fatal(err)
			}

			createResp, err := http.Post(baseAddr+"/v1/admin/restaurants",
				"application/json", bytes.NewBuffer(createReq))
			if err != nil {
				t.Fatal(err)
			}

			createBody, err := ioutil.ReadAll(createResp.Body)
			if err != nil {
				t.Fatal(err)
			}

			err = createResp.Body.Close()
			if err != nil {
				t.Error(err)
			}

			createdRest, err := restaurantapi.DecodeReturnRestaurant(createBody)
			if err != nil {
				t.Fatal(err)
			}

			restaurantID := createdRest.ID

			// Get created restaurant.
			getResp, err := http.Get(baseAddr + "/v1/restaurants")
			if err != nil {
				t.Fatal(err)
			}

			if getResp.StatusCode != http.StatusOK {
				t.Fatalf("Response status: %d", getResp.StatusCode)
			}

			getBody, err := ioutil.ReadAll(getResp.Body)
			if err != nil {
				t.Fatal(err)
			}

			err = getResp.Body.Close()
			if err != nil {
				t.Error(err)
			}

			gotRests, err := restaurantapi.DecodeReturnRestaurantList(getBody)
			if err != nil {
				t.Fatal(err)
			}

			var found bool

			for _, gotRestaurant := range gotRests.List {
				if gotRestaurant.Name == test.restaurant.Name && gotRestaurant.City == test.restaurant.City &&
					gotRestaurant.Address == test.restaurant.Address && gotRestaurant.ID == restaurantID {
					found = true

					break
				}
			}

			if !found {
				t.Fatal("Service did not return created restaurant")
			}

			// Deleting restaurant instance.
			deleter := http.DefaultClient

			delReq, err := http.NewRequest(http.MethodDelete,
				baseAddr+"/v1/admin/restaurants/"+strconv.Itoa(restaurantID), nil)
			if err != nil {
				t.Error(err)
			}

			_, err = deleter.Do(delReq)
			if err != nil {
				t.Errorf("Could not delete created restaurant: %v", err)
			}
		})
	}
}

func TestUpdateRestaurant(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name          string
		restaurant    restaurantapi.RestaurantData
		updRestaurant restaurantapi.RestaurantData
	}{
		{
			name: "mock test 1",
			restaurant: restaurantapi.RestaurantData{
				Name:    "Decent",
				City:    "Holmes",
				Address: "Platforms Road 2077",
			},
			updRestaurant: restaurantapi.RestaurantData{
				Name:    "Reaches",
				City:    "Kosse",
				Address: "Livestock Street 5767",
			},
		},
	}

	for _, currTest := range tests {
		test := currTest

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			// Creating restaurant instance.
			createReq, err := v1.Encode(test.restaurant)
			if err != nil {
				t.Fatal(err)
			}

			createResp, err := http.Post(baseAddr+"/v1/admin/restaurants",
				"application/json", bytes.NewBuffer(createReq))
			if err != nil {
				t.Fatal(err)
			}

			createBody, err := ioutil.ReadAll(createResp.Body)
			if err != nil {
				t.Fatal(err)
			}

			err = createResp.Body.Close()
			if err != nil {
				t.Error(err)
			}

			createdRest, err := restaurantapi.DecodeReturnRestaurant(createBody)
			if err != nil {
				t.Fatal(err)
			}

			restaurantID := createdRest.ID

			// Update created restaurant.
			mockClient := http.DefaultClient

			updRestData, err := v1.Encode(test.updRestaurant)
			if err != nil {
				t.Fatal(err)
			}

			updRequest, err := http.NewRequest(http.MethodPut,
				baseAddr+"/v1/admin/restaurants/"+strconv.Itoa(restaurantID), bytes.NewBuffer(updRestData))
			if err != nil {
				t.Fatal(err)
			}

			updResponse, err := mockClient.Do(updRequest)
			if err != nil {
				t.Fatal(err)
			}

			if updResponse.StatusCode != http.StatusOK {
				t.Fatalf("Response status: %d", updResponse.StatusCode)
			}

			updBody, err := ioutil.ReadAll(updResponse.Body)
			if err != nil {
				t.Fatal(err)
			}

			err = updResponse.Body.Close()
			if err != nil {
				t.Error(err)
			}

			updatedRestaurant, err := restaurantapi.DecodeReturnRestaurant(updBody)
			if err != nil {
				t.Fatal(err)
			}

			if updatedRestaurant.ID != restaurantID {
				t.Errorf("UpdRestaurant ID: expected: %d ; got: %d", restaurantID, updatedRestaurant.ID)
			}

			if updatedRestaurant.Name != test.updRestaurant.Name {
				t.Errorf("UpdRestaurant Name: expected: %s ; got: %s", test.updRestaurant.Name, updatedRestaurant.Name)
			}

			if updatedRestaurant.City != test.updRestaurant.City {
				t.Errorf("UpdRestaurant City: expected: %s ; got: %s", test.updRestaurant.City, updatedRestaurant.City)
			}

			if updatedRestaurant.Address != test.updRestaurant.Address {
				t.Errorf("UpdRestaurant Address: expected: %s ; got: %s", test.updRestaurant.Address,
					updatedRestaurant.Address)
			}

			deleteRequest, err := http.NewRequest(http.MethodDelete,
				baseAddr+"/v1/admin/restaurants/"+strconv.Itoa(restaurantID), nil)
			if err != nil {
				t.Error(err)
			}

			_, err = mockClient.Do(deleteRequest)
			if err != nil {
				t.Errorf("Could not delete created restaurant: %v", err)
			}
		})
	}
}

func TestDeleteRestaurant(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		restaurant restaurantapi.RestaurantData
	}{
		{
			name: "mock test 1",
			restaurant: restaurantapi.RestaurantData{
				Name:    "Micro",
				City:    "Lemhi",
				Address: "Fuel Street 3854",
			},
		},
	}

	for _, currTest := range tests {
		test := currTest

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			// Creating restaurant instance.
			createReq, err := v1.Encode(test.restaurant)
			if err != nil {
				t.Fatal(err)
			}

			createResp, err := http.Post(baseAddr+"/v1/admin/restaurants",
				"application/json", bytes.NewBuffer(createReq))
			if err != nil {
				t.Fatal(err)
			}

			createBody, err := ioutil.ReadAll(createResp.Body)
			if err != nil {
				t.Fatal(err)
			}

			err = createResp.Body.Close()
			if err != nil {
				t.Error(err)
			}

			createdRest, err := restaurantapi.DecodeReturnRestaurant(createBody)
			if err != nil {
				t.Fatal(err)
			}

			restaurantID := createdRest.ID

			// Delete created restaurant.
			delReqMaker := http.DefaultClient

			delReq, err := http.NewRequest(http.MethodDelete,
				baseAddr+"/v1/admin/restaurants/"+strconv.Itoa(restaurantID), nil)
			if err != nil {
				t.Fatal(err)
			}

			delResp, err := delReqMaker.Do(delReq)
			if err != nil {
				t.Fatal(err)
			}

			if delResp.StatusCode != http.StatusOK {
				t.Fatalf("Response status: %d", delResp.StatusCode)
			}

			// Checking deletion
			getResp, err := http.Get(baseAddr + "/v1/restaurants")
			if err != nil {
				t.Fatal(err)
			}

			getBody, err := ioutil.ReadAll(getResp.Body)
			if err != nil {
				t.Fatal(err)
			}

			err = delResp.Body.Close()
			if err != nil {
				t.Error(err)
			}

			gotRests, err := restaurantapi.DecodeReturnRestaurantList(getBody)
			if err != nil {
				t.Fatal(err)
			}

			var found bool

			for _, gotRestaurant := range gotRests.List {
				if gotRestaurant.Name == test.restaurant.Name && gotRestaurant.City == test.restaurant.City &&
					gotRestaurant.Address == test.restaurant.Address && gotRestaurant.ID == restaurantID {
					found = true

					break
				}
			}

			if found {
				t.Fatal("Service did not delete created restaurant")
			}
		})
	}
}

func TestCreateMenu(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		restaurant restaurantapi.RestaurantData
		menu       restaurantapi.MenuData
	}{
		{
			name: "mock test 1",
			restaurant: restaurantapi.RestaurantData{
				Name:    "Notre",
				City:    "Lempster",
				Address: "Booty St 4206",
			},
			menu: restaurantapi.MenuData{
				MenuItems: []restaurantapi.MenuItemData{
					{
						ID:     0,
						Name:   "Happen",
						Course: "Same",
					},
					{
						ID:     0,
						Name:   "Suites",
						Course: "Outlined",
					},
				},
			},
		},
	}

	for _, currTest := range tests {
		test := currTest

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			// Creating restaurant instance.
			createRestReq, err := v1.Encode(test.restaurant)
			if err != nil {
				t.Fatal(err)
			}

			createRestResp, err := http.Post(baseAddr+"/v1/admin/restaurants",
				"application/json", bytes.NewBuffer(createRestReq))
			if err != nil {
				t.Fatal(err)
			}

			createRestResponseBody, err := ioutil.ReadAll(createRestResp.Body)
			if err != nil {
				t.Fatal(err)
			}

			err = createRestResp.Body.Close()
			if err != nil {
				t.Error(err)
			}

			createdRest, err := restaurantapi.DecodeReturnRestaurant(createRestResponseBody)
			if err != nil {
				t.Fatal(err)
			}

			restaurantID := createdRest.ID

			// Creating menu.
			createMenuReq, err := v1.Encode(test.menu)
			if err != nil {
				t.Fatal(err)
			}

			createMenuResp, err := http.Post(baseAddr+"/v1/admin/restaurants/"+strconv.Itoa(restaurantID)+"/menu",
				"application/json", bytes.NewBuffer(createMenuReq))
			if err != nil {
				t.Fatal(err)
			}

			createMenuRespBody, err := ioutil.ReadAll(createMenuResp.Body)
			if err != nil {
				t.Fatal(err)
			}

			err = createRestResp.Body.Close()
			if err != nil {
				t.Error(err)
			}

			createdMenu, err := restaurantapi.DecodeReturnMenu(createMenuRespBody)
			if err != nil {
				t.Fatal(err)
			}

			for _, menuItem := range test.menu.MenuItems {
				var found bool

				for _, createdMenuItem := range createdMenu.Items {
					if menuItem.Name == createdMenuItem.Name && menuItem.Course == createdMenuItem.Course {
						found = true

						break
					}
				}

				if !found {
					t.Errorf("Returned menu does not have an item specified in test: Name: %s, Course: %s",
						menuItem.Name, menuItem.Course)
				}
			}

			for _, createdMenuItem := range createdMenu.Items {
				var found bool

				for _, menuItem := range test.menu.MenuItems {
					if menuItem.Name == createdMenuItem.Name && menuItem.Course == createdMenuItem.Course {
						found = true

						break
					}
				}

				if !found {
					t.Errorf("Returned menu has an item not specified in test: Name: %s, Course: %s",
						createdMenuItem.Name, createdMenuItem.Course)
				}
			}

			// Deleting restaurant instance.
			restaurantDeleter := http.DefaultClient

			delReq, err := http.NewRequest(http.MethodDelete,
				baseAddr+"/v1/admin/restaurants/"+strconv.Itoa(restaurantID), nil)
			if err != nil {
				t.Error(err)
			}

			_, err = restaurantDeleter.Do(delReq)
			if err != nil {
				t.Errorf("Could not delete created restaurant: %v", err)
			}
		})
	}
}

func TestReturnMenu(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		restaurant restaurantapi.RestaurantData
		menu       restaurantapi.MenuData
	}{
		{
			name: "mock test 1",
			restaurant: restaurantapi.RestaurantData{
				Name:    "Norman",
				City:    "Tuan Forest",
				Address: "Tracy Road 6689",
			},
			menu: restaurantapi.MenuData{
				MenuItems: []restaurantapi.MenuItemData{
					{
						ID:     0,
						Name:   "Determined",
						Course: "Under",
					},
					{
						ID:     0,
						Name:   "Named",
						Course: "Edwards",
					},
				},
			},
		},
	}

	for _, currTest := range tests {
		test := currTest

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			// Creating restaurant instance.
			createRestReq, err := v1.Encode(test.restaurant)
			if err != nil {
				t.Fatal(err)
			}

			createRestResp, err := http.Post(baseAddr+"/v1/admin/restaurants",
				"application/json", bytes.NewBuffer(createRestReq))
			if err != nil {
				t.Fatal(err)
			}

			createRestResponseBody, err := ioutil.ReadAll(createRestResp.Body)
			if err != nil {
				t.Fatal(err)
			}

			err = createRestResp.Body.Close()
			if err != nil {
				t.Error(err)
			}

			createdRest, err := restaurantapi.DecodeReturnRestaurant(createRestResponseBody)
			if err != nil {
				t.Fatal(err)
			}

			restaurantID := createdRest.ID

			// Creating menu instance.
			createMenuReq, err := v1.Encode(test.menu)
			if err != nil {
				t.Fatal(err)
			}

			_, err = http.Post(baseAddr+"/v1/admin/restaurants/"+strconv.Itoa(restaurantID)+"/menu",
				"application/json", bytes.NewBuffer(createMenuReq))
			if err != nil {
				t.Fatal(err)
			}

			// Getting menu.
			getMenuResp, err := http.Get(baseAddr + "/v1/restaurants/" + strconv.Itoa(restaurantID) + "/menu")
			if err != nil {
				t.Fatal(err)
			}

			if getMenuResp.StatusCode != http.StatusOK {
				t.Fatalf("Response status: %d", getMenuResp.StatusCode)
			}

			getMenuRespBody, err := ioutil.ReadAll(getMenuResp.Body)
			if err != nil {
				t.Error(err)
			}

			err = getMenuResp.Body.Close()
			if err != nil {
				t.Error(err)
			}

			menu, err := restaurantapi.DecodeReturnMenu(getMenuRespBody)
			if err != nil {
				t.Error(err)
			}

			for _, menuItem := range test.menu.MenuItems {
				var found bool

				for _, createdMenuItem := range menu.Items {
					if menuItem.Name == createdMenuItem.Name && menuItem.Course == createdMenuItem.Course {
						found = true

						break
					}
				}

				if !found {
					t.Errorf("Returned menu does not have an item specified in test: Name: %s, Course: %s",
						menuItem.Name, menuItem.Course)
				}
			}

			for _, createdMenuItem := range menu.Items {
				var found bool

				for _, menuItem := range test.menu.MenuItems {
					if menuItem.Name == createdMenuItem.Name && menuItem.Course == createdMenuItem.Course {
						found = true

						break
					}
				}

				if !found {
					t.Errorf("Returned menu has an item not specified in test: Name: %s, Course: %s",
						createdMenuItem.Name, createdMenuItem.Course)
				}
			}

			// Deleting restaurant instance.
			restaurantDeleter := http.DefaultClient

			delReq, err := http.NewRequest(http.MethodDelete,
				baseAddr+"/v1/admin/restaurants/"+strconv.Itoa(restaurantID), nil)
			if err != nil {
				t.Error(err)
			}

			_, err = restaurantDeleter.Do(delReq)
			if err != nil {
				t.Errorf("Could not delete created restaurant: %v", err)
			}
		})
	}
}

func TestAddMenuItem(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		restaurant  restaurantapi.RestaurantData
		menu        restaurantapi.MenuData
		addMenuItem restaurantapi.MenuItemData
	}{
		{
			name: "mock test 1",
			restaurant: restaurantapi.RestaurantData{
				Name:    "Various",
				City:    "Mount Cory",
				Address: "Stations St 6584",
			},
			menu: restaurantapi.MenuData{
				MenuItems: []restaurantapi.MenuItemData{{
					ID:     0,
					Name:   "Tulsa",
					Course: "Accomplished",
				}},
			},
			addMenuItem: restaurantapi.MenuItemData{
				ID:     0,
				Name:   "Markers",
				Course: "Descending",
			},
		},
	}

	for _, currTest := range tests {
		test := currTest

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			// Creating restaurant instance.
			createRestReq, err := v1.Encode(test.restaurant)
			if err != nil {
				t.Fatal(err)
			}

			createRestResp, err := http.Post(baseAddr+"/v1/admin/restaurants",
				"application/json", bytes.NewBuffer(createRestReq))
			if err != nil {
				t.Fatal(err)
			}

			createRestResponseBody, err := ioutil.ReadAll(createRestResp.Body)
			if err != nil {
				t.Fatal(err)
			}

			err = createRestResp.Body.Close()
			if err != nil {
				t.Error(err)
			}

			createdRest, err := restaurantapi.DecodeReturnRestaurant(createRestResponseBody)
			if err != nil {
				t.Fatal(err)
			}

			restaurantID := createdRest.ID

			// Creating menu instance.
			createMenuReq, err := v1.Encode(test.menu)
			if err != nil {
				t.Fatal(err)
			}

			_, err = http.Post(baseAddr+"/v1/admin/restaurants/"+strconv.Itoa(restaurantID)+"/menu",
				"application/json", bytes.NewBuffer(createMenuReq))
			if err != nil {
				t.Fatal(err)
			}

			// Adding new menu item.
			mockClient := http.DefaultClient

			addMenuItem, err := v1.Encode(test.addMenuItem)
			if err != nil {
				t.Errorf("Could not delete created restaurant: %v", err)
			}

			addItemReq, err := http.NewRequest(http.MethodPut,
				baseAddr+"/v1/admin/restaurants/"+strconv.Itoa(restaurantID)+"/menu", bytes.NewBuffer(addMenuItem))
			if err != nil {
				t.Errorf("Could not delete created restaurant: %v", err)
			}

			addItemResp, err := mockClient.Do(addItemReq)
			if err != nil {
				t.Errorf("Could not delete created restaurant: %v", err)
			}

			if addItemResp.StatusCode != http.StatusOK {
				t.Fatalf("Response status: %d", addItemResp.StatusCode)
			}

			err = addItemResp.Body.Close()
			if err != nil {
				t.Error(err)
			}

			// Checking if item was added successfully.
			getMenuResp, err := http.Get(baseAddr + "/v1/restaurants/" + strconv.Itoa(restaurantID) + "/menu")
			if err != nil {
				t.Fatal(err)
			}

			getMenuRespBody, err := ioutil.ReadAll(getMenuResp.Body)
			if err != nil {
				t.Error(err)
			}

			err = getMenuResp.Body.Close()
			if err != nil {
				t.Error(err)
			}

			menu, err := restaurantapi.DecodeReturnMenu(getMenuRespBody)
			if err != nil {
				t.Error(err)
			}

			var found bool

			for _, createdMenuItem := range menu.Items {
				if test.addMenuItem.Name == createdMenuItem.Name && test.addMenuItem.Course == createdMenuItem.Course {
					found = true

					break
				}
			}

			if !found {
				t.Error("Returned menu does not have an added item")
			}

			// Deleting restaurant instance.
			delReq, err := http.NewRequest(http.MethodDelete,
				baseAddr+"/v1/admin/restaurants/"+strconv.Itoa(restaurantID), nil)
			if err != nil {
				t.Error(err)
			}

			_, err = mockClient.Do(delReq)
			if err != nil {
				t.Errorf("Could not delete created restaurant: %v", err)
			}
		})
	}
}

func TestUpdateMenuItem(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		restaurant restaurantapi.RestaurantData
		menu       restaurantapi.MenuData
		updItem    restaurantapi.MenuItemData
	}{
		{
			name: "mock test 1",
			restaurant: restaurantapi.RestaurantData{
				Name:    "Speeds",
				City:    "Jeromesville",
				Address: "Bacon Road 981",
			},
			menu: restaurantapi.MenuData{
				MenuItems: []restaurantapi.MenuItemData{{
					ID:     0,
					Name:   "Aggregate",
					Course: "Worker",
				}},
			},
			updItem: restaurantapi.MenuItemData{
				ID:     0,
				Name:   "Warned",
				Course: "Offerings",
			},
		},
	}

	for _, currTest := range tests {
		test := currTest

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			// Creating restaurant instance.
			createRestReq, err := v1.Encode(test.restaurant)
			if err != nil {
				t.Fatal(err)
			}

			createRestResp, err := http.Post(baseAddr+"/v1/admin/restaurants",
				"application/json", bytes.NewBuffer(createRestReq))
			if err != nil {
				t.Fatal(err)
			}

			createRestResponseBody, err := ioutil.ReadAll(createRestResp.Body)
			if err != nil {
				t.Fatal(err)
			}

			err = createRestResp.Body.Close()
			if err != nil {
				t.Error(err)
			}

			createdRest, err := restaurantapi.DecodeReturnRestaurant(createRestResponseBody)
			if err != nil {
				t.Fatal(err)
			}

			restaurantID := createdRest.ID

			// Creating menu instance.
			createMenuReq, err := v1.Encode(test.menu)
			if err != nil {
				t.Fatal(err)
			}

			createMenuResp, err := http.Post(baseAddr+"/v1/admin/restaurants/"+strconv.Itoa(restaurantID)+"/menu",
				"application/json", bytes.NewBuffer(createMenuReq))
			if err != nil {
				t.Fatal(err)
			}

			createMenuRespBody, err := ioutil.ReadAll(createMenuResp.Body)
			if err != nil {
				t.Fatal(err)
			}

			err = createMenuResp.Body.Close()
			if err != nil {
				t.Error(err)
			}

			createdMenu, err := restaurantapi.DecodeReturnMenu(createMenuRespBody)
			if err != nil {
				t.Fatal(err)
			}

			restaurantID = createdMenu.RestaurantID
			menuItemID := createdMenu.Items[0].ID

			// Updating menu item.
			mockCLient := http.DefaultClient

			updData, err := v1.Encode(test.updItem)
			if err != nil {
				t.Fatal(err)
			}

			updReq, err := http.NewRequest(http.MethodPatch,
				baseAddr+"/v1/admin/restaurants/"+strconv.Itoa(restaurantID)+"/menu/"+strconv.Itoa(menuItemID),
				bytes.NewBuffer(updData))
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

			updItem, err := restaurantapi.DecodeReturnMenuItem(updRespBody)
			if err != nil {
				t.Fatal(err)
			}

			if updItem.Name != test.updItem.Name {
				t.Errorf("Updated Menu Item Name: Expected: %s, Got: %s", test.updItem.Name, updItem.Name)
			}

			if updItem.Course != test.updItem.Course {
				t.Errorf("Updated Menu Item Course: Expected: %s, Got: %s", test.updItem.Course, updItem.Course)
			}

			// Deleting restaurant instance.
			delReq, err := http.NewRequest(http.MethodDelete,
				baseAddr+"/v1/admin/restaurants/"+strconv.Itoa(restaurantID), nil)
			if err != nil {
				t.Error(err)
			}

			_, err = mockCLient.Do(delReq)
			if err != nil {
				t.Errorf("Could not delete created restaurant: %v", err)
			}
		})
	}
}

func TestDeleteMenuItem(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		restaurant restaurantapi.RestaurantData
		menu       restaurantapi.MenuData
	}{
		{
			name: "mock test 1",
			restaurant: restaurantapi.RestaurantData{
				Name:    "Kansas",
				City:    "Port Costa",
				Address: "Nsw St 7662",
			},
			menu: restaurantapi.MenuData{
				MenuItems: []restaurantapi.MenuItemData{{
					ID:     0,
					Name:   "Encouraging",
					Course: "Global",
				}},
			},
		},
	}

	for _, currTest := range tests {
		test := currTest

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			// Creating restaurant instance.
			createRestReq, err := v1.Encode(test.restaurant)
			if err != nil {
				t.Fatal(err)
			}

			createRestResp, err := http.Post(baseAddr+"/v1/admin/restaurants",
				"application/json", bytes.NewBuffer(createRestReq))
			if err != nil {
				t.Fatal(err)
			}

			createRestResponseBody, err := ioutil.ReadAll(createRestResp.Body)
			if err != nil {
				t.Fatal(err)
			}

			err = createRestResp.Body.Close()
			if err != nil {
				t.Error(err)
			}

			createdRest, err := restaurantapi.DecodeReturnRestaurant(createRestResponseBody)
			if err != nil {
				t.Fatal(err)
			}

			restaurantID := createdRest.ID

			createMenuReq, err := v1.Encode(test.menu)
			if err != nil {
				t.Fatal(err)
			}

			createMenuResp, err := http.Post(baseAddr+"/v1/admin/restaurants/"+strconv.Itoa(restaurantID)+"/menu",
				"application/json", bytes.NewBuffer(createMenuReq))
			if err != nil {
				t.Fatal(err)
			}

			createMenuRespBody, err := ioutil.ReadAll(createMenuResp.Body)
			if err != nil {
				t.Fatal(err)
			}

			err = createMenuResp.Body.Close()
			if err != nil {
				t.Error(err)
			}

			createdMenu, err := restaurantapi.DecodeReturnMenu(createMenuRespBody)
			if err != nil {
				t.Fatal(err)
			}

			restaurantID = createdMenu.RestaurantID
			menuItemID := createdMenu.Items[0].ID

			// Updating menu item.
			mockCLient := http.DefaultClient

			delItemReq, err := http.NewRequest(http.MethodDelete,
				baseAddr+"/v1/admin/restaurants/"+strconv.Itoa(restaurantID)+"/menu/"+strconv.Itoa(menuItemID),
				nil)
			if err != nil {
				t.Fatal(err)
			}

			delItemResp, err := mockCLient.Do(delItemReq)
			if err != nil {
				t.Fatal(err)
			}

			if delItemResp.StatusCode != http.StatusOK {
				t.Fatalf("Response status: %d", delItemResp.StatusCode)
			}

			// Deleting restaurant instance.
			delReq, err := http.NewRequest(http.MethodDelete,
				baseAddr+"/v1/admin/restaurants/"+strconv.Itoa(restaurantID), nil)
			if err != nil {
				t.Error(err)
			}

			_, err = mockCLient.Do(delReq)
			if err != nil {
				t.Errorf("Could not delete created restaurant: %v", err)
			}
		})
	}
}
