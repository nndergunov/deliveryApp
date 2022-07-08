package courierstorage_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/nndergunov/deliveryApp/app/services/courier/pkg/db/dbtest"
	"github.com/nndergunov/deliveryApp/app/services/courier/pkg/docker"
	"github.com/nndergunov/deliveryApp/app/services/courier/pkg/domain"
	"github.com/nndergunov/deliveryApp/app/services/courier/pkg/storage/courierstorage"
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

// equalCourier - function to compare selected fields from struct. Fields which needed
func equalCourier(t *testing.T, get *domain.Courier, want *domain.Courier) {
	if get.Username != want.Username {
		t.Errorf("Username: Expected: %s, Got: %s", want.Username, get.Username)
	}

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

	if get.Available != want.Available {
		t.Errorf("Available: Expected: %v, Got: %v", want.Available, get.Available)
	}
}

func TestInsertCourier(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		in   domain.Courier
		out  *domain.Courier
	}{
		{
			"insert_courier_test",
			domain.Courier{
				Username:  "testUsername",
				Firstname: "testFName",
				Lastname:  "testLName",
				Email:     "test@gmail.com",
				Phone:     "123456789",
				Available: true,
			},
			&domain.Courier{
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

			database, teardown := dbtest.NewUnit(t, c, test.name)
			t.Cleanup(teardown)

			s := courierstorage.NewCourierStorage(courierstorage.Params{DB: database})

			resp, err := s.InsertCourier(test.in)
			require.NoError(t, err)
			require.NotNil(t, resp)

			equalCourier(t, resp, test.out)
		})
	}
}

func TestDeleteCourier(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		in   domain.Courier
		out  error
	}{
		{
			"delete_courier_test",
			domain.Courier{
				Username:  "testUsername",
				Firstname: "testFName",
				Lastname:  "testLName",
				Email:     "test@gmail.com",
				Phone:     "123456789",
				Available: true,
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

			s := courierstorage.NewCourierStorage(courierstorage.Params{DB: database})

			resp, err := s.InsertCourier(test.in)
			require.NoError(t, err)
			require.NotNil(t, resp)

			err = s.DeleteCourier(test.in.ID)
			if err != test.out {
				t.Errorf("DeleteCourier: Expected: %v, Got: %v", test.out, err)
			}
		})
	}
}

func TestUpdateCourier(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		in       domain.Courier
		updateIn domain.Courier
		out      *domain.Courier
	}{
		{
			"update_courier_test",
			domain.Courier{
				Username:  "testUsername",
				Firstname: "testFName",
				Lastname:  "testLName",
				Email:     "test@gmail.com",
				Phone:     "123456789",
				Available: true,
			},
			domain.Courier{
				ID:        1,
				Username:  "testUsername1",
				Firstname: "testFName1",
				Lastname:  "testLName1",
				Email:     "test1@gmail.com",
				Phone:     "1234567891",
				Available: true,
			},
			&domain.Courier{
				ID:        1,
				Username:  "testUsername1",
				Firstname: "testFName1",
				Lastname:  "testLName1",
				Email:     "test1@gmail.com",
				Phone:     "1234567891",
				Available: true,
			},
		},
	}

	for _, currentTest := range tests {
		test := currentTest

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			database, teardown := dbtest.NewUnit(t, c, test.name)
			t.Cleanup(teardown)

			s := courierstorage.NewCourierStorage(courierstorage.Params{DB: database})

			resp, err := s.InsertCourier(test.in)
			require.NoError(t, err)
			require.NotNil(t, resp)

			resp2, err := s.UpdateCourier(test.updateIn)
			require.NoError(t, err)
			require.NotNil(t, resp2)

			equalCourier(t, resp2, test.out)
		})
	}
}

func TestUpdateCourierAvailable(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		inID        int
		InAvailable bool
		out         *domain.Courier
	}{
		{
			"update_courier_available_test",
			1,
			true,
			&domain.Courier{
				ID:        1,
				Username:  "testUsername1",
				Firstname: "testFName1",
				Lastname:  "testLName1",
				Email:     "test1@gmail.com",
				Phone:     "1234567891",
				Available: true,
			},
		},
	}

	for _, currentTest := range tests {
		test := currentTest

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			database, teardown := dbtest.NewUnit(t, c, test.name)
			t.Cleanup(teardown)

			s := courierstorage.NewCourierStorage(courierstorage.Params{DB: database})

			resp, err := s.InsertCourier(*test.out)
			require.NoError(t, err)
			require.NotNil(t, resp)

			resp2, err := s.UpdateCourierAvailable(test.inID, test.InAvailable)
			require.NoError(t, err)
			require.NotNil(t, resp2)

			equalCourier(t, resp2, test.out)
		})
	}
}

func TestGetAllCourier(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		out  []domain.Courier
	}{
		{
			"get_all_courier_test",
			[]domain.Courier{
				{
					ID:        1,
					Username:  "testUsername1",
					Firstname: "testFName1",
					Lastname:  "testLName1",
					Email:     "test1@gmail.com",
					Phone:     "1234567891",
					Available: true,
				},
				{
					ID:        2,
					Username:  "testUsername2",
					Firstname: "testFName2",
					Lastname:  "testLName2",
					Email:     "test2@gmail.com",
					Phone:     "1234567892",
					Available: true,
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

			s := courierstorage.NewCourierStorage(courierstorage.Params{DB: database})

			for _, data := range test.out {
				resp, err := s.InsertCourier(data)
				require.NoError(t, err)
				require.NotNil(t, resp)
			}

			respList, err := s.GetCourierList(domain.SearchParam{})
			require.NoError(t, err)
			require.NotNil(t, respList)
			for i, resp := range respList {
				equalCourier(t, &resp, &test.out[i])
			}
		})
	}
}

func TestGetCourier(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		in   int
		out  *domain.Courier
	}{
		{
			"get_courier_test",
			1,
			&domain.Courier{
				ID:        1,
				Username:  "testUsername1",
				Firstname: "testFName1",
				Lastname:  "testLName1",
				Email:     "test1@gmail.com",
				Phone:     "1234567891",
				Available: true,
			},
		},
	}

	for _, currentTest := range tests {
		test := currentTest

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			database, teardown := dbtest.NewUnit(t, c, test.name)
			t.Cleanup(teardown)

			s := courierstorage.NewCourierStorage(courierstorage.Params{DB: database})

			resp, err := s.InsertCourier(*test.out)
			require.NoError(t, err)
			require.NotNil(t, resp)

			resp2, err := s.GetCourierByID(test.in)
			require.NoError(t, err)
			require.NotNil(t, resp2)

			equalCourier(t, resp2, test.out)
		})
	}
}

func TestInsertCourierLocation(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		in   domain.Location
		out  *domain.Location
	}{
		{
			"insert_courier_location_test",
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

			s := courierstorage.NewCourierStorage(courierstorage.Params{DB: database})

			resp, err := s.InsertLocation(test.in)
			require.NoError(t, err)
			require.NotNil(t, resp)

			assert.Equal(t, resp, test.out)
		})
	}
}

func TestUpdateCourierLocation(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		in       domain.Location
		updateIn domain.Location
		out      *domain.Location
	}{
		{
			"update_courier_location_test",
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

			s := courierstorage.NewCourierStorage(courierstorage.Params{DB: database})

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

func TestGetCourierLocation(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		in   int
		out  *domain.Location
	}{
		{
			"get_courier_location_test",
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

			s := courierstorage.NewCourierStorage(courierstorage.Params{DB: database})

			mockData := *test.out
			resp, err := s.InsertLocation(mockData)
			require.NoError(t, err)
			require.NotNil(t, resp)

			resp2, err := s.GetLocation(test.in)
			require.NoError(t, err)
			require.NotNil(t, resp2)

			assert.Equal(t, resp2, test.out)
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
			"get_courier_location_list_test",
			domain.SearchParam{
				"city": "testCity",
			},
			[]domain.Location{
				{
					UserID:     1,
					Latitude:   "0123456789",
					Longitude:  "0123456789",
					Country:    "TestCountry",
					City:       "testCity",
					Region:     "TestRegion",
					Street:     "TestStreet",
					HomeNumber: "TestHomeNumber",
					Floor:      "TestFloor",
					Door:       "TestDoor",
				},
				{
					UserID:     2,
					Latitude:   "01234567891",
					Longitude:  "01234567891",
					Country:    "TestCountry1",
					City:       "testCity",
					Region:     "TestRegion1",
					Street:     "TestStreet1",
					HomeNumber: "TestHomeNumber1",
					Floor:      "TestFloor1",
					Door:       "TestDoor1",
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

			s := courierstorage.NewCourierStorage(courierstorage.Params{DB: database})

			for _, mockInData := range test.out {
				resp, err := s.InsertLocation(mockInData)
				require.NoError(t, err)
				require.NotNil(t, resp)
			}
			resp, err := s.GetLocationList(test.in)
			require.NoError(t, err)
			require.NotNil(t, resp)

			assert.Equal(t, resp, test.out)
		})
	}
}
