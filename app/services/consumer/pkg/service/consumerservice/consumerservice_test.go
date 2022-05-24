package consumerservice_test

import (
	"bytes"
	"consumer/api/v1/consumerapi"
	v1 "github.com/nndergunov/deliveryApp/app/pkg/api/v1"
	"net/http"
	"strconv"
	"testing"
)

const baseAddr = "http://localhost:8080"

func TestInsertNewConsumerEndpoint(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name         string
		consumerData consumerapi.NewConsumerRequest
	}{
		{
			"Insert consumer simple test",
			consumerapi.NewConsumerRequest{
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
			t.Parallel()

			reqBody, err := v1.Encode(test.consumerData)
			if err != nil {
				t.Fatal(err)
			}

			resp, err := http.Post(baseAddr+"/v1/consumer/",
				"application/json", bytes.NewBuffer(reqBody))
			if err != nil {
				t.Fatal(err)
			}

			if resp.StatusCode != http.StatusOK {
				t.Fatalf("Response status: %d", resp.StatusCode)
			}

			defer func() {
				err := resp.Body.Close()
				if err != nil {
					t.Error(err)
				}
			}()

			if resp.StatusCode != http.StatusOK {
				t.Fatalf("StatusCode: %d", resp.StatusCode)
			}

			respData := consumerapi.ConsumerResponse{}
			if err = consumerapi.DecodeJSON(resp.Body, &respData); err != nil {
				t.Fatal(err)
			}

			if respData.ID < 1 {
				t.Errorf("ID: Expected: > 1, Got: %v", respData.ID)
			}

			defer func() {
				// Deleting restaurant instance.
				deleter := http.DefaultClient

				delReq, err := http.NewRequest(http.MethodDelete,
					baseAddr+"/v1/consumer/"+strconv.Itoa(int(respData.ID)), nil)
				if err != nil {
					t.Error(err)
				}

				_, err = deleter.Do(delReq)
				if err != nil {
					t.Errorf("Could not delete created consumer: %v", err)
				}
			}()

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

		})
	}
}

//
//func TestDeleteConsumerEndpoint(t *testing.T) {
//	t.Parallel()
//
//	t.Run("Delete Consumer simple test", func(t *testing.T) {
//		t.Parallel()
//
//		mockService := new(MockService)
//		log := logger.NewLogger(os.Stdout, "Delete Consumer simple test")
//		handler := consumerhandler.NewConsumerHandler(consumerhandler.Params{
//			Logger:          log,
//			ConsumerService: mockService,
//		})
//
//		resp := httptest.NewRecorder()
//		req := httptest.NewRequest(http.MethodDelete, "/v1/consumer/1", nil)
//
//		handler.ServeHTTP(resp, req)
//		var respData string
//		expData := "consumer deleted"
//
//		if err := consumerapi.DecodeJSON(resp.Body, &respData); err != nil {
//			t.Fatal(err)
//		}
//
//		if resp.Code != http.StatusOK {
//			t.Fatalf("StatusCode: %d", resp.Code)
//		}
//		if respData != expData {
//			t.Errorf("response: Expected: %s, Got: %s", expData, respData)
//		}
//	})
//}
//
//func TestUpdateConsumerEndpoint(t *testing.T) {
//	t.Parallel()
//
//	tests := []struct {
//		name         string
//		consumerData consumerapi.UpdateConsumerRequest
//	}{
//		{
//			"Update consumer simple test",
//			consumerapi.UpdateConsumerRequest{
//				Firstname: "vasya",
//				Lastname:  "",
//				Email:     "vasya@gmail.com",
//				Phone:     "123456789",
//			},
//		},
//	}
//
//	for _, currentTest := range tests {
//		test := currentTest
//
//		t.Run(test.name, func(t *testing.T) {
//			t.Parallel()
//
//			mockService := new(MockService)
//
//			log := logger.NewLogger(os.Stdout, test.name)
//			courierHandler := consumerhandler.NewConsumerHandler(consumerhandler.Params{
//				Logger:          log,
//				ConsumerService: mockService,
//			})
//
//			reqBody, err := v1.Encode(test.consumerData)
//			if err != nil {
//				t.Fatal(err)
//			}
//
//			resp := httptest.NewRecorder()
//			req := httptest.NewRequest(http.MethodPut, "/v1/consumer/1", bytes.NewBuffer(reqBody))
//
//			courierHandler.ServeHTTP(resp, req)
//
//			if resp.Code != http.StatusOK {
//				t.Fatalf("StatusCode: %d", resp.Code)
//			}
//
//			respData := consumerapi.ConsumerResponse{}
//			if err = consumerapi.DecodeJSON(resp.Body, &respData); err != nil {
//				t.Fatal(err)
//			}
//
//			if respData.ID != MockConsumerData.ID {
//				t.Errorf("ConsumerID: Expected: %v, Got: %v", MockConsumerData.ID, respData.ID)
//			}
//
//			if respData.Firstname != MockConsumerData.Firstname {
//				t.Errorf("Firstname: Expected: %s, Got: %s", test.consumerData.Firstname, respData.Firstname)
//			}
//
//			if respData.Lastname != MockConsumerData.Lastname {
//				t.Errorf("Lastname: Expected: %s, Got: %s", test.consumerData.Lastname, respData.Lastname)
//			}
//
//			if respData.Email != MockConsumerData.Email {
//				t.Errorf("Email: Expected: %s, Got: %s", test.consumerData.Email, respData.Email)
//			}
//
//			if respData.Phone != MockConsumerData.Phone {
//				t.Errorf("Phone: Expected: %s, Got: %s", test.consumerData.Phone, respData.Phone)
//			}
//		})
//	}
//}
//
//func TestGetAllConsumerEndpoint(t *testing.T) {
//	t.Parallel()
//
//	t.Run("get all consumer simple test", func(t *testing.T) {
//		t.Parallel()
//		testGetRespList := []*domain.Consumer{MockConsumerData}
//		mockService := new(MockService)
//		log := logger.NewLogger(os.Stdout, "get all consumer simple test")
//		handler := consumerhandler.NewConsumerHandler(consumerhandler.Params{
//			Logger:          log,
//			ConsumerService: mockService,
//		})
//
//		resp := httptest.NewRecorder()
//		req := httptest.NewRequest(http.MethodGet, "/v1/consumer/all/", nil)
//
//		handler.ServeHTTP(resp, req)
//
//		var respDataList consumerapi.ReturnConsumerResponseList
//		if err := consumerapi.DecodeJSON(resp.Body, &respDataList); err != nil {
//			t.Fatal(err)
//		}
//
//		if resp.Code != http.StatusOK {
//			t.Fatalf("StatusCode: %d", resp.Code)
//		}
//		if len(respDataList.ConsumerResponseList) != len(testGetRespList) {
//			t.Errorf("len: Expected: %v, Got: %v", len(testGetRespList), len(respDataList.ConsumerResponseList))
//		}
//		for _, respData := range respDataList.ConsumerResponseList {
//
//			if respData.ID != MockConsumerData.ID {
//				t.Errorf("ConsumerID: Expected: %v, Got: %v", MockConsumerData.ID, respData.ID)
//			}
//
//			if respData.Firstname != MockConsumerData.Firstname {
//				t.Errorf("Firstname: Expected: %s, Got: %s", MockConsumerData.Firstname, respData.Firstname)
//			}
//
//			if respData.Lastname != MockConsumerData.Lastname {
//				t.Errorf("Lastname: Expected: %s, Got: %s", MockConsumerData.Lastname, respData.Lastname)
//			}
//
//			if respData.Email != MockConsumerData.Email {
//				t.Errorf("Email: Expected: %s, Got: %s", MockConsumerData.Email, respData.Email)
//			}
//
//			if respData.Phone != MockConsumerData.Phone {
//				t.Errorf("Phone: Expected: %s, Got: %s", MockConsumerData.Phone, respData.Phone)
//			}
//		}
//	})
//}
//
//func TestGetConsumerEndpoint(t *testing.T) {
//	t.Parallel()
//
//	t.Run("get consumer simple test", func(t *testing.T) {
//		t.Parallel()
//
//		mockService := new(MockService)
//		log := logger.NewLogger(os.Stdout, "get consumer simple test")
//		handler := consumerhandler.NewConsumerHandler(consumerhandler.Params{
//			Logger:          log,
//			ConsumerService: mockService,
//		})
//
//		resp := httptest.NewRecorder()
//		req := httptest.NewRequest(http.MethodGet, "/v1/consumer/1", nil)
//
//		handler.ServeHTTP(resp, req)
//
//		respData := consumerapi.ConsumerResponse{}
//		if err := consumerapi.DecodeJSON(resp.Body, &respData); err != nil {
//			t.Fatal(err)
//		}
//
//		if resp.Code != http.StatusOK {
//			t.Fatalf("StatusCode: %d", resp.Code)
//		}
//		if respData.ID != MockConsumerData.ID {
//			t.Errorf("ConsumerID: Expected: %v, Got: %v", MockConsumerData.ID, respData.ID)
//		}
//
//		if respData.Firstname != MockConsumerData.Firstname {
//			t.Errorf("Firstname: Expected: %s, Got: %s", MockConsumerData.Firstname, respData.Firstname)
//		}
//
//		if respData.Lastname != MockConsumerData.Lastname {
//			t.Errorf("Lastname: Expected: %s, Got: %s", MockConsumerData.Lastname, respData.Lastname)
//		}
//
//		if respData.Email != MockConsumerData.Email {
//			t.Errorf("Email: Expected: %s, Got: %s", MockConsumerData.Email, respData.Email)
//		}
//
//		if respData.Phone != MockConsumerData.Phone {
//			t.Errorf("Phone: Expected: %s, Got: %s", MockConsumerData.Phone, respData.Phone)
//		}
//
//	})
//}
//
//func TestInsertNewConsumerLocationEndpoint(t *testing.T) {
//	t.Parallel()
//
//	tests := []struct {
//		name                 string
//		consumerLocationData consumerapi.NewConsumerLocationRequest
//	}{
//		{
//			"NewConsumerLocation simple test",
//			consumerapi.NewConsumerLocationRequest{
//				Altitude:   "0123456789",
//				Longitude:  "0123456789",
//				Country:    "TestCountry",
//				City:       "Test City",
//				Region:     "",
//				Street:     "",
//				HomeNumber: "",
//				Floor:      "",
//				Door:       "",
//			},
//		},
//	}
//
//	for _, currentTest := range tests {
//		test := currentTest
//
//		t.Run(test.name, func(t *testing.T) {
//			t.Parallel()
//
//			mockService := new(MockService)
//
//			log := logger.NewLogger(os.Stdout, test.name)
//			courierHandler := consumerhandler.NewConsumerHandler(consumerhandler.Params{
//				Logger:          log,
//				ConsumerService: mockService,
//			})
//
//			reqBody, err := v1.Encode(test.consumerLocationData)
//			if err != nil {
//				t.Fatal(err)
//			}
//
//			resp := httptest.NewRecorder()
//			req := httptest.NewRequest(http.MethodPost, "/v1/consumer/location/1", bytes.NewBuffer(reqBody))
//
//			courierHandler.ServeHTTP(resp, req)
//
//			if resp.Code != http.StatusOK {
//				t.Fatalf("StatusCode: %d", resp.Code)
//			}
//
//			respData := consumerapi.ConsumerLocationResponse{}
//			if err = consumerapi.DecodeJSON(resp.Body, &respData); err != nil {
//				t.Fatal(err)
//			}
//
//			if respData.ConsumerID != MockConsumerLocationData.ConsumerID {
//				t.Errorf("ConsumerID: Expected: %v, Got: %v", MockConsumerLocationData.ConsumerID, respData.ConsumerID)
//			}
//
//			if respData.Altitude != MockConsumerLocationData.Altitude {
//				t.Errorf("Altitude: Expected: %s, Got: %s", test.consumerLocationData.Altitude, respData.Altitude)
//			}
//
//			if respData.Longitude != MockConsumerLocationData.Longitude {
//				t.Errorf("Longitude: Expected: %s, Got: %s", test.consumerLocationData.Longitude, respData.Longitude)
//			}
//
//			if respData.Country != MockConsumerLocationData.Country {
//				t.Errorf("Country: Expected: %s, Got: %s", test.consumerLocationData.Country, respData.Country)
//			}
//
//			if respData.City != MockConsumerLocationData.City {
//				t.Errorf("City: Expected: %s, Got: %s", test.consumerLocationData.City, respData.City)
//			}
//
//			if respData.Region != MockConsumerLocationData.Region {
//				t.Errorf("Region: Expected: %s, Got: %s", test.consumerLocationData.Region, respData.Region)
//			}
//		})
//	}
//}
//
//func TestUpdateConsumerLocationEndpoint(t *testing.T) {
//	t.Parallel()
//
//	tests := []struct {
//		name                 string
//		consumerLocationData consumerapi.NewConsumerLocationRequest
//	}{
//		{
//			"UpdateConsumerLocation simple test",
//			consumerapi.NewConsumerLocationRequest{
//				Altitude:   "0123456789",
//				Longitude:  "0123456789",
//				Country:    "TestCountry",
//				City:       "Test City",
//				Region:     "",
//				Street:     "",
//				HomeNumber: "",
//				Floor:      "",
//				Door:       "",
//			},
//		},
//	}
//
//	for _, currentTest := range tests {
//		test := currentTest
//
//		t.Run(test.name, func(t *testing.T) {
//			t.Parallel()
//
//			mockService := new(MockService)
//
//			log := logger.NewLogger(os.Stdout, test.name)
//			courierHandler := consumerhandler.NewConsumerHandler(consumerhandler.Params{
//				Logger:          log,
//				ConsumerService: mockService,
//			})
//
//			reqBody, err := v1.Encode(test.consumerLocationData)
//			if err != nil {
//				t.Fatal(err)
//			}
//
//			resp := httptest.NewRecorder()
//			req := httptest.NewRequest(http.MethodPut, "/v1/consumer/location/1", bytes.NewBuffer(reqBody))
//
//			courierHandler.ServeHTTP(resp, req)
//
//			if resp.Code != http.StatusOK {
//				t.Fatalf("StatusCode: %d", resp.Code)
//			}
//
//			respData := consumerapi.ConsumerLocationResponse{}
//			if err = consumerapi.DecodeJSON(resp.Body, &respData); err != nil {
//				t.Fatal(err)
//			}
//
//			if respData.ConsumerID != MockConsumerLocationData.ConsumerID {
//				t.Errorf("ConsumerID: Expected: %v, Got: %v", MockConsumerLocationData.ConsumerID, respData.ConsumerID)
//			}
//
//			if respData.Altitude != MockConsumerLocationData.Altitude {
//				t.Errorf("Altitude: Expected: %s, Got: %s", test.consumerLocationData.Altitude, respData.Altitude)
//			}
//
//			if respData.Longitude != MockConsumerLocationData.Longitude {
//				t.Errorf("Longitude: Expected: %s, Got: %s", test.consumerLocationData.Longitude, respData.Longitude)
//			}
//
//			if respData.Country != MockConsumerLocationData.Country {
//				t.Errorf("Country: Expected: %s, Got: %s", test.consumerLocationData.Country, respData.Country)
//			}
//
//			if respData.City != MockConsumerLocationData.City {
//				t.Errorf("City: Expected: %s, Got: %s", test.consumerLocationData.City, respData.City)
//			}
//
//			if respData.Region != MockConsumerLocationData.Region {
//				t.Errorf("Region: Expected: %s, Got: %s", test.consumerLocationData.Region, respData.Region)
//			}
//		})
//	}
//}
//
//func TestGetConsumerLocationEndpoint(t *testing.T) {
//	t.Parallel()
//
//	tests := []struct {
//		name string
//	}{
//		{
//			"GetConsumerLocation simple test",
//		},
//	}
//
//	for _, currentTest := range tests {
//		test := currentTest
//
//		t.Run(test.name, func(t *testing.T) {
//			t.Parallel()
//
//			mockService := new(MockService)
//
//			log := logger.NewLogger(os.Stdout, test.name)
//			courierHandler := consumerhandler.NewConsumerHandler(consumerhandler.Params{
//				Logger:          log,
//				ConsumerService: mockService,
//			})
//
//			resp := httptest.NewRecorder()
//			req := httptest.NewRequest(http.MethodGet, "/v1/consumer/location/1", nil)
//
//			courierHandler.ServeHTTP(resp, req)
//
//			if resp.Code != http.StatusOK {
//				t.Fatalf("StatusCode: %d", resp.Code)
//			}
//
//			respData := consumerapi.ConsumerLocationResponse{}
//			if err := consumerapi.DecodeJSON(resp.Body, &respData); err != nil {
//				t.Fatal(err)
//			}
//
//			if respData.ConsumerID != MockConsumerLocationData.ConsumerID {
//				t.Errorf("ConsumerID: Expected: %v, Got: %v", MockConsumerLocationData.ConsumerID, respData.ConsumerID)
//			}
//
//			if respData.Altitude != MockConsumerLocationData.Altitude {
//				t.Errorf("Altitude: Expected: %s, Got: %s", MockConsumerLocationData.Altitude, respData.Altitude)
//			}
//
//			if respData.Longitude != MockConsumerLocationData.Longitude {
//				t.Errorf("Longitude: Expected: %s, Got: %s", MockConsumerLocationData.Longitude, respData.Longitude)
//			}
//
//			if respData.Country != MockConsumerLocationData.Country {
//				t.Errorf("Country: Expected: %s, Got: %s", MockConsumerLocationData.Country, respData.Country)
//			}
//
//			if respData.City != MockConsumerLocationData.City {
//				t.Errorf("City: Expected: %s, Got: %s", MockConsumerLocationData.City, respData.City)
//			}
//
//			if respData.Region != MockConsumerLocationData.Region {
//				t.Errorf("Region: Expected: %s, Got: %s", MockConsumerLocationData.Region, respData.Region)
//			}
//		})
//	}
//}
