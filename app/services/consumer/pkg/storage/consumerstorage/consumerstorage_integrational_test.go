package consumerstorage_test

import (
	"consumer/pkg/db"
	"consumer/pkg/domain"
	"consumer/pkg/storage/consumerstorage"
	"database/sql"
	"github.com/nndergunov/deliveryApp/app/pkg/configreader"
	"os"
	"strings"
	"testing"
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
			t.Parallel()

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
			defer database.Close()

			consumerStorage := consumerstorage.NewConsumerStorage(consumerstorage.Params{DB: database})
			if err != nil {
				t.Fatal(err)
			}

			insertedConsumer, err := consumerStorage.InsertConsumer(test.consumer)
			if err != nil {
				t.Fatal(err)
			}
			defer consumerStorage.DeleteConsumer(insertedConsumer.ID)

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
			t.Parallel()

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
			defer database.Close()

			consumerStorage := consumerstorage.NewConsumerStorage(consumerstorage.Params{DB: database})
			if err != nil {
				t.Fatal(err)
			}

			insertedConsumer, err := consumerStorage.InsertConsumer(test.consumer)
			if err != nil {
				t.Fatal(err)
			}
			defer consumerStorage.DeleteConsumer(insertedConsumer.ID)

			if insertedConsumer == nil {
				t.Errorf("insertedConsumer: Expected: %s, Got: %s", "not nill", "nil")
			}

			err = consumerStorage.DeleteConsumer(insertedConsumer.ID)
			if err != nil {
				t.Fatal(err)
			}

			deletedConsumer, err := consumerStorage.GetConsumer(insertedConsumer.ID, "", "")
			if err != nil && err != sql.ErrNoRows {
				t.Fatal(err)
			}

			if deletedConsumer != nil {
				t.Errorf("deleted Consumer: Expected: %v, Got: %v", nil, deletedConsumer)
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
			t.Parallel()

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
			defer database.Close()

			consumerStorage := consumerstorage.NewConsumerStorage(consumerstorage.Params{DB: database})
			if err != nil {
				t.Fatal(err)
			}

			insertedConsumer, err := consumerStorage.InsertConsumer(test.initialConsumer)
			if err != nil {
				t.Fatal(err)
			}

			defer consumerStorage.DeleteConsumer(insertedConsumer.ID)

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
			initialConsumerList: []domain.Consumer{domain.Consumer{
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
			t.Parallel()

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
			defer database.Close()

			consumerStorage := consumerstorage.NewConsumerStorage(consumerstorage.Params{DB: database})
			if err != nil {
				t.Fatal(err)
			}

			defer consumerStorage.CleanConsumerTable()

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
			t.Parallel()

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
			defer database.Close()

			consumerStorage := consumerstorage.NewConsumerStorage(consumerstorage.Params{DB: database})
			if err != nil {
				t.Fatal(err)
			}

			insertedConsumer, err := consumerStorage.InsertConsumer(test.initialConsumer)
			if err != nil {
				t.Fatal(err)
			}

			defer consumerStorage.DeleteConsumer(insertedConsumer.ID)

			if insertedConsumer == nil {
				t.Errorf("insertedConsumer: Expected: %s, Got: %s", "not nill", "nil")
			}

			gotConsumer, err := consumerStorage.GetConsumer(insertedConsumer.ID, "", "")
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
		})
	}
}

func TestInsertConsumerLocation(t *testing.T) {
	tests := []struct {
		name             string
		consumerLocation domain.ConsumerLocation
	}{
		{
			name: "TestInsertConsumerLocation",
			consumerLocation: domain.ConsumerLocation{
				ConsumerID: 1,
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
			t.Parallel()

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
			defer database.Close()

			consumerStorage := consumerstorage.NewConsumerStorage(consumerstorage.Params{DB: database})
			if err != nil {
				t.Fatal(err)
			}

			insertedConsumerLocation, err := consumerStorage.InsertConsumerLocation(test.consumerLocation)
			if err != nil {
				t.Fatal(err)
			}

			defer consumerStorage.DeleteConsumerLocation(insertedConsumerLocation.ConsumerID)

			if insertedConsumerLocation == nil {
				t.Errorf("insertedConsumerLocation: Expected: %s, Got: %s", "not nill", "nil")
			}

			if insertedConsumerLocation == nil {
				t.Errorf("updatedConsumerLocation: Expected: %s, Got: %s", "not nill", "nil")
			}

			if insertedConsumerLocation.ConsumerID != test.consumerLocation.ConsumerID {
				t.Errorf("insertedConsumerLocation ConsumerID: Expected: %v, Got: %v", test.consumerLocation.ConsumerID, insertedConsumerLocation.ConsumerID)
			}

			if insertedConsumerLocation.Altitude != test.consumerLocation.Altitude {
				t.Errorf("insertedConsumerLocation Altitude: Expected: %v, Got: %v", test.consumerLocation.Altitude, insertedConsumerLocation.Altitude)
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
		})
	}
}

func TestUpdateConsumerLocation(t *testing.T) {
	tests := []struct {
		name                    string
		initialConsumerLocation domain.ConsumerLocation
		updateConsumerLocation  domain.ConsumerLocation
	}{
		{
			name: "Test Update Consumer",
			initialConsumerLocation: domain.ConsumerLocation{
				ConsumerID: 1,
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

			updateConsumerLocation: domain.ConsumerLocation{
				ConsumerID: 1,
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
			t.Parallel()

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
			defer database.Close()

			consumerStorage := consumerstorage.NewConsumerStorage(consumerstorage.Params{DB: database})
			if err != nil {
				t.Fatal(err)
			}

			insertedConsumerLocation, err := consumerStorage.InsertConsumerLocation(test.initialConsumerLocation)
			if err != nil {
				t.Fatal(err)
			}

			defer consumerStorage.DeleteConsumerLocation(insertedConsumerLocation.ConsumerID)

			if insertedConsumerLocation == nil {
				t.Errorf("insertedConsumerLocation: Expected: %s, Got: %s", "not nill", "nil")
			}

			updatedConsumerLocation, err := consumerStorage.UpdateConsumerLocation(test.updateConsumerLocation)
			if err != nil {
				t.Fatal(err)
			}

			if updatedConsumerLocation == nil {
				t.Errorf("updatedConsumerLocation: Expected: %s, Got: %s", "not nill", "nil")
			}

			if updatedConsumerLocation.ConsumerID != test.updateConsumerLocation.ConsumerID {
				t.Errorf("updatedConsumerLocation ConsumerID: Expected: %v, Got: %v", test.updateConsumerLocation.ConsumerID, updatedConsumerLocation.ConsumerID)
			}

			if updatedConsumerLocation.Altitude != test.updateConsumerLocation.Altitude {
				t.Errorf("updatedConsumerLocation Altitude: Expected: %v, Got: %v", test.updateConsumerLocation.Altitude, updatedConsumerLocation.Altitude)
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
		})
	}
}

func TestGetConsumerLocation(t *testing.T) {
	tests := []struct {
		name             string
		consumerLocation domain.ConsumerLocation
	}{
		{
			name: "TestGetConsumerLocation",
			consumerLocation: domain.ConsumerLocation{
				ConsumerID: 1,
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
			t.Parallel()

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
			defer database.Close()

			consumerStorage := consumerstorage.NewConsumerStorage(consumerstorage.Params{DB: database})
			if err != nil {
				t.Fatal(err)
			}

			insertedConsumerLocation, err := consumerStorage.InsertConsumerLocation(test.consumerLocation)
			if err != nil {
				t.Fatal(err)
			}

			defer consumerStorage.DeleteConsumerLocation(insertedConsumerLocation.ConsumerID)

			if insertedConsumerLocation == nil {
				t.Errorf("insertedConsumerLocation: Expected: %s, Got: %s", "not nill", "nil")
			}

			getConsumerLocation, err := consumerStorage.GetConsumerLocation(insertedConsumerLocation.ConsumerID)
			if err != nil {
				t.Fatal(err)
			}

			if getConsumerLocation == nil {
				t.Errorf("getConsumerLocation: Expected: %s, Got: %s", "not nill", "nil")
			}

			if getConsumerLocation.ConsumerID != test.consumerLocation.ConsumerID {
				t.Errorf("getConsumerLocation ConsumerID: Expected: %v, Got: %v", test.consumerLocation.ConsumerID, getConsumerLocation.ConsumerID)
			}

			if getConsumerLocation.Altitude != test.consumerLocation.Altitude {
				t.Errorf("getConsumerLocation Altitude: Expected: %v, Got: %v", test.consumerLocation.Altitude, getConsumerLocation.Altitude)
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
		})
	}
}
