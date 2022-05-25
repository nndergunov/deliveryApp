package db_test

import (
	"courier/pkg/db"
	"courier/pkg/domain"
	"courier/pkg/storage/courierstorage"
	"database/sql"
	"github.com/nndergunov/deliveryApp/app/pkg/configreader"
	"os"
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

			database, err := courierstorage.NewCourierStorage(courierstorage.Params{DB: db})
			if err != nil {
				t.Fatal(err)
			}
			defer database.CleanDB()

			insertedCourier, err := database.InsertCourier(test.courier)
			if err != nil {
				t.Fatal(err)
			}

			if insertedCourier == nil {
				t.Errorf("createCourier: Expected: %s, Got: %s", "not nill", "nil")
			}

			if insertedCourier.Username != test.courier.Username {
				t.Errorf("courierUsername: Expected: %s, Got: %s", test.courier.Username, insertedCourier.Username)
			}

		})
	}
}

func TestRemoveCourier(t *testing.T) {
	tests := []struct {
		name    string
		courier domain.Courier
	}{
		{
			name: "Test Remove Courier",
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

			database, err := courierstorage.NewCourierStorage(courierstorage.Params{DB: db})
			if err != nil {
				t.Fatal(err)
			}
			defer database.CleanDB()

			insertedCourier, err := database.InsertCourier(test.courier)
			if err != nil {
				t.Fatal(err)
			}

			if insertedCourier == nil {
				t.Errorf("createCourier: Expected: %s, Got: %s", "not nill", "nil")
			}

			err = database.RemoveCourier(insertedCourier.ID)
			if err != nil {
				t.Fatal(err)
			}

			foundCourier, err := database.GetCourier(insertedCourier.ID, "", "")
			if err != nil && err != sql.ErrNoRows {
				t.Fatal(err)
			}

			if foundCourier.Status != "nonactive" {
				t.Errorf("foundCourier Status: Expected: %s, Got: %s", "nonactive", foundCourier.Status)
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
			name: "Test Remove Courier",
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

			database, err := courierstorage.NewCourierStorage(courierstorage.Params{DB: db})
			if err != nil {
				t.Fatal(err)
			}
			defer database.CleanDB()

			insertedCourier, err := database.InsertCourier(test.initialCourier)
			if err != nil {
				t.Fatal(err)
			}

			if insertedCourier == nil {
				t.Errorf("insertedCourier: Expected: %s, Got: %s", "not nill", "nil")
			}

			test.updateCourier.ID = insertedCourier.ID
			updatedCourier, err := database.UpdateCourier(test.updateCourier)
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
		})
	}
}

func TestUpdateCourierAvailable(t *testing.T) {
	tests := []struct {
		name           string
		initialCourier domain.Courier
	}{
		{
			name: "Test Remove Courier",
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

			database, err := courierstorage.NewCourierStorage(courierstorage.Params{DB: db})
			if err != nil {
				t.Fatal(err)
			}
			defer database.CleanDB()

			insertedCourier, err := database.InsertCourier(test.initialCourier)
			if err != nil {
				t.Fatal(err)
			}

			if insertedCourier == nil {
				t.Errorf("insertedCourier: Expected: %s, Got: %s", "not nill", "nil")
			}

			updatedCourierAvailable, err := database.UpdateCourierAvailable(insertedCourier.ID, false)
			if err != nil {
				t.Fatal(err)
			}

			if updatedCourierAvailable.Available == insertedCourier.Available {
				t.Errorf("updated Courier Available: Expected: %v, Got: %v", updatedCourierAvailable.Available, insertedCourier.Available)
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

			database, err := courierstorage.NewCourierStorage(courierstorage.Params{DB: db})
			if err != nil {
				t.Fatal(err)
			}
			defer database.CleanDB()

			for _, initialCourier := range test.initialCourierList {
				insertedCourier, err := database.InsertCourier(initialCourier)
				if err != nil {
					t.Fatal(err)
				}
				if insertedCourier == nil {
					t.Errorf("insertedCourier: Expected: %s, Got: %s", "not nill", "nil")
				}

			}

			allCourier, err := database.GetAllCourier()
			if err != nil {
				t.Fatal(err)
			}

			if len(allCourier) != len(test.initialCourierList) {
				t.Errorf("get all coureir len: Expected: %v, Got: %v", len(test.initialCourierList), len(allCourier))

			}
		})
	}
}

func TestGetCourier(t *testing.T) {
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

			database, err := courierstorage.NewCourierStorage(courierstorage.Params{DB: db})
			if err != nil {
				t.Fatal(err)
			}
			defer database.CleanDB()

			insertedCourier, err := database.InsertCourier(test.initialCourier)
			if err != nil {
				t.Fatal(err)
			}
			if insertedCourier == nil {
				t.Errorf("insertedCourier: Expected: %s, Got: %s", "not nill", "nil")
			}

			gotCourier, err := database.GetCourier(insertedCourier.ID, "", "")
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

		})
	}
}
