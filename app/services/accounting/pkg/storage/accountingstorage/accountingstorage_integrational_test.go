package accountingstorage_test

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/nndergunov/deliveryApp/app/pkg/configreader"

	"github.com/nndergunov/deliveryApp/app/services/accounting/pkg/db"
	"github.com/nndergunov/deliveryApp/app/services/accounting/pkg/domain"
	"github.com/nndergunov/deliveryApp/app/services/accounting/pkg/storage/accountingstorage"
)

const configFile = "/config.yaml"

var dbURL = fmt.Sprintf("host=" + configreader.GetString("database.test.host") +
	" port=" + configreader.GetString("database.test.port") +
	" user=" + configreader.GetString("database.test.user") +
	" password=" + configreader.GetString("database.test.password") +
	" dbname=" + configreader.GetString("database.test.dbName") +
	" sslmode=" + configreader.GetString("database.test.sslmode"))

func TestInsertAccount(t *testing.T) {
	tests := []struct {
		name    string
		account domain.Account
	}{
		{
			name: "Test Insert Courier",
			account: domain.Account{
				UserID:   1,
				UserType: "courier",
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
			confPath := strings.TrimSuffix(line, "\\pkg\\storage\\accountingstorage")

			err = configreader.SetConfigFile(confPath + configFile)
			if err != nil {
				t.Fatal(err)
			}

			database, err := db.OpenDB("postgres", dbURL)
			if err != nil {
				t.Fatal(err)
			}

			storage := accountingstorage.NewStorage(accountingstorage.Params{DB: database})
			if err != nil {
				t.Fatal(err)
			}

			resp, err := storage.InsertNewAccount(test.account)
			if err != nil {
				t.Fatal(err)
			}

			if resp == nil {
				t.Errorf("createCourier: Expected: %s, Got: %s", "not nil", "nil")
			}

			if resp.ID != test.account.ID {
				t.Errorf("ID: Expected: %v, Got: %v", test.account.ID, resp.ID)
			}

			if resp.UserID != test.account.UserID {
				t.Errorf("UserID: Expected: %v, Got: %v", test.account.UserID, resp.UserID)
			}

			if resp.UserType != test.account.UserType {
				t.Errorf("UserType: Expected: %s, Got: %s", test.account.UserType, resp.UserType)
			}

			if resp.Balance != test.account.Balance {
				t.Errorf("Balance: Expected: %v, Got: %v", test.account.Balance, resp.Balance)
			}

			if err = storage.DeleteAccount(resp.ID); err != nil {
				t.Error(err)
			}

			if err := database.Close(); err != nil {
				t.Error(err)
			}
		})
	}
}

func TestGetAccountByID(t *testing.T) {
	tests := []struct {
		name    string
		account domain.Account
	}{
		{
			name: "GetAccountByID",
			account: domain.Account{
				UserID:   1,
				UserType: "transaction",
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
			confPath := strings.TrimSuffix(line, "\\pkg\\storage\\accountingstorage")

			err = configreader.SetConfigFile(confPath + configFile)
			if err != nil {
				t.Fatal(err)
			}

			database, err := db.OpenDB("postgres", dbURL)
			if err != nil {
				t.Fatal(err)
			}

			storage := accountingstorage.NewStorage(accountingstorage.Params{DB: database})
			if err != nil {
				t.Fatal(err)
			}

			resp, err := storage.InsertNewAccount(test.account)
			if err != nil {
				t.Fatal(err)
			}

			if resp == nil {
				t.Errorf("createCourier: Expected: %s, Got: %s", "not nil", "nil")
			}
			resp2, err := storage.GetAccountByID(resp.ID)
			if err != nil {
				t.Fatal(err)
			}

			if resp2.ID != test.account.ID {
				t.Errorf("ID: Expected: %v, Got: %v", test.account.ID, resp2.ID)
			}

			if resp2.UserID != test.account.UserID {
				t.Errorf("UserID: Expected: %v, Got: %v", test.account.UserID, resp2.UserID)
			}

			if resp2.UserType != test.account.UserType {
				t.Errorf("UserType: Expected: %s, Got: %s", test.account.UserType, resp2.UserType)
			}

			if resp2.Balance != test.account.Balance {
				t.Errorf("Balance: Expected: %v, Got: %v", test.account.Balance, resp2.Balance)
			}

			if err = storage.DeleteAccount(resp.ID); err != nil {
				t.Error(err)
			}

			if err := database.Close(); err != nil {
				t.Error(err)
			}
		})
	}
}

func TestGetAccountListByParam(t *testing.T) {
	tests := []struct {
		name    string
		account domain.Account
	}{
		{
			name: "TestGetAccountListByParam",
			account: domain.Account{
				UserID:   1,
				UserType: "courier",
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
			confPath := strings.TrimSuffix(line, "\\pkg\\storage\\accountingstorage")

			err = configreader.SetConfigFile(confPath + configFile)
			if err != nil {
				t.Fatal(err)
			}

			database, err := db.OpenDB("postgres", dbURL)
			if err != nil {
				t.Fatal(err)
			}

			storage := accountingstorage.NewStorage(accountingstorage.Params{DB: database})
			if err != nil {
				t.Fatal(err)
			}

			insertedAccountResp, err := storage.InsertNewAccount(test.account)
			if err != nil {
				t.Fatal(err)
			}

			if insertedAccountResp == nil {
				t.Errorf("createCourier: Expected: %s, Got: %s", "not nil", "nil")
			}

			param := domain.SearchParam{}
			param["user_id"] = strconv.Itoa(insertedAccountResp.UserID)
			param["user_type"] = insertedAccountResp.UserType

			resp2List, err := storage.GetAccountListByParam(param)
			if err != nil {
				t.Fatal(err)
			}
			for _, gotByParamAccount := range resp2List {

				if insertedAccountResp.ID != gotByParamAccount.ID {
					t.Errorf("ID: Expected: %v, Got: %v", insertedAccountResp.ID, gotByParamAccount.ID)
				}

				if insertedAccountResp.UserID != gotByParamAccount.UserID {
					t.Errorf("UserID: Expected: %v, Got: %v", insertedAccountResp.UserID, gotByParamAccount.UserID)
				}

				if insertedAccountResp.UserType != gotByParamAccount.UserType {
					t.Errorf("UserType: Expected: %s, Got: %v", insertedAccountResp.UserType, gotByParamAccount.UserID)
				}

				if insertedAccountResp.Balance != gotByParamAccount.Balance {
					t.Errorf("Balance: Expected: %v, Got: %v", insertedAccountResp.Balance, gotByParamAccount.Balance)
				}

				if err = storage.DeleteAccount(insertedAccountResp.ID); err != nil {
					t.Error(err)
				}
			}

			if err := database.Close(); err != nil {
				t.Error(err)
			}
		})
	}
}

func TestAddToAccountBalance(t *testing.T) {
	tests := []struct {
		name        string
		transaction domain.Transaction
	}{
		{
			name: "TestAddToAccountBalance",
			transaction: domain.Transaction{
				FromAccountID: 0,
				ToAccountID:   2,
				Amount:        50,
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
			confPath := strings.TrimSuffix(line, "\\pkg\\storage\\accountingstorage")

			err = configreader.SetConfigFile(confPath + configFile)
			if err != nil {
				t.Fatal(err)
			}

			database, err := db.OpenDB("postgres", dbURL)
			if err != nil {
				t.Fatal(err)
			}

			storage := accountingstorage.NewStorage(accountingstorage.Params{DB: database})
			if err != nil {
				t.Fatal(err)
			}

			respData, err := storage.AddToAccountBalance(test.transaction)
			if err != nil {
				t.Fatal(err)
			}

			if respData == nil {
				t.Errorf("Transaction: Expected: %s, Got: %s", "not nil", "nil")
			}

			if respData.FromAccountID != test.transaction.FromAccountID {
				t.Errorf("FromAccountID: Expected: %v, Got: %v", test.transaction.FromAccountID, respData.FromAccountID)
			}

			if respData.ToAccountID != test.transaction.ToAccountID {
				t.Errorf("ToAccountID: Expected: %v, Got: %v", test.transaction.ToAccountID, respData.ToAccountID)
			}

			if respData.Amount != test.transaction.Amount {
				t.Errorf("Amount: Expected: %v, Got: %v", test.transaction.Amount, respData.Amount)
			}

			if err = storage.DeleteAccount(respData.ID); err != nil {
				t.Error(err)
			}

			if err := database.Close(); err != nil {
				t.Error(err)
			}
		})
	}
}

func TestSubFromAccountBalance(t *testing.T) {
	tests := []struct {
		name        string
		transaction domain.Transaction
	}{
		{
			name: "TestSubFromAccountBalance",
			transaction: domain.Transaction{
				FromAccountID: 1,
				ToAccountID:   0,
				Amount:        50,
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
			confPath := strings.TrimSuffix(line, "\\pkg\\storage\\accountingstorage")

			err = configreader.SetConfigFile(confPath + configFile)
			if err != nil {
				t.Fatal(err)
			}

			database, err := db.OpenDB("postgres", dbURL)
			if err != nil {
				t.Fatal(err)
			}

			storage := accountingstorage.NewStorage(accountingstorage.Params{DB: database})
			if err != nil {
				t.Fatal(err)
			}

			respData, err := storage.SubFromAccountBalance(test.transaction)
			if err != nil {
				t.Fatal(err)
			}

			if respData == nil {
				t.Errorf("Transaction: Expected: %s, Got: %s", "not nil", "nil")
			}

			if respData.FromAccountID != test.transaction.FromAccountID {
				t.Errorf("FromAccountID: Expected: %v, Got: %v", test.transaction.FromAccountID, respData.FromAccountID)
			}

			if respData.ToAccountID != test.transaction.ToAccountID {
				t.Errorf("ToAccountID: Expected: %v, Got: %v", test.transaction.ToAccountID, respData.ToAccountID)
			}

			if respData.Amount != test.transaction.Amount {
				t.Errorf("Amount: Expected: %v, Got: %v", test.transaction.Amount, respData.Amount)
			}

			if err = storage.DeleteAccount(respData.ID); err != nil {
				t.Error(err)
			}

			if err := database.Close(); err != nil {
				t.Error(err)
			}
		})
	}
}

func TestInsertTransaction(t *testing.T) {
	tests := []struct {
		name        string
		transaction domain.Transaction
	}{
		{
			name: "TestSubFromAccountBalance",
			transaction: domain.Transaction{
				FromAccountID: 1,
				ToAccountID:   2,
				Amount:        50,
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
			confPath := strings.TrimSuffix(line, "\\pkg\\storage\\accountingstorage")

			err = configreader.SetConfigFile(confPath + configFile)
			if err != nil {
				t.Fatal(err)
			}

			database, err := db.OpenDB("postgres", dbURL)
			if err != nil {
				t.Fatal(err)
			}

			storage := accountingstorage.NewStorage(accountingstorage.Params{DB: database})
			if err != nil {
				t.Fatal(err)
			}

			respData, err := storage.InsertTransaction(test.transaction)
			if err != nil {
				t.Fatal(err)
			}

			if respData == nil {
				t.Errorf("Transaction: Expected: %s, Got: %s", "not nil", "nil")
			}

			if respData.FromAccountID != test.transaction.FromAccountID {
				t.Errorf("FromAccountID: Expected: %v, Got: %v", test.transaction.FromAccountID, respData.FromAccountID)
			}

			if respData.ToAccountID != test.transaction.ToAccountID {
				t.Errorf("ToAccountID: Expected: %v, Got: %v", test.transaction.ToAccountID, respData.ToAccountID)
			}

			if respData.Amount != test.transaction.Amount {
				t.Errorf("Amount: Expected: %v, Got: %v", test.transaction.Amount, respData.Amount)
			}

			if err = storage.DeleteAccount(respData.ID); err != nil {
				t.Error(err)
			}

			if err := database.Close(); err != nil {
				t.Error(err)
			}
		})
	}
}
