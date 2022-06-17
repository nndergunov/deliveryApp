package accountservice_test

import (
	"bytes"
	"net/http"
	"strconv"
	"testing"

	"github.com/nndergunov/deliveryApp/app/pkg/api/v1/accountingapi"

	"github.com/nndergunov/deliveryApp/app/pkg/api/v1"
)

const baseAddr = "http://localhost:8081"

func TestInsertAccountEndpoint(t *testing.T) {
	tests := []struct {
		name        string
		accountData accountingapi.NewAccountRequest
	}{
		{
			"Insert account simple test",
			accountingapi.NewAccountRequest{
				UserID:   1,
				UserType: "courier",
			},
		},
	}

	for _, currentTest := range tests {
		test := currentTest

		t.Run(test.name, func(t *testing.T) {
			reqBody, err := v1.Encode(test.accountData)
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

			if err := resp.Body.Close(); err != nil {
				t.Error(err)
			}

			if respData.ID < 1 {
				t.Errorf("ID: Expected : > 1, Got: %v", respData.ID)
			}

			if respData.Balance != 0 {
				t.Errorf("Balance: Expected: %v, Got: %v", 0, respData.Balance)
			}

			if respData.UserID != test.accountData.UserID {
				t.Errorf("UserID: Expected: %v, Got: %v", test.accountData.UserID, respData.UserID)
			}

			if respData.UserType != test.accountData.UserType {
				t.Errorf("UserType: Expected: %s, Got: %s", test.accountData.UserType, respData.UserType)
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
		name        string
		accountData accountingapi.NewAccountRequest
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
			reqBody, err := v1.Encode(test.accountData)
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
		name        string
		accountData accountingapi.NewAccountRequest
	}{
		{
			"GetAccount list simple test",
			accountingapi.NewAccountRequest{
				UserID:   1,
				UserType: "courier",
			},
		},
	}

	for _, currentTest := range tests {
		test := currentTest

		t.Run(test.name, func(t *testing.T) {
			reqBody, err := v1.Encode(test.accountData)
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
