package courierhandler_test

import (
	"bytes"
	"github.com/golang/mock/gomock"
	"github.com/nndergunov/deliveryApp/app/pkg/api/v1/courierapi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/nndergunov/deliveryApp/app/services/courier/api/v1/handler/courierhandler"

	"github.com/nndergunov/deliveryApp/app/services/courier/pkg/domain"

	"github.com/nndergunov/deliveryApp/app/pkg/api/v1"
	"github.com/nndergunov/deliveryApp/app/pkg/logger"

	mockservice "github.com/nndergunov/deliveryApp/app/services/courier/pkg/mocks"
)

func TestInsertNewCourierEndpoint(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		in   courierapi.NewCourierRequest
		out  courierapi.CourierResponse
	}{
		{
			"insert_courier_test",
			courierapi.NewCourierRequest{
				Username:  "testUsername",
				Password:  "testPassword",
				Firstname: "testFName",
				Lastname:  "testLName",
				Email:     "test@gmail.com",
				Phone:     "123456789",
			},
			courierapi.CourierResponse{
				ID:        1,
				Username:  "testUsername",
				Firstname: "testFName",
				Lastname:  "testLName",
				Email:     "test@gmail.com",
				Phone:     "123456789",
				Available: true,
			},
		},
	}

	for _, currentTest := range tests {
		test := currentTest

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			ctl := gomock.NewController(t)
			service := mockservice.NewMockCourierService(ctl)

			mockInData := domain.Courier{
				Username:  test.in.Username,
				Password:  test.in.Password,
				Firstname: test.in.Firstname,
				Lastname:  test.in.Lastname,
				Email:     test.in.Email,
				Phone:     test.in.Phone,
			}

			mockOutData := &domain.Courier{
				ID:        test.out.ID,
				Username:  test.out.Username,
				Password:  test.in.Password,
				Firstname: test.out.Firstname,
				Lastname:  test.out.Lastname,
				Email:     test.out.Email,
				Phone:     test.out.Phone,
				Available: test.out.Available,
			}

			service.EXPECT().InsertCourier(mockInData).Return(mockOutData, nil)

			handler := courierhandler.NewCourierHandler(courierhandler.Params{
				Logger:         logger.NewLogger(os.Stdout, test.name),
				CourierService: service,
			})

			reqBody, err := v1.Encode(test.in)
			require.NoError(t, err)

			resp := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPost, "/v1/couriers", bytes.NewBuffer(reqBody))

			handler.ServeHTTP(resp, req)

			if resp.Code != http.StatusOK {
				t.Fatalf("StatusCode: %d", resp.Code)
			}

			respData := courierapi.CourierResponse{}
			err = courierapi.DecodeJSON(resp.Body, &respData)
			require.NoError(t, err)

			assert.Equal(t, test.out, respData)
		})
	}
}

func TestDeleteCourierEndpoint(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		in   string
		out  string
	}{
		{
			"delete_courier_test",
			"1",
			"courier deleted",
		},
	}

	for _, currentTest := range tests {
		test := currentTest

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			ctl := gomock.NewController(t)
			service := mockservice.NewMockCourierService(ctl)

			service.EXPECT().DeleteCourier(test.in).Return(test.out, nil)

			handler := courierhandler.NewCourierHandler(courierhandler.Params{
				Logger:         logger.NewLogger(os.Stdout, test.name),
				CourierService: service,
			})

			resp := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodDelete, "/v1/couriers/"+test.in, nil)

			handler.ServeHTTP(resp, req)

			if resp.Code != http.StatusOK {
				t.Fatalf("StatusCode: %d", resp.Code)
			}

			respData := ""
			err := courierapi.DecodeJSON(resp.Body, &respData)
			require.NoError(t, err)

			assert.Equal(t, test.out, respData)
		})
	}
}

func TestUpdateCourierEndpoint(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		inID   string
		inBody courierapi.UpdateCourierRequest
		out    courierapi.CourierResponse
	}{
		{
			"update_courier_test",
			"1",
			courierapi.UpdateCourierRequest{
				Username:  "testUsername",
				Firstname: "testFName",
				Lastname:  "testLName",
				Email:     "test@gmail.com",
				Phone:     "123456789",
			},
			courierapi.CourierResponse{
				Username:  "testUsername",
				Firstname: "testFName",
				Lastname:  "testLName",
				Email:     "test@gmail.com",
				Phone:     "123456789",
				Available: true,
			},
		},
	}

	for _, currentTest := range tests {
		test := currentTest

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			ctl := gomock.NewController(t)
			service := mockservice.NewMockCourierService(ctl)

			mockInData := domain.Courier{
				Username:  test.inBody.Username,
				Firstname: test.inBody.Firstname,
				Lastname:  test.inBody.Lastname,
				Email:     test.inBody.Email,
				Phone:     test.inBody.Phone,
			}

			mockOutData := &domain.Courier{
				ID:        test.out.ID,
				Username:  test.out.Username,
				Password:  "",
				Firstname: test.out.Firstname,
				Lastname:  test.out.Lastname,
				Email:     test.out.Email,
				Phone:     test.out.Phone,
				Available: test.out.Available,
			}

			service.EXPECT().UpdateCourier(mockInData, test.inID).Return(mockOutData, nil)

			handler := courierhandler.NewCourierHandler(courierhandler.Params{
				Logger:         logger.NewLogger(os.Stdout, test.name),
				CourierService: service,
			})

			reqBody, err := v1.Encode(test.inBody)
			require.NoError(t, err)

			resp := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPut, "/v1/couriers/"+test.inID, bytes.NewBuffer(reqBody))

			handler.ServeHTTP(resp, req)

			if resp.Code != http.StatusOK {
				t.Fatalf("StatusCode: %d", resp.Code)
			}

			respData := courierapi.CourierResponse{}
			err = courierapi.DecodeJSON(resp.Body, &respData)
			require.NoError(t, err)

			assert.Equal(t, test.out, respData)
		})
	}
}

func TestUpdateCourierAvailableEndpoint(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		inID    string
		inParam string
		out     courierapi.CourierResponse
	}{
		{
			"update_courier_test",
			"1",
			"true",
			courierapi.CourierResponse{
				Username:  "testUsername",
				Firstname: "testFName",
				Lastname:  "testLName",
				Email:     "test@gmail.com",
				Phone:     "123456789",
				Available: true,
			},
		},
	}

	for _, currentTest := range tests {
		test := currentTest

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			ctl := gomock.NewController(t)
			service := mockservice.NewMockCourierService(ctl)

			mockOutData := &domain.Courier{
				ID:        test.out.ID,
				Username:  test.out.Username,
				Password:  "",
				Firstname: test.out.Firstname,
				Lastname:  test.out.Lastname,
				Email:     test.out.Email,
				Phone:     test.out.Phone,
				Available: test.out.Available,
			}
			service.EXPECT().UpdateCourierAvailable(test.inID, test.inParam).Return(mockOutData, nil)

			handler := courierhandler.NewCourierHandler(courierhandler.Params{
				Logger:         logger.NewLogger(os.Stdout, test.name),
				CourierService: service,
			})

			resp := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPut, "/v1/couriers-available/"+test.inID+"?available="+test.inParam, nil)

			handler.ServeHTTP(resp, req)

			if resp.Code != http.StatusOK {
				t.Fatalf("StatusCode: %d", resp.Code)
			}

			respData := courierapi.CourierResponse{}
			err := courierapi.DecodeJSON(resp.Body, &respData)
			require.NoError(t, err)

			assert.Equal(t, test.out, respData)
		})
	}
}

func TestGetCourierEndpoint(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		in   string
		out  courierapi.CourierResponse
	}{
		{
			"get_courier_test",
			"1",
			courierapi.CourierResponse{
				Username:  "testUsername",
				Firstname: "testFName",
				Lastname:  "testLName",
				Email:     "test@gmail.com",
				Phone:     "123456789",
				Available: true,
			},
		},
	}

	for _, currentTest := range tests {
		test := currentTest

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			ctl := gomock.NewController(t)
			service := mockservice.NewMockCourierService(ctl)

			mockOutData := &domain.Courier{
				ID:        test.out.ID,
				Username:  test.out.Username,
				Password:  "",
				Firstname: test.out.Firstname,
				Lastname:  test.out.Lastname,
				Email:     test.out.Email,
				Phone:     test.out.Phone,
				Available: test.out.Available,
			}

			service.EXPECT().GetCourier(test.in).Return(mockOutData, nil)

			handler := courierhandler.NewCourierHandler(courierhandler.Params{
				Logger:         logger.NewLogger(os.Stdout, test.name),
				CourierService: service,
			})

			resp := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, "/v1/couriers/"+test.in, nil)

			handler.ServeHTTP(resp, req)

			if resp.Code != http.StatusOK {
				t.Fatalf("StatusCode: %d", resp.Code)
			}

			respData := courierapi.CourierResponse{}
			err := courierapi.DecodeJSON(resp.Body, &respData)
			require.NoError(t, err)

			assert.Equal(t, test.out, respData)
		})
	}
}

func TestGetCourierListEndpoint(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		out  courierapi.CourierResponseList
	}{
		{
			"get_courier_list_test",
			courierapi.CourierResponseList{
				CourierResponseList: []courierapi.CourierResponse{
					courierapi.CourierResponse{
						ID:        1,
						Username:  "test2Username",
						Firstname: "test2FName",
						Lastname:  "test2LName",
						Email:     "test2@gmail.com",
						Phone:     "1234567892",
						Available: false,
					},
					courierapi.CourierResponse{
						ID:        2,
						Username:  "test2Username",
						Firstname: "test2FName",
						Lastname:  "test2LName",
						Email:     "test2@gmail.com",
						Phone:     "1234567892",
						Available: true,
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
			service := mockservice.NewMockCourierService(ctl)

			mockInData := domain.SearchParam{}

			mockOutDataList := []domain.Courier{}
			for _, data := range test.out.CourierResponseList {
				mockOutData := domain.Courier{
					ID:        data.ID,
					Username:  data.Username,
					Password:  "",
					Firstname: data.Firstname,
					Lastname:  data.Lastname,
					Email:     data.Email,
					Phone:     data.Phone,
					Available: data.Available,
				}
				mockOutDataList = append(mockOutDataList, mockOutData)
			}

			service.EXPECT().GetCourierList(mockInData).Return(mockOutDataList, nil)

			handler := courierhandler.NewCourierHandler(courierhandler.Params{
				Logger:         logger.NewLogger(os.Stdout, test.name),
				CourierService: service,
			})

			resp := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, "/v1/couriers", nil)

			handler.ServeHTTP(resp, req)

			if resp.Code != http.StatusOK {
				t.Fatalf("StatusCode: %d", resp.Code)
			}

			respData := courierapi.CourierResponseList{}
			err := courierapi.DecodeJSON(resp.Body, &respData)
			require.NoError(t, err)

			assert.Equal(t, test.out, respData)
		})
	}
}

func TestInsertNewLocationEndpoint(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		in   courierapi.NewLocationRequest
		out  courierapi.LocationResponse
	}{
		{
			"insert_new_location_test",
			courierapi.NewLocationRequest{
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
			courierapi.LocationResponse{
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
			service := mockservice.NewMockCourierService(ctl)

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

			handler := courierhandler.NewCourierHandler(courierhandler.Params{
				Logger:         logger.NewLogger(os.Stdout, test.name),
				CourierService: service,
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

			respData := courierapi.LocationResponse{}
			err = courierapi.DecodeJSON(resp.Body, &respData)
			assert.NoError(t, err)

			assert.Equal(t, respData, test.out)
		})
	}
}

func TestUpdateLocationEndpoint(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		in   courierapi.NewLocationRequest
		out  courierapi.LocationResponse
	}{
		{
			"update_location_test",
			courierapi.NewLocationRequest{
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
			courierapi.LocationResponse{
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
			service := mockservice.NewMockCourierService(ctl)

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

			handler := courierhandler.NewCourierHandler(courierhandler.Params{
				Logger:         logger.NewLogger(os.Stdout, test.name),
				CourierService: service,
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

			respData := courierapi.LocationResponse{}
			err = courierapi.DecodeJSON(resp.Body, &respData)
			assert.NoError(t, err)

			assert.Equal(t, respData, test.out)
		})
	}
}

func TestGetLocationEndpoint(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		in   string
		out  courierapi.LocationResponse
	}{
		{
			"get location test",
			"1",
			courierapi.LocationResponse{
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
			service := mockservice.NewMockCourierService(ctl)

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

			handler := courierhandler.NewCourierHandler(courierhandler.Params{
				Logger:         logger.NewLogger(os.Stdout, test.name),
				CourierService: service,
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

			respData := courierapi.LocationResponse{}
			err = courierapi.DecodeJSON(resp.Body, &respData)
			assert.NoError(t, err)

			assert.Equal(t, respData, test.out)
		})
	}
}

func TestGetLocationListEndpoint(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		in   string
		out  courierapi.LocationResponseList
	}{
		{
			"get location test",
			"testCity",
			courierapi.LocationResponseList{
				[]courierapi.LocationResponse{
					courierapi.LocationResponse{
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
					courierapi.LocationResponse{
						UserID:     2,
						Latitude:   "u01234567892",
						Longitude:  "u01234567892",
						Country:    "uTestCountry2",
						City:       "uTest City2",
						Region:     "uTestRegion2",
						Street:     "uTestStreet2",
						HomeNumber: "uTestHomeNumber2",
						Floor:      "uTestFloor2",
						Door:       "uTestDoor2",
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
			service := mockservice.NewMockCourierService(ctl)

			mockOutDataList := []domain.Location{}
			for _, data := range test.out.LocationResponseList {
				mockOutData := domain.Location{
					UserID:     data.UserID,
					Latitude:   data.Latitude,
					Longitude:  data.Longitude,
					Country:    data.Country,
					City:       data.City,
					Region:     data.Region,
					Street:     data.Street,
					HomeNumber: data.HomeNumber,
					Floor:      data.Floor,
					Door:       data.Door,
				}
				mockOutDataList = append(mockOutDataList, mockOutData)
			}

			service.EXPECT().GetLocationList(domain.SearchParam{"city": "testCity"}).Return(mockOutDataList, nil)

			handler := courierhandler.NewCourierHandler(courierhandler.Params{
				Logger:         logger.NewLogger(os.Stdout, test.name),
				CourierService: service,
			})

			reqBody, err := v1.Encode(test.in)
			if err != nil {
				t.Fatal(err)
			}

			resp := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, "/v1/locations?city="+test.in, bytes.NewBuffer(reqBody))

			handler.ServeHTTP(resp, req)

			if resp.Code != http.StatusOK {
				t.Fatalf("StatusCode: %d", resp.Code)
			}

			respData := courierapi.LocationResponseList{}
			err = courierapi.DecodeJSON(resp.Body, &respData)
			assert.NoError(t, err)

			assert.Equal(t, respData, test.out)
		})
	}
}
