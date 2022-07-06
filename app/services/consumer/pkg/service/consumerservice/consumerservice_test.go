package consumerservice_test

import (
	"os"
	"strconv"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/nndergunov/deliveryApp/app/pkg/logger"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/nndergunov/deliveryApp/app/services/consumer/pkg/domain"
	mockstorage "github.com/nndergunov/deliveryApp/app/services/consumer/pkg/mocks"
	"github.com/nndergunov/deliveryApp/app/services/consumer/pkg/service/consumerservice"
)

func TestInsertNewConsumer(t *testing.T) {
	tests := []struct {
		name string
		in   domain.Consumer
		out  *domain.Consumer
	}{
		{
			"Insert consumer test",
			domain.Consumer{
				Firstname: "testFName",
				Lastname:  "testLName",
				Email:     "test@gmail.com",
				Phone:     "123456789",
			},
			&domain.Consumer{
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
			storage := mockstorage.NewMockConsumerStorage(ctl)
			storage.EXPECT().GetConsumerDuplicateByParam(domain.SearchParam{"email": test.in.Email, "phone": test.in.Phone}).Return(nil, nil)
			storage.EXPECT().InsertConsumer(test.in).Return(test.out, nil)

			service := consumerservice.NewConsumerService(consumerservice.Params{ConsumerStorage: storage, Logger: logger.NewLogger(os.Stdout, "service: ")})
			newAccount, err := service.InsertConsumer(test.in)
			require.NoError(t, err)

			assert.Equal(t, test.out, newAccount)
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
			"Consumer deleted",
		},
	}

	for _, currentTest := range tests {
		test := currentTest

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			ctl := gomock.NewController(t)
			storage := mockstorage.NewMockConsumerStorage(ctl)
			mockInData := 1
			mockOutData := &domain.Consumer{}
			storage.EXPECT().GetConsumerByID(mockInData).Return(mockOutData, nil)
			storage.EXPECT().DeleteConsumer(mockInData).Return(nil)
			storage.EXPECT().DeleteLocation(mockInData).Return(nil)

			service := consumerservice.NewConsumerService(consumerservice.Params{ConsumerStorage: storage, Logger: logger.NewLogger(os.Stdout, "service: ")})
			resp, err := service.DeleteConsumer(test.in)
			require.NoError(t, err)

			assert.Equal(t, test.out, resp)
		})
	}
}

func TestUpdateConsumer(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		in   domain.Consumer
		out  *domain.Consumer
	}{
		{
			"update consumer test",
			domain.Consumer{
				ID:        1,
				Firstname: "testFName",
				Lastname:  "testLName",
				Email:     "test@gmail.com",
				Phone:     "123456789",
			},
			&domain.Consumer{
				ID:        1,
				Firstname: "utestFName",
				Lastname:  "utestLName",
				Email:     "utest@gmail.com",
				Phone:     "1234567899",
			},
		},
	}

	for _, currentTest := range tests {
		test := currentTest

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			ctl := gomock.NewController(t)
			storage := mockstorage.NewMockConsumerStorage(ctl)
			intStr := strconv.Itoa(test.in.ID)
			storage.EXPECT().GetConsumerDuplicateByParam(domain.SearchParam{"id": intStr, "email": test.in.Email, "phone": test.in.Phone}).Return(nil, nil)
			storage.EXPECT().UpdateConsumer(test.in).Return(test.out, nil)

			service := consumerservice.NewConsumerService(consumerservice.Params{ConsumerStorage: storage, Logger: logger.NewLogger(os.Stdout, "service: ")})
			resp, err := service.UpdateConsumer(test.in, intStr)
			require.NoError(t, err)

			assert.Equal(t, test.out, resp)
		})
	}
}

func TestGetAllConsumer(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		out  []domain.Consumer
	}{
		{
			"get all consumer test",
			[]domain.Consumer{
				{
					ID:        1,
					Firstname: "test1FName",
					Lastname:  "test1LName",
					Email:     "test1@gmail.com",
					Phone:     "1234567891",
				},
				{
					ID:        2,
					Firstname: "test2FName",
					Lastname:  "test2LName",
					Email:     "test2@gmail.com",
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
			storage := mockstorage.NewMockConsumerStorage(ctl)
			storage.EXPECT().GetAllConsumer().Return(test.out, nil)

			service := consumerservice.NewConsumerService(consumerservice.Params{ConsumerStorage: storage, Logger: logger.NewLogger(os.Stdout, "service: ")})
			resp, err := service.GetAllConsumer()
			require.NoError(t, err)

			assert.Equal(t, test.out, resp)
		})
	}
}

func TestGetConsumer(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		in   string
		out  *domain.Consumer
	}{
		{
			"get consumer test",
			"1",
			&domain.Consumer{
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
			storage := mockstorage.NewMockConsumerStorage(ctl)
			id, err := strconv.Atoi(test.in)
			require.NoError(t, err)

			storage.EXPECT().GetConsumerByID(id).Return(test.out, nil)

			service := consumerservice.NewConsumerService(consumerservice.Params{ConsumerStorage: storage, Logger: logger.NewLogger(os.Stdout, "service: ")})
			resp, err := service.GetConsumer(test.in)
			require.NoError(t, err)

			assert.Equal(t, test.out, resp)
		})
	}
}

func TestInsertNewConsumerLocation(t *testing.T) {
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
			storage := mockstorage.NewMockConsumerStorage(ctl)

			mockOutData := &domain.Consumer{}
			storage.EXPECT().GetConsumerByID(test.in.UserID).Return(mockOutData, nil)
			storage.EXPECT().GetLocation(test.in.UserID).Return(nil, nil)
			storage.EXPECT().InsertLocation(test.in).Return(test.out, nil)

			service := consumerservice.NewConsumerService(consumerservice.Params{ConsumerStorage: storage, Logger: logger.NewLogger(os.Stdout, "service: ")})
			resp, err := service.InsertLocation(test.in, strconv.Itoa(test.in.UserID))
			require.NoError(t, err)

			assert.Equal(t, test.out, resp)
		})
	}
}

func TestUpdateConsumerLocation(t *testing.T) {
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
			"update location test",
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
			storage := mockstorage.NewMockConsumerStorage(ctl)

			id, err := strconv.Atoi(test.in.userID)
			require.NoError(t, err)
			test.in.location.UserID = id

			storage.EXPECT().UpdateLocation(test.in.location).Return(test.out, nil)

			service := consumerservice.NewConsumerService(consumerservice.Params{ConsumerStorage: storage, Logger: logger.NewLogger(os.Stdout, "service: ")})
			resp, err := service.UpdateLocation(test.in.location, test.in.userID)
			require.NoError(t, err)

			assert.Equal(t, test.out, resp)
		})
	}
}

func TestGetConsumerLocationEndpoint(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		in   string
		out  *domain.Location
	}{
		{
			"get location test",
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
			storage := mockstorage.NewMockConsumerStorage(ctl)

			mockInData, err := strconv.Atoi(test.in)
			require.NoError(t, err)

			storage.EXPECT().GetLocation(mockInData).Return(test.out, nil)

			service := consumerservice.NewConsumerService(consumerservice.Params{ConsumerStorage: storage, Logger: logger.NewLogger(os.Stdout, "service: ")})
			resp, err := service.GetLocation(test.in)
			require.NoError(t, err)

			assert.Equal(t, test.out, resp)
		})
	}
}
