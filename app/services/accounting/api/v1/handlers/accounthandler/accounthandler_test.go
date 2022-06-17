package accounthandler_test

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

	"github.com/nndergunov/delivryApp/app/services/accounting/api/v1/handlers/accounthandler"
	"github.com/nndergunov/delivryApp/app/services/accounting/pkg/domain"
)

var (
	MockAccountData = &domain.Account{
		UserID:    1,
		UserType:  "courier",
		Balance:   50,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	MockTransactionData = &domain.Transaction{
		ID:            1,
		FromAccountID: 1,
		ToAccountID:   2,
		Amount:        50,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
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

func (m MockService) Transact(transaction domain.Transaction) (*domain.Transaction, error) {
	return MockTransactionData, nil
}

func TestInsertNewAccountEndpoint(t *testing.T) {
	t.Parallel()

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
			t.Parallel()

			mockService := new(MockService)

			log := logger.NewLogger(os.Stdout, test.name)
			handler := accounthandler.NewAccountHandler(accounthandler.Params{
				Logger:         log,
				AccountService: mockService,
			})

			reqBody, err := v1.Encode(test.accountData)
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

			if respData.ID != MockAccountData.ID {
				t.Errorf("ID: Expected: %v, Got: %v", MockAccountData.ID, respData.ID)
			}

			if respData.UserID != MockAccountData.UserID {
				t.Errorf("UserID: Expected: %v, Got: %v", MockAccountData.UserID, respData.UserID)
			}

			if respData.UserType != MockAccountData.UserType {
				t.Errorf("UserType: Expected: %s, Got: %s", MockAccountData.UserType, respData.UserType)
			}

			if respData.Balance != MockAccountData.Balance {
				t.Errorf("Balance: Expected: %v, Got: %v", MockAccountData.Balance, respData.Balance)
			}
		})
	}
}

func TestGetAccountEndpoint(t *testing.T) {
	t.Parallel()

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
			t.Parallel()

			mockService := new(MockService)

			log := logger.NewLogger(os.Stdout, test.name)
			handler := accounthandler.NewAccountHandler(accounthandler.Params{
				Logger:         log,
				AccountService: mockService,
			})

			reqBody, err := v1.Encode(test.accountData)
			if err != nil {
				t.Fatal(err)
			}

			resp := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPost, "/v1/accounts", bytes.NewBuffer(reqBody))

			handler.ServeHTTP(resp, req)

			if resp.Code != http.StatusOK {
				t.Fatalf("StatusCode: %d", resp.Code)
			}

			insertedAccount := accountingapi.AccountResponse{}
			if err = accountingapi.DecodeJSON(resp.Body, &insertedAccount); err != nil {
				t.Fatal(err)
			}

			accountIDStr := strconv.Itoa(insertedAccount.ID)

			resp2 := httptest.NewRecorder()
			req2 := httptest.NewRequest(http.MethodGet, "/v1/accounts/"+accountIDStr, nil)

			handler.ServeHTTP(resp2, req2)

			if resp2.Code != http.StatusOK {
				t.Fatalf("StatusCode: %d", resp2.Code)
			}

			gotAccount := accountingapi.AccountResponse{}
			if err = accountingapi.DecodeJSON(resp.Body, &gotAccount); err != nil {
				t.Fatal(err)
			}

			if gotAccount.ID != MockAccountData.ID {
				t.Errorf("ID: Expected: %v, Got: %v", MockAccountData.ID, gotAccount.ID)
			}

			if gotAccount.UserID != MockAccountData.UserID {
				t.Errorf("UserID: Expected: %v, Got: %v", MockAccountData.UserID, gotAccount.UserID)
			}

			if gotAccount.UserType != MockAccountData.UserType {
				t.Errorf("UserType: Expected: %s, Got: %s", MockAccountData.UserType, gotAccount.UserType)
			}

			if gotAccount.Balance != MockAccountData.Balance {
				t.Errorf("Balance: Expected: %v, Got: %v", MockAccountData.Balance, gotAccount.Balance)
			}
		})
	}
}
