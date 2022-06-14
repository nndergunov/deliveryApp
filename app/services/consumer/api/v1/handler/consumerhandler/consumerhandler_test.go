package consumerhandler_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/nndergunov/deliveryApp/app/pkg/api/v1"
	"github.com/nndergunov/deliveryApp/app/pkg/logger"

	"consumer/api/v1/consumerapi"
	"consumer/api/v1/handler/consumerhandler"
	"consumer/pkg/domain"
)

var (
	MockConsumerData = &domain.Consumer{
		ID:        1,
		Firstname: "vasya",
		Lastname:  "secret",
		Email:     "vasya@gmail.com",
		Phone:     "123456789",
	}

	MockLocationData = &domain.Location{
		UserID:     1,
		Latitude:   "0123456789",
		Longitude:  "0123456789",
		Country:    "TestCountry",
		City:       "Test City",
		Region:     "",
		Street:     "",
		HomeNumber: "",
		Floor:      "",
		Door:       "",
	}
)

type MockService struct{}

func (m MockService) InsertConsumer(_ domain.Consumer) (*domain.Consumer, error) {
	return MockConsumerData, nil
}

func (m MockService) DeleteConsumer(_ string) (data string, err error) {
	return "consumer deleted", nil
}

func (m MockService) UpdateConsumer(_ domain.Consumer, _ string) (*domain.Consumer, error) {
	return MockConsumerData, nil
}

func (m MockService) GetAllConsumer() ([]domain.Consumer, error) {
	return []domain.Consumer{*MockConsumerData}, nil
}

func (m MockService) GetConsumer(_ string) (*domain.Consumer, error) {
	return MockConsumerData, nil
}

func (m MockService) InsertLocation(_ domain.Location, id string) (*domain.Location, error) {
	return MockLocationData, nil
}

func (m MockService) GetLocation(_ string) (*domain.Location, error) {
	return MockLocationData, nil
}

func (m MockService) UpdateLocation(_ domain.Location, id string) (*domain.Location, error) {
	return MockLocationData, nil
}

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

			mockService := new(MockService)

			log := logger.NewLogger(os.Stdout, test.name)
			consumerHandler := consumerhandler.NewConsumerHandler(consumerhandler.Params{
				Logger:          log,
				ConsumerService: mockService,
			})

			reqBody, err := v1.Encode(test.consumerData)
			if err != nil {
				t.Fatal(err)
			}

			resp := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPost, "/v1/consumers", bytes.NewBuffer(reqBody))

			consumerHandler.ServeHTTP(resp, req)

			if resp.Code != http.StatusOK {
				t.Fatalf("StatusCode: %d", resp.Code)
			}

			respData := consumerapi.ConsumerResponse{}
			if err = consumerapi.DecodeJSON(resp.Body, &respData); err != nil {
				t.Fatal(err)
			}

			if respData.ID != MockConsumerData.ID {
				t.Errorf("UserID: Expected: %v, Got: %v", MockConsumerData.ID, respData.ID)
			}

			if respData.Firstname != MockConsumerData.Firstname {
				t.Errorf("Firstname: Expected: %s, Got: %s", test.consumerData.Firstname, respData.Firstname)
			}

			if respData.Lastname != MockConsumerData.Lastname {
				t.Errorf("Lastname: Expected: %s, Got: %s", test.consumerData.Lastname, respData.Lastname)
			}

			if respData.Email != MockConsumerData.Email {
				t.Errorf("Email: Expected: %s, Got: %s", test.consumerData.Email, respData.Email)
			}

			if respData.Phone != MockConsumerData.Phone {
				t.Errorf("Phone: Expected: %s, Got: %s", test.consumerData.Phone, respData.Phone)
			}
		})
	}
}

func TestDeleteConsumerEndpoint(t *testing.T) {
	t.Parallel()

	t.Run("Delete Consumer simple test", func(t *testing.T) {
		t.Parallel()

		mockService := new(MockService)
		log := logger.NewLogger(os.Stdout, "Delete Consumer simple test")
		handler := consumerhandler.NewConsumerHandler(consumerhandler.Params{
			Logger:          log,
			ConsumerService: mockService,
		})

		resp := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodDelete, "/v1/consumers/1", nil)

		handler.ServeHTTP(resp, req)
		var respData string
		expData := "consumer deleted"

		if err := consumerapi.DecodeJSON(resp.Body, &respData); err != nil {
			t.Fatal(err)
		}

		if resp.Code != http.StatusOK {
			t.Fatalf("StatusCode: %d", resp.Code)
		}
		if respData != expData {
			t.Errorf("response: Expected: %s, Got: %s", expData, respData)
		}
	})
}

func TestUpdateConsumerEndpoint(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name         string
		consumerData consumerapi.UpdateConsumerRequest
	}{
		{
			"Update consumer simple test",
			consumerapi.UpdateConsumerRequest{
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

			mockService := new(MockService)

			log := logger.NewLogger(os.Stdout, test.name)
			consumerHandler := consumerhandler.NewConsumerHandler(consumerhandler.Params{
				Logger:          log,
				ConsumerService: mockService,
			})

			reqBody, err := v1.Encode(test.consumerData)
			if err != nil {
				t.Fatal(err)
			}

			resp := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPut, "/v1/consumers/1", bytes.NewBuffer(reqBody))

			consumerHandler.ServeHTTP(resp, req)

			if resp.Code != http.StatusOK {
				t.Fatalf("StatusCode: %d", resp.Code)
			}

			respData := consumerapi.ConsumerResponse{}
			if err = consumerapi.DecodeJSON(resp.Body, &respData); err != nil {
				t.Fatal(err)
			}

			if respData.ID != MockConsumerData.ID {
				t.Errorf("UserID: Expected: %v, Got: %v", MockConsumerData.ID, respData.ID)
			}

			if respData.Firstname != MockConsumerData.Firstname {
				t.Errorf("Firstname: Expected: %s, Got: %s", test.consumerData.Firstname, respData.Firstname)
			}

			if respData.Lastname != MockConsumerData.Lastname {
				t.Errorf("Lastname: Expected: %s, Got: %s", test.consumerData.Lastname, respData.Lastname)
			}

			if respData.Email != MockConsumerData.Email {
				t.Errorf("Email: Expected: %s, Got: %s", test.consumerData.Email, respData.Email)
			}

			if respData.Phone != MockConsumerData.Phone {
				t.Errorf("Phone: Expected: %s, Got: %s", test.consumerData.Phone, respData.Phone)
			}
		})
	}
}

func TestGetAllConsumerEndpoint(t *testing.T) {
	t.Parallel()

	t.Run("get all consumer simple test", func(t *testing.T) {
		t.Parallel()
		testGetRespList := []*domain.Consumer{MockConsumerData}
		mockService := new(MockService)
		log := logger.NewLogger(os.Stdout, "get all consumer simple test")
		handler := consumerhandler.NewConsumerHandler(consumerhandler.Params{
			Logger:          log,
			ConsumerService: mockService,
		})

		resp := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/v1/consumers", nil)

		handler.ServeHTTP(resp, req)

		var respDataList consumerapi.ReturnConsumerResponseList
		if err := consumerapi.DecodeJSON(resp.Body, &respDataList); err != nil {
			t.Fatal(err)
		}

		if resp.Code != http.StatusOK {
			t.Fatalf("StatusCode: %d", resp.Code)
		}
		if len(respDataList.ConsumerResponseList) != len(testGetRespList) {
			t.Errorf("len: Expected: %v, Got: %v", len(testGetRespList), len(respDataList.ConsumerResponseList))
		}
		for _, respData := range respDataList.ConsumerResponseList {

			if respData.ID != MockConsumerData.ID {
				t.Errorf("UserID: Expected: %v, Got: %v", MockConsumerData.ID, respData.ID)
			}

			if respData.Firstname != MockConsumerData.Firstname {
				t.Errorf("Firstname: Expected: %s, Got: %s", MockConsumerData.Firstname, respData.Firstname)
			}

			if respData.Lastname != MockConsumerData.Lastname {
				t.Errorf("Lastname: Expected: %s, Got: %s", MockConsumerData.Lastname, respData.Lastname)
			}

			if respData.Email != MockConsumerData.Email {
				t.Errorf("Email: Expected: %s, Got: %s", MockConsumerData.Email, respData.Email)
			}

			if respData.Phone != MockConsumerData.Phone {
				t.Errorf("Phone: Expected: %s, Got: %s", MockConsumerData.Phone, respData.Phone)
			}
		}
	})
}

func TestGetConsumerEndpoint(t *testing.T) {
	t.Parallel()

	t.Run("get consumer simple test", func(t *testing.T) {
		t.Parallel()

		mockService := new(MockService)
		log := logger.NewLogger(os.Stdout, "get consumer simple test")
		handler := consumerhandler.NewConsumerHandler(consumerhandler.Params{
			Logger:          log,
			ConsumerService: mockService,
		})

		resp := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/v1/consumers/1", nil)

		handler.ServeHTTP(resp, req)

		respData := consumerapi.ConsumerResponse{}
		if err := consumerapi.DecodeJSON(resp.Body, &respData); err != nil {
			t.Fatal(err)
		}

		if resp.Code != http.StatusOK {
			t.Fatalf("StatusCode: %d", resp.Code)
		}
		if respData.ID != MockConsumerData.ID {
			t.Errorf("UserID: Expected: %v, Got: %v", MockConsumerData.ID, respData.ID)
		}

		if respData.Firstname != MockConsumerData.Firstname {
			t.Errorf("Firstname: Expected: %s, Got: %s", MockConsumerData.Firstname, respData.Firstname)
		}

		if respData.Lastname != MockConsumerData.Lastname {
			t.Errorf("Lastname: Expected: %s, Got: %s", MockConsumerData.Lastname, respData.Lastname)
		}

		if respData.Email != MockConsumerData.Email {
			t.Errorf("Email: Expected: %s, Got: %s", MockConsumerData.Email, respData.Email)
		}

		if respData.Phone != MockConsumerData.Phone {
			t.Errorf("Phone: Expected: %s, Got: %s", MockConsumerData.Phone, respData.Phone)
		}
	})
}

func TestInsertNewConsumerLocationEndpoint(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name                 string
		consumerLocationData consumerapi.NewLocationRequest
	}{
		{
			"NewConsumerLocation simple test",
			consumerapi.NewLocationRequest{
				Latitude:   "0123456789",
				Longitude:  "0123456789",
				Country:    "TestCountry",
				City:       "Test City",
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
			t.Parallel()

			mockService := new(MockService)

			log := logger.NewLogger(os.Stdout, test.name)
			consumerHandler := consumerhandler.NewConsumerHandler(consumerhandler.Params{
				Logger:          log,
				ConsumerService: mockService,
			})

			reqBody, err := v1.Encode(test.consumerLocationData)
			if err != nil {
				t.Fatal(err)
			}

			resp := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPost, "/v1/consumers/1/location", bytes.NewBuffer(reqBody))

			consumerHandler.ServeHTTP(resp, req)

			if resp.Code != http.StatusOK {
				t.Fatalf("StatusCode: %d", resp.Code)
			}

			respData := consumerapi.LocationResponse{}
			if err = consumerapi.DecodeJSON(resp.Body, &respData); err != nil {
				t.Fatal(err)
			}

			if respData.UserID != MockLocationData.UserID {
				t.Errorf("UserID: Expected: %v, Got: %v", MockLocationData.UserID, respData.UserID)
			}

			if respData.Altitude != MockLocationData.Latitude {
				t.Errorf("Latitude: Expected: %s, Got: %s", test.consumerLocationData.Latitude, respData.Altitude)
			}

			if respData.Longitude != MockLocationData.Longitude {
				t.Errorf("Longitude: Expected: %s, Got: %s", test.consumerLocationData.Longitude, respData.Longitude)
			}

			if respData.Country != MockLocationData.Country {
				t.Errorf("Country: Expected: %s, Got: %s", test.consumerLocationData.Country, respData.Country)
			}

			if respData.City != MockLocationData.City {
				t.Errorf("City: Expected: %s, Got: %s", test.consumerLocationData.City, respData.City)
			}

			if respData.Region != MockLocationData.Region {
				t.Errorf("Region: Expected: %s, Got: %s", test.consumerLocationData.Region, respData.Region)
			}
		})
	}
}

func TestUpdateConsumerLocationEndpoint(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name                 string
		consumerLocationData consumerapi.NewLocationRequest
	}{
		{
			"UpdateLocation simple test",
			consumerapi.NewLocationRequest{
				Latitude:   "0123456789",
				Longitude:  "0123456789",
				Country:    "TestCountry",
				City:       "Test City",
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
			t.Parallel()

			mockService := new(MockService)

			log := logger.NewLogger(os.Stdout, test.name)
			consumerHandler := consumerhandler.NewConsumerHandler(consumerhandler.Params{
				Logger:          log,
				ConsumerService: mockService,
			})

			reqBody, err := v1.Encode(test.consumerLocationData)
			if err != nil {
				t.Fatal(err)
			}

			resp := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPut, "/v1/consumers/1/location", bytes.NewBuffer(reqBody))

			consumerHandler.ServeHTTP(resp, req)

			if resp.Code != http.StatusOK {
				t.Fatalf("StatusCode: %d", resp.Code)
			}

			respData := consumerapi.LocationResponse{}
			if err = consumerapi.DecodeJSON(resp.Body, &respData); err != nil {
				t.Fatal(err)
			}

			if respData.UserID != MockLocationData.UserID {
				t.Errorf("UserID: Expected: %v, Got: %v", MockLocationData.UserID, respData.UserID)
			}

			if respData.Altitude != MockLocationData.Latitude {
				t.Errorf("Latitude: Expected: %s, Got: %s", test.consumerLocationData.Latitude, respData.Altitude)
			}

			if respData.Longitude != MockLocationData.Longitude {
				t.Errorf("Longitude: Expected: %s, Got: %s", test.consumerLocationData.Longitude, respData.Longitude)
			}

			if respData.Country != MockLocationData.Country {
				t.Errorf("Country: Expected: %s, Got: %s", test.consumerLocationData.Country, respData.Country)
			}

			if respData.City != MockLocationData.City {
				t.Errorf("City: Expected: %s, Got: %s", test.consumerLocationData.City, respData.City)
			}

			if respData.Region != MockLocationData.Region {
				t.Errorf("Region: Expected: %s, Got: %s", test.consumerLocationData.Region, respData.Region)
			}
		})
	}
}

func TestGetConsumerLocationEndpoint(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
	}{
		{
			"GetLocation simple test",
		},
	}

	for _, currentTest := range tests {
		test := currentTest

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			mockService := new(MockService)

			log := logger.NewLogger(os.Stdout, test.name)
			consumerHandler := consumerhandler.NewConsumerHandler(consumerhandler.Params{
				Logger:          log,
				ConsumerService: mockService,
			})

			resp := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, "/v1/consumers/1/location", nil)

			consumerHandler.ServeHTTP(resp, req)

			if resp.Code != http.StatusOK {
				t.Fatalf("StatusCode: %d", resp.Code)
			}

			respData := consumerapi.LocationResponse{}
			if err := consumerapi.DecodeJSON(resp.Body, &respData); err != nil {
				t.Fatal(err)
			}

			if respData.UserID != MockLocationData.UserID {
				t.Errorf("UserID: Expected: %v, Got: %v", MockLocationData.UserID, respData.UserID)
			}

			if respData.Altitude != MockLocationData.Latitude {
				t.Errorf("Latitude: Expected: %s, Got: %s", MockLocationData.Latitude, respData.Altitude)
			}

			if respData.Longitude != MockLocationData.Longitude {
				t.Errorf("Longitude: Expected: %s, Got: %s", MockLocationData.Longitude, respData.Longitude)
			}

			if respData.Country != MockLocationData.Country {
				t.Errorf("Country: Expected: %s, Got: %s", MockLocationData.Country, respData.Country)
			}

			if respData.City != MockLocationData.City {
				t.Errorf("City: Expected: %s, Got: %s", MockLocationData.City, respData.City)
			}

			if respData.Region != MockLocationData.Region {
				t.Errorf("Region: Expected: %s, Got: %s", MockLocationData.Region, respData.Region)
			}
		})
	}
}
