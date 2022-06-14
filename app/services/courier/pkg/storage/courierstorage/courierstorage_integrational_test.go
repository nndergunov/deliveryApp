package courierstorage_test

import (
	"database/sql"
	"github.com/nndergunov/deliveryApp/app/pkg/configreader"
	"os"
	"strconv"
	"strings"
	"testing"

	"courier/pkg/db"
	"courier/pkg/domain"
	"courier/pkg/storage/courierstorage"
)

const configFile = "/config.yaml"

func TestInsertCourier(t *testing.T) {
	tests := []struct {
		name    string
		courier domain.Courier
	}{
		{
			name: "Test Insert Courier",
			courier: domain.Courier{
				Username:  "vasyauser",
				Firstname: "vasya",
				Lastname:  "",
				Email:     "vasya@gmail.com",
				Phone:     "123456789",
				Available: false,
			},
		},
	}

	for _, currentTest := range tests {
		test := currentTest

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			line, err := os.Getwd()
			if err != nil {
				t.Fatal(err)

			}
			confPath := strings.TrimSuffix(line, "\\pkg\\storage\\courierstorage")

			err = configreader.SetConfigFile(confPath + configFile)
			if err != nil {
				t.Fatal(err)
			}

			database, err := db.OpenDB("postgres", configreader.GetString("DB.test"))
			if err != nil {
				t.Fatal(err)
			}

			courierStorage := courierstorage.NewCourierStorage(courierstorage.Params{DB: database})
			if err != nil {
				t.Fatal(err)
			}

			insertedCourier, err := courierStorage.InsertCourier(test.courier)
			if err != nil {
				t.Fatal(err)
			}

			if insertedCourier == nil {
				t.Errorf("createCourier: Expected: %s, Got: %s", "not nil", "nil")
			}

			if insertedCourier.Username != test.courier.Username {
				t.Errorf("Username: Expected: %s, Got: %s", test.courier.Username, insertedCourier.Username)
			}

			if err = courierStorage.DeleteCourier(insertedCourier.ID); err != nil {
				t.Error(err)
			}

			if err := database.Close(); err != nil {
				t.Error(err)
			}
		})
	}
}

func TestDeleteCourier(t *testing.T) {
	tests := []struct {
		name    string
		courier domain.Courier
	}{
		{
			name: "Test Delete Courier",
			courier: domain.Courier{
				Username:  "vasyauser",
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

			line, err := os.Getwd()
			if err != nil {
				t.Fatal(err)

			}
			confPath := strings.TrimSuffix(line, "\\pkg\\storage\\courierstorage")

			err = configreader.SetConfigFile(confPath + configFile)
			if err != nil {
				t.Fatal(err)
			}

			database, err := db.OpenDB("postgres", configreader.GetString("DB.test"))
			if err != nil {
				t.Fatal(err)
			}

			couierStorage := courierstorage.NewCourierStorage(courierstorage.Params{DB: database})

			defer couierStorage.CleanCourierTable()

			insertedCourier, err := couierStorage.InsertCourier(test.courier)
			if err != nil {
				t.Fatal(err)
			}

			if insertedCourier == nil {
				t.Errorf("createCourier: Expected: %s, Got: %s", "not nill", "nil")
			}

			err = couierStorage.DeleteCourier(insertedCourier.ID)
			if err != nil {
				t.Fatal(err)
			}

			foundCourier, err := couierStorage.GetCourierByID(insertedCourier.ID)
			if err != nil && err != sql.ErrNoRows {
				t.Fatal(err)
			}

			if foundCourier != nil {
				t.Errorf("deletedCourier: Expected: %s, Got: %s", "nil", "not .+nil")
			}

			if err := database.Close(); err != nil {
				t.Error(err)
			}
		})
	}
}

func TestUpdateCourier(t *testing.T) {
	tests := []struct {
		name           string
		initialCourier domain.Courier
		updateCourier  domain.Courier
	}{
		{
			name: "Test Update Courier",
			initialCourier: domain.Courier{
				Username:  "vasyauser",
				Firstname: "vasya",
				Lastname:  "",
				Email:     "vasya@gmail.com",
				Phone:     "123456789",
			},
			updateCourier: domain.Courier{
				Username:  "updatedvasyauser",
				Firstname: "updatedvasya",
				Lastname:  "vasyavov",
				Email:     "updatedvasya@gmail.com",
				Phone:     "123456789",
			},
		},
	}

	for _, currentTest := range tests {
		test := currentTest

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			line, err := os.Getwd()
			if err != nil {
				t.Fatal(err)

			}
			confPath := strings.TrimSuffix(line, "\\pkg\\storage\\courierstorage")

			err = configreader.SetConfigFile(confPath + configFile)
			if err != nil {
				t.Fatal(err)
			}

			database, err := db.OpenDB("postgres", configreader.GetString("DB.test"))
			if err != nil {
				t.Fatal(err)
			}

			courierStorage := courierstorage.NewCourierStorage(courierstorage.Params{DB: database})

			insertedCourier, err := courierStorage.InsertCourier(test.initialCourier)
			if err != nil {
				t.Fatal(err)
			}

			if insertedCourier == nil {
				t.Errorf("insertedCourier: Expected: %s, Got: %s", "not nill", "nil")
			}

			test.updateCourier.ID = insertedCourier.ID
			updatedCourier, err := courierStorage.UpdateCourier(test.updateCourier)
			if err != nil {
				t.Fatal(err)
			}

			if updatedCourier.Username != test.updateCourier.Username {
				t.Errorf("updatedCourier username: Expected: %s, Got: %s", test.updateCourier.Username, updatedCourier.Username)
			}

			if updatedCourier.Firstname != test.updateCourier.Firstname {
				t.Errorf("updatedCourier Firstname: Expected: %s, Got: %s", test.updateCourier.Firstname, updatedCourier.Firstname)
			}

			if updatedCourier.Lastname != test.updateCourier.Lastname {
				t.Errorf("updatedCourier Lastname: Expected: %s, Got: %s", test.updateCourier.Lastname, updatedCourier.Lastname)
			}

			if updatedCourier.Email != test.updateCourier.Email {
				t.Errorf("updatedCourier Email: Expected: %s, Got: %s", test.updateCourier.Email, updatedCourier.Email)
			}

			if err = courierStorage.DeleteCourier(insertedCourier.ID); err != nil {
				t.Error(err)
			}

			if err := database.Close(); err != nil {
				t.Error(err)
			}
		})
	}
}

func TestUpdateCourierAvailable(t *testing.T) {
	tests := []struct {
		name           string
		initialCourier domain.Courier
	}{
		{
			name: "Test update available Courier",
			initialCourier: domain.Courier{
				Username:  "vasyauser",
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

			line, err := os.Getwd()
			if err != nil {
				t.Fatal(err)

			}
			confPath := strings.TrimSuffix(line, "\\pkg\\storage\\courierstorage")

			err = configreader.SetConfigFile(confPath + configFile)
			if err != nil {
				t.Fatal(err)
			}
			database, err := db.OpenDB("postgres", configreader.GetString("DB.test"))
			if err != nil {
				t.Fatal(err)
			}

			courierStorage := courierstorage.NewCourierStorage(courierstorage.Params{DB: database})

			insertedCourier, err := courierStorage.InsertCourier(test.initialCourier)
			if err != nil {
				t.Fatal(err)
			}

			if insertedCourier == nil {
				t.Errorf("insertedCourier: Expected: %s, Got: %s", "not nill", "nil")
			}

			updatedCourierAvailable, err := courierStorage.UpdateCourierAvailable(insertedCourier.ID, !insertedCourier.Available)
			if err != nil {
				t.Fatal(err)
			}

			if updatedCourierAvailable.Available == insertedCourier.Available {
				t.Errorf("updated Courier Available: Expected: %v, Got: %v", updatedCourierAvailable.Available, insertedCourier.Available)
			}

			if err = courierStorage.DeleteCourier(insertedCourier.ID); err != nil {
				t.Error(err)
			}

			if err := database.Close(); err != nil {
				t.Error(err)
			}
		})
	}
}

func TestGetAllCourier(t *testing.T) {
	tests := []struct {
		name               string
		initialCourierList []domain.Courier
	}{
		{
			name: "Test Get ALl Courier",
			initialCourierList: []domain.Courier{domain.Courier{
				Username:  "vasyauser",
				Firstname: "vasya",
				Lastname:  "",
				Email:     "vasya@gmail.com",
				Phone:     "123456789",
				Available: true,
			},
			},
		},
	}

	for _, currentTest := range tests {
		test := currentTest

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			line, err := os.Getwd()
			if err != nil {
				t.Fatal(err)

			}
			confPath := strings.TrimSuffix(line, "\\pkg\\storage\\courierstorage")

			err = configreader.SetConfigFile(confPath + configFile)
			if err != nil {
				t.Fatal(err)
			}

			database, err := db.OpenDB("postgres", configreader.GetString("DB.test"))
			if err != nil {
				t.Fatal(err)
			}

			courierStorage := courierstorage.NewCourierStorage(courierstorage.Params{DB: database})

			for _, initialCourier := range test.initialCourierList {
				insertedCourier, err := courierStorage.InsertCourier(initialCourier)
				if err != nil {
					t.Fatal(err)
				}
				if insertedCourier == nil {
					t.Errorf("insertedCourier: Expected: %s, Got: %s", "not nill", "nil")
				}

			}
			param := domain.SearchParam{}
			param["available"] = "true"

			allCourier, err := courierStorage.GetAllCourier(param)
			if err != nil {
				t.Fatal(err)
			}

			if len(allCourier) != len(test.initialCourierList) {
				t.Errorf("get all coureir len: Expected: %v, Got: %v", len(test.initialCourierList), len(allCourier))

			}

			if err := courierStorage.CleanCourierTable(); err != nil {
				t.Error(err)
			}

			if err := database.Close(); err != nil {
				t.Error(err)
			}
		})
	}
}

func TestGetCourierByID(t *testing.T) {
	tests := []struct {
		name           string
		initialCourier domain.Courier
	}{
		{
			name: "Test Get Courier",
			initialCourier: domain.Courier{
				Username:  "vasyauser",
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

			line, err := os.Getwd()
			if err != nil {
				t.Fatal(err)

			}
			confPath := strings.TrimSuffix(line, "\\pkg\\storage\\courierstorage")

			err = configreader.SetConfigFile(confPath + configFile)
			if err != nil {
				t.Fatal(err)
			}

			database, err := db.OpenDB("postgres", configreader.GetString("DB.test"))
			if err != nil {
				t.Fatal(err)
			}

			courierStorage := courierstorage.NewCourierStorage(courierstorage.Params{DB: database})

			insertedCourier, err := courierStorage.InsertCourier(test.initialCourier)
			if err != nil {
				t.Fatal(err)
			}
			if insertedCourier == nil {
				t.Errorf("insertedCourier: Expected: %s, Got: %s", "not nill", "nil")
			}

			gotCourier, err := courierStorage.GetCourierByID(insertedCourier.ID)
			if err != nil {
				t.Fatal(err)
			}

			if gotCourier.Username != test.initialCourier.Username {
				t.Errorf("get courier username: Expected: %s, Got: %s", test.initialCourier.Username, gotCourier.Username)
			}

			if gotCourier.Firstname != test.initialCourier.Firstname {
				t.Errorf("get courier Firstname: Expected: %s, Got: %s", test.initialCourier.Firstname, gotCourier.Firstname)
			}

			if gotCourier.Lastname != test.initialCourier.Lastname {
				t.Errorf("get courier Lastname: Expected: %s, Got: %s", test.initialCourier.Lastname, gotCourier.Lastname)
			}

			if gotCourier.Phone != test.initialCourier.Phone {
				t.Errorf("get courier Phone: Expected: %s, Got: %s", test.initialCourier.Phone, gotCourier.Phone)
			}

			if gotCourier.Email != test.initialCourier.Email {
				t.Errorf("get courier Email: Expected: %s, Got: %s", test.initialCourier.Email, gotCourier.Email)
			}

			if err = courierStorage.DeleteCourier(insertedCourier.ID); err != nil {
				t.Error(err)
			}

			if err := database.Close(); err != nil {
				t.Error(err)
			}
		})
	}
}

func TestGetCourierDuplicateByParam(t *testing.T) {
	tests := []struct {
		name           string
		initialCourier domain.Courier
	}{
		{
			name: "Test GetCourierDuplicateByParam",
			initialCourier: domain.Courier{
				Username:  "vasyauser",
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

			line, err := os.Getwd()
			if err != nil {
				t.Fatal(err)

			}
			confPath := strings.TrimSuffix(line, "\\pkg\\storage\\courierstorage")

			err = configreader.SetConfigFile(confPath + configFile)
			if err != nil {
				t.Fatal(err)
			}

			database, err := db.OpenDB("postgres", configreader.GetString("DB.test"))
			if err != nil {
				t.Fatal(err)
			}

			courierStorage := courierstorage.NewCourierStorage(courierstorage.Params{DB: database})

			insertedCourier, err := courierStorage.InsertCourier(test.initialCourier)
			if err != nil {
				t.Fatal(err)
			}
			if insertedCourier == nil {
				t.Errorf("insertedCourier: Expected: %s, Got: %s", "not nill", "nil")
			}

			param := domain.SearchParam{}

			param["username"] = test.initialCourier.Username
			param["id"] = strconv.Itoa(int(insertedCourier.ID))

			gotCourier, err := courierStorage.GetCourierDuplicateByParam(param)
			if err != nil && err != sql.ErrNoRows {
				t.Fatal(err)
			}

			if gotCourier != nil {
				t.Errorf("gotCourier: Expected: %s, Got: %s", "nil", "not nil")
			}

			if err = courierStorage.DeleteCourier(insertedCourier.ID); err != nil {
				t.Error(err)
			}

			if err := database.Close(); err != nil {
				t.Error(err)
			}
		})
	}
}

func TestInsertLocation(t *testing.T) {
	tests := []struct {
		name     string
		location domain.Location
	}{
		{
			name: "TestInsertLocation",
			location: domain.Location{
				UserID:     1,
				Latitude:   "0123456789",
				Longitude:  "0123456789",
				Country:    "TestCountry",
				City:       "TestCity",
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

			line, err := os.Getwd()
			if err != nil {
				t.Fatal(err)

			}
			confPath := strings.TrimSuffix(line, "\\pkg\\storage\\courierstorage")

			err = configreader.SetConfigFile(confPath + configFile)
			if err != nil {
				t.Fatal(err)
			}

			database, err := db.OpenDB("postgres", configreader.GetString("DB.test"))
			if err != nil {
				t.Fatal(err)
			}

			courierStorage := courierstorage.NewCourierStorage(courierstorage.Params{DB: database})
			if err != nil {
				t.Fatal(err)
			}

			insertedLocation, err := courierStorage.InsertLocation(test.location)
			if err != nil {
				t.Fatal(err)
			}

			if insertedLocation == nil {
				t.Errorf("insertedLocation: Expected: %s, Got: %s", "not nill", "nil")
			}

			if insertedLocation == nil {
				t.Errorf("updatedLocation: Expected: %s, Got: %s", "not nill", "nil")
			}

			if insertedLocation.UserID != test.location.UserID {
				t.Errorf("insertedLocation UserID: Expected: %v, Got: %v", test.location.UserID, insertedLocation.UserID)
			}

			if insertedLocation.Latitude != test.location.Latitude {
				t.Errorf("insertedLocation Latitude: Expected: %v, Got: %v", test.location.Latitude, insertedLocation.Latitude)
			}

			if insertedLocation.Longitude != test.location.Longitude {
				t.Errorf("insertedLocation Longitude: Expected: %v, Got: %v", test.location.Longitude, insertedLocation.Longitude)
			}

			if insertedLocation.Country != test.location.Country {
				t.Errorf("insertedLocation Country: Expected: %v, Got: %v", test.location.Country, insertedLocation.Country)
			}

			if insertedLocation.City != test.location.City {
				t.Errorf("insertedLocation City: Expected: %v, Got: %v", test.location.City, insertedLocation.City)
			}

			if insertedLocation.Region != test.location.Region {
				t.Errorf("insertedLocation Region: Expected: %v, Got: %v", test.location.Region, insertedLocation.Region)
			}

			if err = courierStorage.DeleteLocation(insertedLocation.UserID); err != nil {
				t.Error(err)
			}

			if err := database.Close(); err != nil {
				t.Error(err)
			}
		})
	}
}

func TestUpdateLocation(t *testing.T) {
	tests := []struct {
		name            string
		initialLocation domain.Location
		updateLocation  domain.Location
	}{
		{
			name: "Test Update Courier",
			initialLocation: domain.Location{
				UserID:     1,
				Latitude:   "0123456789",
				Longitude:  "0123456789",
				Country:    "TestCountry",
				City:       "TestCity",
				Region:     "",
				Street:     "",
				HomeNumber: "",
				Floor:      "",
				Door:       "",
			},

			updateLocation: domain.Location{
				UserID:     1,
				Latitude:   "9876543210",
				Longitude:  "9876543210",
				Country:    "CountryTest",
				City:       "CityTest",
				Region:     "TestRegion",
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

			line, err := os.Getwd()
			if err != nil {
				t.Fatal(err)

			}
			confPath := strings.TrimSuffix(line, "\\pkg\\storage\\courierstorage")

			err = configreader.SetConfigFile(confPath + configFile)
			if err != nil {
				t.Fatal(err)
			}

			database, err := db.OpenDB("postgres", configreader.GetString("DB.test"))
			if err != nil {
				t.Fatal(err)
			}

			courierStorage := courierstorage.NewCourierStorage(courierstorage.Params{DB: database})
			if err != nil {
				t.Fatal(err)
			}

			insertedLocation, err := courierStorage.InsertLocation(test.initialLocation)
			if err != nil {
				t.Fatal(err)
			}

			if insertedLocation == nil {
				t.Errorf("insertedLocation: Expected: %s, Got: %s", "not nill", "nil")
			}

			updatedLocation, err := courierStorage.UpdateLocation(test.updateLocation)
			if err != nil {
				t.Fatal(err)
			}

			if updatedLocation == nil {
				t.Errorf("updatedLocation: Expected: %s, Got: %s", "not nill", "nil")
			}

			if updatedLocation.UserID != test.updateLocation.UserID {
				t.Errorf("updatedLocation UserID: Expected: %v, Got: %v", test.updateLocation.UserID, updatedLocation.UserID)
			}

			if updatedLocation.Latitude != test.updateLocation.Latitude {
				t.Errorf("updatedLocation Latitude: Expected: %v, Got: %v", test.updateLocation.Latitude, updatedLocation.Latitude)
			}

			if updatedLocation.Longitude != test.updateLocation.Longitude {
				t.Errorf("updatedLocation Longitude: Expected: %v, Got: %v", test.updateLocation.Longitude, updatedLocation.Longitude)
			}

			if updatedLocation.Country != test.updateLocation.Country {
				t.Errorf("updatedLocation Country: Expected: %v, Got: %v", test.updateLocation.Country, updatedLocation.Country)
			}

			if updatedLocation.City != test.updateLocation.City {
				t.Errorf("updatedLocation City: Expected: %v, Got: %v", test.updateLocation.City, updatedLocation.City)
			}

			if updatedLocation.Region != test.updateLocation.Region {
				t.Errorf("updatedLocation Region: Expected: %v, Got: %v", test.updateLocation.Region, updatedLocation.Region)
			}

			if err = courierStorage.DeleteLocation(insertedLocation.UserID); err != nil {
				t.Error(err)
			}

			if err := database.Close(); err != nil {
				t.Error(err)
			}
		})
	}
}

func TestGetLocation(t *testing.T) {
	tests := []struct {
		name     string
		location domain.Location
	}{
		{
			name: "TestGetLocation",
			location: domain.Location{
				UserID:     1,
				Latitude:   "0123456789",
				Longitude:  "0123456789",
				Country:    "TestCountry",
				City:       "TestCity",
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

			line, err := os.Getwd()
			if err != nil {
				t.Fatal(err)

			}
			confPath := strings.TrimSuffix(line, "\\pkg\\storage\\courierstorage")

			err = configreader.SetConfigFile(confPath + configFile)
			if err != nil {
				t.Fatal(err)
			}

			database, err := db.OpenDB("postgres", configreader.GetString("DB.test"))
			if err != nil {
				t.Fatal(err)
			}

			courierStorage := courierstorage.NewCourierStorage(courierstorage.Params{DB: database})
			if err != nil {
				t.Fatal(err)
			}

			insertedLocation, err := courierStorage.InsertLocation(test.location)
			if err != nil {
				t.Fatal(err)
			}

			if insertedLocation == nil {
				t.Errorf("insertedLocation: Expected: %s, Got: %s", "not nill", "nil")
			}

			getLocation, err := courierStorage.GetLocation(insertedLocation.UserID)
			if err != nil {
				t.Fatal(err)
			}

			if getLocation == nil {
				t.Errorf("getLocation: Expected: %s, Got: %s", "not nill", "nil")
			}

			if getLocation.UserID != test.location.UserID {
				t.Errorf("getLocation UserID: Expected: %v, Got: %v", test.location.UserID, getLocation.UserID)
			}

			if getLocation.Latitude != test.location.Latitude {
				t.Errorf("getLocation Latitude: Expected: %v, Got: %v", test.location.Latitude, getLocation.Latitude)
			}

			if getLocation.Longitude != test.location.Longitude {
				t.Errorf("getLocation Longitude: Expected: %v, Got: %v", test.location.Longitude, getLocation.Longitude)
			}

			if getLocation.Country != test.location.Country {
				t.Errorf("getLocation Country: Expected: %v, Got: %v", test.location.Country, getLocation.Country)
			}

			if getLocation.City != test.location.City {
				t.Errorf("getLocation City: Expected: %v, Got: %v", test.location.City, getLocation.City)
			}

			if getLocation.Region != test.location.Region {
				t.Errorf("getLocation Region: Expected: %v, Got: %v", test.location.Region, getLocation.Region)
			}

			if err = courierStorage.DeleteLocation(insertedLocation.UserID); err != nil {
				t.Error(err)
			}

			if err := database.Close(); err != nil {
				t.Error(err)
			}
		})
	}
}
