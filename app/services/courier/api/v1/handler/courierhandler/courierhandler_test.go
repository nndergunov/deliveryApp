package courierhandler_test

import (
	"bytes"
	"courier/api/v1/courierapi"
	"courier/api/v1/handler/courierhandler"
	"courier/pkg/domain"
	v1 "github.com/nndergunov/deliveryApp/app/pkg/api/v1"
	"github.com/nndergunov/deliveryApp/app/pkg/logger"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
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

	MockCourierLocationData = &domain.CourierLocation{
		CourierID:  1,
		Altitude:   "0123456789",
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

func (m MockService) UpdateCourier(courier domain.Courier, id string) (*domain.Courier, error) {
	return MockCourierData, nil
}

func (m MockService) UpdateCourierAvailable(id, available string) (*domain.Courier, error) {
	return MockCourierData, nil
}

func (m MockService) GetAllCourier(_ map[string]string) ([]domain.Courier, error) {
	return []domain.Courier{*MockCourierData}, nil
}

func (m MockService) GetCourier(id string) (*domain.Courier, error) {
	return MockCourierData, nil
}

func (m MockService) InsertCourierLocation(_ domain.CourierLocation, id string) (*domain.CourierLocation, error) {
	return MockCourierLocationData, nil
}

func (m MockService) GetCourierLocation(id string) (*domain.CourierLocation, error) {
	return MockCourierLocationData, nil
}

func (m MockService) UpdateCourierLocation(courier domain.CourierLocation, id string) (*domain.CourierLocation, error) {
	return MockCourierLocationData, nil
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
			req := httptest.NewRequest(http.MethodPost, "/v1/courier", bytes.NewBuffer(reqBody))

			courierHandler.ServeHTTP(resp, req)

			if resp.Code != http.StatusOK {
				t.Fatalf("StatusCode: %d", resp.Code)
			}

			respData := courierapi.CourierResponse{}
			if err = courierapi.DecodeJSON(resp.Body, &respData); err != nil {
				t.Fatal(err)
			}

			if respData.ID != MockCourierData.ID {
				t.Errorf("ID: Expected: %v, Got: %v", MockCourierData.ID, respData.ID)
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
		req := httptest.NewRequest(http.MethodDelete, "/v1/courier/1", nil)

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
			req := httptest.NewRequest(http.MethodPut, "/v1/courier/1", bytes.NewBuffer(reqBody))

			courierHandler.ServeHTTP(resp, req)

			if resp.Code != http.StatusOK {
				t.Fatalf("StatusCode: %d", resp.Code)
			}

			respData := courierapi.CourierResponse{}
			if err = courierapi.DecodeJSON(resp.Body, &respData); err != nil {
				t.Fatal(err)
			}

			if respData.ID != MockCourierData.ID {
				t.Errorf("ID: Expected: %v, Got: %v", MockCourierData.ID, respData.ID)
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
		req := httptest.NewRequest(http.MethodPut, "/v1/courier/available/1?available=true", nil)

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
		req := httptest.NewRequest(http.MethodGet, "/v1/courier/1", nil)

		handler.ServeHTTP(resp, req)

		respData := courierapi.CourierResponse{}
		if err := courierapi.DecodeJSON(resp.Body, &respData); err != nil {
			t.Fatal(err)
		}

		if resp.Code != http.StatusOK {
			t.Fatalf("StatusCode: %d", resp.Code)
		}
		if respData.ID != MockCourierData.ID {
			t.Errorf("ID: Expected: %v, Got: %v", MockCourierData.ID, respData.ID)
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
		req := httptest.NewRequest(http.MethodGet, "/v1/courier/all", nil)

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
				t.Errorf("ID: Expected: %v, Got: %v", MockCourierData.ID, respData.ID)
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

func TestInsertNewCourierLocationEndpoint(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name                string
		courierLocationData courierapi.NewCourierLocationRequest
	}{
		{
			"NewCourierLocation simple test",
			courierapi.NewCourierLocationRequest{
				Altitude:   "0123456789",
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

			reqBody, err := v1.Encode(test.courierLocationData)
			if err != nil {
				t.Fatal(err)
			}

			resp := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPost, "/v1/courier/location/1", bytes.NewBuffer(reqBody))

			courierHandler.ServeHTTP(resp, req)

			if resp.Code != http.StatusOK {
				t.Fatalf("StatusCode: %d", resp.Code)
			}

			respData := courierapi.CourierLocationResponse{}
			if err = courierapi.DecodeJSON(resp.Body, &respData); err != nil {
				t.Fatal(err)
			}

			if respData.CourierID != MockCourierLocationData.CourierID {
				t.Errorf("CourierID: Expected: %v, Got: %v", MockCourierLocationData.CourierID, respData.CourierID)
			}

			if respData.Altitude != MockCourierLocationData.Altitude {
				t.Errorf("Altitude: Expected: %s, Got: %s", test.courierLocationData.Altitude, respData.Altitude)
			}

			if respData.Longitude != MockCourierLocationData.Longitude {
				t.Errorf("Longitude: Expected: %s, Got: %s", test.courierLocationData.Longitude, respData.Longitude)
			}

			if respData.Country != MockCourierLocationData.Country {
				t.Errorf("Country: Expected: %s, Got: %s", test.courierLocationData.Country, respData.Country)
			}

			if respData.City != MockCourierLocationData.City {
				t.Errorf("City: Expected: %s, Got: %s", test.courierLocationData.City, respData.City)
			}

			if respData.Region != MockCourierLocationData.Region {
				t.Errorf("Region: Expected: %s, Got: %s", test.courierLocationData.Region, respData.Region)
			}
		})
	}
}

func TestUpdateCourierLocationEndpoint(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name                string
		courierLocationData courierapi.NewCourierLocationRequest
	}{
		{
			"UpdateCourierLocation simple test",
			courierapi.NewCourierLocationRequest{
				Altitude:   "0123456789",
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

			reqBody, err := v1.Encode(test.courierLocationData)
			if err != nil {
				t.Fatal(err)
			}

			resp := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPut, "/v1/courier/location/1", bytes.NewBuffer(reqBody))

			courierHandler.ServeHTTP(resp, req)

			if resp.Code != http.StatusOK {
				t.Fatalf("StatusCode: %d", resp.Code)
			}

			respData := courierapi.CourierLocationResponse{}
			if err = courierapi.DecodeJSON(resp.Body, &respData); err != nil {
				t.Fatal(err)
			}

			if respData.CourierID != MockCourierLocationData.CourierID {
				t.Errorf("CourierID: Expected: %v, Got: %v", MockCourierLocationData.CourierID, respData.CourierID)
			}

			if respData.Altitude != MockCourierLocationData.Altitude {
				t.Errorf("Altitude: Expected: %s, Got: %s", test.courierLocationData.Altitude, respData.Altitude)
			}

			if respData.Longitude != MockCourierLocationData.Longitude {
				t.Errorf("Longitude: Expected: %s, Got: %s", test.courierLocationData.Longitude, respData.Longitude)
			}

			if respData.Country != MockCourierLocationData.Country {
				t.Errorf("Country: Expected: %s, Got: %s", test.courierLocationData.Country, respData.Country)
			}

			if respData.City != MockCourierLocationData.City {
				t.Errorf("City: Expected: %s, Got: %s", test.courierLocationData.City, respData.City)
			}

			if respData.Region != MockCourierLocationData.Region {
				t.Errorf("Region: Expected: %s, Got: %s", test.courierLocationData.Region, respData.Region)
			}
		})
	}
}

func TestGetCourierLocationEndpoint(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
	}{
		{
			"GetCourierLocation simple test",
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
			req := httptest.NewRequest(http.MethodGet, "/v1/courier/location/1", nil)

			courierHandler.ServeHTTP(resp, req)

			if resp.Code != http.StatusOK {
				t.Fatalf("StatusCode: %d", resp.Code)
			}

			respData := courierapi.CourierLocationResponse{}
			if err := courierapi.DecodeJSON(resp.Body, &respData); err != nil {
				t.Fatal(err)
			}

			if respData.CourierID != MockCourierLocationData.CourierID {
				t.Errorf("CourierID: Expected: %v, Got: %v", MockCourierLocationData.CourierID, respData.CourierID)
			}

			if respData.Altitude != MockCourierLocationData.Altitude {
				t.Errorf("Altitude: Expected: %s, Got: %s", MockCourierLocationData.Altitude, respData.Altitude)
			}

			if respData.Longitude != MockCourierLocationData.Longitude {
				t.Errorf("Longitude: Expected: %s, Got: %s", MockCourierLocationData.Longitude, respData.Longitude)
			}

			if respData.Country != MockCourierLocationData.Country {
				t.Errorf("Country: Expected: %s, Got: %s", MockCourierLocationData.Country, respData.Country)
			}

			if respData.City != MockCourierLocationData.City {
				t.Errorf("City: Expected: %s, Got: %s", MockCourierLocationData.City, respData.City)
			}

			if respData.Region != MockCourierLocationData.Region {
				t.Errorf("Region: Expected: %s, Got: %s", MockCourierLocationData.Region, respData.Region)
			}
		})
	}
}
