package accountingstorage_test

import (
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/nndergunov/deliveryApp/app/pkg/configreader"

	"github.com/nndergunov/delivryApp/app/services/accounting/pkg/db"
	"github.com/nndergunov/delivryApp/app/services/accounting/pkg/domain"
	"github.com/nndergunov/delivryApp/app/services/accounting/pkg/storage/accountingstorage"
)

const configFile = "/config.yaml"

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

			database, err := db.OpenDB("postgres", configreader.GetString("DB.test"))
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

			database, err := db.OpenDB("postgres", configreader.GetString("DB.test"))
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

			database, err := db.OpenDB("postgres", configreader.GetString("DB.test"))
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

				if insertedAccountResp.UserID != gotByParamAccount.UserID {
					t.Errorf("UserID: Expected: %v, Got: %v", insertedAccountResp.UserID, gotByParamAccount.UserID)
				}

				if insertedAccountResp.UserType != gotByParamAccount.UserType {
					t.Errorf("UserType: Expected: %s, Got: %v", insertedAccountResp.UserType, gotByParamAccount.UserType)
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
		account     domain.Account
		transaction domain.Transaction
	}{
		{
			name: "TestAddToAccountBalance",
			account: domain.Account{
				UserID:   1,
				UserType: "consumer",
			},

			transaction: domain.Transaction{
				Amount: 50,
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

			database, err := db.OpenDB("postgres", configreader.GetString("DB.test"))
			if err != nil {
				t.Fatal(err)
			}

			storage := accountingstorage.NewStorage(accountingstorage.Params{DB: database})
			if err != nil {
				t.Fatal(err)
			}

			account, err := storage.InsertNewAccount(test.account)
			if err != nil {
				t.Fatal(err)
			}

			test.transaction.ToAccountID = account.ID

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

			if err = storage.DeleteAccount(account.ID); err != nil {
				t.Error(err)
			}

			if err = storage.DeleteTransaction(respData.ID); err != nil {
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
		account     domain.Account
		transaction domain.Transaction
	}{
		{
			name: "TestSubFromAccountBalance",
			account: domain.Account{
				UserID:   1,
				UserType: "consumer",
			},

			transaction: domain.Transaction{
				Amount: 50,
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

			database, err := db.OpenDB("postgres", configreader.GetString("DB.test"))
			if err != nil {
				t.Fatal(err)
			}

			storage := accountingstorage.NewStorage(accountingstorage.Params{DB: database})
			if err != nil {
				t.Fatal(err)
			}

			account, err := storage.InsertNewAccount(test.account)
			if err != nil {
				t.Fatal(err)
			}

			test.transaction.FromAccountID = account.ID

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

			if err = storage.DeleteAccount(account.ID); err != nil {
				t.Error(err)
			}

			if err = storage.DeleteTransaction(respData.ID); err != nil {
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
		fromAccount domain.Account
		toAccount   domain.Account
		transaction domain.Transaction
	}{
		{
			name: "TestInsertTransaction",
			fromAccount: domain.Account{
				UserID:   1,
				UserType: "consumer",
			},

			toAccount: domain.Account{
				UserID:   1,
				UserType: "courier",
			},

			transaction: domain.Transaction{
				Amount: 50,
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

			database, err := db.OpenDB("postgres", configreader.GetString("DB.test"))
			if err != nil {
				t.Fatal(err)
			}

			storage := accountingstorage.NewStorage(accountingstorage.Params{DB: database})
			if err != nil {
				t.Fatal(err)
			}
			// insert 1 account
			account1, err := storage.InsertNewAccount(test.fromAccount)
			if err != nil {
				t.Fatal(err)
			}
			// add balance to account 1
			test.transaction.ToAccountID = account1.ID
			trAddBalance, err := storage.AddToAccountBalance(test.transaction)
			if err != nil {
				t.Fatal(err)
			}

			// insert 2 account
			account2, err := storage.InsertNewAccount(test.toAccount)
			if err != nil {
				t.Fatal(err)
			}

			// transact from account 1 to account 2

			test.transaction.FromAccountID = account1.ID
			test.transaction.ToAccountID = account2.ID

			trFromAccountToAccount, err := storage.InsertTransaction(test.transaction)
			if err != nil {
				t.Fatal(err)
			}

			if trFromAccountToAccount == nil {
				t.Errorf("Transaction: Expected: %s, Got: %s", "not nil", "nil")
			}

			if trFromAccountToAccount.FromAccountID != test.transaction.FromAccountID {
				t.Errorf("FromAccountID: Expected: %v, Got: %v", test.transaction.FromAccountID, trFromAccountToAccount.FromAccountID)
			}

			if trFromAccountToAccount.ToAccountID != test.transaction.ToAccountID {
				t.Errorf("ToAccountID: Expected: %v, Got: %v", test.transaction.ToAccountID, trFromAccountToAccount.ToAccountID)
			}

			if trFromAccountToAccount.Amount != test.transaction.Amount {
				t.Errorf("Amount: Expected: %v, Got: %v", test.transaction.Amount, trFromAccountToAccount.Amount)
			}

			if err = storage.DeleteAccount(account1.ID); err != nil {
				t.Error(err)
			}

			if err = storage.DeleteAccount(account2.ID); err != nil {
				t.Error(err)
			}

			if err = storage.DeleteTransaction(trAddBalance.ID); err != nil {
				t.Error(err)
			}

			if err = storage.DeleteTransaction(trFromAccountToAccount.ID); err != nil {
				t.Error(err)
			}

			if err := database.Close(); err != nil {
				t.Error(err)
			}
		})
	}
}
