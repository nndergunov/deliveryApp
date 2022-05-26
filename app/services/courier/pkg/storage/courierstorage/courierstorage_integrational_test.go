package courierstorage_test

import (
	"courier/pkg/db"
	"courier/pkg/domain"
	"courier/pkg/storage/courierstorage"
	"database/sql"
	"github.com/nndergunov/deliveryApp/app/pkg/configreader"
	"os"
	"strconv"
	"strings"
	"testing"
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

func TestInsertCourierLocation(t *testing.T) {
	tests := []struct {
		name            string
		courierLocation domain.CourierLocation
	}{
		{
			name: "TestInsertCourierLocation",
			courierLocation: domain.CourierLocation{
				CourierID:  1,
				Altitude:   "0123456789",
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

			insertedCourierLocation, err := courierStorage.InsertCourierLocation(test.courierLocation)
			if err != nil {
				t.Fatal(err)
			}

			if insertedCourierLocation == nil {
				t.Errorf("insertedCourierLocation: Expected: %s, Got: %s", "not nill", "nil")
			}

			if insertedCourierLocation == nil {
				t.Errorf("updatedCourierLocation: Expected: %s, Got: %s", "not nill", "nil")
			}

			if insertedCourierLocation.CourierID != test.courierLocation.CourierID {
				t.Errorf("insertedCourierLocation CourierID: Expected: %v, Got: %v", test.courierLocation.CourierID, insertedCourierLocation.CourierID)
			}

			if insertedCourierLocation.Altitude != test.courierLocation.Altitude {
				t.Errorf("insertedCourierLocation Altitude: Expected: %v, Got: %v", test.courierLocation.Altitude, insertedCourierLocation.Altitude)
			}

			if insertedCourierLocation.Longitude != test.courierLocation.Longitude {
				t.Errorf("insertedCourierLocation Longitude: Expected: %v, Got: %v", test.courierLocation.Longitude, insertedCourierLocation.Longitude)
			}

			if insertedCourierLocation.Country != test.courierLocation.Country {
				t.Errorf("insertedCourierLocation Country: Expected: %v, Got: %v", test.courierLocation.Country, insertedCourierLocation.Country)
			}

			if insertedCourierLocation.City != test.courierLocation.City {
				t.Errorf("insertedCourierLocation City: Expected: %v, Got: %v", test.courierLocation.City, insertedCourierLocation.City)
			}

			if insertedCourierLocation.Region != test.courierLocation.Region {
				t.Errorf("insertedCourierLocation Region: Expected: %v, Got: %v", test.courierLocation.Region, insertedCourierLocation.Region)
			}

			if err = courierStorage.DeleteCourierLocation(insertedCourierLocation.CourierID); err != nil {
				t.Error(err)
			}

			if err := database.Close(); err != nil {
				t.Error(err)
			}
		})
	}
}

func TestUpdateCourierLocation(t *testing.T) {
	tests := []struct {
		name                   string
		initialCourierLocation domain.CourierLocation
		updateCourierLocation  domain.CourierLocation
	}{
		{
			name: "Test Update Courier",
			initialCourierLocation: domain.CourierLocation{
				CourierID:  1,
				Altitude:   "0123456789",
				Longitude:  "0123456789",
				Country:    "TestCountry",
				City:       "TestCity",
				Region:     "",
				Street:     "",
				HomeNumber: "",
				Floor:      "",
				Door:       "",
			},

			updateCourierLocation: domain.CourierLocation{
				CourierID:  1,
				Altitude:   "9876543210",
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

			insertedCourierLocation, err := courierStorage.InsertCourierLocation(test.initialCourierLocation)
			if err != nil {
				t.Fatal(err)
			}

			if insertedCourierLocation == nil {
				t.Errorf("insertedCourierLocation: Expected: %s, Got: %s", "not nill", "nil")
			}

			updatedCourierLocation, err := courierStorage.UpdateCourierLocation(test.updateCourierLocation)
			if err != nil {
				t.Fatal(err)
			}

			if updatedCourierLocation == nil {
				t.Errorf("updatedCourierLocation: Expected: %s, Got: %s", "not nill", "nil")
			}

			if updatedCourierLocation.CourierID != test.updateCourierLocation.CourierID {
				t.Errorf("updatedCourierLocation CourierID: Expected: %v, Got: %v", test.updateCourierLocation.CourierID, updatedCourierLocation.CourierID)
			}

			if updatedCourierLocation.Altitude != test.updateCourierLocation.Altitude {
				t.Errorf("updatedCourierLocation Altitude: Expected: %v, Got: %v", test.updateCourierLocation.Altitude, updatedCourierLocation.Altitude)
			}

			if updatedCourierLocation.Longitude != test.updateCourierLocation.Longitude {
				t.Errorf("updatedCourierLocation Longitude: Expected: %v, Got: %v", test.updateCourierLocation.Longitude, updatedCourierLocation.Longitude)
			}

			if updatedCourierLocation.Country != test.updateCourierLocation.Country {
				t.Errorf("updatedCourierLocation Country: Expected: %v, Got: %v", test.updateCourierLocation.Country, updatedCourierLocation.Country)
			}

			if updatedCourierLocation.City != test.updateCourierLocation.City {
				t.Errorf("updatedCourierLocation City: Expected: %v, Got: %v", test.updateCourierLocation.City, updatedCourierLocation.City)
			}

			if updatedCourierLocation.Region != test.updateCourierLocation.Region {
				t.Errorf("updatedCourierLocation Region: Expected: %v, Got: %v", test.updateCourierLocation.Region, updatedCourierLocation.Region)
			}

			if err = courierStorage.DeleteCourierLocation(insertedCourierLocation.CourierID); err != nil {
				t.Error(err)
			}

			if err := database.Close(); err != nil {
				t.Error(err)
			}
		})
	}
}

func TestGetCourierLocation(t *testing.T) {
	tests := []struct {
		name            string
		courierLocation domain.CourierLocation
	}{
		{
			name: "TestGetCourierLocation",
			courierLocation: domain.CourierLocation{
				CourierID:  1,
				Altitude:   "0123456789",
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

			insertedCourierLocation, err := courierStorage.InsertCourierLocation(test.courierLocation)
			if err != nil {
				t.Fatal(err)
			}

			if insertedCourierLocation == nil {
				t.Errorf("insertedCourierLocation: Expected: %s, Got: %s", "not nill", "nil")
			}

			getCourierLocation, err := courierStorage.GetCourierLocation(insertedCourierLocation.CourierID)
			if err != nil {
				t.Fatal(err)
			}

			if getCourierLocation == nil {
				t.Errorf("getCourierLocation: Expected: %s, Got: %s", "not nill", "nil")
			}

			if getCourierLocation.CourierID != test.courierLocation.CourierID {
				t.Errorf("getCourierLocation CourierID: Expected: %v, Got: %v", test.courierLocation.CourierID, getCourierLocation.CourierID)
			}

			if getCourierLocation.Altitude != test.courierLocation.Altitude {
				t.Errorf("getCourierLocation Altitude: Expected: %v, Got: %v", test.courierLocation.Altitude, getCourierLocation.Altitude)
			}

			if getCourierLocation.Longitude != test.courierLocation.Longitude {
				t.Errorf("getCourierLocation Longitude: Expected: %v, Got: %v", test.courierLocation.Longitude, getCourierLocation.Longitude)
			}

			if getCourierLocation.Country != test.courierLocation.Country {
				t.Errorf("getCourierLocation Country: Expected: %v, Got: %v", test.courierLocation.Country, getCourierLocation.Country)
			}

			if getCourierLocation.City != test.courierLocation.City {
				t.Errorf("getCourierLocation City: Expected: %v, Got: %v", test.courierLocation.City, getCourierLocation.City)
			}

			if getCourierLocation.Region != test.courierLocation.Region {
				t.Errorf("getCourierLocation Region: Expected: %v, Got: %v", test.courierLocation.Region, getCourierLocation.Region)
			}

			if err = courierStorage.DeleteCourierLocation(insertedCourierLocation.CourierID); err != nil {
				t.Error(err)
			}

			if err := database.Close(); err != nil {
				t.Error(err)
			}
		})
	}
}
