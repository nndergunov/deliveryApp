package courierservice_test

import (
	"github.com/nndergunov/deliveryApp/app/pkg/api/v1"

	"bytes"
	"net/http"
	"strconv"
	"testing"

	"courier/api/v1/courierapi"
)

const baseAddr = "http://localhost:8081"

func TestInsertCourierEndpoint(t *testing.T) {

	tests := []struct {
		name        string
		courierData courierapi.NewCourierRequest
	}{
		{
			"Insert courier simple test",
			courierapi.NewCourierRequest{
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

			insertCourierResp, err := http.Post(baseAddr+"/v1/courier", "application/json", bytes.NewBuffer(reqBody))
			if err != nil {
				t.Fatal(err)
			}

			if insertCourierResp.StatusCode != http.StatusOK {
				t.Fatalf("Response status: %d", insertCourierResp.StatusCode)
			}

			if insertCourierResp.StatusCode != http.StatusOK {
				t.Fatalf("StatusCode: %d", insertCourierResp.StatusCode)
			}

			insertCourierRespData := courierapi.CourierResponse{}
			if err = courierapi.DecodeJSON(insertCourierResp.Body, &insertCourierRespData); err != nil {
				t.Fatal(err)
			}

			if err := insertCourierResp.Body.Close(); err != nil {
				t.Error(err)
			}

			if insertCourierRespData.ID < 1 {
				t.Errorf("ID: Expected : > 1, Got: %v", insertCourierRespData.ID)
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
				baseAddr+"/v1/courier/"+strconv.Itoa(int(insertCourierRespData.ID)), nil)
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
		courierData courierapi.NewCourierRequest
		delRespData string
	}{
		{
			"delete courier simple test",
			courierapi.NewCourierRequest{
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

			resp, err := http.Post(baseAddr+"/v1/courier", "application/json", bytes.NewBuffer(reqBody))
			if err != nil {
				t.Fatal(err)
			}

			if resp.StatusCode != http.StatusOK {
				t.Fatalf("Response status: %d", resp.StatusCode)
			}

			respData := courierapi.CourierResponse{}
			if err = courierapi.DecodeJSON(resp.Body, &respData); err != nil {
				t.Fatal(err)
			}

			if err := resp.Body.Close(); err != nil {
				t.Error(err)
			}

			if respData.ID < 1 {
				t.Errorf("ID: Expected : > 1, Got: %v", respData.ID)
			}
			// Deleting courier instance.
			deleter := http.DefaultClient

			delReq, err := http.NewRequest(http.MethodDelete,
				baseAddr+"/v1/courier/"+strconv.Itoa(int(respData.ID)), nil)
			if err != nil {
				t.Error(err)
			}

			delResp, err := deleter.Do(delReq)
			if err != nil {
				t.Errorf("Could not delete Insert courier: %v", err)
			}

			delRespData := ""
			if err = courierapi.DecodeJSON(delResp.Body, &delRespData); err != nil {
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
		initialCourierData courierapi.NewCourierRequest
		UpdatedCourierData courierapi.UpdateCourierRequest
	}{
		{
			"Update courier simple test",
			courierapi.NewCourierRequest{
				Username:  "TestUsername",
				Password:  "TestPassword",
				Firstname: "vasya",
				Lastname:  "",
				Email:     "vasya@gmail.com",
				Phone:     "123456789",
			},

			courierapi.UpdateCourierRequest{
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

			InsertCourierResp, err := http.Post(baseAddr+"/v1/courier", "application/json", bytes.NewBuffer(reqBody))
			if err != nil {
				t.Fatal(err)
			}

			if InsertCourierResp.StatusCode != http.StatusOK {
				t.Fatalf("Response status: %d", InsertCourierResp.StatusCode)
			}

			InsertCourierData := courierapi.CourierResponse{}
			if err = courierapi.DecodeJSON(InsertCourierResp.Body, &InsertCourierData); err != nil {
				t.Fatal(err)
			}

			if err := InsertCourierResp.Body.Close(); err != nil {
				t.Error(err)
			}

			if InsertCourierData.ID < 1 {
				t.Errorf("ID: Expected : > 1, Got: %v", InsertCourierData.ID)
			}

			updateCourierReqBody, err := v1.Encode(test.UpdatedCourierData)
			if err != nil {
				t.Fatal(err)
			}

			client2 := http.DefaultClient

			req, err := http.NewRequest(http.MethodPut,
				baseAddr+"/v1/courier/"+strconv.Itoa(int(InsertCourierData.ID)), bytes.NewBuffer(updateCourierReqBody))
			if err != nil {
				t.Error(err)
			}

			updateCourierResp, err := client2.Do(req)
			if err != nil {
				t.Errorf("Could not update courier: %v", err)
			}

			updatedCourier := courierapi.CourierResponse{}
			if err = courierapi.DecodeJSON(updateCourierResp.Body, &updatedCourier); err != nil {
				t.Fatal(err)
			}

			if err := updateCourierResp.Body.Close(); err != nil {
				t.Error(err)
			}

			if updatedCourier.ID != InsertCourierData.ID {
				t.Errorf("ID: Expected: %v, Got: %v", InsertCourierData.ID, updatedCourier.Firstname)
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
				baseAddr+"/v1/courier/"+strconv.Itoa(int(InsertCourierData.ID)), nil)
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
		initialCourierData courierapi.NewCourierRequest
	}{
		{
			"Update courier available simple test",
			courierapi.NewCourierRequest{
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

			InsertCourierResp, err := http.Post(baseAddr+"/v1/courier", "application/json", bytes.NewBuffer(InsertCourierReqBody))
			if err != nil {
				t.Fatal(err)
			}

			if InsertCourierResp.StatusCode != http.StatusOK {
				t.Fatalf("Response status: %d", InsertCourierResp.StatusCode)
			}

			InsertCourier := courierapi.CourierResponse{}
			if err = courierapi.DecodeJSON(InsertCourierResp.Body, &InsertCourier); err != nil {
				t.Fatal(err)
			}

			if err := InsertCourierResp.Body.Close(); err != nil {
				t.Error(err)
			}

			if InsertCourier.ID < 1 {
				t.Errorf("ID: Expected : > 1, Got: %v", InsertCourier.ID)
			}

			client2 := http.DefaultClient

			req, err := http.NewRequest(http.MethodPut,
				baseAddr+"/v1/courier/available/"+strconv.Itoa(int(InsertCourier.ID))+"?available="+strconv.FormatBool(!InsertCourier.Available), nil)
			if err != nil {
				t.Error(err)
			}

			updateCourierResp, err := client2.Do(req)
			if err != nil {
				t.Errorf("Could not update courier: %v", err)
			}

			updatedCourier := courierapi.CourierResponse{}
			if err = courierapi.DecodeJSON(updateCourierResp.Body, &updatedCourier); err != nil {
				t.Fatal(err)
			}

			if err := updateCourierResp.Body.Close(); err != nil {
				t.Error(err)
			}

			if updatedCourier.ID != InsertCourier.ID {
				t.Errorf("ID: Expected: %v, Got: %v", InsertCourier.ID, updatedCourier.Firstname)
			}

			if updatedCourier.Available == InsertCourier.Available {
				t.Errorf("Available: Expected: %v, Got: %v", !InsertCourier.Available, updatedCourier.Available)
			}

			// Deleting courier instance.
			deleter := http.DefaultClient

			delReq, err := http.NewRequest(http.MethodDelete,
				baseAddr+"/v1/courier/"+strconv.Itoa(int(InsertCourier.ID)), nil)
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

func TestGetAllCourierEndpoint(t *testing.T) {

	tests := []struct {
		name            string
		courierDataList []courierapi.NewCourierRequest
	}{
		{
			"TestGetAllCourierEndpoint test",
			[]courierapi.NewCourierRequest{
				courierapi.NewCourierRequest{
					Username:  "TestUsername",
					Password:  "TestPassword",
					Firstname: "courier1FName",
					Lastname:  "courier1LName",
					Email:     "courier1@gmail.com",
					Phone:     "111111111",
				},

				courierapi.NewCourierRequest{
					Username:  "TestUsername2",
					Password:  "TestPassword",
					Firstname: "courier2FName",
					Lastname:  "courier2LName",
					Email:     "courier2@gmail.com",
					Phone:     "222222222",
				},
			}},
	}

	for _, currentTest := range tests {
		test := currentTest

		t.Run(test.name, func(t *testing.T) {

			var InsertCourierList []courierapi.CourierResponse

			for _, courier := range test.courierDataList {

				InsertCourierReqBody, err := v1.Encode(courier)
				if err != nil {
					t.Fatal(err)
				}

				InsertCourierResp, err := http.Post(baseAddr+"/v1/courier", "application/json", bytes.NewBuffer(InsertCourierReqBody))
				if err != nil {
					t.Fatal(err)
				}

				if InsertCourierResp.StatusCode != http.StatusOK {
					t.Fatalf("Response status: %d", InsertCourierResp.StatusCode)
				}

				InsertCourier := courierapi.CourierResponse{}
				if err = courierapi.DecodeJSON(InsertCourierResp.Body, &InsertCourier); err != nil {
					t.Fatal(err)
				}

				if err := InsertCourierResp.Body.Close(); err != nil {
					t.Error(err)
				}

				InsertCourierList = append(InsertCourierList, InsertCourier)
			}

			getAllCourierResp, err := http.Get(baseAddr + "/v1/courier/all")
			if err != nil {
				t.Fatal(err)
			}

			if getAllCourierResp.StatusCode != http.StatusOK {
				t.Fatalf("Response status: %d", getAllCourierResp.StatusCode)
			}

			getAllCourier := courierapi.ReturnCourierResponseList{}
			if err = courierapi.DecodeJSON(getAllCourierResp.Body, &getAllCourier); err != nil {
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
					baseAddr+"/v1/courier/"+strconv.Itoa(int(InsertCourier.ID)), nil)
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
		courierData courierapi.NewCourierRequest
	}{
		{
			"TestGetAllCourierEndpoint test",
			courierapi.NewCourierRequest{
				Username:  "TestUsername",
				Password:  "TestPassword",
				Firstname: "courier1FName",
				Lastname:  "courier1LName",
				Email:     "courier1@gmail.com",
				Phone:     "111111111",
			}},
	}

	for _, currentTest := range tests {
		test := currentTest

		t.Run(test.name, func(t *testing.T) {

			InsertCoruierReqBody, err := v1.Encode(test.courierData)
			if err != nil {
				t.Fatal(err)
			}

			InsertCourierResp, err := http.Post(baseAddr+"/v1/courier", "application/json", bytes.NewBuffer(InsertCoruierReqBody))
			if err != nil {
				t.Fatal(err)
			}

			if InsertCourierResp.StatusCode != http.StatusOK {
				t.Fatalf("Response status: %d", InsertCourierResp.StatusCode)
			}

			InsertCourier := courierapi.CourierResponse{}
			if err = courierapi.DecodeJSON(InsertCourierResp.Body, &InsertCourier); err != nil {
				t.Fatal(err)
			}

			if err := InsertCourierResp.Body.Close(); err != nil {
				t.Error(err)
			}

			getCourierResp, err := http.Get(baseAddr + "/v1/courier/" + strconv.Itoa(int(InsertCourier.ID)))
			if err != nil {
				t.Fatal(err)
			}

			if getCourierResp.StatusCode != http.StatusOK {
				t.Fatalf("Response status: %d", getCourierResp.StatusCode)
			}

			getCourier := courierapi.CourierResponse{}
			if err = courierapi.DecodeJSON(getCourierResp.Body, &getCourier); err != nil {
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
				baseAddr+"/v1/courier/"+strconv.Itoa(int(InsertCourier.ID)), nil)
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

func TestInsertCourierLocationEndpoint(t *testing.T) {

	tests := []struct {
		name                string
		courierData         courierapi.NewCourierRequest
		courierLocationData courierapi.NewCourierLocationRequest
	}{
		{
			"InsertCourierLocationEndpoint simple test",
			courierapi.NewCourierRequest{
				Username:  "TestUsername",
				Password:  "TestPassword",
				Firstname: "TestFName",
				Lastname:  "TestLName",
				Email:     "TestEmail",
				Phone:     "TestPhone",
			},
			courierapi.NewCourierLocationRequest{
				Altitude:   "987654321",
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

			InsertCourierResp, err := http.Post(baseAddr+"/v1/courier", "application/json", bytes.NewBuffer(InsertCourierReqBody))
			if err != nil {
				t.Fatal(err)
			}

			if InsertCourierResp.StatusCode != http.StatusOK {
				t.Fatalf("Response status: %d", InsertCourierResp.StatusCode)
			}

			courierInsert := courierapi.CourierResponse{}
			if err = courierapi.DecodeJSON(InsertCourierResp.Body, &courierInsert); err != nil {
				t.Fatal(err)

			}

			courierLocationReqBody, err := v1.Encode(test.courierLocationData)
			if err != nil {
				t.Fatal(err)
			}

			courierLocationResp, err := http.Post(baseAddr+"/v1/courier/location/"+strconv.Itoa(int(courierInsert.ID)),
				"application/json", bytes.NewBuffer(courierLocationReqBody))

			if err != nil {
				t.Fatal(err)
			}

			if courierLocationResp.StatusCode != http.StatusOK {
				t.Fatalf("Response status: %d", courierLocationResp.StatusCode)
			}

			courierLocation := courierapi.CourierLocationResponse{}
			if err = courierapi.DecodeJSON(courierLocationResp.Body, &courierLocation); err != nil {
				t.Fatal(err)
			}

			if err := courierLocationResp.Body.Close(); err != nil {
				t.Error(err)
			}

			if courierLocation.CourierID < 1 {
				t.Errorf("CourierID: Expected : > 1, Got: %v", courierLocation.CourierID)
			}

			if courierLocation.Altitude != test.courierLocationData.Altitude {
				t.Errorf("Altitude: Expected: %s, Got: %s", test.courierLocationData.Altitude, courierLocation.Altitude)
			}

			if courierLocation.Longitude != test.courierLocationData.Longitude {
				t.Errorf("Longitude: Expected: %s, Got: %s", test.courierLocationData.Longitude, courierLocation.Longitude)
			}

			if courierLocation.Country != test.courierLocationData.Country {
				t.Errorf("Country: Expected: %s, Got: %s", test.courierLocationData.Country, courierLocation.Country)
			}

			if courierLocation.Region != test.courierLocationData.Region {
				t.Errorf("Region: Expected: %s, Got: %s", test.courierLocationData.Region, courierLocation.Region)
			}

			if courierLocation.Street != test.courierLocationData.Street {
				t.Errorf("Street: Expected: %s, Got: %s", test.courierLocationData.Street, courierLocation.Street)
			}

			if courierLocation.HomeNumber != test.courierLocationData.HomeNumber {
				t.Errorf("HomeNumber: Expected: %s, Got: %s", test.courierLocationData.HomeNumber, courierLocation.HomeNumber)
			}

			if courierLocation.Floor != test.courierLocationData.Floor {
				t.Errorf("Floor: Expected: %s, Got: %s", test.courierLocationData.Floor, courierLocation.Floor)
			}

			if courierLocation.Door != test.courierLocationData.Door {
				t.Errorf("Door: Expected: %s, Got: %s", test.courierLocationData.Door, courierLocation.Door)
			}

			// Deleting courier instance.
			deleter := http.DefaultClient

			delReq, err := http.NewRequest(http.MethodDelete,
				baseAddr+"/v1/courier/"+strconv.Itoa(int(courierInsert.ID)), nil)
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

func TestUpdateCourierLocationEndpoint(t *testing.T) {

	tests := []struct {
		name                       string
		courierData                courierapi.NewCourierRequest
		courierLocationInitialData courierapi.NewCourierLocationRequest
		courierLocationUpdatedData courierapi.UpdateCourierLocationRequest
	}{
		{
			"UpdateCourierLocationEndpoint simple test",
			courierapi.NewCourierRequest{
				Username:  "TestUsername",
				Password:  "TestPassword",
				Firstname: "TestFName",
				Lastname:  "TestLName",
				Email:     "TestEmail",
				Phone:     "TestPhone",
			},
			courierapi.NewCourierLocationRequest{
				Altitude:   "987654321",
				Longitude:  "123456789",
				Country:    "TestCountry",
				City:       "TestCity",
				Region:     "TestRegion",
				Street:     "TestStreet",
				HomeNumber: "TestHomeNumber",
				Floor:      "TestFloor",
				Door:       "TestDoor",
			},
			courierapi.UpdateCourierLocationRequest{
				Altitude:   "123456789",
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

			InsertCourierResp, err := http.Post(baseAddr+"/v1/courier", "application/json", bytes.NewBuffer(InsertCourierReqBody))
			if err != nil {
				t.Fatal(err)
			}

			if InsertCourierResp.StatusCode != http.StatusOK {
				t.Fatalf("Response status: %d", InsertCourierResp.StatusCode)
			}

			courierInsert := courierapi.CourierResponse{}
			if err = courierapi.DecodeJSON(InsertCourierResp.Body, &courierInsert); err != nil {
				t.Fatal(err)

			}

			courierLocationReqBody, err := v1.Encode(test.courierLocationInitialData)
			if err != nil {
				t.Fatal(err)
			}

			courierLocationResp, err := http.Post(baseAddr+"/v1/courier/location/"+strconv.Itoa(int(courierInsert.ID)),
				"application/json", bytes.NewBuffer(courierLocationReqBody))

			if err != nil {
				t.Fatal(err)
			}

			if courierLocationResp.StatusCode != http.StatusOK {
				t.Fatalf("Response status: %d", courierLocationResp.StatusCode)
			}

			if err := courierLocationResp.Body.Close(); err != nil {
				t.Error(err)
			}

			courierLocationUpdateReqBody, err := v1.Encode(test.courierLocationUpdatedData)
			if err != nil {
				t.Fatal(err)
			}

			client := http.Client{}

			req, err := http.NewRequest(http.MethodPut, baseAddr+"/v1/courier/location/"+strconv.Itoa(int(courierInsert.ID)), bytes.NewBuffer(courierLocationUpdateReqBody))

			if err != nil {
				t.Fatal(err)
			}
			courierLocationUpdateResp, err := client.Do(req)

			if err != nil {
				t.Fatal(err)
			}

			if courierLocationUpdateResp.StatusCode != http.StatusOK {
				t.Fatalf("Response status: %d", courierLocationUpdateResp.StatusCode)
			}

			courierLocationUpdated := courierapi.CourierLocationResponse{}
			if err = courierapi.DecodeJSON(courierLocationUpdateResp.Body, &courierLocationUpdated); err != nil {
				t.Fatal(err)

			}

			if err := courierLocationUpdateResp.Body.Close(); err != nil {
				t.Error(err)
			}

			if courierLocationUpdated.CourierID != courierInsert.ID {
				t.Errorf("CourierID: Expected :  %v , Got: %v", courierInsert.ID, courierLocationUpdated.CourierID)
			}

			if courierLocationUpdated.Altitude != test.courierLocationUpdatedData.Altitude {
				t.Errorf("Altitude: Expected: %s, Got: %s", test.courierLocationUpdatedData.Altitude, courierLocationUpdated.Altitude)
			}

			if courierLocationUpdated.Longitude != test.courierLocationUpdatedData.Longitude {
				t.Errorf("Longitude: Expected: %s, Got: %s", test.courierLocationUpdatedData.Longitude, courierLocationUpdated.Longitude)
			}

			if courierLocationUpdated.Country != test.courierLocationUpdatedData.Country {
				t.Errorf("Country: Expected: %s, Got: %s", test.courierLocationUpdatedData.Country, courierLocationUpdated.Country)
			}

			if courierLocationUpdated.Region != test.courierLocationUpdatedData.Region {
				t.Errorf("Region: Expected: %s, Got: %s", test.courierLocationUpdatedData.Region, courierLocationUpdated.Region)
			}

			if courierLocationUpdated.Street != test.courierLocationUpdatedData.Street {
				t.Errorf("Street: Expected: %s, Got: %s", test.courierLocationUpdatedData.Street, courierLocationUpdated.Street)
			}

			if courierLocationUpdated.HomeNumber != test.courierLocationUpdatedData.HomeNumber {
				t.Errorf("HomeNumber: Expected: %s, Got: %s", test.courierLocationUpdatedData.HomeNumber, courierLocationUpdated.HomeNumber)
			}

			if courierLocationUpdated.Floor != test.courierLocationUpdatedData.Floor {
				t.Errorf("Floor: Expected: %s, Got: %s", test.courierLocationUpdatedData.Floor, courierLocationUpdated.Floor)
			}

			if courierLocationUpdated.Door != test.courierLocationUpdatedData.Door {
				t.Errorf("Door: Expected: %s, Got: %s", test.courierLocationUpdatedData.Door, courierLocationUpdated.Door)
			}

			// Deleting courier instance.
			deleter := http.DefaultClient

			delReq, err := http.NewRequest(http.MethodDelete,
				baseAddr+"/v1/courier/"+strconv.Itoa(int(courierInsert.ID)), nil)
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

func TestGetCourierLocationEndpoint(t *testing.T) {

	tests := []struct {
		name                string
		courierData         courierapi.NewCourierRequest
		courierLocationData courierapi.NewCourierLocationRequest
	}{
		{
			"GetCourierLocationEndpoint simple test",
			courierapi.NewCourierRequest{
				Username:  "TestUsername",
				Password:  "TestPassword",
				Firstname: "TestFName",
				Lastname:  "TestLName",
				Email:     "TestEmail",
				Phone:     "TestPhone",
			},
			courierapi.NewCourierLocationRequest{
				Altitude:   "987654321",
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

			InsertCourierResp, err := http.Post(baseAddr+"/v1/courier", "application/json", bytes.NewBuffer(InsertCourierReqBody))
			if err != nil {
				t.Fatal(err)
			}

			if InsertCourierResp.StatusCode != http.StatusOK {
				t.Fatalf("Response status: %d", InsertCourierResp.StatusCode)
			}

			courierInsert := courierapi.CourierResponse{}
			if err = courierapi.DecodeJSON(InsertCourierResp.Body, &courierInsert); err != nil {
				t.Fatal(err)

			}

			InsertCourierLocationReqBody, err := v1.Encode(test.courierLocationData)
			if err != nil {
				t.Fatal(err)
			}

			InsertCourierLocationResp, err := http.Post(baseAddr+"/v1/courier/location/"+strconv.Itoa(int(courierInsert.ID)),
				"application/json", bytes.NewBuffer(InsertCourierLocationReqBody))

			if err != nil {
				t.Fatal(err)
			}

			if InsertCourierLocationResp.StatusCode != http.StatusOK {
				t.Fatalf("Response status: %d", InsertCourierLocationResp.StatusCode)
			}

			getCoruierLocationResp, err := http.Get(baseAddr + "/v1/courier/location/" + strconv.Itoa(int(courierInsert.ID)))

			if err != nil {
				t.Fatal(err)
			}

			if getCoruierLocationResp.StatusCode != http.StatusOK {
				t.Fatalf("Response status: %d", InsertCourierLocationResp.StatusCode)
			}

			getCourierLocation := courierapi.CourierLocationResponse{}
			if err = courierapi.DecodeJSON(getCoruierLocationResp.Body, &getCourierLocation); err != nil {
				t.Fatal(err)
			}

			if err := getCoruierLocationResp.Body.Close(); err != nil {
				t.Error(err)
			}

			if getCourierLocation.CourierID != courierInsert.ID {
				t.Errorf("CourierID: Expected : %v, Got: %v", courierInsert.ID, getCourierLocation.CourierID)
			}

			if getCourierLocation.Altitude != test.courierLocationData.Altitude {
				t.Errorf("Altitude: Expected: %s, Got: %s", test.courierLocationData.Altitude, getCourierLocation.Altitude)
			}

			if getCourierLocation.Longitude != test.courierLocationData.Longitude {
				t.Errorf("Longitude: Expected: %s, Got: %s", test.courierLocationData.Longitude, getCourierLocation.Longitude)
			}

			if getCourierLocation.Country != test.courierLocationData.Country {
				t.Errorf("Country: Expected: %s, Got: %s", test.courierLocationData.Country, getCourierLocation.Country)
			}

			if getCourierLocation.Region != test.courierLocationData.Region {
				t.Errorf("Region: Expected: %s, Got: %s", test.courierLocationData.Region, getCourierLocation.Region)
			}

			if getCourierLocation.Street != test.courierLocationData.Street {
				t.Errorf("Street: Expected: %s, Got: %s", test.courierLocationData.Street, getCourierLocation.Street)
			}

			if getCourierLocation.HomeNumber != test.courierLocationData.HomeNumber {
				t.Errorf("HomeNumber: Expected: %s, Got: %s", test.courierLocationData.HomeNumber, getCourierLocation.HomeNumber)
			}

			if getCourierLocation.Floor != test.courierLocationData.Floor {
				t.Errorf("Floor: Expected: %s, Got: %s", test.courierLocationData.Floor, getCourierLocation.Floor)
			}

			if getCourierLocation.Door != test.courierLocationData.Door {
				t.Errorf("Door: Expected: %s, Got: %s", test.courierLocationData.Door, getCourierLocation.Door)
			}

			// Deleting courier instance.
			deleter := http.DefaultClient

			delReq, err := http.NewRequest(http.MethodDelete,
				baseAddr+"/v1/courier/"+strconv.Itoa(int(courierInsert.ID)), nil)
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
