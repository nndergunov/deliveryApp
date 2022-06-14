package courierhandler_test

import (
	"bytes"
	v1 "github.com/nndergunov/deliveryApp/app/pkg/api/v1"
	"github.com/nndergunov/deliveryApp/app/pkg/logger"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"courier/api/v1/courierapi"
	"courier/api/v1/handler/courierhandler"
	"courier/pkg/domain"
)

var (
	MockCourierData = &domain.Courier{
		ID:        1,
		Username:  "TestUsername",
		Firstname: "TestFName",
		Lastname:  "TestLName",
		Email:     "Test@gmail.com",
		Phone:     "123456789",
		Available: true,
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

func (m MockService) InsertCourier(_ domain.Courier) (*domain.Courier, error) {
	return MockCourierData, nil
}

func (m MockService) DeleteCourier(_ string) (data any, err error) {
	return "courier deleted", nil
}

func (m MockService) UpdateCourier(_ domain.Courier, _ string) (*domain.Courier, error) {
	return MockCourierData, nil
}

func (m MockService) UpdateCourierAvailable(_, _ string) (*domain.Courier, error) {
	return MockCourierData, nil
}

func (m MockService) GetAllCourier(_ map[string]string) ([]domain.Courier, error) {
	return []domain.Courier{*MockCourierData}, nil
}

func (m MockService) GetCourier(id string) (*domain.Courier, error) {
	return MockCourierData, nil
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

func TestInsertNewCourierEndpoint(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		courierData courierapi.NewCourierRequest
	}{
		{
			"Insert courier simple test",
			courierapi.NewCourierRequest{
				Username:  "TestUsername",
				Firstname: "TestFName",
				Lastname:  "TestLName",
				Email:     "Test@gmail.com",
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
			courierHandler := courierhandler.NewCourierHandler(courierhandler.Params{
				Logger:         log,
				CourierService: mockService,
			})

			reqBody, err := v1.Encode(test.courierData)
			if err != nil {
				t.Fatal(err)
			}

			resp := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPost, "/v1/couriers", bytes.NewBuffer(reqBody))

			courierHandler.ServeHTTP(resp, req)

			if resp.Code != http.StatusOK {
				t.Fatalf("StatusCode: %d", resp.Code)
			}

			respData := courierapi.CourierResponse{}
			if err = courierapi.DecodeJSON(resp.Body, &respData); err != nil {
				t.Fatal(err)
			}

			if respData.ID != MockCourierData.ID {
				t.Errorf("UserID: Expected: %v, Got: %v", MockCourierData.ID, respData.ID)
			}

			if respData.Username != MockCourierData.Username {
				t.Errorf("Username: Expected: %v, Got: %v", test.courierData.Username, respData.Username)
			}

			if respData.Firstname != MockCourierData.Firstname {
				t.Errorf("Firstname: Expected: %s, Got: %s", test.courierData.Firstname, respData.Firstname)
			}

			if respData.Lastname != MockCourierData.Lastname {
				t.Errorf("Lastname: Expected: %s, Got: %s", test.courierData.Lastname, respData.Lastname)
			}

			if respData.Email != MockCourierData.Email {
				t.Errorf("Email: Expected: %s, Got: %s", test.courierData.Email, respData.Email)
			}

			if respData.Phone != MockCourierData.Phone {
				t.Errorf("Phone: Expected: %s, Got: %s", test.courierData.Phone, respData.Phone)
			}

			if respData.Available != MockCourierData.Available {
				t.Errorf("Available: Expected: %s, Got: %s", test.courierData.Phone, respData.Phone)
			}

		})
	}
}

func TestDeleteCourierEndpoint(t *testing.T) {
	t.Parallel()

	t.Run("Delete Courier simple test", func(t *testing.T) {
		t.Parallel()

		mockService := new(MockService)
		log := logger.NewLogger(os.Stdout, "Delete Courier simple test")
		handler := courierhandler.NewCourierHandler(courierhandler.Params{
			Logger:         log,
			CourierService: mockService,
		})

		resp := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodDelete, "/v1/couriers/1", nil)

		handler.ServeHTTP(resp, req)
		var respData string
		expData := "courier deleted"

		if err := courierapi.DecodeJSON(resp.Body, &respData); err != nil {
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

func TestUpdateCourierEndpoint(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		courierData courierapi.UpdateCourierRequest
	}{
		{
			"Update courier simple test",
			courierapi.UpdateCourierRequest{
				Username:  "TestUsername",
				Firstname: "TestFName",
				Lastname:  "TestLName",
				Email:     "Test@gmail.com",
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
			courierHandler := courierhandler.NewCourierHandler(courierhandler.Params{
				Logger:         log,
				CourierService: mockService,
			})

			reqBody, err := v1.Encode(test.courierData)
			if err != nil {
				t.Fatal(err)
			}

			resp := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPut, "/v1/couriers/1", bytes.NewBuffer(reqBody))

			courierHandler.ServeHTTP(resp, req)

			if resp.Code != http.StatusOK {
				t.Fatalf("StatusCode: %d", resp.Code)
			}

			respData := courierapi.CourierResponse{}
			if err = courierapi.DecodeJSON(resp.Body, &respData); err != nil {
				t.Fatal(err)
			}

			if respData.ID != MockCourierData.ID {
				t.Errorf("UserID: Expected: %v, Got: %v", MockCourierData.ID, respData.ID)
			}

			if respData.Username != MockCourierData.Username {
				t.Errorf("Username: Expected: %v, Got: %v", test.courierData.Username, respData.Username)
			}

			if respData.Firstname != MockCourierData.Firstname {
				t.Errorf("Firstname: Expected: %s, Got: %s", test.courierData.Firstname, respData.Firstname)
			}

			if respData.Lastname != MockCourierData.Lastname {
				t.Errorf("Lastname: Expected: %s, Got: %s", test.courierData.Lastname, respData.Lastname)
			}

			if respData.Email != MockCourierData.Email {
				t.Errorf("Email: Expected: %s, Got: %s", test.courierData.Email, respData.Email)
			}

			if respData.Phone != MockCourierData.Phone {
				t.Errorf("Phone: Expected: %s, Got: %s", test.courierData.Phone, respData.Phone)
			}

			if respData.Available != MockCourierData.Available {
				t.Errorf("Available: Expected: %s, Got: %s", test.courierData.Phone, respData.Phone)
			}

		})
	}
}

func TestUpdateCourierAvailableEndpoint(t *testing.T) {
	t.Parallel()

	t.Run("update courier-available simple test", func(t *testing.T) {
		t.Parallel()

		mockService := new(MockService)
		log := logger.NewLogger(os.Stdout, "update courier-available simple test")
		handler := courierhandler.NewCourierHandler(courierhandler.Params{
			Logger:         log,
			CourierService: mockService,
		})

		resp := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPut, "/v1/couriers-available/1?available=true", nil)

		handler.ServeHTTP(resp, req)

		respData := courierapi.CourierResponse{}

		if err := courierapi.DecodeJSON(resp.Body, &respData); err != nil {
			t.Fatal(err)
		}

		if resp.Code != http.StatusOK {
			t.Fatalf("StatusCode: %d", resp.Code)
		}
		if !respData.Available {
			t.Errorf("response: Expected: %v, Got: %v", true, respData.Available)
		}
	})
}

func TestGetCourierEndpoint(t *testing.T) {
	t.Parallel()

	t.Run("get courier simple test", func(t *testing.T) {
		t.Parallel()

		mockService := new(MockService)
		log := logger.NewLogger(os.Stdout, "get courier simple test")
		handler := courierhandler.NewCourierHandler(courierhandler.Params{
			Logger:         log,
			CourierService: mockService,
		})

		resp := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/v1/couriers/1", nil)

		handler.ServeHTTP(resp, req)

		respData := courierapi.CourierResponse{}
		if err := courierapi.DecodeJSON(resp.Body, &respData); err != nil {
			t.Fatal(err)
		}

		if resp.Code != http.StatusOK {
			t.Fatalf("StatusCode: %d", resp.Code)
		}
		if respData.ID != MockCourierData.ID {
			t.Errorf("UserID: Expected: %v, Got: %v", MockCourierData.ID, respData.ID)
		}

		if respData.Username != MockCourierData.Username {
			t.Errorf("Username: Expected: %v, Got: %v", MockCourierData.Username, respData.Username)
		}

		if respData.Firstname != MockCourierData.Firstname {
			t.Errorf("Firstname: Expected: %s, Got: %s", MockCourierData.Firstname, respData.Firstname)
		}

		if respData.Lastname != MockCourierData.Lastname {
			t.Errorf("Lastname: Expected: %s, Got: %s", MockCourierData.Lastname, respData.Lastname)
		}

		if respData.Email != MockCourierData.Email {
			t.Errorf("Email: Expected: %s, Got: %s", MockCourierData.Email, respData.Email)
		}

		if respData.Phone != MockCourierData.Phone {
			t.Errorf("Phone: Expected: %s, Got: %s", MockCourierData.Phone, respData.Phone)
		}

		if respData.Available != MockCourierData.Available {
			t.Errorf("Available: Expected: %s, Got: %s", MockCourierData.Phone, respData.Phone)
		}

	})
}

func TestGetCourierAllEndpoint(t *testing.T) {
	t.Parallel()

	t.Run("get courier all simple test", func(t *testing.T) {
		t.Parallel()
		testGetRespList := []*domain.Courier{MockCourierData}
		mockService := new(MockService)
		log := logger.NewLogger(os.Stdout, "get-all courier simple test")
		handler := courierhandler.NewCourierHandler(courierhandler.Params{
			Logger:         log,
			CourierService: mockService,
		})

		resp := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/v1/couriers", nil)

		handler.ServeHTTP(resp, req)

		if resp.Code != http.StatusOK {
			t.Fatalf("StatusCode: %d", resp.Code)
		}

		respDataList := courierapi.ReturnCourierResponseList{}
		if err := courierapi.DecodeJSON(resp.Body, &respDataList); err != nil {
			t.Fatal(err)
		}

		if len(respDataList.CourierResponseList) != len(testGetRespList) {
			t.Errorf("len: Expected: %v, Got: %v", len(testGetRespList), len(respDataList.CourierResponseList))
		}

		for _, respData := range respDataList.CourierResponseList {

			if respData.ID != MockCourierData.ID {
				t.Errorf("UserID: Expected: %v, Got: %v", MockCourierData.ID, respData.ID)
			}

			if respData.Username != MockCourierData.Username {
				t.Errorf("Username: Expected: %v, Got: %v", MockCourierData.Username, respData.Username)
			}

			if respData.Firstname != MockCourierData.Firstname {
				t.Errorf("Firstname: Expected: %s, Got: %s", MockCourierData.Firstname, respData.Firstname)
			}

			if respData.Lastname != MockCourierData.Lastname {
				t.Errorf("Lastname: Expected: %s, Got: %s", MockCourierData.Lastname, respData.Lastname)
			}

			if respData.Email != MockCourierData.Email {
				t.Errorf("Email: Expected: %s, Got: %s", MockCourierData.Email, respData.Email)
			}

			if respData.Phone != MockCourierData.Phone {
				t.Errorf("Phone: Expected: %s, Got: %s", MockCourierData.Phone, respData.Phone)
			}

			if respData.Available != MockCourierData.Available {
				t.Errorf("Available: Expected: %s, Got: %s", MockCourierData.Phone, respData.Phone)
			}

		}

	})
}

func TestInsertNewLocationEndpoint(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name         string
		locationData courierapi.NewLocationRequest
	}{
		{
			"New Location simple test",
			courierapi.NewLocationRequest{
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
			courierHandler := courierhandler.NewCourierHandler(courierhandler.Params{
				Logger:         log,
				CourierService: mockService,
			})

			reqBody, err := v1.Encode(test.locationData)
			if err != nil {
				t.Fatal(err)
			}

			resp := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPost, "/v1/couriers/1/location", bytes.NewBuffer(reqBody))

			courierHandler.ServeHTTP(resp, req)

			if resp.Code != http.StatusOK {
				t.Fatalf("StatusCode: %d", resp.Code)
			}

			respData := courierapi.LocationResponse{}
			if err = courierapi.DecodeJSON(resp.Body, &respData); err != nil {
				t.Fatal(err)
			}

			if respData.UserID != MockLocationData.UserID {
				t.Errorf("UserID: Expected: %v, Got: %v", MockLocationData.UserID, respData.UserID)
			}

			if respData.Latitude != MockLocationData.Latitude {
				t.Errorf("Latitude: Expected: %s, Got: %s", test.locationData.Latitude, respData.Latitude)
			}

			if respData.Longitude != MockLocationData.Longitude {
				t.Errorf("Longitude: Expected: %s, Got: %s", test.locationData.Longitude, respData.Longitude)
			}

			if respData.Country != MockLocationData.Country {
				t.Errorf("Country: Expected: %s, Got: %s", test.locationData.Country, respData.Country)
			}

			if respData.City != MockLocationData.City {
				t.Errorf("City: Expected: %s, Got: %s", test.locationData.City, respData.City)
			}

			if respData.Region != MockLocationData.Region {
				t.Errorf("Region: Expected: %s, Got: %s", test.locationData.Region, respData.Region)
			}
		})
	}
}

func TestUpdateLocationEndpoint(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name         string
		locationData courierapi.NewLocationRequest
	}{
		{
			"UpdateLocation simple test",
			courierapi.NewLocationRequest{
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
			courierHandler := courierhandler.NewCourierHandler(courierhandler.Params{
				Logger:         log,
				CourierService: mockService,
			})

			reqBody, err := v1.Encode(test.locationData)
			if err != nil {
				t.Fatal(err)
			}

			resp := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPut, "/v1/couriers/1/location", bytes.NewBuffer(reqBody))

			courierHandler.ServeHTTP(resp, req)

			if resp.Code != http.StatusOK {
				t.Fatalf("StatusCode: %d", resp.Code)
			}

			respData := courierapi.LocationResponse{}
			if err = courierapi.DecodeJSON(resp.Body, &respData); err != nil {
				t.Fatal(err)
			}

			if respData.UserID != MockLocationData.UserID {
				t.Errorf("UserID: Expected: %v, Got: %v", MockLocationData.UserID, respData.UserID)
			}

			if respData.Latitude != MockLocationData.Latitude {
				t.Errorf("Latitude: Expected: %s, Got: %s", test.locationData.Latitude, respData.Latitude)
			}

			if respData.Longitude != MockLocationData.Longitude {
				t.Errorf("Longitude: Expected: %s, Got: %s", test.locationData.Longitude, respData.Longitude)
			}

			if respData.Country != MockLocationData.Country {
				t.Errorf("Country: Expected: %s, Got: %s", test.locationData.Country, respData.Country)
			}

			if respData.City != MockLocationData.City {
				t.Errorf("City: Expected: %s, Got: %s", test.locationData.City, respData.City)
			}

			if respData.Region != MockLocationData.Region {
				t.Errorf("Region: Expected: %s, Got: %s", test.locationData.Region, respData.Region)
			}
		})
	}
}

func TestGetLocationEndpoint(t *testing.T) {
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
			courierHandler := courierhandler.NewCourierHandler(courierhandler.Params{
				Logger:         log,
				CourierService: mockService,
			})

			resp := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, "/v1/couriers/1/location", nil)

			courierHandler.ServeHTTP(resp, req)

			if resp.Code != http.StatusOK {
				t.Fatalf("StatusCode: %d", resp.Code)
			}

			respData := courierapi.LocationResponse{}
			if err := courierapi.DecodeJSON(resp.Body, &respData); err != nil {
				t.Fatal(err)
			}

			if respData.UserID != MockLocationData.UserID {
				t.Errorf("UserID: Expected: %v, Got: %v", MockLocationData.UserID, respData.UserID)
			}

			if respData.Latitude != MockLocationData.Latitude {
				t.Errorf("Latitude: Expected: %s, Got: %s", MockLocationData.Latitude, respData.Latitude)
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
