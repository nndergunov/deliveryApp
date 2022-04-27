package tests

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"testing"

	v1 "github.com/nndergunov/deliveryApp/app/pkg/api/v1"
	"github.com/nndergunov/deliveryApp/app/pkg/api/v1/restaurantapi"
)

const baseAddr = "http://localhost:8082"

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

			defer func(Body io.ReadCloser) {
				err := Body.Close()
				if err != nil {
					t.Error(err)
				}
			}(resp.Body)

			if resp.StatusCode != http.StatusOK {
				t.Fatalf("Response status: %d", resp.StatusCode)
			}

			respBody, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				t.Fatal(err)
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

			defer func(Body io.ReadCloser) {
				err := Body.Close()
				if err != nil {
					t.Error(err)
				}
			}(createResp.Body)

			createBody, err := ioutil.ReadAll(createResp.Body)
			if err != nil {
				t.Fatal(err)
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

			defer func(Body io.ReadCloser) {
				err := Body.Close()
				if err != nil {
					t.Error(err)
				}
			}(getResp.Body)

			if getResp.StatusCode != http.StatusOK {
				t.Fatalf("Response status: %d", getResp.StatusCode)
			}

			getBody, err := ioutil.ReadAll(getResp.Body)
			if err != nil {
				t.Fatal(err)
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
				}
			}

			if !found {
				t.Fatal("Service did not return created restaurant")
			}

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

			defer func(Body io.ReadCloser) {
				err := Body.Close()
				if err != nil {
					t.Error(err)
				}
			}(createResp.Body)

			createBody, err := ioutil.ReadAll(createResp.Body)
			if err != nil {
				t.Fatal(err)
			}

			createdRest, err := restaurantapi.DecodeReturnRestaurant(createBody)
			if err != nil {
				t.Fatal(err)
			}

			restaurantID := createdRest.ID

			// Update created restaurant.
			updater := http.DefaultClient

			updRestData, err := v1.Encode(test.updRestaurant)

			updRequest, err := http.NewRequest(http.MethodPut,
				baseAddr+"/v1/admin/restaurants/"+strconv.Itoa(restaurantID), bytes.NewBuffer(updRestData))
			if err != nil {
				t.Fatal(err)
			}

			updResponse, err := updater.Do(updRequest)
			if err != nil {
				t.Fatal(err)
			}

			defer func(Body io.ReadCloser) {
				err := Body.Close()
				if err != nil {
					t.Error(err)
				}
			}(updResponse.Body)

			if updResponse.StatusCode != http.StatusOK {
				t.Fatalf("Response status: %d", updResponse.StatusCode)
			}

			updBody, err := ioutil.ReadAll(updResponse.Body)
			if err != nil {
				t.Fatal(err)
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

			_, err = updater.Do(deleteRequest)
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

			defer func(Body io.ReadCloser) {
				err := Body.Close()
				if err != nil {
					t.Error(err)
				}
			}(createResp.Body)

			createBody, err := ioutil.ReadAll(createResp.Body)
			if err != nil {
				t.Fatal(err)
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

			defer func(Body io.ReadCloser) {
				err := Body.Close()
				if err != nil {
					t.Error(err)
				}
			}(getResp.Body)

			getBody, err := ioutil.ReadAll(getResp.Body)
			if err != nil {
				t.Fatal(err)
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
				}
			}

			if found {
				t.Fatal("Service did not delete created restaurant")
			}
		})
	}
}
