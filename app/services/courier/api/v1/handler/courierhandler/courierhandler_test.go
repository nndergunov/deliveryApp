package courierhandler_test

import (
	"bytes"
	"courier/api/v1/courierapi"
	"courier/api/v1/handler/courierhandler"
	"courier/domain"
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
		Username:  "vasya123",
		Firstname: "vasya",
		Lastname:  "",
		Email:     "vasya@gmail.com",
		Phone:     "123456789",
		Status:    "active",
		Available: true,
	}
)

type MockService struct{}

func (m MockService) InsertCourier(courier domain.Courier) (*domain.Courier, error) {
	return MockCourierData, nil
}

func (m MockService) RemoveCourier(id string) (data any, err error) {
	return "courier removed", nil
}

func (m MockService) UpdateCourier(courier domain.Courier, id string) (*domain.Courier, error) {
	return MockCourierData, nil
}

func (m MockService) UpdateCourierAvailable(_, _ string) (*domain.Courier, error) {
	return MockCourierData, nil
}

func (m MockService) GetAllCourier() ([]domain.Courier, error) {
	return []domain.Courier{*MockCourierData}, nil
}

func (m MockService) GetCourier(_ string) (*domain.Courier, error) {
	return MockCourierData, nil
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
				Username:  "vasya123",
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
			courierHandler := courierhandler.NewCourierHandler(courierhandler.Params{
				Logger:         log,
				CourierService: mockService,
			})

			reqBody, err := v1.Encode(test.courierData)
			if err != nil {
				t.Fatal(err)
			}

			resp := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPost, "/v1/courier/new", bytes.NewBuffer(reqBody))

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

func TestRemoveCourierEndpoint(t *testing.T) {
	t.Parallel()

	t.Run("Remove Courier simple test", func(t *testing.T) {
		t.Parallel()

		mockService := new(MockService)
		log := logger.NewLogger(os.Stdout, "Remove Courier simple test")
		handler := courierhandler.NewCourierHandler(courierhandler.Params{
			Logger:         log,
			CourierService: mockService,
		})

		resp := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/v1/courier/remove?id=1", nil)

		handler.ServeHTTP(resp, req)
		var respData string
		expData := "courier removed"

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
				Username:  "vasya123",
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
			courierHandler := courierhandler.NewCourierHandler(courierhandler.Params{
				Logger:         log,
				CourierService: mockService,
			})

			reqBody, err := v1.Encode(test.courierData)
			if err != nil {
				t.Fatal(err)
			}

			resp := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPut, "/v1/courier/update?id=1", bytes.NewBuffer(reqBody))

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
		req := httptest.NewRequest(http.MethodPut, "/v1/courier/update-available?id=1&available=true", nil)

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

func TestGetAllCourierEndpoint(t *testing.T) {
	t.Parallel()

	t.Run("get-all courier simple test", func(t *testing.T) {
		t.Parallel()
		testGetRespList := []*domain.Courier{MockCourierData}
		mockService := new(MockService)
		log := logger.NewLogger(os.Stdout, "get-all courier simple test")
		handler := courierhandler.NewCourierHandler(courierhandler.Params{
			Logger:         log,
			CourierService: mockService,
		})

		resp := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/v1/courier/get-all", nil)

		handler.ServeHTTP(resp, req)

		respDataList := courierapi.ReturnCourierList{}
		if err := courierapi.DecodeJSON(resp.Body, &respDataList); err != nil {
			t.Fatal(err)
		}

		if resp.Code != http.StatusOK {
			t.Fatalf("StatusCode: %d", resp.Code)
		}
		if len(respDataList.CourierList) != len(testGetRespList) {
			t.Errorf("len: Expected: %v, Got: %v", len(testGetRespList), len(respDataList.CourierList))
		}
		for _, respData := range respDataList.CourierList {

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
		req := httptest.NewRequest(http.MethodGet, "/v1/courier/get?id=1", nil)

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
