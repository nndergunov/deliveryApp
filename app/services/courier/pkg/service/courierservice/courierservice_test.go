package courierservice_test

import (
	"bytes"
	courierapi2 "github.com/nndergunov/deliveryApp/app/pkg/api/v1/courierapi"
	"net/http"
	"strconv"
	"testing"

	"github.com/nndergunov/deliveryApp/app/pkg/api/v1"
)

const baseAddr = "http://localhost:8081"

func TestInsertCourierEndpoint(t *testing.T) {
	tests := []struct {
		name        string
		courierData courierapi2.NewCourierRequest
	}{
		{
			"Insert courier simple test",
			courierapi2.NewCourierRequest{
				Username:  "TestUsername",
				Password:  "TestPassword",
				Firstname: "vasya",
				Lastname:  "testLastname",
				Email:     "vasya@gmail.com",
				Phone:     "123456789",
			},
		},
	}

	for _, currentTest := range tests {
		test := currentTest

		t.Run(test.name, func(t *testing.T) {
			reqBody, err := v1.Encode(test.courierData)
			if err != nil {
				t.Fatal(err)
			}

			insertCourierResp, err := http.Post(baseAddr+"/v1/couriers", "application/json", bytes.NewBuffer(reqBody))
			if err != nil {
				t.Fatal(err)
			}

			if insertCourierResp.StatusCode != http.StatusOK {
				t.Fatalf("Response status: %d", insertCourierResp.StatusCode)
			}

			if insertCourierResp.StatusCode != http.StatusOK {
				t.Fatalf("StatusCode: %d", insertCourierResp.StatusCode)
			}

			insertCourierRespData := courierapi2.CourierResponse{}
			if err = courierapi2.DecodeJSON(insertCourierResp.Body, &insertCourierRespData); err != nil {
				t.Fatal(err)
			}

			if err := insertCourierResp.Body.Close(); err != nil {
				t.Error(err)
			}

			if insertCourierRespData.ID < 1 {
				t.Errorf("UserID: Expected : > 1, Got: %v", insertCourierRespData.ID)
			}

			if insertCourierRespData.Firstname != test.courierData.Firstname {
				t.Errorf("Firstname: Expected: %s, Got: %s", test.courierData.Firstname, insertCourierRespData.Firstname)
			}

			if insertCourierRespData.Lastname != test.courierData.Lastname {
				t.Errorf("Lastname: Expected: %s, Got: %s", test.courierData.Lastname, insertCourierRespData.Lastname)
			}

			if insertCourierRespData.Email != test.courierData.Email {
				t.Errorf("Email: Expected: %s, Got: %s", test.courierData.Email, insertCourierRespData.Email)
			}

			if insertCourierRespData.Phone != test.courierData.Phone {
				t.Errorf("Phone: Expected: %s, Got: %s", test.courierData.Phone, insertCourierRespData.Phone)
			}

			// Deleting courier instance.
			deleter := http.DefaultClient

			delReq, err := http.NewRequest(http.MethodDelete,
				baseAddr+"/v1/couriers/"+strconv.Itoa(int(insertCourierRespData.ID)), nil)
			if err != nil {
				t.Error(err)
			}

			_, err = deleter.Do(delReq)
			if err != nil {
				t.Errorf("Could not delete Insert courier: %v", err)
			}
		})
	}
}

func TestDeleteCourierEndpoint(t *testing.T) {
	tests := []struct {
		name        string
		courierData courierapi2.NewCourierRequest
		delRespData string
	}{
		{
			"delete courier simple test",
			courierapi2.NewCourierRequest{
				Username:  "TestUsername",
				Password:  "TestPassword",
				Firstname: "vasya",
				Lastname:  "testLastname",
				Email:     "vasya@gmail.com",
				Phone:     "123456789",
			},
			"courier deleted",
		},
	}

	for _, currentTest := range tests {
		test := currentTest

		t.Run(test.name, func(t *testing.T) {
			reqBody, err := v1.Encode(test.courierData)
			if err != nil {
				t.Fatal(err)
			}

			resp, err := http.Post(baseAddr+"/v1/couriers", "application/json", bytes.NewBuffer(reqBody))
			if err != nil {
				t.Fatal(err)
			}

			if resp.StatusCode != http.StatusOK {
				t.Fatalf("Response status: %d", resp.StatusCode)
			}

			respData := courierapi2.CourierResponse{}
			if err = courierapi2.DecodeJSON(resp.Body, &respData); err != nil {
				t.Fatal(err)
			}

			if err := resp.Body.Close(); err != nil {
				t.Error(err)
			}

			if respData.ID < 1 {
				t.Errorf("UserID: Expected : > 1, Got: %v", respData.ID)
			}
			// Deleting courier instance.
			deleter := http.DefaultClient

			delReq, err := http.NewRequest(http.MethodDelete,
				baseAddr+"/v1/couriers/"+strconv.Itoa(respData.ID), nil)
			if err != nil {
				t.Error(err)
			}

			delResp, err := deleter.Do(delReq)
			if err != nil {
				t.Errorf("Could not delete Insert courier: %v", err)
			}

			delRespData := ""
			if err = courierapi2.DecodeJSON(delResp.Body, &delRespData); err != nil {
				t.Fatal(err)
			}
			if delRespData != test.delRespData {
				t.Errorf("delRespData: Expected: %s, Got: %s", test.delRespData, delRespData)
			}
		})
	}
}

func TestUpdateCourierEndpoint(t *testing.T) {
	tests := []struct {
		name               string
		initialCourierData courierapi2.NewCourierRequest
		UpdatedCourierData courierapi2.UpdateCourierRequest
	}{
		{
			"Update courier simple test",
			courierapi2.NewCourierRequest{
				Username:  "TestUsername",
				Password:  "TestPassword",
				Firstname: "vasya",
				Lastname:  "",
				Email:     "vasya@gmail.com",
				Phone:     "123456789",
			},

			courierapi2.UpdateCourierRequest{
				Firstname: "updatedFName",
				Lastname:  "updatedLName",
				Email:     "updatedVasya@gmail.com",
				Phone:     "987654321",
			},
		},
	}

	for _, currentTest := range tests {
		test := currentTest

		t.Run(test.name, func(t *testing.T) {
			reqBody, err := v1.Encode(test.initialCourierData)
			if err != nil {
				t.Fatal(err)
			}

			InsertCourierResp, err := http.Post(baseAddr+"/v1/couriers", "application/json", bytes.NewBuffer(reqBody))
			if err != nil {
				t.Fatal(err)
			}

			if InsertCourierResp.StatusCode != http.StatusOK {
				t.Fatalf("Response status: %d", InsertCourierResp.StatusCode)
			}

			InsertCourierData := courierapi2.CourierResponse{}
			if err = courierapi2.DecodeJSON(InsertCourierResp.Body, &InsertCourierData); err != nil {
				t.Fatal(err)
			}

			if err := InsertCourierResp.Body.Close(); err != nil {
				t.Error(err)
			}

			if InsertCourierData.ID < 1 {
				t.Errorf("UserID: Expected : > 1, Got: %v", InsertCourierData.ID)
			}

			updateCourierReqBody, err := v1.Encode(test.UpdatedCourierData)
			if err != nil {
				t.Fatal(err)
			}

			client2 := http.DefaultClient

			req, err := http.NewRequest(http.MethodPut,
				baseAddr+"/v1/couriers/"+strconv.Itoa(InsertCourierData.ID), bytes.NewBuffer(updateCourierReqBody))
			if err != nil {
				t.Error(err)
			}

			updateCourierResp, err := client2.Do(req)
			if err != nil {
				t.Errorf("Could not update courier: %v", err)
			}

			updatedCourier := courierapi2.CourierResponse{}
			if err = courierapi2.DecodeJSON(updateCourierResp.Body, &updatedCourier); err != nil {
				t.Fatal(err)
			}

			if err := updateCourierResp.Body.Close(); err != nil {
				t.Error(err)
			}

			if updatedCourier.ID != InsertCourierData.ID {
				t.Errorf("UserID: Expected: %v, Got: %v", InsertCourierData.ID, updatedCourier.Firstname)
			}

			if updatedCourier.Firstname != test.UpdatedCourierData.Firstname {
				t.Errorf("Firstname: Expected: %s, Got: %s", test.UpdatedCourierData.Firstname, updatedCourier.Firstname)
			}

			if updatedCourier.Lastname != test.UpdatedCourierData.Lastname {
				t.Errorf("Lastname: Expected: %s, Got: %s", test.UpdatedCourierData.Lastname, updatedCourier.Lastname)
			}

			if updatedCourier.Email != test.UpdatedCourierData.Email {
				t.Errorf("Email: Expected: %s, Got: %s", test.UpdatedCourierData.Email, updatedCourier.Email)
			}

			if updatedCourier.Phone != test.UpdatedCourierData.Phone {
				t.Errorf("Phone: Expected: %s, Got: %s", test.UpdatedCourierData.Phone, updatedCourier.Phone)
			}

			// Deleting courier instance.
			deleter := http.DefaultClient

			delReq, err := http.NewRequest(http.MethodDelete,
				baseAddr+"/v1/couriers/"+strconv.Itoa(InsertCourierData.ID), nil)
			if err != nil {
				t.Error(err)
			}

			_, err = deleter.Do(delReq)
			if err != nil {
				t.Errorf("Could not delete Insert courier: %v", err)
			}
		})
	}
}

func TestUpdateCourierAvailableEndpoint(t *testing.T) {
	tests := []struct {
		name               string
		initialCourierData courierapi2.NewCourierRequest
	}{
		{
			"Update courier available simple test",
			courierapi2.NewCourierRequest{
				Username:  "TestUsername",
				Password:  "TestPassword",
				Firstname: "vasya",
				Lastname:  "",
				Email:     "vasya@gmail.com",
				Phone:     "123456789",
			},
		},
	}

	for _, currentTest := range tests {
		test := currentTest

		t.Run(test.name, func(t *testing.T) {
			InsertCourierReqBody, err := v1.Encode(test.initialCourierData)
			if err != nil {
				t.Fatal(err)
			}

			InsertCourierResp, err := http.Post(baseAddr+"/v1/couriers", "application/json", bytes.NewBuffer(InsertCourierReqBody))
			if err != nil {
				t.Fatal(err)
			}

			if InsertCourierResp.StatusCode != http.StatusOK {
				t.Fatalf("Response status: %d", InsertCourierResp.StatusCode)
			}

			InsertCourier := courierapi2.CourierResponse{}
			if err = courierapi2.DecodeJSON(InsertCourierResp.Body, &InsertCourier); err != nil {
				t.Fatal(err)
			}

			if err := InsertCourierResp.Body.Close(); err != nil {
				t.Error(err)
			}

			if InsertCourier.ID < 1 {
				t.Errorf("UserID: Expected : > 1, Got: %v", InsertCourier.ID)
			}

			client2 := http.DefaultClient

			req, err := http.NewRequest(http.MethodPut,
				baseAddr+"/v1/couriers-available/"+strconv.Itoa(InsertCourier.ID)+"?available="+strconv.FormatBool(!InsertCourier.Available), nil)
			if err != nil {
				t.Error(err)
			}

			updateCourierResp, err := client2.Do(req)
			if err != nil {
				t.Errorf("Could not update courier: %v", err)
			}

			updatedCourier := courierapi2.CourierResponse{}
			if err = courierapi2.DecodeJSON(updateCourierResp.Body, &updatedCourier); err != nil {
				t.Fatal(err)
			}

			if err := updateCourierResp.Body.Close(); err != nil {
				t.Error(err)
			}

			if updatedCourier.ID != InsertCourier.ID {
				t.Errorf("UserID: Expected: %v, Got: %v", InsertCourier.ID, updatedCourier.Firstname)
			}

			if updatedCourier.Available == InsertCourier.Available {
				t.Errorf("Available: Expected: %v, Got: %v", !InsertCourier.Available, updatedCourier.Available)
			}

			// Deleting courier instance.
			deleter := http.DefaultClient

			delReq, err := http.NewRequest(http.MethodDelete,
				baseAddr+"/v1/couriers/"+strconv.Itoa(InsertCourier.ID), nil)
			if err != nil {
				t.Error(err)
			}

			_, err = deleter.Do(delReq)
			if err != nil {
				t.Errorf("Could not delete Insert courier: %v", err)
			}
		})
	}
}

func TestGetCourierListEndpoint(t *testing.T) {
	tests := []struct {
		name            string
		courierDataList []courierapi2.NewCourierRequest
	}{
		{
			"TestGetCourierListEndpoint test",
			[]courierapi2.NewCourierRequest{
				{
					Username:  "TestUsername",
					Password:  "TestPassword",
					Firstname: "courier1FName",
					Lastname:  "courier1LName",
					Email:     "courier1@gmail.com",
					Phone:     "111111111",
				},

				{
					Username:  "TestUsername2",
					Password:  "TestPassword",
					Firstname: "courier2FName",
					Lastname:  "courier2LName",
					Email:     "courier2@gmail.com",
					Phone:     "222222222",
				},
			},
		},
	}

	for _, currentTest := range tests {
		test := currentTest

		t.Run(test.name, func(t *testing.T) {
			var InsertCourierList []courierapi2.CourierResponse

			for _, courier := range test.courierDataList {

				InsertCourierReqBody, err := v1.Encode(courier)
				if err != nil {
					t.Fatal(err)
				}

				InsertCourierResp, err := http.Post(baseAddr+"/v1/couriers", "application/json", bytes.NewBuffer(InsertCourierReqBody))
				if err != nil {
					t.Fatal(err)
				}

				if InsertCourierResp.StatusCode != http.StatusOK {
					t.Fatalf("Response status: %d", InsertCourierResp.StatusCode)
				}

				InsertCourier := courierapi2.CourierResponse{}
				if err = courierapi2.DecodeJSON(InsertCourierResp.Body, &InsertCourier); err != nil {
					t.Fatal(err)
				}

				if err := InsertCourierResp.Body.Close(); err != nil {
					t.Error(err)
				}

				InsertCourierList = append(InsertCourierList, InsertCourier)
			}

			getAllCourierResp, err := http.Get(baseAddr + "/v1/couriers")
			if err != nil {
				t.Fatal(err)
			}

			if getAllCourierResp.StatusCode != http.StatusOK {
				t.Fatalf("Response status: %d", getAllCourierResp.StatusCode)
			}

			getAllCourier := courierapi2.CourierResponseList{}
			if err = courierapi2.DecodeJSON(getAllCourierResp.Body, &getAllCourier); err != nil {
				t.Fatal(err)
			}

			if len(getAllCourier.CourierResponseList) != len(test.courierDataList) {
				t.Errorf("len: Expected: %v, Got: %v", len(test.courierDataList), len(getAllCourier.CourierResponseList))
			}

			if err := getAllCourierResp.Body.Close(); err != nil {
				t.Error(err)
			}

			// delete all courier we have Insert
			for _, InsertCourier := range InsertCourierList {

				// Deleting courier instance.
				deleter := http.DefaultClient

				delReq, err := http.NewRequest(http.MethodDelete,
					baseAddr+"/v1/couriers/"+strconv.Itoa(InsertCourier.ID), nil)
				if err != nil {
					t.Error(err)
				}

				_, err = deleter.Do(delReq)
				if err != nil {
					t.Errorf("Could not delete Insert courier: %v", err)
				}

			}
		})
	}
}

func TestGetCourierEndpoint(t *testing.T) {
	tests := []struct {
		name        string
		courierData courierapi2.NewCourierRequest
	}{
		{
			"TestGetCourierListEndpoint test",
			courierapi2.NewCourierRequest{
				Username:  "TestUsername",
				Password:  "TestPassword",
				Firstname: "courier1FName",
				Lastname:  "courier1LName",
				Email:     "courier1@gmail.com",
				Phone:     "111111111",
			},
		},
	}

	for _, currentTest := range tests {
		test := currentTest

		t.Run(test.name, func(t *testing.T) {
			InsertCoruierReqBody, err := v1.Encode(test.courierData)
			if err != nil {
				t.Fatal(err)
			}

			InsertCourierResp, err := http.Post(baseAddr+"/v1/couriers", "application/json", bytes.NewBuffer(InsertCoruierReqBody))
			if err != nil {
				t.Fatal(err)
			}

			if InsertCourierResp.StatusCode != http.StatusOK {
				t.Fatalf("Response status: %d", InsertCourierResp.StatusCode)
			}

			InsertCourier := courierapi2.CourierResponse{}
			if err = courierapi2.DecodeJSON(InsertCourierResp.Body, &InsertCourier); err != nil {
				t.Fatal(err)
			}

			if err := InsertCourierResp.Body.Close(); err != nil {
				t.Error(err)
			}

			getCourierResp, err := http.Get(baseAddr + "/v1/couriers/" + strconv.Itoa(InsertCourier.ID))
			if err != nil {
				t.Fatal(err)
			}

			if getCourierResp.StatusCode != http.StatusOK {
				t.Fatalf("Response status: %d", getCourierResp.StatusCode)
			}

			getCourier := courierapi2.CourierResponse{}
			if err = courierapi2.DecodeJSON(getCourierResp.Body, &getCourier); err != nil {
				t.Fatal(err)
			}

			if err := getCourierResp.Body.Close(); err != nil {
				t.Error(err)
			}

			if getCourier.Firstname != test.courierData.Firstname {
				t.Errorf("Firstname: Expected: %s, Got: %s", test.courierData.Firstname, getCourier.Firstname)
			}

			if getCourier.Lastname != test.courierData.Lastname {
				t.Errorf("Lastname: Expected: %s, Got: %s", test.courierData.Lastname, getCourier.Lastname)
			}

			if getCourier.Email != test.courierData.Email {
				t.Errorf("Email: Expected: %s, Got: %s", test.courierData.Email, getCourier.Email)
			}

			if getCourier.Phone != test.courierData.Phone {
				t.Errorf("Phone: Expected: %s, Got: %s", test.courierData.Phone, getCourier.Phone)
			}

			// Deleting courier instance.
			deleter := http.DefaultClient

			delReq, err := http.NewRequest(http.MethodDelete,
				baseAddr+"/v1/couriers/"+strconv.Itoa(InsertCourier.ID), nil)
			if err != nil {
				t.Error(err)
			}

			_, err = deleter.Do(delReq)
			if err != nil {
				t.Errorf("Could not delete Insert courier: %v", err)
			}
		})
	}
}

func TestInsertLocationEndpoint(t *testing.T) {
	tests := []struct {
		name         string
		courierData  courierapi2.NewCourierRequest
		locationData courierapi2.NewLocationRequest
	}{
		{
			"InsertLocationEndpoint simple test",
			courierapi2.NewCourierRequest{
				Username:  "TestUsername",
				Password:  "TestPassword",
				Firstname: "TestFName",
				Lastname:  "TestLName",
				Email:     "TestEmail",
				Phone:     "TestPhone",
			},
			courierapi2.NewLocationRequest{
				Latitude:   "987654321",
				Longitude:  "123456789",
				Country:    "TestCountry",
				City:       "TestCity",
				Region:     "TestRegion",
				Street:     "TestStreet",
				HomeNumber: "TestHomeNumber",
				Floor:      "TestFloor",
				Door:       "TestDoor",
			},
		},
	}

	for _, currentTest := range tests {
		test := currentTest

		t.Run(test.name, func(t *testing.T) {
			InsertCourierReqBody, err := v1.Encode(test.courierData)
			if err != nil {
				t.Fatal(err)
			}

			InsertCourierResp, err := http.Post(baseAddr+"/v1/couriers", "application/json", bytes.NewBuffer(InsertCourierReqBody))
			if err != nil {
				t.Fatal(err)
			}

			if InsertCourierResp.StatusCode != http.StatusOK {
				t.Fatalf("Response status: %d", InsertCourierResp.StatusCode)
			}

			courierInsert := courierapi2.CourierResponse{}
			if err = courierapi2.DecodeJSON(InsertCourierResp.Body, &courierInsert); err != nil {
				t.Fatal(err)
			}

			locationReqBody, err := v1.Encode(test.locationData)
			if err != nil {
				t.Fatal(err)
			}

			locationResp, err := http.Post(baseAddr+"/v1/locations/"+strconv.Itoa(courierInsert.ID),
				"application/json", bytes.NewBuffer(locationReqBody))
			if err != nil {
				t.Fatal(err)
			}

			if locationResp.StatusCode != http.StatusOK {
				t.Fatalf("Response status: %d", locationResp.StatusCode)
			}

			location := courierapi2.LocationResponse{}
			if err = courierapi2.DecodeJSON(locationResp.Body, &location); err != nil {
				t.Fatal(err)
			}

			if err := locationResp.Body.Close(); err != nil {
				t.Error(err)
			}

			if location.UserID < 1 {
				t.Errorf("UserID: Expected : > 1, Got: %v", location.UserID)
			}

			if location.Latitude != test.locationData.Latitude {
				t.Errorf("Latitude: Expected: %s, Got: %s", test.locationData.Latitude, location.Latitude)
			}

			if location.Longitude != test.locationData.Longitude {
				t.Errorf("Longitude: Expected: %s, Got: %s", test.locationData.Longitude, location.Longitude)
			}

			if location.Country != test.locationData.Country {
				t.Errorf("Country: Expected: %s, Got: %s", test.locationData.Country, location.Country)
			}

			if location.Region != test.locationData.Region {
				t.Errorf("Region: Expected: %s, Got: %s", test.locationData.Region, location.Region)
			}

			if location.Street != test.locationData.Street {
				t.Errorf("Street: Expected: %s, Got: %s", test.locationData.Street, location.Street)
			}

			if location.HomeNumber != test.locationData.HomeNumber {
				t.Errorf("HomeNumber: Expected: %s, Got: %s", test.locationData.HomeNumber, location.HomeNumber)
			}

			if location.Floor != test.locationData.Floor {
				t.Errorf("Floor: Expected: %s, Got: %s", test.locationData.Floor, location.Floor)
			}

			if location.Door != test.locationData.Door {
				t.Errorf("Door: Expected: %s, Got: %s", test.locationData.Door, location.Door)
			}

			// Deleting courier instance.
			deleter := http.DefaultClient

			delReq, err := http.NewRequest(http.MethodDelete,
				baseAddr+"/v1/couriers/"+strconv.Itoa(courierInsert.ID), nil)
			if err != nil {
				t.Error(err)
			}

			_, err = deleter.Do(delReq)
			if err != nil {
				t.Errorf("Could not delete Insert courier: %v", err)
			}
		})
	}
}

func TestUpdateLocationEndpoint(t *testing.T) {
	tests := []struct {
		name                string
		courierData         courierapi2.NewCourierRequest
		locationInitialData courierapi2.NewLocationRequest
		locationUpdatedData courierapi2.UpdateLocationRequest
	}{
		{
			"UpdateLocationEndpoint simple test",
			courierapi2.NewCourierRequest{
				Username:  "TestUsername",
				Password:  "TestPassword",
				Firstname: "TestFName",
				Lastname:  "TestLName",
				Email:     "TestEmail",
				Phone:     "TestPhone",
			},
			courierapi2.NewLocationRequest{
				Latitude:   "987654321",
				Longitude:  "123456789",
				Country:    "TestCountry",
				City:       "TestCity",
				Region:     "TestRegion",
				Street:     "TestStreet",
				HomeNumber: "TestHomeNumber",
				Floor:      "TestFloor",
				Door:       "TestDoor",
			},
			courierapi2.UpdateLocationRequest{
				Latitude:   "123456789",
				Longitude:  "987654321",
				Country:    "UpdatedTestCountry",
				City:       "UpdatedTestCity",
				Region:     "UpdatedTestRegion",
				Street:     "UpdatedTestStreet",
				HomeNumber: "UpdatedTestHomeNumber",
				Floor:      "UpdatedTestFloor",
				Door:       "UpdatedTestDoor",
			},
		},
	}

	for _, currentTest := range tests {
		test := currentTest

		t.Run(test.name, func(t *testing.T) {
			InsertCourierReqBody, err := v1.Encode(test.courierData)
			if err != nil {
				t.Fatal(err)
			}

			InsertCourierResp, err := http.Post(baseAddr+"/v1/couriers", "application/json", bytes.NewBuffer(InsertCourierReqBody))
			if err != nil {
				t.Fatal(err)
			}

			if InsertCourierResp.StatusCode != http.StatusOK {
				t.Fatalf("Response status: %d", InsertCourierResp.StatusCode)
			}

			courierInsert := courierapi2.CourierResponse{}
			if err = courierapi2.DecodeJSON(InsertCourierResp.Body, &courierInsert); err != nil {
				t.Fatal(err)
			}

			locationReqBody, err := v1.Encode(test.locationInitialData)
			if err != nil {
				t.Fatal(err)
			}

			locationResp, err := http.Post(baseAddr+"/v1/locations/"+strconv.Itoa(courierInsert.ID),
				"application/json", bytes.NewBuffer(locationReqBody))
			if err != nil {
				t.Fatal(err)
			}

			if locationResp.StatusCode != http.StatusOK {
				t.Fatalf("Response status: %d", locationResp.StatusCode)
			}

			if err := locationResp.Body.Close(); err != nil {
				t.Error(err)
			}

			locationUpdateReqBody, err := v1.Encode(test.locationUpdatedData)
			if err != nil {
				t.Fatal(err)
			}

			client := http.Client{}

			req, err := http.NewRequest(http.MethodPut, baseAddr+"/v1/locations/"+strconv.Itoa(courierInsert.ID), bytes.NewBuffer(locationUpdateReqBody))
			if err != nil {
				t.Fatal(err)
			}
			locationUpdateResp, err := client.Do(req)
			if err != nil {
				t.Fatal(err)
			}

			if locationUpdateResp.StatusCode != http.StatusOK {
				t.Fatalf("Response status: %d", locationUpdateResp.StatusCode)
			}

			locationUpdated := courierapi2.LocationResponse{}
			if err = courierapi2.DecodeJSON(locationUpdateResp.Body, &locationUpdated); err != nil {
				t.Fatal(err)
			}

			if err := locationUpdateResp.Body.Close(); err != nil {
				t.Error(err)
			}

			if locationUpdated.UserID != courierInsert.ID {
				t.Errorf("UserID: Expected :  %v , Got: %v", courierInsert.ID, locationUpdated.UserID)
			}

			if locationUpdated.Latitude != test.locationUpdatedData.Latitude {
				t.Errorf("Latitude: Expected: %s, Got: %s", test.locationUpdatedData.Latitude, locationUpdated.Latitude)
			}

			if locationUpdated.Longitude != test.locationUpdatedData.Longitude {
				t.Errorf("Longitude: Expected: %s, Got: %s", test.locationUpdatedData.Longitude, locationUpdated.Longitude)
			}

			if locationUpdated.Country != test.locationUpdatedData.Country {
				t.Errorf("Country: Expected: %s, Got: %s", test.locationUpdatedData.Country, locationUpdated.Country)
			}

			if locationUpdated.Region != test.locationUpdatedData.Region {
				t.Errorf("Region: Expected: %s, Got: %s", test.locationUpdatedData.Region, locationUpdated.Region)
			}

			if locationUpdated.Street != test.locationUpdatedData.Street {
				t.Errorf("Street: Expected: %s, Got: %s", test.locationUpdatedData.Street, locationUpdated.Street)
			}

			if locationUpdated.HomeNumber != test.locationUpdatedData.HomeNumber {
				t.Errorf("HomeNumber: Expected: %s, Got: %s", test.locationUpdatedData.HomeNumber, locationUpdated.HomeNumber)
			}

			if locationUpdated.Floor != test.locationUpdatedData.Floor {
				t.Errorf("Floor: Expected: %s, Got: %s", test.locationUpdatedData.Floor, locationUpdated.Floor)
			}

			if locationUpdated.Door != test.locationUpdatedData.Door {
				t.Errorf("Door: Expected: %s, Got: %s", test.locationUpdatedData.Door, locationUpdated.Door)
			}

			// Deleting courier instance.
			deleter := http.DefaultClient

			delReq, err := http.NewRequest(http.MethodDelete,
				baseAddr+"/v1/couriers/"+strconv.Itoa(courierInsert.ID), nil)
			if err != nil {
				t.Error(err)
			}

			_, err = deleter.Do(delReq)
			if err != nil {
				t.Errorf("Could not delete Insert courier: %v", err)
			}
		})
	}
}

func TestGetLocationEndpoint(t *testing.T) {
	tests := []struct {
		name         string
		courierData  courierapi2.NewCourierRequest
		locationData courierapi2.NewLocationRequest
	}{
		{
			"GetLocationEndpoint simple test",
			courierapi2.NewCourierRequest{
				Username:  "TestUsername",
				Password:  "TestPassword",
				Firstname: "TestFName",
				Lastname:  "TestLName",
				Email:     "TestEmail",
				Phone:     "TestPhone",
			},
			courierapi2.NewLocationRequest{
				Latitude:   "987654321",
				Longitude:  "123456789",
				Country:    "TestCountry",
				City:       "TestCity",
				Region:     "TestRegion",
				Street:     "TestStreet",
				HomeNumber: "TestHomeNumber",
				Floor:      "TestFloor",
				Door:       "TestDoor",
			},
		},
	}

	for _, currentTest := range tests {
		test := currentTest

		t.Run(test.name, func(t *testing.T) {
			InsertCourierReqBody, err := v1.Encode(test.courierData)
			if err != nil {
				t.Fatal(err)
			}

			InsertCourierResp, err := http.Post(baseAddr+"/v1/couriers", "application/json", bytes.NewBuffer(InsertCourierReqBody))
			if err != nil {
				t.Fatal(err)
			}

			if InsertCourierResp.StatusCode != http.StatusOK {
				t.Fatalf("Response status: %d", InsertCourierResp.StatusCode)
			}

			courierInsert := courierapi2.CourierResponse{}
			if err = courierapi2.DecodeJSON(InsertCourierResp.Body, &courierInsert); err != nil {
				t.Fatal(err)
			}

			InsertLocationReqBody, err := v1.Encode(test.locationData)
			if err != nil {
				t.Fatal(err)
			}

			InsertLocationResp, err := http.Post(baseAddr+"/v1/locations/"+strconv.Itoa(courierInsert.ID),
				"application/json", bytes.NewBuffer(InsertLocationReqBody))
			if err != nil {
				t.Fatal(err)
			}

			if InsertLocationResp.StatusCode != http.StatusOK {
				t.Fatalf("Response status: %d", InsertLocationResp.StatusCode)
			}

			getCourierLocationResp, err := http.Get(baseAddr + "/v1/locations/" + strconv.Itoa(courierInsert.ID))
			if err != nil {
				t.Fatal(err)
			}

			if getCourierLocationResp.StatusCode != http.StatusOK {
				t.Fatalf("Response status: %d", getCourierLocationResp.StatusCode)
			}

			getLocation := courierapi2.LocationResponse{}
			if err = courierapi2.DecodeJSON(getCourierLocationResp.Body, &getLocation); err != nil {
				t.Fatal(err)
			}

			if err := getCourierLocationResp.Body.Close(); err != nil {
				t.Error(err)
			}

			if getLocation.UserID != courierInsert.ID {
				t.Errorf("UserID: Expected : %v, Got: %v", courierInsert.ID, getLocation.UserID)
			}

			if getLocation.Latitude != test.locationData.Latitude {
				t.Errorf("Latitude: Expected: %s, Got: %s", test.locationData.Latitude, getLocation.Latitude)
			}

			if getLocation.Longitude != test.locationData.Longitude {
				t.Errorf("Longitude: Expected: %s, Got: %s", test.locationData.Longitude, getLocation.Longitude)
			}

			if getLocation.Country != test.locationData.Country {
				t.Errorf("Country: Expected: %s, Got: %s", test.locationData.Country, getLocation.Country)
			}

			if getLocation.Region != test.locationData.Region {
				t.Errorf("Region: Expected: %s, Got: %s", test.locationData.Region, getLocation.Region)
			}

			if getLocation.Street != test.locationData.Street {
				t.Errorf("Street: Expected: %s, Got: %s", test.locationData.Street, getLocation.Street)
			}

			if getLocation.HomeNumber != test.locationData.HomeNumber {
				t.Errorf("HomeNumber: Expected: %s, Got: %s", test.locationData.HomeNumber, getLocation.HomeNumber)
			}

			if getLocation.Floor != test.locationData.Floor {
				t.Errorf("Floor: Expected: %s, Got: %s", test.locationData.Floor, getLocation.Floor)
			}

			if getLocation.Door != test.locationData.Door {
				t.Errorf("Door: Expected: %s, Got: %s", test.locationData.Door, getLocation.Door)
			}

			// Deleting courier instance.
			deleter := http.DefaultClient

			delReq, err := http.NewRequest(http.MethodDelete,
				baseAddr+"/v1/couriers/"+strconv.Itoa(courierInsert.ID), nil)
			if err != nil {
				t.Error(err)
			}

			_, err = deleter.Do(delReq)
			if err != nil {
				t.Errorf("Could not delete Insert courier: %v", err)
			}
		})
	}
}

func TestGetLocationListEndpoint(t *testing.T) {
	tests := []struct {
		name         string
		courierData  courierapi2.NewCourierRequest
		locationData courierapi2.NewLocationRequest
	}{
		{
			"GetLocationListEndpoint simple test",
			courierapi2.NewCourierRequest{
				Username:  "TestUsername",
				Password:  "TestPassword",
				Firstname: "TestFName",
				Lastname:  "TestLName",
				Email:     "TestEmail",
				Phone:     "TestPhone",
			},
			courierapi2.NewLocationRequest{
				Latitude:   "987654321",
				Longitude:  "123456789",
				Country:    "TestCountry",
				City:       "TestCity",
				Region:     "",
				Street:     "",
				HomeNumber: "",
				Floor:      "",
				Door:       "",
			},
		},
	}

	for _, currentTest := range tests {
		test := currentTest

		t.Run(test.name, func(t *testing.T) {
			InsertCourierReqBody, err := v1.Encode(test.courierData)
			if err != nil {
				t.Fatal(err)
			}

			InsertCourierResp, err := http.Post(baseAddr+"/v1/couriers", "application/json", bytes.NewBuffer(InsertCourierReqBody))
			if err != nil {
				t.Fatal(err)
			}

			if InsertCourierResp.StatusCode != http.StatusOK {
				t.Fatalf("Response status: %d", InsertCourierResp.StatusCode)
			}

			courierInsert := courierapi2.CourierResponse{}
			if err = courierapi2.DecodeJSON(InsertCourierResp.Body, &courierInsert); err != nil {
				t.Fatal(err)
			}

			InsertLocationReqBody, err := v1.Encode(test.locationData)
			if err != nil {
				t.Fatal(err)
			}

			InsertLocationResp, err := http.Post(baseAddr+"/v1/locations/"+strconv.Itoa(courierInsert.ID),
				"application/json", bytes.NewBuffer(InsertLocationReqBody))
			if err != nil {
				t.Fatal(err)
			}

			if InsertLocationResp.StatusCode != http.StatusOK {
				t.Fatalf("Response status: %d", InsertLocationResp.StatusCode)
			}

			getLocationListResp, err := http.Get(baseAddr + "/v1/locations?city=" + test.locationData.City)
			if err != nil {
				t.Fatal(err)
			}

			if getLocationListResp.StatusCode != http.StatusOK {
				t.Fatalf("Response status: %d", getLocationListResp.StatusCode)
			}

			getLocationList := courierapi2.LocationResponseList{}
			if err = courierapi2.DecodeJSON(getLocationListResp.Body, &getLocationList); err != nil {
				t.Fatal(err)
			}

			if err := getLocationListResp.Body.Close(); err != nil {
				t.Error(err)
			}

			for _, getLocation := range getLocationList.LocationResponseList {

				if getLocation.UserID != courierInsert.ID {
					t.Errorf("UserID: Expected : %v, Got: %v", courierInsert.ID, getLocation.UserID)
				}

				if getLocation.Latitude != test.locationData.Latitude {
					t.Errorf("Latitude: Expected: %s, Got: %s", test.locationData.Latitude, getLocation.Latitude)
				}

				if getLocation.Longitude != test.locationData.Longitude {
					t.Errorf("Longitude: Expected: %s, Got: %s", test.locationData.Longitude, getLocation.Longitude)
				}

				if getLocation.Country != test.locationData.Country {
					t.Errorf("Country: Expected: %s, Got: %s", test.locationData.Country, getLocation.Country)
				}

				if getLocation.Region != test.locationData.Region {
					t.Errorf("Region: Expected: %s, Got: %s", test.locationData.Region, getLocation.Region)
				}

				if getLocation.Street != test.locationData.Street {
					t.Errorf("Street: Expected: %s, Got: %s", test.locationData.Street, getLocation.Street)
				}

				if getLocation.HomeNumber != test.locationData.HomeNumber {
					t.Errorf("HomeNumber: Expected: %s, Got: %s", test.locationData.HomeNumber, getLocation.HomeNumber)
				}

				if getLocation.Floor != test.locationData.Floor {
					t.Errorf("Floor: Expected: %s, Got: %s", test.locationData.Floor, getLocation.Floor)
				}

				if getLocation.Door != test.locationData.Door {
					t.Errorf("Door: Expected: %s, Got: %s", test.locationData.Door, getLocation.Door)
				}
			}

			// Deleting courier instance.
			deleter := http.DefaultClient

			delReq, err := http.NewRequest(http.MethodDelete,
				baseAddr+"/v1/couriers/"+strconv.Itoa(courierInsert.ID), nil)
			if err != nil {
				t.Error(err)
			}

			_, err = deleter.Do(delReq)
			if err != nil {
				t.Errorf("Could not delete Insert courier: %v", err)
			}
		})
	}
}
