package consumerservice_test

import (
	"bytes"
	"net/http"
	"strconv"
	"testing"

	"github.com/nndergunov/deliveryApp/app/pkg/api/v1"

	"consumer/api/v1/consumerapi"
)

const baseAddr = "http://localhost:8080"

func TestInsertNewConsumerEndpoint(t *testing.T) {
	tests := []struct {
		name         string
		consumerData consumerapi.NewConsumerRequest
	}{
		{
			"Insert consumer simple test",
			consumerapi.NewConsumerRequest{
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
			reqBody, err := v1.Encode(test.consumerData)
			if err != nil {
				t.Fatal(err)
			}

			resp, err := http.Post(baseAddr+"/v1/consumers", "application/json", bytes.NewBuffer(reqBody))
			if err != nil {
				t.Fatal(err)
			}

			if resp.StatusCode != http.StatusOK {
				t.Fatalf("Response status: %d", resp.StatusCode)
			}

			if resp.StatusCode != http.StatusOK {
				t.Fatalf("StatusCode: %d", resp.StatusCode)
			}

			respData := consumerapi.ConsumerResponse{}
			if err = consumerapi.DecodeJSON(resp.Body, &respData); err != nil {
				t.Fatal(err)
			}

			if err := resp.Body.Close(); err != nil {
				t.Error(err)
			}

			if respData.ID < 1 {
				t.Errorf("ID: Expected : > 1, Got: %v", respData.ID)
			}

			if respData.Firstname != test.consumerData.Firstname {
				t.Errorf("Firstname: Expected: %s, Got: %s", test.consumerData.Firstname, respData.Firstname)
			}

			if respData.Lastname != test.consumerData.Lastname {
				t.Errorf("Lastname: Expected: %s, Got: %s", test.consumerData.Lastname, respData.Lastname)
			}

			if respData.Email != test.consumerData.Email {
				t.Errorf("Email: Expected: %s, Got: %s", test.consumerData.Email, respData.Email)
			}

			if respData.Phone != test.consumerData.Phone {
				t.Errorf("Phone: Expected: %s, Got: %s", test.consumerData.Phone, respData.Phone)
			}

			// Deleting consumer instance.
			deleter := http.DefaultClient

			delReq, err := http.NewRequest(http.MethodDelete,
				baseAddr+"/v1/consumers/"+strconv.Itoa(respData.ID), nil)
			if err != nil {
				t.Error(err)
			}

			_, err = deleter.Do(delReq)
			if err != nil {
				t.Errorf("Could not delete created consumer: %v", err)
			}
		})
	}
}

func TestDeleteConsumerEndpoint(t *testing.T) {
	tests := []struct {
		name         string
		consumerData consumerapi.NewConsumerRequest
		delRespData  string
	}{
		{
			"Insert consumer simple test",
			consumerapi.NewConsumerRequest{
				Firstname: "vasya",
				Lastname:  "testLastname",
				Email:     "vasya@gmail.com",
				Phone:     "123456789",
			},
			"Consumer deleted",
		},
	}

	for _, currentTest := range tests {
		test := currentTest

		t.Run(test.name, func(t *testing.T) {
			reqBody, err := v1.Encode(test.consumerData)
			if err != nil {
				t.Fatal(err)
			}

			resp, err := http.Post(baseAddr+"/v1/consumers", "application/json", bytes.NewBuffer(reqBody))
			if err != nil {
				t.Fatal(err)
			}

			if resp.StatusCode != http.StatusOK {
				t.Fatalf("Response status: %d", resp.StatusCode)
			}

			if resp.StatusCode != http.StatusOK {
				t.Fatalf("StatusCode: %d", resp.StatusCode)
			}

			respData := consumerapi.ConsumerResponse{}
			if err = consumerapi.DecodeJSON(resp.Body, &respData); err != nil {
				t.Fatal(err)
			}

			if err := resp.Body.Close(); err != nil {
				t.Error(err)
			}

			if respData.ID < 1 {
				t.Errorf("ID: Expected : > 1, Got: %v", respData.ID)
			}

			// Deleting consumer instance.
			deleter := http.DefaultClient

			delReq, err := http.NewRequest(http.MethodDelete,
				baseAddr+"/v1/consumers/"+strconv.Itoa(int(respData.ID)), nil)
			if err != nil {
				t.Error(err)
			}

			delResp, err := deleter.Do(delReq)
			if err != nil {
				t.Errorf("Could not delete created consumer: %v", err)
			}

			delRespData := ""
			if err = consumerapi.DecodeJSON(delResp.Body, &delRespData); err != nil {
				t.Fatal(err)
			}
			if delRespData != test.delRespData {
				t.Errorf("delRespData: Expected: %s, Got: %s", test.delRespData, delRespData)
			}
		})
	}
}

func TestUpdateConsumerEndpoint(t *testing.T) {
	tests := []struct {
		name                string
		initialConsumerData consumerapi.NewConsumerRequest
		UpdatedConsumerData consumerapi.UpdateConsumerRequest
	}{
		{
			"Update consumer simple test",
			consumerapi.NewConsumerRequest{
				Firstname: "vasya",
				Lastname:  "",
				Email:     "vasya@gmail.com",
				Phone:     "123456789",
			},

			consumerapi.UpdateConsumerRequest{
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
			reqBody, err := v1.Encode(test.initialConsumerData)
			if err != nil {
				t.Fatal(err)
			}

			resp1, err := http.Post(baseAddr+"/v1/consumers", "application/json", bytes.NewBuffer(reqBody))
			if err != nil {
				t.Fatal(err)
			}

			if resp1.StatusCode != http.StatusOK {
				t.Fatalf("Response status: %d", resp1.StatusCode)
			}

			respData := consumerapi.ConsumerResponse{}
			if err = consumerapi.DecodeJSON(resp1.Body, &respData); err != nil {
				t.Fatal(err)
			}

			if err := resp1.Body.Close(); err != nil {
				t.Error(err)
			}

			if respData.ID < 1 {
				t.Errorf("ID: Expected : > 1, Got: %v", respData.ID)
			}

			reqBody2, err := v1.Encode(test.UpdatedConsumerData)
			if err != nil {
				t.Fatal(err)
			}

			client2 := http.DefaultClient

			req, err := http.NewRequest(http.MethodPut,
				baseAddr+"/v1/consumers/"+strconv.Itoa(int(respData.ID)), bytes.NewBuffer(reqBody2))
			if err != nil {
				t.Error(err)
			}

			resp2, err := client2.Do(req)
			if err != nil {
				t.Errorf("Could not update consumer: %v", err)
			}

			updatedConsumerResponse := consumerapi.ConsumerResponse{}
			if err = consumerapi.DecodeJSON(resp2.Body, &updatedConsumerResponse); err != nil {
				t.Fatal(err)
			}

			if err := resp2.Body.Close(); err != nil {
				t.Error(err)
			}

			if updatedConsumerResponse.ID != respData.ID {
				t.Errorf("ID: Expected: %v, Got: %v", respData.ID, updatedConsumerResponse.Firstname)
			}

			if updatedConsumerResponse.Firstname != test.UpdatedConsumerData.Firstname {
				t.Errorf("Firstname: Expected: %s, Got: %s", test.UpdatedConsumerData.Firstname, updatedConsumerResponse.Firstname)
			}

			if updatedConsumerResponse.Lastname != test.UpdatedConsumerData.Lastname {
				t.Errorf("Lastname: Expected: %s, Got: %s", test.UpdatedConsumerData.Lastname, updatedConsumerResponse.Lastname)
			}

			if updatedConsumerResponse.Email != test.UpdatedConsumerData.Email {
				t.Errorf("Email: Expected: %s, Got: %s", test.UpdatedConsumerData.Email, updatedConsumerResponse.Email)
			}

			if updatedConsumerResponse.Phone != test.UpdatedConsumerData.Phone {
				t.Errorf("Phone: Expected: %s, Got: %s", test.UpdatedConsumerData.Phone, updatedConsumerResponse.Phone)
			}

			// Deleting consumer instance.
			deleter := http.DefaultClient

			delReq, err := http.NewRequest(http.MethodDelete,
				baseAddr+"/v1/consumers/"+strconv.Itoa(int(respData.ID)), nil)
			if err != nil {
				t.Error(err)
			}

			_, err = deleter.Do(delReq)
			if err != nil {
				t.Errorf("Could not delete created consumer: %v", err)
			}
		})
	}
}

func TestGetAllConsumerEndpoint(t *testing.T) {
	tests := []struct {
		name             string
		consumerDataList []consumerapi.NewConsumerRequest
	}{
		{
			"TestGetAllConsumerEndpoint test",
			[]consumerapi.NewConsumerRequest{
				{
					Firstname: "consumer1FName",
					Lastname:  "consumer1LName",
					Email:     "consumer1@gmail.com",
					Phone:     "111111111",
				},

				{
					Firstname: "consumer2FName",
					Lastname:  "consumer2LName",
					Email:     "consumer2@gmail.com",
					Phone:     "222222222",
				},
			},
		},
	}

	for _, currentTest := range tests {
		test := currentTest

		t.Run(test.name, func(t *testing.T) {
			var createdConsumerList []consumerapi.ConsumerResponse

			for _, consumer := range test.consumerDataList {

				reqBody, err := v1.Encode(consumer)
				if err != nil {
					t.Fatal(err)
				}

				resp1, err := http.Post(baseAddr+"/v1/consumers", "application/json", bytes.NewBuffer(reqBody))
				if err != nil {
					t.Fatal(err)
				}

				if resp1.StatusCode != http.StatusOK {
					t.Fatalf("Response status: %d", resp1.StatusCode)
				}

				createdConsumer := consumerapi.ConsumerResponse{}
				if err = consumerapi.DecodeJSON(resp1.Body, &createdConsumer); err != nil {
					t.Fatal(err)
				}

				if err := resp1.Body.Close(); err != nil {
					t.Error(err)
				}

				createdConsumerList = append(createdConsumerList, createdConsumer)
			}

			resp2, err := http.Get(baseAddr + "/v1/consumers")
			if err != nil {
				t.Fatal(err)
			}

			if resp2.StatusCode != http.StatusOK {
				t.Fatalf("Response status: %d", resp2.StatusCode)
			}

			respDataList2 := consumerapi.ReturnConsumerResponseList{}
			if err = consumerapi.DecodeJSON(resp2.Body, &respDataList2); err != nil {
				t.Fatal(err)
			}

			if len(respDataList2.ConsumerResponseList) != len(test.consumerDataList) {
				t.Errorf("len: Expected: %v, Got: %v", len(test.consumerDataList), len(respDataList2.ConsumerResponseList))
			}

			if err := resp2.Body.Close(); err != nil {
				t.Error(err)
			}

			// delete all consumer we have created
			for _, createdConsumer := range createdConsumerList {

				// Deleting consumer instance.
				deleter := http.DefaultClient

				delReq, err := http.NewRequest(http.MethodDelete,
					baseAddr+"/v1/consumers/"+strconv.Itoa(int(createdConsumer.ID)), nil)
				if err != nil {
					t.Error(err)
				}

				_, err = deleter.Do(delReq)
				if err != nil {
					t.Errorf("Could not delete created consumer: %v", err)
				}

			}
		})
	}
}

func TestGetConsumerEndpoint(t *testing.T) {
	tests := []struct {
		name         string
		consumerData consumerapi.NewConsumerRequest
	}{
		{
			"TestGetAllConsumerEndpoint test",
			consumerapi.NewConsumerRequest{
				Firstname: "consumer1FName",
				Lastname:  "consumer1LName",
				Email:     "consumer1@gmail.com",
				Phone:     "111111111",
			},
		},
	}

	for _, currentTest := range tests {
		test := currentTest

		t.Run(test.name, func(t *testing.T) {
			reqBody, err := v1.Encode(test.consumerData)
			if err != nil {
				t.Fatal(err)
			}

			respPostConsumer, err := http.Post(baseAddr+"/v1/consumers", "application/json", bytes.NewBuffer(reqBody))
			if err != nil {
				t.Fatal(err)
			}

			if respPostConsumer.StatusCode != http.StatusOK {
				t.Fatalf("Response status: %d", respPostConsumer.StatusCode)
			}

			createdConsumerResponse := consumerapi.ConsumerResponse{}
			if err = consumerapi.DecodeJSON(respPostConsumer.Body, &createdConsumerResponse); err != nil {
				t.Fatal(err)
			}

			if err := respPostConsumer.Body.Close(); err != nil {
				t.Error(err)
			}

			respGetConsumer, err := http.Get(baseAddr + "/v1/consumers/" + strconv.Itoa(int(createdConsumerResponse.ID)))
			if err != nil {
				t.Fatal(err)
			}

			if respGetConsumer.StatusCode != http.StatusOK {
				t.Fatalf("Response status: %d", respGetConsumer.StatusCode)
			}

			gotConsumerResponse := consumerapi.ConsumerResponse{}
			if err = consumerapi.DecodeJSON(respGetConsumer.Body, &gotConsumerResponse); err != nil {
				t.Fatal(err)
			}

			if err := respGetConsumer.Body.Close(); err != nil {
				t.Error(err)
			}

			if gotConsumerResponse.Firstname != test.consumerData.Firstname {
				t.Errorf("Firstname: Expected: %s, Got: %s", test.consumerData.Firstname, gotConsumerResponse.Firstname)
			}

			if gotConsumerResponse.Lastname != test.consumerData.Lastname {
				t.Errorf("Lastname: Expected: %s, Got: %s", test.consumerData.Lastname, gotConsumerResponse.Lastname)
			}

			if gotConsumerResponse.Email != test.consumerData.Email {
				t.Errorf("Email: Expected: %s, Got: %s", test.consumerData.Email, gotConsumerResponse.Email)
			}

			if gotConsumerResponse.Phone != test.consumerData.Phone {
				t.Errorf("Phone: Expected: %s, Got: %s", test.consumerData.Phone, gotConsumerResponse.Phone)
			}

			// Deleting consumer instance.
			deleter := http.DefaultClient

			delReq, err := http.NewRequest(http.MethodDelete,
				baseAddr+"/v1/consumers/"+strconv.Itoa(int(createdConsumerResponse.ID)), nil)
			if err != nil {
				t.Error(err)
			}

			_, err = deleter.Do(delReq)
			if err != nil {
				t.Errorf("Could not delete created consumer: %v", err)
			}
		})
	}
}

func TestInsertNewConsumerLocationEndpoint(t *testing.T) {
	tests := []struct {
		name                 string
		consumerData         consumerapi.NewConsumerRequest
		consumerLocationData consumerapi.NewLocationRequest
	}{
		{
			"InsertNewConsumerLocationEndpoint simple test",
			consumerapi.NewConsumerRequest{
				Firstname: "TestFName",
				Lastname:  "TestLName",
				Email:     "TestEmail",
				Phone:     "TestPhone",
			},
			consumerapi.NewLocationRequest{
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
			reqBody1, err := v1.Encode(test.consumerData)
			if err != nil {
				t.Fatal(err)
			}

			resp1, err := http.Post(baseAddr+"/v1/consumers", "application/json", bytes.NewBuffer(reqBody1))
			if err != nil {
				t.Fatal(err)
			}

			if resp1.StatusCode != http.StatusOK {
				t.Fatalf("Response status: %d", resp1.StatusCode)
			}

			consumerCreatedRespData := consumerapi.ConsumerResponse{}
			if err = consumerapi.DecodeJSON(resp1.Body, &consumerCreatedRespData); err != nil {
				t.Fatal(err)
			}

			reqBody2, err := v1.Encode(test.consumerLocationData)
			if err != nil {
				t.Fatal(err)
			}

			resp2, err := http.Post(baseAddr+"/v1/locations/"+strconv.Itoa(consumerCreatedRespData.ID),
				"application/json", bytes.NewBuffer(reqBody2))
			if err != nil {
				t.Fatal(err)
			}

			if resp2.StatusCode != http.StatusOK {
				t.Fatalf("Response status: %d", resp2.StatusCode)
			}

			consumerLocationRespData := consumerapi.LocationResponse{}
			if err = consumerapi.DecodeJSON(resp2.Body, &consumerLocationRespData); err != nil {
				t.Fatal(err)
			}

			if err := resp2.Body.Close(); err != nil {
				t.Error(err)
			}

			if consumerLocationRespData.UserID < 1 {
				t.Errorf("UserID: Expected : > 1, Got: %v", consumerLocationRespData.UserID)
			}

			if consumerLocationRespData.Altitude != test.consumerLocationData.Latitude {
				t.Errorf("Latitude: Expected: %s, Got: %s", test.consumerLocationData.Latitude, consumerLocationRespData.Altitude)
			}

			if consumerLocationRespData.Longitude != test.consumerLocationData.Longitude {
				t.Errorf("Longitude: Expected: %s, Got: %s", test.consumerLocationData.Longitude, consumerLocationRespData.Longitude)
			}

			if consumerLocationRespData.Country != test.consumerLocationData.Country {
				t.Errorf("Country: Expected: %s, Got: %s", test.consumerLocationData.Country, consumerLocationRespData.Country)
			}

			if consumerLocationRespData.Region != test.consumerLocationData.Region {
				t.Errorf("Region: Expected: %s, Got: %s", test.consumerLocationData.Region, consumerLocationRespData.Region)
			}

			if consumerLocationRespData.Street != test.consumerLocationData.Street {
				t.Errorf("Street: Expected: %s, Got: %s", test.consumerLocationData.Street, consumerLocationRespData.Street)
			}

			if consumerLocationRespData.HomeNumber != test.consumerLocationData.HomeNumber {
				t.Errorf("HomeNumber: Expected: %s, Got: %s", test.consumerLocationData.HomeNumber, consumerLocationRespData.HomeNumber)
			}

			if consumerLocationRespData.Floor != test.consumerLocationData.Floor {
				t.Errorf("Floor: Expected: %s, Got: %s", test.consumerLocationData.Floor, consumerLocationRespData.Floor)
			}

			if consumerLocationRespData.Door != test.consumerLocationData.Door {
				t.Errorf("Door: Expected: %s, Got: %s", test.consumerLocationData.Door, consumerLocationRespData.Door)
			}

			// Deleting consumer instance.
			deleter := http.DefaultClient

			delReq, err := http.NewRequest(http.MethodDelete,
				baseAddr+"/v1/consumers/"+strconv.Itoa(consumerCreatedRespData.ID), nil)
			if err != nil {
				t.Error(err)
			}

			_, err = deleter.Do(delReq)
			if err != nil {
				t.Errorf("Could not delete created consumer: %v", err)
			}
		})
	}
}

func TestUpdateConsumerLocationEndpoint(t *testing.T) {
	tests := []struct {
		name                        string
		consumerData                consumerapi.NewConsumerRequest
		consumerLocationInitialData consumerapi.NewLocationRequest
		consumerLocationUpdatedData consumerapi.UpdateLocationRequest
	}{
		{
			"UpdateConsumerLocationEndpoint simple test",
			consumerapi.NewConsumerRequest{
				Firstname: "TestFName",
				Lastname:  "TestLName",
				Email:     "TestEmail",
				Phone:     "TestPhone",
			},
			consumerapi.NewLocationRequest{
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
			consumerapi.UpdateLocationRequest{
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
			reqBody1, err := v1.Encode(test.consumerData)
			if err != nil {
				t.Fatal(err)
			}

			resp1, err := http.Post(baseAddr+"/v1/consumers", "application/json", bytes.NewBuffer(reqBody1))
			if err != nil {
				t.Fatal(err)
			}

			if resp1.StatusCode != http.StatusOK {
				t.Fatalf("Response status: %d", resp1.StatusCode)
			}

			consumerCreatedRespData := consumerapi.ConsumerResponse{}
			if err = consumerapi.DecodeJSON(resp1.Body, &consumerCreatedRespData); err != nil {
				t.Fatal(err)
			}

			reqBody2, err := v1.Encode(test.consumerLocationInitialData)
			if err != nil {
				t.Fatal(err)
			}

			resp2, err := http.Post(baseAddr+"/v1/locations/"+strconv.Itoa(consumerCreatedRespData.ID),
				"application/json", bytes.NewBuffer(reqBody2))
			if err != nil {
				t.Fatal(err)
			}

			if resp2.StatusCode != http.StatusOK {
				t.Fatalf("Response status: %d", resp2.StatusCode)
			}

			if err := resp2.Body.Close(); err != nil {
				t.Error(err)
			}

			reqBody3, err := v1.Encode(test.consumerLocationUpdatedData)
			if err != nil {
				t.Fatal(err)
			}

			client := http.Client{}

			req, err := http.NewRequest(http.MethodPut, baseAddr+"/v1/locations/"+strconv.Itoa(consumerCreatedRespData.ID), bytes.NewBuffer(reqBody3))
			if err != nil {
				t.Fatal(err)
			}
			resp3, err := client.Do(req)
			if err != nil {
				t.Fatal(err)
			}

			if resp3.StatusCode != http.StatusOK {
				t.Fatalf("Response status: %d", resp3.StatusCode)
			}

			consumerLocationUpdatedRespData := consumerapi.LocationResponse{}
			if err = consumerapi.DecodeJSON(resp3.Body, &consumerLocationUpdatedRespData); err != nil {
				t.Fatal(err)
			}

			if err := resp3.Body.Close(); err != nil {
				t.Error(err)
			}

			if consumerLocationUpdatedRespData.UserID != consumerCreatedRespData.ID {
				t.Errorf("UserID: Expected :  %v , Got: %v", consumerCreatedRespData.ID, consumerLocationUpdatedRespData.UserID)
			}

			if consumerLocationUpdatedRespData.Altitude != test.consumerLocationUpdatedData.Latitude {
				t.Errorf("Latitude: Expected: %s, Got: %s", test.consumerLocationUpdatedData.Latitude, consumerLocationUpdatedRespData.Altitude)
			}

			if consumerLocationUpdatedRespData.Longitude != test.consumerLocationUpdatedData.Longitude {
				t.Errorf("Longitude: Expected: %s, Got: %s", test.consumerLocationUpdatedData.Longitude, consumerLocationUpdatedRespData.Longitude)
			}

			if consumerLocationUpdatedRespData.Country != test.consumerLocationUpdatedData.Country {
				t.Errorf("Country: Expected: %s, Got: %s", test.consumerLocationUpdatedData.Country, consumerLocationUpdatedRespData.Country)
			}

			if consumerLocationUpdatedRespData.Region != test.consumerLocationUpdatedData.Region {
				t.Errorf("Region: Expected: %s, Got: %s", test.consumerLocationUpdatedData.Region, consumerLocationUpdatedRespData.Region)
			}

			if consumerLocationUpdatedRespData.Street != test.consumerLocationUpdatedData.Street {
				t.Errorf("Street: Expected: %s, Got: %s", test.consumerLocationUpdatedData.Street, consumerLocationUpdatedRespData.Street)
			}

			if consumerLocationUpdatedRespData.HomeNumber != test.consumerLocationUpdatedData.HomeNumber {
				t.Errorf("HomeNumber: Expected: %s, Got: %s", test.consumerLocationUpdatedData.HomeNumber, consumerLocationUpdatedRespData.HomeNumber)
			}

			if consumerLocationUpdatedRespData.Floor != test.consumerLocationUpdatedData.Floor {
				t.Errorf("Floor: Expected: %s, Got: %s", test.consumerLocationUpdatedData.Floor, consumerLocationUpdatedRespData.Floor)
			}

			if consumerLocationUpdatedRespData.Door != test.consumerLocationUpdatedData.Door {
				t.Errorf("Door: Expected: %s, Got: %s", test.consumerLocationUpdatedData.Door, consumerLocationUpdatedRespData.Door)
			}

			// Deleting consumer instance.
			deleter := http.DefaultClient

			delReq, err := http.NewRequest(http.MethodDelete,
				baseAddr+"/v1/consumers/"+strconv.Itoa(consumerCreatedRespData.ID), nil)
			if err != nil {
				t.Error(err)
			}

			_, err = deleter.Do(delReq)
			if err != nil {
				t.Errorf("Could not delete created consumer: %v", err)
			}
		})
	}
}

func TestGetConsumerLocationEndpoint(t *testing.T) {
	tests := []struct {
		name                 string
		consumerData         consumerapi.NewConsumerRequest
		consumerLocationData consumerapi.NewLocationRequest
	}{
		{
			"GetConsumerLocationEndpoint simple test",
			consumerapi.NewConsumerRequest{
				Firstname: "TestFName",
				Lastname:  "TestLName",
				Email:     "TestEmail",
				Phone:     "TestPhone",
			},
			consumerapi.NewLocationRequest{
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
			reqBody1, err := v1.Encode(test.consumerData)
			if err != nil {
				t.Fatal(err)
			}

			resp1, err := http.Post(baseAddr+"/v1/consumers", "application/json", bytes.NewBuffer(reqBody1))
			if err != nil {
				t.Fatal(err)
			}

			if resp1.StatusCode != http.StatusOK {
				t.Fatalf("Response status: %d", resp1.StatusCode)
			}

			consumerCreatedRespData := consumerapi.ConsumerResponse{}
			if err = consumerapi.DecodeJSON(resp1.Body, &consumerCreatedRespData); err != nil {
				t.Fatal(err)
			}

			reqBody2, err := v1.Encode(test.consumerLocationData)
			if err != nil {
				t.Fatal(err)
			}

			resp2, err := http.Post(baseAddr+"/v1/locations/"+strconv.Itoa(consumerCreatedRespData.ID),
				"application/json", bytes.NewBuffer(reqBody2))
			if err != nil {
				t.Fatal(err)
			}

			if resp2.StatusCode != http.StatusOK {
				t.Fatalf("Response status: %d", resp2.StatusCode)
			}

			resp3, err := http.Get(baseAddr + "/v1/locations/" + strconv.Itoa(consumerCreatedRespData.ID))
			if err != nil {
				t.Fatal(err)
			}

			if resp3.StatusCode != http.StatusOK {
				t.Fatalf("Response status: %d", resp2.StatusCode)
			}

			consumerLocationRespData := consumerapi.LocationResponse{}
			if err = consumerapi.DecodeJSON(resp3.Body, &consumerLocationRespData); err != nil {
				t.Fatal(err)
			}

			if err := resp3.Body.Close(); err != nil {
				t.Error(err)
			}

			if consumerLocationRespData.UserID != consumerCreatedRespData.ID {
				t.Errorf("UserID: Expected : %v, Got: %v", consumerCreatedRespData.ID, consumerLocationRespData.UserID)
			}

			if consumerLocationRespData.Altitude != test.consumerLocationData.Latitude {
				t.Errorf("Latitude: Expected: %s, Got: %s", test.consumerLocationData.Latitude, consumerLocationRespData.Altitude)
			}

			if consumerLocationRespData.Longitude != test.consumerLocationData.Longitude {
				t.Errorf("Longitude: Expected: %s, Got: %s", test.consumerLocationData.Longitude, consumerLocationRespData.Longitude)
			}

			if consumerLocationRespData.Country != test.consumerLocationData.Country {
				t.Errorf("Country: Expected: %s, Got: %s", test.consumerLocationData.Country, consumerLocationRespData.Country)
			}

			if consumerLocationRespData.Region != test.consumerLocationData.Region {
				t.Errorf("Region: Expected: %s, Got: %s", test.consumerLocationData.Region, consumerLocationRespData.Region)
			}

			if consumerLocationRespData.Street != test.consumerLocationData.Street {
				t.Errorf("Street: Expected: %s, Got: %s", test.consumerLocationData.Street, consumerLocationRespData.Street)
			}

			if consumerLocationRespData.HomeNumber != test.consumerLocationData.HomeNumber {
				t.Errorf("HomeNumber: Expected: %s, Got: %s", test.consumerLocationData.HomeNumber, consumerLocationRespData.HomeNumber)
			}

			if consumerLocationRespData.Floor != test.consumerLocationData.Floor {
				t.Errorf("Floor: Expected: %s, Got: %s", test.consumerLocationData.Floor, consumerLocationRespData.Floor)
			}

			if consumerLocationRespData.Door != test.consumerLocationData.Door {
				t.Errorf("Door: Expected: %s, Got: %s", test.consumerLocationData.Door, consumerLocationRespData.Door)
			}

			// Deleting consumer instance.
			deleter := http.DefaultClient

			delReq, err := http.NewRequest(http.MethodDelete,
				baseAddr+"/v1/consumers/"+strconv.Itoa(int(consumerCreatedRespData.ID)), nil)
			if err != nil {
				t.Error(err)
			}

			_, err = deleter.Do(delReq)
			if err != nil {
				t.Errorf("Could not delete created consumer: %v", err)
			}
		})
	}
}
