package accountinghandler_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/nndergunov/deliveryApp/app/pkg/api/v1/accountingapi"

	"github.com/nndergunov/deliveryApp/app/pkg/api/v1"
	"github.com/nndergunov/deliveryApp/app/pkg/logger"

	"github.com/nndergunov/deliveryApp/app/services/accounting/api/v1/handlers/accountinghandler"
	"github.com/nndergunov/deliveryApp/app/services/accounting/pkg/domain"

)

var (
	MockAccountData = &domain.Account{
		UserID:    1,
		UserType:  "courier",
		Balance:   50,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}

	MockTransactionFromAccountToAccountData = &domain.Transaction{
		ID:            1,
		FromAccountID: 1,
		ToAccountID:   2,
		Amount:        50,
		CreatedAt:     time.Time{},
		UpdatedAt:     time.Time{},
		Valid:         true,
	}

	MockTransactionFromAccountData = &domain.Transaction{
		ID:            1,
		FromAccountID: 1,
		ToAccountID:   0,
		Amount:        50,
		CreatedAt:     time.Time{},
		UpdatedAt:     time.Time{},
		Valid:         true,
	}

	MockTransactionToAccountData = &domain.Transaction{
		ID:            1,
		FromAccountID: 0,
		ToAccountID:   2,
		Amount:        50,
		CreatedAt:     time.Time{},
		UpdatedAt:     time.Time{},
		Valid:         true,
	}
)

type MockService struct{}

func (m MockService) InsertNewAccount(_ domain.Account) (*domain.Account, error) {
	return MockAccountData, nil
}

func (m MockService) GetAccountByID(_ string) (*domain.Account, error) {
	return MockAccountData, nil
}

func (m MockService) GetAccountListByParam(_ domain.SearchParam) ([]domain.Account, error) {
	return []domain.Account{*MockAccountData}, nil
}

func (m MockService) DeleteAccount(_ string) (string, error) {
	return "account deleted", nil
}

func (m MockService) InsertTransaction(transaction domain.Transaction) (*domain.Transaction, error) {
	if transaction.FromAccountID == 0 {
		return MockTransactionToAccountData, nil
	}
	if transaction.ToAccountID == 0 {
		return MockTransactionFromAccountData, nil
	}
	return MockTransactionFromAccountToAccountData, nil
}

func (m MockService) DeleteTransaction(id string) (string, error) {
	return "transaction deleted", nil
}

func TestInsertNewAccountEndpointSuccess(t *testing.T) {
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

			mockService := new(MockService)

			log := logger.NewLogger(os.Stdout, test.name)
			handler := accountinghandler.NewHandler(accountinghandler.Params{

				Logger:         log,
				AccountService: mockService,
			})

			reqBody, err := v1.Encode(test.newAccountRequest)
			if err != nil {
				t.Fatal(err)
			}

			resp := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPost, "/v1/accounts", bytes.NewBuffer(reqBody))

			handler.ServeHTTP(resp, req)

			if resp.Code != http.StatusOK {
				t.Fatalf("StatusCode: %d", resp.Code)
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
		})
	}
}

func TestGetAccountEndpointSuccess(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name            string
		accountResponse accountingapi.AccountResponse
	}{
		{
			"get account simple test",
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

			mockService := new(MockService)

			log := logger.NewLogger(os.Stdout, test.name)
			handler := accountinghandler.NewHandler(accountinghandler.Params{

				Logger:         log,
				AccountService: mockService,
			})

			accountIDStr := strconv.Itoa(MockAccountData.ID)

			resp := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, "/v1/accounts/"+accountIDStr, nil)

			handler.ServeHTTP(resp, req)

			if resp.Code != http.StatusOK {
				t.Fatalf("StatusCode: %d", resp.Code)
			}

			respData := accountingapi.AccountResponse{}
			if err := accountingapi.DecodeJSON(resp.Body, &respData); err != nil {
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
		})
	}
}

func TestGetAccountListEndpointSuccess(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name                string
		accountResponseList accountingapi.AccountListResponse
	}{
		{
			"TestGetAccountListEndpointSuccess",
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
			t.Parallel()

			mockService := new(MockService)

			log := logger.NewLogger(os.Stdout, test.name)
			handler := accountinghandler.NewHandler(accountinghandler.Params{

				Logger:         log,
				AccountService: mockService,
			})

			resp := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, "/v1/accounts", nil)

			handler.ServeHTTP(resp, req)

			if resp.Code != http.StatusOK {
				t.Fatalf("StatusCode: %d", resp.Code)
			}

			respDataList := accountingapi.AccountListResponse{}
			if err := accountingapi.DecodeJSON(resp.Body, &respDataList); err != nil {
				t.Fatal(err)
			}

			if len(respDataList.AccountList) < len(test.accountResponseList.AccountList) {
				t.Errorf("len: Expected: %v, Got: %v", len(test.accountResponseList.AccountList), len(respDataList.AccountList))
			}

			for _, respData := range respDataList.AccountList {
				for _, testAccountResponse := range test.accountResponseList.AccountList {

					if respData.ID != testAccountResponse.ID {
						t.Errorf("ID: Expected: %v, Got: %v", testAccountResponse.ID, respData.ID)
					}

					if respData.UserID != testAccountResponse.UserID {
						t.Errorf("UserID: Expected: %v, Got: %v", testAccountResponse.UserID, respData.UserID)
					}

					if respData.UserType != testAccountResponse.UserType {
						t.Errorf("UserType: Expected: %s, Got: %s", testAccountResponse.UserType, respData.UserType)
					}

					if respData.Balance != testAccountResponse.Balance {
						t.Errorf("Balance: Expected: %v, Got: %v", testAccountResponse.Balance, respData.Balance)
					}
				}
			}
		})
	}
}

func TestDeleteEndpointSuccess(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name                   string
		accountDeletedResponse string
	}{
		{
			"delete account simple test",
			"account deleted",
		},
	}
	for _, currentTest := range tests {
		test := currentTest

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			mockService := new(MockService)

			log := logger.NewLogger(os.Stdout, test.name)
			handler := accountinghandler.NewHandler(accountinghandler.Params{

				Logger:         log,
				AccountService: mockService,
			})

			resp := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodDelete, "/v1/accounts/1", nil)

			handler.ServeHTTP(resp, req)

			if resp.Code != http.StatusOK {
				t.Fatalf("StatusCode: %d", resp.Code)
			}

			accountDeletedResponse := ""
			if err := accountingapi.DecodeJSON(resp.Body, &accountDeletedResponse); err != nil {
				t.Fatal(err)
			}

			if accountDeletedResponse != test.accountDeletedResponse {
				t.Errorf("accountDeletedResponse: Expected: %v, Got: %v", test.accountDeletedResponse, accountDeletedResponse)
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
				CreatedAt:     time.Time{},
				UpdatedAt:     time.Time{},
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
				CreatedAt:     time.Time{},
				UpdatedAt:     time.Time{},
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
				CreatedAt:     time.Time{},
				UpdatedAt:     time.Time{},
				Valid:         true,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			mockService := new(MockService)

			log := logger.NewLogger(os.Stdout, test.name)
			handler := accountinghandler.NewHandler(accountinghandler.Params{

				Logger:         log,
				AccountService: mockService,
			})

			reqBody, err := v1.Encode(test)
			if err != nil {
				t.Fatal(err)
			}

			resp := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPost, "/v1/transactions", bytes.NewBuffer(reqBody))

			handler.ServeHTTP(resp, req)

			if resp.Code != http.StatusOK {
				t.Fatalf("StatusCode: %d", resp.Code)
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
		})
	}
}
