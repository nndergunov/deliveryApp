package consumerhandler_test

import (
	"bytes"
	"consumer/api/v1/consumerapi"
	"consumer/api/v1/handler/consumerhandler"
	"consumer/domain"
	v1 "github.com/nndergunov/deliveryApp/app/pkg/api/v1"
	"github.com/nndergunov/deliveryApp/app/pkg/logger"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var (
	MockConsumerData = &domain.Consumer{
		ID:               1,
		Firstname:        "vasya",
		Lastname:         "",
		Email:            "vasya@gmail.com",
		Phone:            "123456789",
		ConsumerLocation: *MockConsumerLocationData,
	}

	MockConsumerLocationData = &domain.ConsumerLocation{
		ID:          1,
		LocationAlt: "0123456789",
		LocationLat: "0123456789",
		Country:     "",
		City:        "",
		Region:      "",
		Street:      "",
		HomeNumber:  "",
		Floor:       "",
		Door:        "",
	}
)

type MockService struct{}

func (m MockService) InsertConsumer(consumer domain.Consumer) (*domain.Consumer, error) {
	return MockConsumerData, nil
}

func (m MockService) DeleteConsumer(id string) (data any, err error) {
	return "consumer deleted", nil
}

func (m MockService) UpdateConsumer(consumer domain.Consumer, id string) (*domain.Consumer, error) {
	return MockConsumerData, nil
}

func (m MockService) GetAllConsumer() ([]domain.Consumer, error) {
	return []domain.Consumer{*MockConsumerData}, nil
}

func (m MockService) GetConsumer(id string) (*domain.Consumer, error) {
	return MockConsumerData, nil
}

func (m MockService) UpdateConsumerLocation(consumer domain.ConsumerLocation, id string) (*domain.ConsumerLocation, error) {
	return MockConsumerLocationData, nil
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
			courierHandler := consumerhandler.NewConsumerHandler(consumerhandler.Params{
				Logger:          log,
				ConsumerService: mockService,
			})

			reqBody, err := v1.Encode(test.consumerData)
			if err != nil {
				t.Fatal(err)
			}

			resp := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPost, "/v1/consumer/new", bytes.NewBuffer(reqBody))

			courierHandler.ServeHTTP(resp, req)

			if resp.Code != http.StatusOK {
				t.Fatalf("StatusCode: %d", resp.Code)
			}

			respData := consumerapi.ConsumerResponse{}
			if err = consumerapi.DecodeJSON(resp.Body, &respData); err != nil {
				t.Fatal(err)
			}

			if respData.ID != MockConsumerData.ID {
				t.Errorf("ID: Expected: %v, Got: %v", MockConsumerData.ID, respData.ID)
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
		req := httptest.NewRequest(http.MethodPost, "/v1/consumer/delete?id=1", nil)

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
			courierHandler := consumerhandler.NewConsumerHandler(consumerhandler.Params{
				Logger:          log,
				ConsumerService: mockService,
			})

			reqBody, err := v1.Encode(test.consumerData)
			if err != nil {
				t.Fatal(err)
			}

			resp := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPut, "/v1/consumer/update?id=1", bytes.NewBuffer(reqBody))

			courierHandler.ServeHTTP(resp, req)

			if resp.Code != http.StatusOK {
				t.Fatalf("StatusCode: %d", resp.Code)
			}

			respData := consumerapi.ConsumerResponse{}
			if err = consumerapi.DecodeJSON(resp.Body, &respData); err != nil {
				t.Fatal(err)
			}

			if respData.ID != MockConsumerData.ID {
				t.Errorf("ID: Expected: %v, Got: %v", MockConsumerData.ID, respData.ID)
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

	t.Run("get-all consumer simple test", func(t *testing.T) {
		t.Parallel()
		testGetRespList := []*domain.Consumer{MockConsumerData}
		mockService := new(MockService)
		log := logger.NewLogger(os.Stdout, "get-all consumer simple test")
		handler := consumerhandler.NewConsumerHandler(consumerhandler.Params{
			Logger:          log,
			ConsumerService: mockService,
		})

		resp := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/v1/consumer/get-all", nil)

		handler.ServeHTTP(resp, req)

		respDataList := consumerapi.ReturnConsumerList{}
		if err := consumerapi.DecodeJSON(resp.Body, &respDataList); err != nil {
			t.Fatal(err)
		}

		if resp.Code != http.StatusOK {
			t.Fatalf("StatusCode: %d", resp.Code)
		}
		if len(respDataList.ConsumerList) != len(testGetRespList) {
			t.Errorf("len: Expected: %v, Got: %v", len(testGetRespList), len(respDataList.ConsumerList))
		}
		for _, respData := range respDataList.ConsumerList {

			if respData.ID != MockConsumerData.ID {
				t.Errorf("ID: Expected: %v, Got: %v", MockConsumerData.ID, respData.ID)
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
		req := httptest.NewRequest(http.MethodGet, "/v1/consumer/get?id=1", nil)

		handler.ServeHTTP(resp, req)

		respData := consumerapi.ConsumerResponse{}
		if err := consumerapi.DecodeJSON(resp.Body, &respData); err != nil {
			t.Fatal(err)
		}

		if resp.Code != http.StatusOK {
			t.Fatalf("StatusCode: %d", resp.Code)
		}
		if respData.ID != MockConsumerData.ID {
			t.Errorf("ID: Expected: %v, Got: %v", MockConsumerData.ID, respData.ID)
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
