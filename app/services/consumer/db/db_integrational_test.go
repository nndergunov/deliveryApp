package db_test

import (
	"consumer/db"
	"consumer/db/storage"
	"consumer/domain"
	"database/sql"
	"github.com/nndergunov/deliveryApp/app/pkg/configreader"
	"os"
	"strings"
	"testing"
)

const configFile = "/config.yaml"

func TestInsertCourier(t *testing.T) {
	tests := []struct {
		name     string
		consumer domain.Consumer
	}{
		{
			name: "Test Insert Courier",
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
			confPath := strings.TrimSuffix(line, "db")

			err = configreader.SetConfigFile(confPath + configFile)
			if err != nil {
				t.Fatal(err)
			}

			db, err := db.Open("postgres", configreader.GetString("DB.test"))
			if err != nil {
				t.Fatal(err)
			}
			defer db.Close()

			database, err := storage.NewConsumerStorage(storage.Params{DB: db})
			if err != nil {
				t.Fatal(err)
			}
			defer database.CleanConsumerTable()

			insertedCourier, err := database.InsertConsumer(test.consumer)
			if err != nil {
				t.Fatal(err)
			}

			if insertedCourier == nil {
				t.Errorf("inserted Consumer: Expected: %s, Got: %s", "not nill", "nil")
			}

			if insertedCourier.Firstname != test.consumer.Firstname {
				t.Errorf("inserted Consumer: Expected: %s, Got: %s", test.consumer.Firstname, insertedCourier.Firstname)
			}

		})
	}
}

func TestConsumerCourier(t *testing.T) {
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
			confPath := strings.TrimSuffix(line, "db")

			err = configreader.SetConfigFile(confPath + configFile)
			if err != nil {
				t.Fatal(err)
			}

			db, err := db.Open("postgres", configreader.GetString("DB.test"))
			if err != nil {
				t.Fatal(err)
			}
			defer db.Close()

			database, err := storage.NewConsumerStorage(storage.Params{DB: db})
			if err != nil {
				t.Fatal(err)
			}
			defer database.CleanConsumerTable()

			insertedCourier, err := database.InsertConsumer(test.consumer)
			if err != nil {
				t.Fatal(err)
			}

			if insertedCourier == nil {
				t.Errorf("createCourier: Expected: %s, Got: %s", "not nill", "nil")
			}

			err = database.DeleteConsumer(insertedCourier.ID)
			if err != nil {
				t.Fatal(err)
			}

			deletedConsumer, err := database.GetConsumer(insertedCourier.ID, "", "")
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
		name           string
		initialCourier domain.Consumer
		updateConsumer domain.Consumer
	}{
		{
			name: "Test Update Consumer",
			initialCourier: domain.Consumer{
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
			confPath := strings.TrimSuffix(line, "db")

			err = configreader.SetConfigFile(confPath + configFile)
			if err != nil {
				t.Fatal(err)
			}

			db, err := db.Open("postgres", configreader.GetString("DB.test"))
			if err != nil {
				t.Fatal(err)
			}
			defer db.Close()

			database, err := storage.NewConsumerStorage(storage.Params{DB: db})
			if err != nil {
				t.Fatal(err)
			}
			defer database.CleanConsumerTable()
			defer database.CleanConsumerLocationTable()

			insertedConsumer, err := database.InsertConsumer(test.initialCourier)
			if err != nil {
				t.Fatal(err)
			}

			if insertedConsumer == nil {
				t.Errorf("insertedConsumer: Expected: %s, Got: %s", "not nill", "nil")
			}

			test.updateConsumer.ID = insertedConsumer.ID
			updatedConsumer, err := database.UpdateConsumer(test.updateConsumer)
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
			confPath := strings.TrimSuffix(line, "db")

			err = configreader.SetConfigFile(confPath + configFile)
			if err != nil {
				t.Fatal(err)
			}

			db, err := db.Open("postgres", configreader.GetString("DB.test"))
			if err != nil {
				t.Fatal(err)
			}
			defer db.Close()

			database, err := storage.NewConsumerStorage(storage.Params{DB: db})
			if err != nil {
				t.Fatal(err)
			}
			defer database.CleanConsumerTable()
			defer database.CleanConsumerLocationTable()

			for _, initialConsumer := range test.initialConsumerList {
				insertedConsumer, err := database.InsertConsumer(initialConsumer)
				if err != nil {
					t.Fatal(err)
				}
				if insertedConsumer == nil {
					t.Errorf("insertedConsumer: Expected: %s, Got: %s", "not nill", "nil")
				}

			}

			allConsumer, err := database.GetAllConsumer()
			if err != nil {
				t.Fatal(err)
			}

			if len(allConsumer) != len(test.initialConsumerList) {
				t.Errorf("allConsumer len: Expected: %v, Got: %v", len(test.initialConsumerList), len(allConsumer))

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
			confPath := strings.TrimSuffix(line, "db")

			err = configreader.SetConfigFile(confPath + configFile)
			if err != nil {
				t.Fatal(err)
			}

			db, err := db.Open("postgres", configreader.GetString("DB.test"))
			if err != nil {
				t.Fatal(err)
			}
			defer db.Close()

			database, err := storage.NewConsumerStorage(storage.Params{DB: db})
			if err != nil {
				t.Fatal(err)
			}
			defer database.CleanConsumerTable()
			defer database.CleanConsumerLocationTable()

			insertedConsumer, err := database.InsertConsumer(test.initialConsumer)
			if err != nil {
				t.Fatal(err)
			}
			if insertedConsumer == nil {
				t.Errorf("insertedConsumer: Expected: %s, Got: %s", "not nill", "nil")
			}

			gotCourier, err := database.GetConsumer(insertedConsumer.ID, "", "")
			if err != nil {
				t.Fatal(err)
			}

			if gotCourier.Firstname != test.initialConsumer.Firstname {
				t.Errorf("get consumer Firstname: Expected: %s, Got: %s", test.initialConsumer.Firstname, gotCourier.Firstname)
			}
			if gotCourier.Lastname != test.initialConsumer.Lastname {
				t.Errorf("get consumer Lastname: Expected: %s, Got: %s", test.initialConsumer.Lastname, gotCourier.Lastname)
			}
			if gotCourier.Phone != test.initialConsumer.Phone {
				t.Errorf("get consumer Phone: Expected: %s, Got: %s", test.initialConsumer.Phone, gotCourier.Phone)
			}
			if gotCourier.Email != test.initialConsumer.Email {
				t.Errorf("get consumer Email: Expected: %s, Got: %s", test.initialConsumer.Email, gotCourier.Email)
			}

		})
	}
}
