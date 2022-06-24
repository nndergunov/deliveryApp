package accountingservice_test

import (
	"bytes"
	"net/http"
	"strconv"
	"testing"
	"time"

	"github.com/nndergunov/deliveryApp/app/pkg/api/v1/accountingapi"

	"github.com/nndergunov/deliveryApp/app/pkg/api/v1"
)

const baseAddr = "http://localhost:8081"

func TestInsertAccountEndpoint(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name              string
		newAccountRequest accountingapi.NewAccountRequest
		accountResponse   accountingapi.AccountResponse
	}{
		{
			"Insert account simple test",
			accountingapi.NewAccountRequest{
				UserID:   1,
				UserType: "courier",
			},
			accountingapi.AccountResponse{
				UserID:    1,
				UserType:  "courier",
				Balance:   50,
				CreatedAt: time.Time{},
				UpdatedAt: time.Time{},
			},
		},
	}

	for _, currentTest := range tests {
		test := currentTest

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			reqBody, err := v1.Encode(test.newAccountRequest)
			if err != nil {
				t.Fatal(err)
			}

			resp, err := http.Post(http.MethodPost, "/v1/accounts", bytes.NewBuffer(reqBody))
			if err != nil {
				t.Fatal(err)
			}

			if resp.StatusCode != http.StatusOK {
				t.Fatalf("Response status: %d", resp.StatusCode)
			}

			respData := accountingapi.AccountResponse{}
			if err = accountingapi.DecodeJSON(resp.Body, &respData); err != nil {
				t.Fatal(err)
			}

			if respData.ID != test.accountResponse.ID {
				t.Errorf("ID: Expected: %v, Got: %v", test.accountResponse.ID, respData.ID)
			}

			if respData.UserID != test.accountResponse.UserID {
				t.Errorf("UserID: Expected: %v, Got: %v", test.accountResponse.UserID, respData.UserID)
			}

			if respData.UserType != test.accountResponse.UserType {
				t.Errorf("UserType: Expected: %s, Got: %s", test.accountResponse.UserType, respData.UserType)
			}

			if respData.Balance != test.accountResponse.Balance {
				t.Errorf("Balance: Expected: %v, Got: %v", test.accountResponse.Balance, respData.Balance)
			}

			// Deleting instance.
			deleter := http.DefaultClient

			delReq, err := http.NewRequest(http.MethodDelete,
				baseAddr+"/v1/accounts/"+strconv.Itoa(respData.ID), nil)
			if err != nil {
				t.Error(err)
			}

			_, err = deleter.Do(delReq)
			if err != nil {
				t.Errorf("Could not delete: %v", err)
			}
		})
	}
}

func TestGetAccountEndpoint(t *testing.T) {
	tests := []struct {
		name            string
		accountDataList accountingapi.NewAccountRequest
	}{
		{
			"GetAccount simple test",
			accountingapi.NewAccountRequest{
				UserID:   1,
				UserType: "courier",
			},
		},
	}

	for _, currentTest := range tests {
		test := currentTest

		t.Run(test.name, func(t *testing.T) {
			reqBody, err := v1.Encode(test.accountDataList)
			if err != nil {
				t.Fatal(err)
			}

			resp, err := http.Post(baseAddr+"/v1/accounts", "application/json", bytes.NewBuffer(reqBody))
			if err != nil {
				t.Fatal(err)
			}

			if resp.StatusCode != http.StatusOK {
				t.Fatalf("Response status: %d", resp.StatusCode)
			}

			gotAccount := accountingapi.AccountResponse{}
			if err = accountingapi.DecodeJSON(resp.Body, &gotAccount); err != nil {
				t.Fatal(err)
			}

			accountIDStr := strconv.Itoa(gotAccount.ID)

			resp2, err := http.Get("/v1/accounts/" + accountIDStr)
			if err != nil {
				t.Fatal(err)
			}

			if resp2.StatusCode != http.StatusOK {
				t.Fatalf("StatusCode: %d", resp2.StatusCode)
			}

			insertedAccount := accountingapi.AccountResponse{}
			if err = accountingapi.DecodeJSON(resp2.Body, &insertedAccount); err != nil {
				t.Fatal(err)
			}

			if insertedAccount.ID != gotAccount.ID {
				t.Errorf("ID: Expected: %v, Got: %v", gotAccount.ID, insertedAccount.ID)
			}

			if insertedAccount.UserID != gotAccount.UserID {
				t.Errorf("UserID: Expected: %v, Got: %v", gotAccount.UserID, insertedAccount.UserID)
			}

			if insertedAccount.UserType != gotAccount.UserType {
				t.Errorf("UserType: Expected: %s, Got: %s", gotAccount.UserType, insertedAccount.UserType)
			}

			if insertedAccount.Balance != gotAccount.Balance {
				t.Errorf("Balance: Expected: %v, Got: %v", gotAccount.Balance, insertedAccount.Balance)
			}

			// Deleting instance.
			deleter := http.DefaultClient

			delReq, err := http.NewRequest(http.MethodDelete,
				baseAddr+"/v1/accounts/"+strconv.Itoa(gotAccount.ID), nil)
			if err != nil {
				t.Error(err)
			}

			_, err = deleter.Do(delReq)
			if err != nil {
				t.Errorf("Could not delete: %v", err)
			}
		})
	}
}

func TestGetAccountListEndpoint(t *testing.T) {
	tests := []struct {
		name                string
		accountInsertData   accountingapi.NewAccountRequest
		accountResponseList accountingapi.AccountListResponse
	}{
		{
			"TestGetAccountListEndpointSuccess",
			accountingapi.NewAccountRequest{
				UserID:   1,
				UserType: "courier",
			},
			accountingapi.AccountListResponse{
				AccountList: []accountingapi.AccountResponse{
					{
						UserID:    1,
						UserType:  "courier",
						Balance:   50,
						CreatedAt: time.Time{},
						UpdatedAt: time.Time{},
					},
				},
			},
		},
	}

	for _, currentTest := range tests {
		test := currentTest

		t.Run(test.name, func(t *testing.T) {
			reqBody, err := v1.Encode(test.accountInsertData)
			if err != nil {
				t.Fatal(err)
			}

			resp, err := http.Post(baseAddr+"/v1/accounts", "application/json", bytes.NewBuffer(reqBody))
			if err != nil {
				t.Fatal(err)
			}

			if resp.StatusCode != http.StatusOK {
				t.Fatalf("Response status: %d", resp.StatusCode)
			}

			resp2, err := http.Get("/v1/accounts")
			if err != nil {
				t.Fatal(err)
			}

			if resp2.StatusCode != http.StatusOK {
				t.Fatalf("StatusCode: %d", resp2.StatusCode)
			}

			gotAccountList := accountingapi.AccountListResponse{}
			if err = accountingapi.DecodeJSON(resp.Body, &gotAccountList); err != nil {
				t.Fatal(err)
			}

			if len(gotAccountList.AccountList) != len(test.accountResponseList.AccountList) {
				t.Errorf("len: Expected: %v, Got: %v", len(test.accountResponseList.AccountList), len(gotAccountList.AccountList))
			}

			for _, gotAccount := range gotAccountList.AccountList {
				for _, testAccountResp := range test.accountResponseList.AccountList {

					if gotAccount.ID != testAccountResp.ID {
						t.Errorf("ID: Expected: %v, Got: %v", testAccountResp.ID, gotAccount.ID)
					}

					if gotAccount.UserID != testAccountResp.UserID {
						t.Errorf("UserID: Expected: %v, Got: %v", testAccountResp.UserID, gotAccount.UserID)
					}

					if gotAccount.UserType != testAccountResp.UserType {
						t.Errorf("UserType: Expected: %s, Got: %s", testAccountResp.UserType, gotAccount.UserType)
					}

					if gotAccount.Balance != testAccountResp.Balance {
						t.Errorf("Balance: Expected: %v, Got: %v", testAccountResp.Balance, gotAccount.Balance)
					}
				}

				// Deleting instance.
				deleter := http.DefaultClient

				delReq, err := http.NewRequest(http.MethodDelete,
					baseAddr+"/v1/accounts/"+strconv.Itoa(gotAccount.ID), nil)
				if err != nil {
					t.Error(err)
				}

				_, err = deleter.Do(delReq)
				if err != nil {
					t.Errorf("Could not delete: %v", err)
				}
			}
		})
	}
}

func TestInsertTransactionsEndpointSuccess(t *testing.T) {
	t.Parallel()
	type test struct {
		name                string
		transactionRequest  accountingapi.TransactionRequest
		transactionResponse accountingapi.TransactionResponse
	}
	tests := []test{
		{
			name: "from account to account",
			transactionRequest: accountingapi.TransactionRequest{
				FromAccountID: 1,
				ToAccountID:   2,
				Amount:        50,
			},
			transactionResponse: accountingapi.TransactionResponse{
				ID:            1,
				FromAccountID: 1,
				ToAccountID:   2,
				Amount:        50,
				CreatedAt:     time.Now(),
				UpdatedAt:     time.Now(),
				Valid:         true,
			},
		},
		{
			name: "from account",
			transactionRequest: accountingapi.TransactionRequest{
				FromAccountID: 1,
				Amount:        50,
			},
			transactionResponse: accountingapi.TransactionResponse{
				ID:            1,
				FromAccountID: 1,
				ToAccountID:   0,
				Amount:        50,
				CreatedAt:     time.Now(),
				UpdatedAt:     time.Now(),
				Valid:         true,
			},
		},
		{
			name: "to account",
			transactionRequest: accountingapi.TransactionRequest{
				ToAccountID: 2,
				Amount:      50,
			},
			transactionResponse: accountingapi.TransactionResponse{
				ID:            1,
				FromAccountID: 0,
				ToAccountID:   2,
				Amount:        50,
				CreatedAt:     time.Now(),
				UpdatedAt:     time.Now(),
				Valid:         true,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			reqBody, err := v1.Encode(test)
			if err != nil {
				t.Fatal(err)
			}

			resp, err := http.Post(http.MethodPost, "/v1/transactions", bytes.NewBuffer(reqBody))
			if err != nil {
				t.Fatal(err)
			}

			if resp.StatusCode != http.StatusOK {
				t.Fatalf("StatusCode: %d", resp.StatusCode)
			}

			respData := accountingapi.TransactionResponse{}
			if err := accountingapi.DecodeJSON(resp.Body, &respData); err != nil {
				t.Fatal(err)
			}

			if respData.ID != test.transactionResponse.ID {
				t.Errorf("ID: Expected: %v, Got: %v", test.transactionResponse.ID, respData.ID)
			}

			if respData.FromAccountID != test.transactionResponse.FromAccountID {
				t.Errorf("FromAccountID: Expected: %v, Got: %v", test.transactionResponse.FromAccountID, respData.FromAccountID)
			}

			if respData.ToAccountID != test.transactionResponse.ToAccountID {
				t.Errorf("ToAccountID: Expected: %v, Got: %v", test.transactionResponse.ToAccountID, respData.ToAccountID)
			}

			if respData.Amount != test.transactionResponse.Amount {
				t.Errorf("Amount: Expected: %v, Got: %v", test.transactionResponse.Amount, respData.Amount)
			}

			if respData.Valid != test.transactionResponse.Valid {
				t.Errorf("Valid: Expected: %v, Got: %v", test.transactionResponse.Valid, respData.Valid)
			}

			if respData.CreatedAt != test.transactionResponse.CreatedAt {
				t.Errorf("CreatedAt: Expected: %v, Got: %v", test.transactionResponse.CreatedAt, respData.CreatedAt)
			}

			if respData.UpdatedAt != test.transactionResponse.UpdatedAt {
				t.Errorf("UpdatedAt: Expected: %v, Got: %v", test.transactionResponse.UpdatedAt, respData.UpdatedAt)
			}

			// Deleting instance.
			deleter := http.DefaultClient

			delReq, err := http.NewRequest(http.MethodDelete,
				baseAddr+"/v1/transactions/"+strconv.Itoa(respData.ID), nil)
			if err != nil {
				t.Error(err)
			}

			_, err = deleter.Do(delReq)
			if err != nil {
				t.Errorf("Could not delete: %v", err)
			}
		})
	}
}
