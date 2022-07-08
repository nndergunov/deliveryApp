package consumerstorage_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/nndergunov/deliveryApp/app/services/consumer/pkg/db/dbtest"
	"github.com/nndergunov/deliveryApp/app/services/consumer/pkg/docker"
	"github.com/nndergunov/deliveryApp/app/services/consumer/pkg/domain"
	"github.com/nndergunov/deliveryApp/app/services/consumer/pkg/storage/consumerstorage"
)

var c *docker.Container

func TestMain(m *testing.M) {
	var err error
	c, err = dbtest.StartDB()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer dbtest.StopDB(c)

	m.Run()
}

// equalConsumer - function to compare selected fields from struct. Fields which needed
func equalConsumer(t *testing.T, get *domain.Consumer, want *domain.Consumer) {
	if get.Firstname != want.Firstname {
		t.Errorf("Firstname: Expected: %s, Got: %s", want.Firstname, get.Firstname)
	}
	if get.Lastname != want.Lastname {
		t.Errorf("Lastname: Expected: %s, Got: %s", want.Lastname, get.Lastname)
	}

	if get.Email != want.Email {
		t.Errorf("Email: Expected: %s, Got: %s", want.Email, get.Email)
	}

	if get.Phone != want.Phone {
		t.Errorf("Phone: Expected: %s, Got: %s", want.Phone, get.Phone)
	}
}

func TestInsertConsumer(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		in   domain.Consumer
		out  *domain.Consumer
	}{
		{
			"insert_consumer_test",
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

			database, teardown := dbtest.NewUnit(t, c, test.name)
			t.Cleanup(teardown)

			s := consumerstorage.NewConsumerStorage(consumerstorage.Params{DB: database})

			resp, err := s.InsertConsumer(test.in)
			require.NoError(t, err)
			require.NotNil(t, resp)

			equalConsumer(t, resp, test.out)
		})
	}
}

func TestDeleteConsumer(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		in   domain.Consumer
		out  error
	}{
		{
			"delete_consumer_test",
			domain.Consumer{
				ID:        1,
				Firstname: "testFName",
				Lastname:  "testLName",
				Email:     "test@gmail.com",
				Phone:     "123456789",
			},
			nil,
		},
	}

	for _, currentTest := range tests {
		test := currentTest

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			database, teardown := dbtest.NewUnit(t, c, test.name)
			t.Cleanup(teardown)

			s := consumerstorage.NewConsumerStorage(consumerstorage.Params{DB: database})

			resp, err := s.InsertConsumer(test.in)
			require.NoError(t, err)
			require.NotNil(t, resp)

			err = s.DeleteConsumer(test.in.ID)
			if err != test.out {
				t.Errorf("DeleteConsumer: Expected: %v, Got: %v", test.out, err)
			}
		})
	}
}

func TestUpdateConsumer(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		in       domain.Consumer
		updateIn domain.Consumer
		out      *domain.Consumer
	}{
		{
			"update_consumer_test",
			domain.Consumer{
				ID:        1,
				Firstname: "testFName",
				Lastname:  "testLName",
				Email:     "test@gmail.com",
				Phone:     "123456789",
			},
			domain.Consumer{
				ID:        1,
				Firstname: "utestFName",
				Lastname:  "utestLName",
				Email:     "utest@gmail.com",
				Phone:     "1234567899",
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

			database, teardown := dbtest.NewUnit(t, c, test.name)
			t.Cleanup(teardown)

			s := consumerstorage.NewConsumerStorage(consumerstorage.Params{DB: database})

			resp, err := s.InsertConsumer(test.in)
			require.NoError(t, err)
			require.NotNil(t, resp)

			resp2, err := s.UpdateConsumer(test.updateIn)
			require.NoError(t, err)
			require.NotNil(t, resp)

			equalConsumer(t, resp2, test.out)
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
			"get_all_consumer_test",
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

			database, teardown := dbtest.NewUnit(t, c, test.name)
			t.Cleanup(teardown)

			s := consumerstorage.NewConsumerStorage(consumerstorage.Params{DB: database})

			for _, data := range test.out {
				resp, err := s.InsertConsumer(data)
				require.NoError(t, err)
				require.NotNil(t, resp)
			}

			respList, err := s.GetAllConsumer()
			require.NoError(t, err)
			require.NotNil(t, respList)
			for i, resp := range respList {
				equalConsumer(t, &resp, &test.out[i])
			}
		})
	}
}

func TestGetConsumer(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		in   int
		out  *domain.Consumer
	}{
		{
			"get_consumer_test",
			1,
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

			database, teardown := dbtest.NewUnit(t, c, test.name)
			t.Cleanup(teardown)

			s := consumerstorage.NewConsumerStorage(consumerstorage.Params{DB: database})

			resp, err := s.InsertConsumer(*test.out)
			require.NoError(t, err)
			require.NotNil(t, resp)

			resp2, err := s.GetConsumerByID(test.in)
			require.NoError(t, err)
			require.NotNil(t, resp2)

			equalConsumer(t, resp2, test.out)
		})
	}
}

func TestInsertConsumerLocation(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		in   domain.Location
		out  *domain.Location
	}{
		{
			"insert_consumer_location_test",
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

			database, teardown := dbtest.NewUnit(t, c, test.name)
			t.Cleanup(teardown)

			s := consumerstorage.NewConsumerStorage(consumerstorage.Params{DB: database})

			resp, err := s.InsertLocation(test.in)
			require.NoError(t, err)
			require.NotNil(t, resp)

			assert.Equal(t, resp, test.out)
		})
	}
}

func TestUpdateConsumerLocation(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		in       domain.Location
		updateIn domain.Location
		out      *domain.Location
	}{
		{
			"update_consumer_location_test",
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
			domain.Location{
				UserID:     1,
				Latitude:   "01234567892",
				Longitude:  "01234567892",
				Country:    "TestCountry2",
				City:       "Test City2",
				Region:     "TestRegion2",
				Street:     "TestStreet2",
				HomeNumber: "TestHomeNumber2",
				Floor:      "TestFloor2",
				Door:       "TestDoor2",
			},
			&domain.Location{
				UserID:     1,
				Latitude:   "01234567892",
				Longitude:  "01234567892",
				Country:    "TestCountry2",
				City:       "Test City2",
				Region:     "TestRegion2",
				Street:     "TestStreet2",
				HomeNumber: "TestHomeNumber2",
				Floor:      "TestFloor2",
				Door:       "TestDoor2",
			},
		},
	}

	for _, currentTest := range tests {
		test := currentTest

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			database, teardown := dbtest.NewUnit(t, c, test.name)
			t.Cleanup(teardown)

			s := consumerstorage.NewConsumerStorage(consumerstorage.Params{DB: database})

			resp, err := s.InsertLocation(test.in)
			require.NoError(t, err)
			require.NotNil(t, resp)

			resp2, err := s.UpdateLocation(test.updateIn)
			require.NoError(t, err)
			require.NotNil(t, resp)

			assert.Equal(t, resp2, test.out)
		})
	}
}

func TestGetConsumerLocation(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		in   int
		out  *domain.Location
	}{
		{
			"get_consumer_location_test",
			1,
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

			database, teardown := dbtest.NewUnit(t, c, test.name)
			t.Cleanup(teardown)

			s := consumerstorage.NewConsumerStorage(consumerstorage.Params{DB: database})

			mockData := *test.out
			resp, err := s.InsertLocation(mockData)
			require.NoError(t, err)
			require.NotNil(t, resp)

			resp2, err := s.GetLocation(test.in)
			require.NoError(t, err)
			require.NotNil(t, resp)

			assert.Equal(t, resp2, test.out)
		})
	}
}
