package courierservice_test

import (
	"os"
	"strconv"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/nndergunov/deliveryApp/app/pkg/logger"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/nndergunov/deliveryApp/app/services/courier/pkg/domain"
	mockstorage "github.com/nndergunov/deliveryApp/app/services/courier/pkg/mocks"
	"github.com/nndergunov/deliveryApp/app/services/courier/pkg/service/courierservice"
)

func TestInsertNewCourier(t *testing.T) {
	tests := []struct {
		name string
		in   domain.Courier
		out  *domain.Courier
	}{
		{
			"insert_courier_test",
			domain.Courier{
				Username:  "testUsername",
				Password:  "testPassword",
				Firstname: "testFName",
				Lastname:  "testLName",
				Email:     "test@gmail.com",
				Phone:     "123456789",
			},
			&domain.Courier{
				ID:        1,
				Username:  "testUsername",
				Password:  "testPassword",
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
			storage := mockstorage.NewMockCourierStorage(ctl)
			storage.EXPECT().GetCourierDuplicateByParam(domain.SearchParam{"username": test.in.Username}).Return(nil, nil)
			storage.EXPECT().InsertCourier(test.in).Return(test.out, nil)

			service := courierservice.NewService(courierservice.Params{Storage: storage, Logger: logger.NewLogger(os.Stdout, "service: ")})
			newAccount, err := service.InsertCourier(test.in)
			require.NoError(t, err)

			assert.Equal(t, test.out, newAccount)
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
			storage := mockstorage.NewMockCourierStorage(ctl)

			mockInData := 1
			mockOutData := &domain.Courier{}
			storage.EXPECT().GetCourierByID(mockInData).Return(mockOutData, nil)
			storage.EXPECT().DeleteCourier(mockInData).Return(nil)
			storage.EXPECT().DeleteLocation(mockInData).Return(nil)

			service := courierservice.NewService(courierservice.Params{Storage: storage, Logger: logger.NewLogger(os.Stdout, "service: ")})
			resp, err := service.DeleteCourier(test.in)
			require.NoError(t, err)

			assert.Equal(t, test.out, resp)
		})
	}
}

func TestUpdateCourier(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		in   domain.Courier
		out  *domain.Courier
	}{
		{
			"update_courier_test",
			domain.Courier{
				ID:        1,
				Username:  "testUsername",
				Firstname: "testFName",
				Lastname:  "testLName",
				Email:     "",
				Phone:     "",
			},
			&domain.Courier{
				ID:        1,
				Username:  "testUsername",
				Firstname: "testFName",
				Lastname:  "testLName",
				Email:     "",
				Phone:     "",
			},
		},
	}

	for _, currentTest := range tests {
		test := currentTest

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			ctl := gomock.NewController(t)
			storage := mockstorage.NewMockCourierStorage(ctl)
			intStr := strconv.Itoa(test.in.ID)

			storage.EXPECT().GetCourierDuplicateByParam(domain.SearchParam{"id": intStr, "username": test.in.Username}).Return(nil, nil).Times(1)

			storage.EXPECT().UpdateCourier(test.in).Return(test.out, nil)

			service := courierservice.NewService(courierservice.Params{Storage: storage, Logger: logger.NewLogger(os.Stdout, "service: ")})
			resp, err := service.UpdateCourier(test.in, intStr)
			require.NoError(t, err)

			assert.Equal(t, test.out, resp)
		})
	}
}

func TestUpdateCourierAvailable(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name        string
		inID        string
		InAvailable string
		out         *domain.Courier
	}{
		{
			"update_courier_available_test",
			"1",
			"true",
			&domain.Courier{
				ID:        1,
				Username:  "testUsername",
				Password:  "testPassword",
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
			storage := mockstorage.NewMockCourierStorage(ctl)

			idInt, err := strconv.Atoi(test.inID)
			require.NoError(t, err)

			storage.EXPECT().GetCourierByID(idInt).Return(test.out, nil)

			availableBool, err := strconv.ParseBool(test.InAvailable)
			require.NoError(t, err)

			storage.EXPECT().UpdateCourierAvailable(idInt, availableBool).Return(test.out, nil)

			service := courierservice.NewService(courierservice.Params{Storage: storage, Logger: logger.NewLogger(os.Stdout, "service: ")})
			resp, err := service.UpdateCourierAvailable(test.inID, test.InAvailable)
			require.NoError(t, err)

			assert.Equal(t, test.out, resp)
		})
	}
}

func TestGetAllCourier(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		in   domain.SearchParam
		out  []domain.Courier
	}{
		{
			"get all courier test",
			domain.SearchParam{
				"available": "true",
			},
			[]domain.Courier{
				{
					ID:        1,
					Username:  "testUsername",
					Password:  "testPassword",
					Firstname: "testFName",
					Lastname:  "testLName",
					Email:     "test@gmail.com",
					Phone:     "123456789",
				},
				{
					ID:        2,
					Username:  "testUsername2",
					Password:  "testPassword2",
					Firstname: "testFName2",
					Lastname:  "testLName2",
					Email:     "test@gmail.com2",
					Phone:     "1234567892",
				},
			},
		},
	}

	for _, currentTest := range tests {
		test := currentTest

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			ctl := gomock.NewController(t)
			storage := mockstorage.NewMockCourierStorage(ctl)
			storage.EXPECT().GetCourierList(test.in).Return(test.out, nil)

			service := courierservice.NewService(courierservice.Params{Storage: storage, Logger: logger.NewLogger(os.Stdout, "service: ")})
			resp, err := service.GetCourierList(test.in)
			require.NoError(t, err)

			assert.Equal(t, test.out, resp)
		})
	}
}

func TestGetCourier(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		in   string
		out  *domain.Courier
	}{
		{
			"get courier test",
			"1",
			&domain.Courier{
				ID:        1,
				Firstname: "test1FName",
				Lastname:  "test1LName",
				Email:     "test1@gmail.com",
				Phone:     "1234567891",
			},
		},
	}

	for _, currentTest := range tests {
		test := currentTest

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			ctl := gomock.NewController(t)
			storage := mockstorage.NewMockCourierStorage(ctl)
			id, err := strconv.Atoi(test.in)
			require.NoError(t, err)

			storage.EXPECT().GetCourierByID(id).Return(test.out, nil)

			service := courierservice.NewService(courierservice.Params{Storage: storage, Logger: logger.NewLogger(os.Stdout, "service: ")})
			resp, err := service.GetCourier(test.in)
			require.NoError(t, err)

			assert.Equal(t, test.out, resp)
		})
	}
}

func TestInsertNewCourierLocation(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		in   domain.Location
		out  *domain.Location
	}{
		{
			"insert new location test",
			domain.Location{
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
			&domain.Location{
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
			storage := mockstorage.NewMockCourierStorage(ctl)

			mockOutData := &domain.Courier{}
			storage.EXPECT().GetCourierByID(test.in.UserID).Return(mockOutData, nil)
			storage.EXPECT().GetLocation(test.in.UserID).Return(nil, nil)
			storage.EXPECT().InsertLocation(test.in).Return(test.out, nil)

			service := courierservice.NewService(courierservice.Params{Storage: storage, Logger: logger.NewLogger(os.Stdout, "service: ")})
			resp, err := service.InsertLocation(test.in, strconv.Itoa(test.in.UserID))
			require.NoError(t, err)

			assert.Equal(t, test.out, resp)
		})
	}
}

func TestUpdateCourierLocation(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		in   struct {
			location domain.Location
			userID   string
		}
		out *domain.Location
	}{
		{
			"update_location_test",
			struct {
				location domain.Location
				userID   string
			}{
				location: domain.Location{
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
				userID: "1",
			},

			&domain.Location{
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
			storage := mockstorage.NewMockCourierStorage(ctl)

			id, err := strconv.Atoi(test.in.userID)
			require.NoError(t, err)
			test.in.location.UserID = id

			userIDStr, err := strconv.Atoi(test.in.userID)
			require.NoError(t, err)

			storage.EXPECT().GetLocation(userIDStr).Return(test.out, nil)
			storage.EXPECT().UpdateLocation(test.in.location).Return(test.out, nil)

			service := courierservice.NewService(courierservice.Params{Storage: storage, Logger: logger.NewLogger(os.Stdout, "service: ")})
			resp, err := service.UpdateLocation(test.in.location, test.in.userID)
			require.NoError(t, err)

			assert.Equal(t, test.out, resp)
		})
	}
}

func TestGetCourierLocation(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		in   string
		out  *domain.Location
	}{
		{
			"get_location_test",
			"1",
			&domain.Location{
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
			storage := mockstorage.NewMockCourierStorage(ctl)

			mockInData, err := strconv.Atoi(test.in)
			require.NoError(t, err)

			storage.EXPECT().GetLocation(mockInData).Return(test.out, nil)

			service := courierservice.NewService(courierservice.Params{Storage: storage, Logger: logger.NewLogger(os.Stdout, "service: ")})
			resp, err := service.GetLocation(test.in)
			require.NoError(t, err)

			assert.Equal(t, test.out, resp)
		})
	}
}

func TestGetCourierLocationList(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		in   domain.SearchParam
		out  []domain.Location
	}{
		{
			"get_location_list_test",
			domain.SearchParam{
				"city": "testIstanbul",
			},
			[]domain.Location{
				{
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
		},
	}

	for _, currentTest := range tests {
		test := currentTest

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			ctl := gomock.NewController(t)
			storage := mockstorage.NewMockCourierStorage(ctl)

			storage.EXPECT().GetLocationList(test.in).Return(test.out, nil)

			service := courierservice.NewService(courierservice.Params{Storage: storage, Logger: logger.NewLogger(os.Stdout, "service: ")})
			resp, err := service.GetLocationList(test.in)
			require.NoError(t, err)

			assert.Equal(t, test.out, resp)
		})
	}
}
