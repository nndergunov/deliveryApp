package consumerhandler_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/nndergunov/deliveryApp/app/pkg/api/v1"
	"github.com/nndergunov/deliveryApp/app/pkg/api/v1/consumerapi"
	"github.com/nndergunov/deliveryApp/app/pkg/logger"

	"github.com/nndergunov/deliveryApp/app/services/consumer/api/v1/rest/handler/consumerhandler"
	"github.com/nndergunov/deliveryApp/app/services/consumer/pkg/domain"
	mockservice "github.com/nndergunov/deliveryApp/app/services/consumer/pkg/mocks"
)

func TestInsertNewConsumerEndpoint(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		in   consumerapi.NewConsumerRequest
		out  consumerapi.ConsumerResponse
	}{
		{
			"Insert consumer simple test",
			consumerapi.NewConsumerRequest{
				Firstname: "testFName",
				Lastname:  "testLName",
				Email:     "test@gmail.com",
				Phone:     "123456789",
			},
			consumerapi.ConsumerResponse{
				ID:        1,
				Firstname: "testFName",
				Lastname:  "testLName",
				Email:     "test@gmail.com",
				Phone:     "123456789",
			},
		},
	}

	for _, currentTest := range tests {
		test := currentTest

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			ctl := gomock.NewController(t)
			service := mockservice.NewMockConsumerService(ctl)

			mockInData := domain.Consumer{
				Firstname: test.in.Firstname,
				Lastname:  test.in.Lastname,
				Email:     test.in.Email,
				Phone:     test.in.Phone,
			}

			mockOutData := &domain.Consumer{
				ID:        test.out.ID,
				Firstname: test.out.Firstname,
				Lastname:  test.out.Lastname,
				Email:     test.out.Email,
				Phone:     test.out.Phone,
			}

			service.EXPECT().InsertConsumer(mockInData).Return(mockOutData, nil)

			handler := consumerhandler.NewConsumerHandler(consumerhandler.Params{
				Logger:          logger.NewLogger(os.Stdout, test.name),
				ConsumerService: service,
			})

			reqBody, err := v1.Encode(test.in)
			require.NoError(t, err)

			resp := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPost, "/v1/consumers", bytes.NewBuffer(reqBody))

			handler.ServeHTTP(resp, req)

			if resp.Code != http.StatusOK {
				t.Fatalf("StatusCode: %d", resp.Code)
			}

			respData := consumerapi.ConsumerResponse{}
			err = consumerapi.DecodeJSON(resp.Body, &respData)
			require.NoError(t, err)

			assert.Equal(t, test.out, respData)
		})
	}
}

func TestDeleteConsumerEndpoint(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		in   string
		out  string
	}{
		{
			"delete consumer test",
			"1",
			"consumer deleted",
		},
	}

	for _, currentTest := range tests {
		test := currentTest

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			ctl := gomock.NewController(t)
			service := mockservice.NewMockConsumerService(ctl)

			service.EXPECT().DeleteConsumer(test.in).Return(test.out, nil)

			handler := consumerhandler.NewConsumerHandler(consumerhandler.Params{
				Logger:          logger.NewLogger(os.Stdout, test.name),
				ConsumerService: service,
			})

			resp := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodDelete, "/v1/consumers/"+test.in, nil)

			handler.ServeHTTP(resp, req)
			var respData string
			err := consumerapi.DecodeJSON(resp.Body, &respData)
			require.NoError(t, err)

			if resp.Code != http.StatusOK {
				t.Fatalf("StatusCode: %d", resp.Code)
			}

			assert.Equal(t, test.out, respData)
		})
	}
}

func TestUpdateConsumerEndpoint(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		in   consumerapi.NewConsumerRequest
		out  consumerapi.ConsumerResponse
	}{
		{
			"Update consumer test",
			consumerapi.NewConsumerRequest{
				Firstname: "testFName",
				Lastname:  "testLName",
				Email:     "test@gmail.com",
				Phone:     "123456789",
			},

			consumerapi.ConsumerResponse{
				ID:        1,
				Firstname: "UTestFName",
				Lastname:  "UTestLName",
				Email:     "UTest@gmail.com",
				Phone:     "123456788",
			},
		},
	}

	for _, currentTest := range tests {
		test := currentTest

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			ctl := gomock.NewController(t)
			service := mockservice.NewMockConsumerService(ctl)

			mockInData := domain.Consumer{
				Firstname: test.in.Firstname,
				Lastname:  test.in.Lastname,
				Email:     test.in.Email,
				Phone:     test.in.Phone,
			}

			mockOutData := &domain.Consumer{
				ID:        test.out.ID,
				Firstname: test.out.Firstname,
				Lastname:  test.out.Lastname,
				Email:     test.out.Email,
				Phone:     test.out.Phone,
			}

			service.EXPECT().UpdateConsumer(mockInData, "1").Return(mockOutData, nil)

			handler := consumerhandler.NewConsumerHandler(consumerhandler.Params{
				Logger:          logger.NewLogger(os.Stdout, test.name),
				ConsumerService: service,
			})

			reqBody, err := v1.Encode(test.in)
			require.NoError(t, err)

			resp := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPut, "/v1/consumers/1", bytes.NewBuffer(reqBody))

			handler.ServeHTTP(resp, req)

			if resp.Code != http.StatusOK {
				t.Fatalf("StatusCode: %d", resp.Code)
			}

			respData := consumerapi.ConsumerResponse{}
			err = consumerapi.DecodeJSON(resp.Body, &respData)
			require.NoError(t, err)

			assert.Equal(t, respData, test.out)
		})
	}
}

func TestGetAllConsumerEndpoint(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		out  consumerapi.ReturnConsumerResponseList
	}{
		{
			name: "Get all consumer test",
			out: consumerapi.ReturnConsumerResponseList{
				ConsumerResponseList: []consumerapi.ConsumerResponse{
					{
						ID:        1,
						Firstname: "test1FName",
						Lastname:  "test1LName",
						Email:     "test1@gmail.com",
						Phone:     "111111111",
					},
					{
						ID:        2,
						Firstname: "test2FName",
						Lastname:  "test2LName",
						Email:     "test2@gmail.com",
						Phone:     "222222222",
					},
				},
			},
		},
	}
	for _, currentTest := range tests {
		test := currentTest

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			ctl := gomock.NewController(t)
			service := mockservice.NewMockConsumerService(ctl)

			var mockOutDataList []domain.Consumer
			for _, data := range test.out.ConsumerResponseList {
				mockOutData := domain.Consumer{
					ID:        data.ID,
					Firstname: data.Firstname,
					Lastname:  data.Lastname,
					Email:     data.Email,
					Phone:     data.Phone,
				}
				mockOutDataList = append(mockOutDataList, mockOutData)
			}

			service.EXPECT().GetAllConsumer().Return(mockOutDataList, nil)

			handler := consumerhandler.NewConsumerHandler(consumerhandler.Params{
				Logger:          logger.NewLogger(os.Stdout, test.name),
				ConsumerService: service,
			})

			resp := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, "/v1/consumers", nil)

			handler.ServeHTTP(resp, req)

			if resp.Code != http.StatusOK {
				t.Fatalf("StatusCode: %d", resp.Code)
			}

			respData := consumerapi.ReturnConsumerResponseList{}
			err := consumerapi.DecodeJSON(resp.Body, &respData)
			require.NoError(t, err)

			assert.Equal(t, respData, test.out)
		})
	}
}

func TestGetConsumerEndpoint(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		in   string
		out  consumerapi.ConsumerResponse
	}{
		{
			"Update consumer test",
			"1",

			consumerapi.ConsumerResponse{
				ID:        1,
				Firstname: "test1FName",
				Lastname:  "test1LName",
				Email:     "test1@gmail.com",
				Phone:     "111111111",
			},
		},
	}

	for _, currentTest := range tests {
		test := currentTest

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			ctl := gomock.NewController(t)
			service := mockservice.NewMockConsumerService(ctl)

			mockOutData := &domain.Consumer{
				ID:        test.out.ID,
				Firstname: test.out.Firstname,
				Lastname:  test.out.Lastname,
				Email:     test.out.Email,
				Phone:     test.out.Phone,
			}

			service.EXPECT().GetConsumer("1").Return(mockOutData, nil)

			handler := consumerhandler.NewConsumerHandler(consumerhandler.Params{
				Logger:          logger.NewLogger(os.Stdout, test.name),
				ConsumerService: service,
			})

			reqBody, err := v1.Encode(test.in)
			require.NoError(t, err)

			resp := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, "/v1/consumers/"+test.in, bytes.NewBuffer(reqBody))

			handler.ServeHTTP(resp, req)

			if resp.Code != http.StatusOK {
				t.Fatalf("StatusCode: %d", resp.Code)
			}

			respData := consumerapi.ConsumerResponse{}
			err = consumerapi.DecodeJSON(resp.Body, &respData)
			require.NoError(t, err)

			assert.Equal(t, respData, test.out)
		})
	}
}

func TestInsertNewConsumerLocationEndpoint(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		in   consumerapi.NewLocationRequest
		out  consumerapi.LocationResponse
	}{
		{
			"insert new location simple test",
			consumerapi.NewLocationRequest{
				Latitude:   "0123456789",
				Longitude:  "0123456789",
				Country:    "TestCountry",
				City:       "Test City",
				Region:     "TestRegion",
				Street:     "TestStreet",
				HomeNumber: "TestHomeNumber",
				Floor:      "TestFloor",
				Door:       "TestDoor",
			},
			consumerapi.LocationResponse{
				UserID:     1,
				Latitude:   "0123456789",
				Longitude:  "0123456789",
				Country:    "TestCountry",
				City:       "Test City",
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
			t.Parallel()

			ctl := gomock.NewController(t)
			service := mockservice.NewMockConsumerService(ctl)

			mockInData := domain.Location{
				Latitude:   test.in.Latitude,
				Longitude:  test.in.Longitude,
				Country:    test.in.Country,
				City:       test.in.City,
				Region:     test.in.Region,
				Street:     test.in.Street,
				HomeNumber: test.in.HomeNumber,
				Floor:      test.in.Floor,
				Door:       test.in.Door,
			}

			mockOutData := &domain.Location{
				UserID:     test.out.UserID,
				Latitude:   test.out.Latitude,
				Longitude:  test.out.Longitude,
				Country:    test.out.Country,
				City:       test.out.City,
				Region:     test.out.Region,
				Street:     test.out.Street,
				HomeNumber: test.out.HomeNumber,
				Floor:      test.out.Floor,
				Door:       test.out.Door,
			}

			service.EXPECT().InsertLocation(mockInData, "1").Return(mockOutData, nil)

			handler := consumerhandler.NewConsumerHandler(consumerhandler.Params{
				Logger:          logger.NewLogger(os.Stdout, test.name),
				ConsumerService: service,
			})

			reqBody, err := v1.Encode(test.in)
			if err != nil {
				t.Fatal(err)
			}

			resp := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPost, "/v1/locations/1", bytes.NewBuffer(reqBody))

			handler.ServeHTTP(resp, req)

			if resp.Code != http.StatusOK {
				t.Fatalf("StatusCode: %d", resp.Code)
			}

			respData := consumerapi.LocationResponse{}
			err = consumerapi.DecodeJSON(resp.Body, &respData)
			assert.NoError(t, err)

			assert.Equal(t, respData, test.out)
		})
	}
}

func TestUpdateConsumerLocationEndpoint(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		in   consumerapi.NewLocationRequest
		out  consumerapi.LocationResponse
	}{
		{
			"update location test",
			consumerapi.NewLocationRequest{
				Latitude:   "0123456789",
				Longitude:  "0123456789",
				Country:    "TestCountry",
				City:       "Test City",
				Region:     "TestRegion",
				Street:     "TestStreet",
				HomeNumber: "TestHomeNumber",
				Floor:      "TestFloor",
				Door:       "TestDoor",
			},
			consumerapi.LocationResponse{
				UserID:     1,
				Latitude:   "u0123456789",
				Longitude:  "u0123456789",
				Country:    "uTestCountry",
				City:       "uTest City",
				Region:     "uTestRegion",
				Street:     "uTestStreet",
				HomeNumber: "uTestHomeNumber",
				Floor:      "uTestFloor",
				Door:       "uTestDoor",
			},
		},
	}

	for _, currentTest := range tests {
		test := currentTest

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			ctl := gomock.NewController(t)
			service := mockservice.NewMockConsumerService(ctl)

			mockInData := domain.Location{
				Latitude:   test.in.Latitude,
				Longitude:  test.in.Longitude,
				Country:    test.in.Country,
				City:       test.in.City,
				Region:     test.in.Region,
				Street:     test.in.Street,
				HomeNumber: test.in.HomeNumber,
				Floor:      test.in.Floor,
				Door:       test.in.Door,
			}

			mockOutData := &domain.Location{
				UserID:     test.out.UserID,
				Latitude:   test.out.Latitude,
				Longitude:  test.out.Longitude,
				Country:    test.out.Country,
				City:       test.out.City,
				Region:     test.out.Region,
				Street:     test.out.Street,
				HomeNumber: test.out.HomeNumber,
				Floor:      test.out.Floor,
				Door:       test.out.Door,
			}

			service.EXPECT().UpdateLocation(mockInData, "1").Return(mockOutData, nil)

			handler := consumerhandler.NewConsumerHandler(consumerhandler.Params{
				Logger:          logger.NewLogger(os.Stdout, test.name),
				ConsumerService: service,
			})

			reqBody, err := v1.Encode(test.in)
			if err != nil {
				t.Fatal(err)
			}

			resp := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPut, "/v1/locations/1", bytes.NewBuffer(reqBody))

			handler.ServeHTTP(resp, req)

			if resp.Code != http.StatusOK {
				t.Fatalf("StatusCode: %d", resp.Code)
			}

			respData := consumerapi.LocationResponse{}
			err = consumerapi.DecodeJSON(resp.Body, &respData)
			assert.NoError(t, err)

			assert.Equal(t, respData, test.out)
		})
	}
}

func TestGetConsumerLocationEndpoint(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		in   string
		out  consumerapi.LocationResponse
	}{
		{
			"get location test",
			"1",
			consumerapi.LocationResponse{
				UserID:     1,
				Latitude:   "u0123456789",
				Longitude:  "u0123456789",
				Country:    "uTestCountry",
				City:       "uTest City",
				Region:     "uTestRegion",
				Street:     "uTestStreet",
				HomeNumber: "uTestHomeNumber",
				Floor:      "uTestFloor",
				Door:       "uTestDoor",
			},
		},
	}

	for _, currentTest := range tests {
		test := currentTest

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			ctl := gomock.NewController(t)
			service := mockservice.NewMockConsumerService(ctl)

			mockOutData := &domain.Location{
				UserID:     test.out.UserID,
				Latitude:   test.out.Latitude,
				Longitude:  test.out.Longitude,
				Country:    test.out.Country,
				City:       test.out.City,
				Region:     test.out.Region,
				Street:     test.out.Street,
				HomeNumber: test.out.HomeNumber,
				Floor:      test.out.Floor,
				Door:       test.out.Door,
			}

			service.EXPECT().GetLocation(test.in).Return(mockOutData, nil)

			handler := consumerhandler.NewConsumerHandler(consumerhandler.Params{
				Logger:          logger.NewLogger(os.Stdout, test.name),
				ConsumerService: service,
			})

			reqBody, err := v1.Encode(test.in)
			if err != nil {
				t.Fatal(err)
			}

			resp := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, "/v1/locations/1", bytes.NewBuffer(reqBody))

			handler.ServeHTTP(resp, req)

			if resp.Code != http.StatusOK {
				t.Fatalf("StatusCode: %d", resp.Code)
			}

			respData := consumerapi.LocationResponse{}
			err = consumerapi.DecodeJSON(resp.Body, &respData)
			assert.NoError(t, err)

			assert.Equal(t, respData, test.out)
		})
	}
}
