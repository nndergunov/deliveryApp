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
				Balance:   0,
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

			resp, err := http.Post(baseAddr+"/v1/accounts", "application/json", bytes.NewBuffer(reqBody))
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

			resp2, err := http.Get(baseAddr + "/v1/accounts/" + accountIDStr)
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
						Balance:   0,
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

			resp2, err := http.Get(baseAddr + "/v1/accounts?" + "user_id=" + strconv.Itoa(test.accountInsertData.UserID) +
				"&user_type=" + test.accountInsertData.UserType)
			if err != nil {
				t.Fatal(err)
			}

			if resp2.StatusCode != http.StatusOK {
				t.Fatalf("StatusCode: %d", resp2.StatusCode)
			}

			gotAccountList := accountingapi.AccountListResponse{}
			if err = accountingapi.DecodeJSON(resp2.Body, &gotAccountList); err != nil {
				t.Fatal(err)
			}

			if len(gotAccountList.AccountList) != len(test.accountResponseList.AccountList) {
				t.Errorf("len: Expected: %v, Got: %v", len(test.accountResponseList.AccountList), len(gotAccountList.AccountList))
			}

			for _, gotAccount := range gotAccountList.AccountList {
				for _, testAccountResp := range test.accountResponseList.AccountList {

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

func TestInsertTransactionsFromAccountToAccountEndpointSuccess(t *testing.T) {
	t.Parallel()
	type test struct {
		name                string
		fromAccount         accountingapi.NewAccountRequest
		toAccount           accountingapi.NewAccountRequest
		transactionRequest  accountingapi.TransactionRequest
		transactionResponse accountingapi.TransactionResponse
	}
	tests := []test{
		{
			name: "from account to account",
			fromAccount: accountingapi.NewAccountRequest{
				UserID:   1,
				UserType: "consumer",
			},

			toAccount: accountingapi.NewAccountRequest{
				UserID:   1,
				UserType: "courier",
			},

			transactionRequest: accountingapi.TransactionRequest{
				Amount: 50,
			},

			transactionResponse: accountingapi.TransactionResponse{
				Amount: 50,
				Valid:  true,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// insert from account1
			account1ReqBody, err := v1.Encode(test.fromAccount)
			if err != nil {
				t.Fatal(err)
			}

			account1Resp, err := http.Post(baseAddr+"/v1/accounts", "application/json", bytes.NewBuffer(account1ReqBody))
			if err != nil {
				t.Fatal(err)
			}

			if account1Resp.StatusCode != http.StatusOK {
				t.Fatalf("StatusCode: %d", account1Resp.StatusCode)
			}

			account1RespData := accountingapi.AccountResponse{}
			if err := accountingapi.DecodeJSON(account1Resp.Body, &account1RespData); err != nil {
				t.Fatal(err)
			}

			// add balance to account1
			test.transactionRequest.ToAccountID = account1RespData.ID

			trAddBalanceAccount1ReqBody, err := v1.Encode(test.transactionRequest)
			if err != nil {
				t.Fatal(err)
			}

			trAddBalanceAccount1Resp, err := http.Post(baseAddr+"/v1/transactions", "application/json", bytes.NewBuffer(trAddBalanceAccount1ReqBody))
			if err != nil {
				t.Fatal(err)
			}

			if trAddBalanceAccount1Resp.StatusCode != http.StatusOK {
				t.Fatalf("StatusCode: %d", account1Resp.StatusCode)
			}

			trAddBalanceAccount1RespData := accountingapi.TransactionResponse{}
			if err := accountingapi.DecodeJSON(trAddBalanceAccount1Resp.Body, &trAddBalanceAccount1RespData); err != nil {
				t.Fatal(err)
			}

			if trAddBalanceAccount1RespData.Amount != test.transactionResponse.Amount {
				t.Errorf("trAddBalanceAccount1RespData: Expected: %v, Got: %v", test.transactionResponse.Amount, trAddBalanceAccount1RespData.Amount)
			}

			// insert account 2
			account2ReqBody2, err := v1.Encode(test.toAccount)
			if err != nil {
				t.Fatal(err)
			}

			respAccount2, err := http.Post(baseAddr+"/v1/accounts", "application/json", bytes.NewBuffer(account2ReqBody2))
			if err != nil {
				t.Fatal(err)
			}

			if respAccount2.StatusCode != http.StatusOK {
				t.Fatalf("StatusCode: %d", respAccount2.StatusCode)
			}

			account2RespData := accountingapi.AccountResponse{}
			if err := accountingapi.DecodeJSON(respAccount2.Body, &account2RespData); err != nil {
				t.Fatal(err)
			}

			// transaction from account 1 to account 2
			test.transactionRequest.FromAccountID = account1RespData.ID
			test.transactionRequest.ToAccountID = account2RespData.ID

			test.transactionResponse.FromAccountID = account1RespData.ID
			test.transactionResponse.ToAccountID = account2RespData.ID

			trReqBody, err := v1.Encode(test.transactionRequest)
			if err != nil {
				t.Fatal(err)
			}

			trResp, err := http.Post(baseAddr+"/v1/transactions", "application/json", bytes.NewBuffer(trReqBody))
			if err != nil {
				t.Fatal(err)
			}

			if trResp.StatusCode != http.StatusOK {
				t.Fatalf("StatusCode: %d", trResp.StatusCode)
			}

			trRespData := accountingapi.TransactionResponse{}
			if err := accountingapi.DecodeJSON(trResp.Body, &trRespData); err != nil {
				t.Fatal(err)
			}

			if trRespData.FromAccountID != test.transactionResponse.FromAccountID {
				t.Errorf("FromAccountID: Expected: %v, Got: %v", test.transactionResponse.FromAccountID, trRespData.FromAccountID)
			}

			if trRespData.ToAccountID != test.transactionResponse.ToAccountID {
				t.Errorf("ToAccountID: Expected: %v, Got: %v", test.transactionResponse.ToAccountID, trRespData.ToAccountID)
			}

			if trRespData.Amount != test.transactionResponse.Amount {
				t.Errorf("Amount: Expected: %v, Got: %v", test.transactionResponse.Amount, trRespData.Amount)
			}

			if trRespData.Valid != test.transactionResponse.Valid {
				t.Errorf("Valid: Expected: %v, Got: %v", test.transactionResponse.Valid, trRespData.Valid)
			}

			// Deleting instance.
			deleter := http.DefaultClient

			delReq, err := http.NewRequest(http.MethodDelete,
				baseAddr+"/v1/accounts/"+strconv.Itoa(account1RespData.ID), nil)
			if err != nil {
				t.Error(err)
			}

			_, err = deleter.Do(delReq)
			if err != nil {
				t.Errorf("Could not delete: %v", err)
			}

			// Deleting trAddBalance.
			deleter2 := http.DefaultClient

			delReq2, err := http.NewRequest(http.MethodDelete,
				baseAddr+"/v1/transactions/"+strconv.Itoa(trAddBalanceAccount1RespData.ID), nil)
			if err != nil {
				t.Error(err)
			}

			_, err = deleter2.Do(delReq2)
			if err != nil {
				t.Errorf("Could not delete: %v", err)
			}

			// Deleting account 2.
			deleter3 := http.DefaultClient

			delReq3, err := http.NewRequest(http.MethodDelete,
				baseAddr+"/v1/accounts/"+strconv.Itoa(account2RespData.ID), nil)
			if err != nil {
				t.Error(err)
			}

			_, err = deleter3.Do(delReq3)
			if err != nil {
				t.Errorf("Could not delete: %v", err)
			}

			// Deleting trFromAccount1ToAccount2.
			deleter4 := http.DefaultClient

			delReq4, err := http.NewRequest(http.MethodDelete,
				baseAddr+"/v1/transactions/"+strconv.Itoa(trRespData.ID), nil)
			if err != nil {
				t.Error(err)
			}

			_, err = deleter4.Do(delReq4)
			if err != nil {
				t.Errorf("Could not delete: %v", err)
			}
		})
	}
}
