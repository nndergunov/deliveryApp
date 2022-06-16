package consumerstorage_test

import (
	"database/sql"
	"os"
	"strings"
	"testing"

	"github.com/nndergunov/deliveryApp/app/pkg/configreader"

	"consumer/pkg/db"
	"consumer/pkg/domain"
	"consumer/pkg/storage/consumerstorage"
)

const configFile = "/config.yaml"

func TestInsertConsumer(t *testing.T) {
	tests := []struct {
		name     string
		consumer domain.Consumer
	}{
		{
			name: "Test Insert Consumer",
			consumer: domain.Consumer{
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
			line, err := os.Getwd()
			if err != nil {
				t.Fatal(err)
			}
			confPath := strings.TrimSuffix(line, "\\pkg\\storage\\consumerstorage")

			err = configreader.SetConfigFile(confPath + configFile)
			if err != nil {
				t.Fatal(err)
			}

			database, err := db.OpenDB("postgres", configreader.GetString("DB.test"))
			if err != nil {
				t.Fatal(err)
			}

			consumerStorage := consumerstorage.NewConsumerStorage(consumerstorage.Params{DB: database})
			if err != nil {
				t.Fatal(err)
			}

			insertedConsumer, err := consumerStorage.InsertConsumer(test.consumer)
			if err != nil {
				t.Fatal(err)
			}

			if insertedConsumer == nil {
				t.Errorf("inserted Consumer: Expected: %s, Got: %s", "not nill", "nil")
			}

			if insertedConsumer.Firstname != test.consumer.Firstname {
				t.Errorf("inserted Consumer Firstname: Expected: %s, Got: %s", test.consumer.Firstname, insertedConsumer.Firstname)
			}
			if insertedConsumer.Lastname != test.consumer.Lastname {
				t.Errorf("inserted Consumer: Expected Lastname: %s, Got: %s", test.consumer.Lastname, insertedConsumer.Lastname)
			}

			if insertedConsumer.Email != test.consumer.Email {
				t.Errorf("inserted Consumer Email: Expected: %s, Got: %s", test.consumer.Email, insertedConsumer.Email)
			}

			if insertedConsumer.Phone != test.consumer.Phone {
				t.Errorf("inserted Consumer Phone: Expected: %s, Got: %s", test.consumer.Phone, insertedConsumer.Phone)
			}

			if err := consumerStorage.DeleteConsumer(insertedConsumer.ID); err != nil {
				t.Error(err)
			}

			if err := database.Close(); err != nil {
				t.Error(err)
			}
		})
	}
}

func TestDeleteConsumer(t *testing.T) {
	tests := []struct {
		name     string
		consumer domain.Consumer
	}{
		{
			name: "Test Delete Consumer",
			consumer: domain.Consumer{
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
			line, err := os.Getwd()
			if err != nil {
				t.Fatal(err)
			}
			confPath := strings.TrimSuffix(line, "\\pkg\\storage\\consumerstorage")

			err = configreader.SetConfigFile(confPath + configFile)
			if err != nil {
				t.Fatal(err)
			}

			database, err := db.OpenDB("postgres", configreader.GetString("DB.test"))
			if err != nil {
				t.Fatal(err)
			}

			consumerStorage := consumerstorage.NewConsumerStorage(consumerstorage.Params{DB: database})
			if err != nil {
				t.Fatal(err)
			}

			insertedConsumer, err := consumerStorage.InsertConsumer(test.consumer)
			if err != nil {
				t.Fatal(err)
			}

			if insertedConsumer == nil {
				t.Errorf("insertedConsumer: Expected: %s, Got: %s", "not nill", "nil")
			}

			err = consumerStorage.DeleteConsumer(insertedConsumer.ID)
			if err != nil {
				t.Fatal(err)
			}

			deletedConsumer, err := consumerStorage.GetConsumerByID(insertedConsumer.ID)
			if err != nil && err != sql.ErrNoRows {
				t.Fatal(err)
			}

			if deletedConsumer != nil {
				t.Errorf("deleted Consumer: Expected: %v, Got: %v", nil, deletedConsumer)
			}

			if err := database.Close(); err != nil {
				t.Error(err)
			}
		})
	}
}

func TestUpdateConsumer(t *testing.T) {
	tests := []struct {
		name            string
		initialConsumer domain.Consumer
		updateConsumer  domain.Consumer
	}{
		{
			name: "Test Update Consumer",
			initialConsumer: domain.Consumer{
				Firstname: "vasya",
				Lastname:  "",
				Email:     "vasya@gmail.com",
				Phone:     "123456789",
			},
			updateConsumer: domain.Consumer{
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
			line, err := os.Getwd()
			if err != nil {
				t.Fatal(err)
			}
			confPath := strings.TrimSuffix(line, "\\pkg\\storage\\consumerstorage")

			err = configreader.SetConfigFile(confPath + configFile)
			if err != nil {
				t.Fatal(err)
			}

			database, err := db.OpenDB("postgres", configreader.GetString("DB.test"))
			if err != nil {
				t.Fatal(err)
			}

			consumerStorage := consumerstorage.NewConsumerStorage(consumerstorage.Params{DB: database})
			if err != nil {
				t.Fatal(err)
			}

			insertedConsumer, err := consumerStorage.InsertConsumer(test.initialConsumer)
			if err != nil {
				t.Fatal(err)
			}

			if insertedConsumer == nil {
				t.Errorf("insertedConsumer: Expected: %s, Got: %s", "not nill", "nil")
			}

			test.updateConsumer.ID = insertedConsumer.ID
			updatedConsumer, err := consumerStorage.UpdateConsumer(test.updateConsumer)
			if err != nil {
				t.Fatal(err)
			}

			if updatedConsumer.Firstname != test.updateConsumer.Firstname {
				t.Errorf("updatedConsumer Firstname: Expected: %s, Got: %s", test.updateConsumer.Firstname, updatedConsumer.Firstname)
			}

			if updatedConsumer.Lastname != test.updateConsumer.Lastname {
				t.Errorf("updatedConsumer Lastname: Expected: %s, Got: %s", test.updateConsumer.Lastname, updatedConsumer.Lastname)
			}

			if updatedConsumer.Email != test.updateConsumer.Email {
				t.Errorf("updatedConsumer Email: Expected: %s, Got: %s", test.updateConsumer.Email, updatedConsumer.Email)
			}

			err = consumerStorage.DeleteConsumer(insertedConsumer.ID)
			if err != nil {
				t.Error(err)
			}

			if err := database.Close(); err != nil {
				t.Error(err)
			}
		})
	}
}

func TestGetAllConsumer(t *testing.T) {
	tests := []struct {
		name                string
		initialConsumerList []domain.Consumer
	}{
		{
			name: "Test Get ALl Consumer",
			initialConsumerList: []domain.Consumer{
				{
					Firstname: "vasya",
					Lastname:  "",
					Email:     "vasya@gmail.com",
					Phone:     "123456789",
				},
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
			confPath := strings.TrimSuffix(line, "\\pkg\\storage\\consumerstorage")

			err = configreader.SetConfigFile(confPath + configFile)
			if err != nil {
				t.Fatal(err)
			}

			database, err := db.OpenDB("postgres", configreader.GetString("DB.test"))
			if err != nil {
				t.Fatal(err)
			}

			consumerStorage := consumerstorage.NewConsumerStorage(consumerstorage.Params{DB: database})
			if err != nil {
				t.Fatal(err)
			}

			for _, initialConsumer := range test.initialConsumerList {
				insertedConsumer, err := consumerStorage.InsertConsumer(initialConsumer)
				if err != nil {
					t.Fatal(err)
				}
				if insertedConsumer == nil {
					t.Errorf("insertedConsumer: Expected: %s, Got: %s", "not nill", "nil")
				}
			}

			allConsumer, err := consumerStorage.GetAllConsumer()
			if err != nil {
				t.Fatal(err)
			}

			if len(allConsumer) != len(test.initialConsumerList) {
				t.Errorf("GetaAllConsumer len: Expected: %v, Got: %v", len(test.initialConsumerList), len(allConsumer))
			}

			if err = consumerStorage.CleanConsumerTable(); err != nil {
				t.Error(err)
			}

			if err := database.Close(); err != nil {
				t.Error(err)
			}
		})
	}
}

func TestGetConsumer(t *testing.T) {
	tests := []struct {
		name            string
		initialConsumer domain.Consumer
	}{
		{
			name: "Test Get Consumer",
			initialConsumer: domain.Consumer{
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
			line, err := os.Getwd()
			if err != nil {
				t.Fatal(err)
			}
			confPath := strings.TrimSuffix(line, "\\pkg\\storage\\consumerstorage")

			err = configreader.SetConfigFile(confPath + configFile)
			if err != nil {
				t.Fatal(err)
			}

			database, err := db.OpenDB("postgres", configreader.GetString("DB.test"))
			if err != nil {
				t.Fatal(err)
			}

			consumerStorage := consumerstorage.NewConsumerStorage(consumerstorage.Params{DB: database})
			if err != nil {
				t.Fatal(err)
			}

			insertedConsumer, err := consumerStorage.InsertConsumer(test.initialConsumer)
			if err != nil {
				t.Fatal(err)
			}

			if insertedConsumer == nil {
				t.Errorf("insertedConsumer: Expected: %s, Got: %s", "not nill", "nil")
			}

			gotConsumer, err := consumerStorage.GetConsumerByID(insertedConsumer.ID)
			if err != nil {
				t.Fatal(err)
			}

			if gotConsumer.Firstname != test.initialConsumer.Firstname {
				t.Errorf("get consumer Firstname: Expected: %s, Got: %s", test.initialConsumer.Firstname, gotConsumer.Firstname)
			}

			if gotConsumer.Lastname != test.initialConsumer.Lastname {
				t.Errorf("get consumer Lastname: Expected: %s, Got: %s", test.initialConsumer.Lastname, gotConsumer.Lastname)
			}

			if gotConsumer.Phone != test.initialConsumer.Phone {
				t.Errorf("get consumer Phone: Expected: %s, Got: %s", test.initialConsumer.Phone, gotConsumer.Phone)
			}

			if gotConsumer.Email != test.initialConsumer.Email {
				t.Errorf("get consumer Email: Expected: %s, Got: %s", test.initialConsumer.Email, gotConsumer.Email)
			}

			if err = consumerStorage.DeleteConsumer(insertedConsumer.ID); err != nil {
				t.Error(err)
			}

			if err := database.Close(); err != nil {
				t.Error(err)
			}
		})
	}
}

func TestInsertConsumerLocation(t *testing.T) {
	tests := []struct {
		name             string
		consumerLocation domain.Location
	}{
		{
			name: "TestInsertConsumerLocation",
			consumerLocation: domain.Location{
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
			confPath := strings.TrimSuffix(line, "\\pkg\\storage\\consumerstorage")

			err = configreader.SetConfigFile(confPath + configFile)
			if err != nil {
				t.Fatal(err)
			}

			database, err := db.OpenDB("postgres", configreader.GetString("DB.test"))
			if err != nil {
				t.Fatal(err)
			}

			consumerStorage := consumerstorage.NewConsumerStorage(consumerstorage.Params{DB: database})
			if err != nil {
				t.Fatal(err)
			}

			insertedConsumerLocation, err := consumerStorage.InsertLocation(test.consumerLocation)
			if err != nil {
				t.Fatal(err)
			}

			if insertedConsumerLocation == nil {
				t.Errorf("insertedConsumerLocation: Expected: %s, Got: %s", "not nill", "nil")
			}

			if insertedConsumerLocation == nil {
				t.Errorf("updatedConsumerLocation: Expected: %s, Got: %s", "not nill", "nil")
			}

			if insertedConsumerLocation.UserID != test.consumerLocation.UserID {
				t.Errorf("insertedConsumerLocation UserID: Expected: %v, Got: %v", test.consumerLocation.UserID, insertedConsumerLocation.UserID)
			}

			if insertedConsumerLocation.Latitude != test.consumerLocation.Latitude {
				t.Errorf("insertedConsumerLocation Latitude: Expected: %v, Got: %v", test.consumerLocation.Latitude, insertedConsumerLocation.Latitude)
			}

			if insertedConsumerLocation.Longitude != test.consumerLocation.Longitude {
				t.Errorf("insertedConsumerLocation Longitude: Expected: %v, Got: %v", test.consumerLocation.Longitude, insertedConsumerLocation.Longitude)
			}

			if insertedConsumerLocation.Country != test.consumerLocation.Country {
				t.Errorf("insertedConsumerLocation Country: Expected: %v, Got: %v", test.consumerLocation.Country, insertedConsumerLocation.Country)
			}

			if insertedConsumerLocation.City != test.consumerLocation.City {
				t.Errorf("insertedConsumerLocation City: Expected: %v, Got: %v", test.consumerLocation.City, insertedConsumerLocation.City)
			}

			if insertedConsumerLocation.Region != test.consumerLocation.Region {
				t.Errorf("insertedConsumerLocation Region: Expected: %v, Got: %v", test.consumerLocation.Region, insertedConsumerLocation.Region)
			}

			if err = consumerStorage.DeleteLocation(insertedConsumerLocation.UserID); err != nil {
				t.Error(err)
			}

			if err := database.Close(); err != nil {
				t.Error(err)
			}
		})
	}
}

func TestUpdateConsumerLocation(t *testing.T) {
	tests := []struct {
		name                    string
		initialConsumerLocation domain.Location
		updateConsumerLocation  domain.Location
	}{
		{
			name: "Test Update Consumer",
			initialConsumerLocation: domain.Location{
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

			updateConsumerLocation: domain.Location{
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
			confPath := strings.TrimSuffix(line, "\\pkg\\storage\\consumerstorage")

			err = configreader.SetConfigFile(confPath + configFile)
			if err != nil {
				t.Fatal(err)
			}

			database, err := db.OpenDB("postgres", configreader.GetString("DB.test"))
			if err != nil {
				t.Fatal(err)
			}

			consumerStorage := consumerstorage.NewConsumerStorage(consumerstorage.Params{DB: database})
			if err != nil {
				t.Fatal(err)
			}

			insertedConsumerLocation, err := consumerStorage.InsertLocation(test.initialConsumerLocation)
			if err != nil {
				t.Fatal(err)
			}

			if insertedConsumerLocation == nil {
				t.Errorf("insertedConsumerLocation: Expected: %s, Got: %s", "not nill", "nil")
			}

			updatedConsumerLocation, err := consumerStorage.UpdateLocation(test.updateConsumerLocation)
			if err != nil {
				t.Fatal(err)
			}

			if updatedConsumerLocation == nil {
				t.Errorf("updatedConsumerLocation: Expected: %s, Got: %s", "not nill", "nil")
			}

			if updatedConsumerLocation.UserID != test.updateConsumerLocation.UserID {
				t.Errorf("updatedConsumerLocation UserID: Expected: %v, Got: %v", test.updateConsumerLocation.UserID, updatedConsumerLocation.UserID)
			}

			if updatedConsumerLocation.Latitude != test.updateConsumerLocation.Latitude {
				t.Errorf("updatedConsumerLocation Latitude: Expected: %v, Got: %v", test.updateConsumerLocation.Latitude, updatedConsumerLocation.Latitude)
			}

			if updatedConsumerLocation.Longitude != test.updateConsumerLocation.Longitude {
				t.Errorf("updatedConsumerLocation Longitude: Expected: %v, Got: %v", test.updateConsumerLocation.Longitude, updatedConsumerLocation.Longitude)
			}

			if updatedConsumerLocation.Country != test.updateConsumerLocation.Country {
				t.Errorf("updatedConsumerLocation Country: Expected: %v, Got: %v", test.updateConsumerLocation.Country, updatedConsumerLocation.Country)
			}

			if updatedConsumerLocation.City != test.updateConsumerLocation.City {
				t.Errorf("updatedConsumerLocation City: Expected: %v, Got: %v", test.updateConsumerLocation.City, updatedConsumerLocation.City)
			}

			if updatedConsumerLocation.Region != test.updateConsumerLocation.Region {
				t.Errorf("updatedConsumerLocation Region: Expected: %v, Got: %v", test.updateConsumerLocation.Region, updatedConsumerLocation.Region)
			}

			if err = consumerStorage.DeleteLocation(insertedConsumerLocation.UserID); err != nil {
				t.Error(err)
			}

			if err := database.Close(); err != nil {
				t.Error(err)
			}
		})
	}
}

func TestGetConsumerLocation(t *testing.T) {
	tests := []struct {
		name             string
		consumerLocation domain.Location
	}{
		{
			name: "TestGetConsumerLocation",
			consumerLocation: domain.Location{
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
			confPath := strings.TrimSuffix(line, "\\pkg\\storage\\consumerstorage")

			err = configreader.SetConfigFile(confPath + configFile)
			if err != nil {
				t.Fatal(err)
			}

			database, err := db.OpenDB("postgres", configreader.GetString("DB.test"))
			if err != nil {
				t.Fatal(err)
			}

			consumerStorage := consumerstorage.NewConsumerStorage(consumerstorage.Params{DB: database})
			if err != nil {
				t.Fatal(err)
			}

			insertedConsumerLocation, err := consumerStorage.InsertLocation(test.consumerLocation)
			if err != nil {
				t.Fatal(err)
			}

			if insertedConsumerLocation == nil {
				t.Errorf("insertedConsumerLocation: Expected: %s, Got: %s", "not nill", "nil")
			}

			getConsumerLocation, err := consumerStorage.GetLocation(insertedConsumerLocation.UserID)
			if err != nil {
				t.Fatal(err)
			}

			if getConsumerLocation == nil {
				t.Errorf("getConsumerLocation: Expected: %s, Got: %s", "not nill", "nil")
			}

			if getConsumerLocation.UserID != test.consumerLocation.UserID {
				t.Errorf("getConsumerLocation UserID: Expected: %v, Got: %v", test.consumerLocation.UserID, getConsumerLocation.UserID)
			}

			if getConsumerLocation.Latitude != test.consumerLocation.Latitude {
				t.Errorf("getConsumerLocation Latitude: Expected: %v, Got: %v", test.consumerLocation.Latitude, getConsumerLocation.Latitude)
			}

			if getConsumerLocation.Longitude != test.consumerLocation.Longitude {
				t.Errorf("getConsumerLocation Longitude: Expected: %v, Got: %v", test.consumerLocation.Longitude, getConsumerLocation.Longitude)
			}

			if getConsumerLocation.Country != test.consumerLocation.Country {
				t.Errorf("getConsumerLocation Country: Expected: %v, Got: %v", test.consumerLocation.Country, getConsumerLocation.Country)
			}

			if getConsumerLocation.City != test.consumerLocation.City {
				t.Errorf("getConsumerLocation City: Expected: %v, Got: %v", test.consumerLocation.City, getConsumerLocation.City)
			}

			if getConsumerLocation.Region != test.consumerLocation.Region {
				t.Errorf("getConsumerLocation Region: Expected: %v, Got: %v", test.consumerLocation.Region, getConsumerLocation.Region)
			}

			if err = consumerStorage.DeleteLocation(insertedConsumerLocation.UserID); err != nil {
				t.Error(err)
			}

			if err := database.Close(); err != nil {
				t.Error(err)
			}
		})
	}
}
