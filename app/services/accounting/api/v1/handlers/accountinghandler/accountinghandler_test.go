package accountinghandler_test

import (
	"bytes"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/nndergunov/deliveryApp/app/pkg/api/v1/accountingapi"

	"github.com/nndergunov/deliveryApp/app/pkg/api/v1"
	"github.com/nndergunov/deliveryApp/app/pkg/logger"

	"github.com/nndergunov/deliveryApp/app/services/accounting/api/v1/handlers/accountinghandler"
	"github.com/nndergunov/deliveryApp/app/services/accounting/pkg/domain"
	mockservice "github.com/nndergunov/deliveryApp/app/services/accounting/pkg/mocks"
)

func TestInsertNewAccountEndpointSuccess(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		in   accountingapi.NewAccountRequest
		out  accountingapi.AccountResponse
	}{
		{
			"Insert account test",
			accountingapi.NewAccountRequest{
				UserID:   1,
				UserType: "courier",
			},
			accountingapi.AccountResponse{
				ID:       1,
				UserID:   1,
				UserType: "courier",
				Balance:  50,
			},
		},
	}

	for _, currentTest := range tests {
		test := currentTest

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			ctl := gomock.NewController(t)
			service := mockservice.NewMockAccountService(ctl)

			mockInData := domain.Account{
				UserID:   test.in.UserID,
				UserType: test.in.UserType,
			}

			mockOutData := &domain.Account{
				ID:       test.out.ID,
				UserID:   test.out.UserID,
				UserType: test.out.UserType,
				Balance:  test.out.Balance,
			}

			service.EXPECT().InsertNewAccount(mockInData).Return(mockOutData, nil)

			handler := accountinghandler.NewHandler(accountinghandler.Params{
				Logger:         logger.NewLogger(os.Stdout, test.name),
				AccountService: service,
			})

			reqBody, err := v1.Encode(test.in)
			require.NoError(t, err)

			resp := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPost, "/v1/accounts", bytes.NewBuffer(reqBody))

			handler.ServeHTTP(resp, req)

			if resp.Code != http.StatusOK {
				t.Errorf("StatusCode: %d", resp.Code)
			}

			respData := accountingapi.AccountResponse{}
			err = accountingapi.DecodeJSON(resp.Body, &respData)
			require.NoError(t, err)

			if respData.ID != test.out.ID {
				t.Errorf("ID: Expected: %v, Got: %v", test.out.ID, respData.ID)
			}

			if respData.UserID != test.out.UserID {
				t.Errorf("UserID: Expected: %v, Got: %v", test.out.UserID, respData.UserID)
			}

			if respData.UserType != test.out.UserType {
				t.Errorf("UserType: Expected: %s, Got: %s", test.out.UserType, respData.UserType)
			}

			if respData.Balance != test.out.Balance {
				t.Errorf("Balance: Expected: %v, Got: %v", test.out.Balance, respData.Balance)
			}
		})
	}
}

func TestGetAccountEndpointSuccess(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		in   string
		out  accountingapi.AccountResponse
	}{
		{
			"get account test",
			"1",
			accountingapi.AccountResponse{
				ID:       1,
				UserID:   1,
				UserType: "courier",
				Balance:  50,
			},
		},
	}

	for _, currentTest := range tests {
		test := currentTest

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			ctl := gomock.NewController(t)
			service := mockservice.NewMockAccountService(ctl)

			mockOutData := &domain.Account{
				ID:       test.out.ID,
				UserID:   test.out.UserID,
				UserType: test.out.UserType,
				Balance:  test.out.Balance,
			}

			service.EXPECT().GetAccountByID(test.in).Return(mockOutData, nil)

			handler := accountinghandler.NewHandler(accountinghandler.Params{
				Logger:         logger.NewLogger(os.Stdout, test.name),
				AccountService: service,
			})

			resp := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, "/v1/accounts/"+test.in, nil)

			handler.ServeHTTP(resp, req)

			if resp.Code != http.StatusOK {
				t.Errorf("StatusCode: %d", resp.Code)
			}

			respData := accountingapi.AccountResponse{}
			err := accountingapi.DecodeJSON(resp.Body, &respData)
			require.NoError(t, err)

			if respData.ID != test.out.ID {
				t.Errorf("ID: Expected: %v, Got: %v", test.out.ID, respData.ID)
			}

			if respData.UserID != test.out.UserID {
				t.Errorf("UserID: Expected: %v, Got: %v", test.out.UserID, respData.UserID)
			}

			if respData.UserType != test.out.UserType {
				t.Errorf("UserType: Expected: %s, Got: %s", test.out.UserType, respData.UserType)
			}

			if respData.Balance != test.out.Balance {
				t.Errorf("Balance: Expected: %v, Got: %v", test.out.Balance, respData.Balance)
			}
		})
	}
}

func TestGetAccountListEndpointSuccess(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		in   domain.SearchParam
		out  accountingapi.AccountListResponse
	}{
		{
			"Get AccountList test",
			domain.SearchParam{"user_id": "1", "user_type": "courier"},
			accountingapi.AccountListResponse{
				AccountList: []accountingapi.AccountResponse{
					{
						ID:       1,
						UserID:   1,
						UserType: "courier",
						Balance:  50,
					},
					{
						ID:       2,
						UserID:   1,
						UserType: "courier",
						Balance:  30,
					},
				},
			},
		},
	}
	for _, currentTest := range tests {
		test := currentTest

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			ctl := gomock.NewController(t)
			service := mockservice.NewMockAccountService(ctl)

			var mockOutDataList []domain.Account
			for _, mockOutData := range test.out.AccountList {

				account := domain.Account{
					ID:       mockOutData.ID,
					UserID:   mockOutData.UserID,
					UserType: mockOutData.UserType,
					Balance:  mockOutData.Balance,
				}
				mockOutDataList = append(mockOutDataList, account)
			}

			service.EXPECT().GetAccountListByParam(test.in).Return(mockOutDataList, nil)

			handler := accountinghandler.NewHandler(accountinghandler.Params{
				Logger:         logger.NewLogger(os.Stdout, test.name),
				AccountService: service,
			})

			resp := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, "/v1/accounts?user_id="+test.in["user_id"]+"&"+"user_type="+test.in["user_type"], nil)

			handler.ServeHTTP(resp, req)

			if resp.Code != http.StatusOK {
				t.Errorf("StatusCode: %d", resp.Code)
			}

			respDataList := accountingapi.AccountListResponse{}
			err := accountingapi.DecodeJSON(resp.Body, &respDataList)
			require.NoError(t, err)

			if len(respDataList.AccountList) < len(test.out.AccountList) {
				t.Errorf("len: Expected: %v, Got: %v", len(test.out.AccountList), len(respDataList.AccountList))
			}

			for i, respData := range respDataList.AccountList {
				if respData.ID != test.out.AccountList[i].ID {
					t.Errorf("ID: Expected: %v, Got: %v", test.out.AccountList[i].ID, respData.ID)
				}

				if respData.UserID != test.out.AccountList[i].UserID {
					t.Errorf("UserID: Expected: %v, Got: %v", test.out.AccountList[i].UserID, respData.UserID)
				}

				if respData.UserType != test.out.AccountList[i].UserType {
					t.Errorf("UserType: Expected: %s, Got: %s", test.out.AccountList[i].UserType, respData.UserType)
				}

				if respData.Balance != test.out.AccountList[i].Balance {
					t.Errorf("Balance: Expected: %v, Got: %v", test.out.AccountList[i].Balance, respData.Balance)
				}
			}
		})
	}
}

func TestDeleteEndpointSuccess(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		in   string
		out  string
	}{
		{
			"delete account test",
			"1",
			"account deleted",
		},
	}
	for _, currentTest := range tests {
		test := currentTest

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			ctl := gomock.NewController(t)
			service := mockservice.NewMockAccountService(ctl)

			service.EXPECT().DeleteAccount(test.in).Return(test.out, nil)

			handler := accountinghandler.NewHandler(accountinghandler.Params{
				Logger:         logger.NewLogger(os.Stdout, test.name),
				AccountService: service,
			})

			resp := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodDelete, "/v1/accounts/"+test.in, nil)

			handler.ServeHTTP(resp, req)

			if resp.Code != http.StatusOK {
				t.Errorf("StatusCode: %d", resp.Code)
			}

			accountDeletedResponse := ""
			err := accountingapi.DecodeJSON(resp.Body, &accountDeletedResponse)
			require.NoError(t, err)

			if accountDeletedResponse != test.out {
				t.Errorf("out: Expected: %v, Got: %v", test.out, accountDeletedResponse)
			}
		})
	}
}

func TestInsertTransactionsEndpointSuccess(t *testing.T) {
	t.Parallel()
	type test struct {
		name string
		in   accountingapi.TransactionRequest
		out  accountingapi.TransactionResponse
	}
	tests := []test{
		{
			name: "from account to account",
			in: accountingapi.TransactionRequest{
				FromAccountID: 1,
				ToAccountID:   2,
				Amount:        50,
			},
			out: accountingapi.TransactionResponse{
				ID:            1,
				FromAccountID: 1,
				ToAccountID:   2,
				Amount:        50,
				Valid:         true,
			},
		},
		{
			name: "from account",
			in: accountingapi.TransactionRequest{
				FromAccountID: 1,
				Amount:        50,
			},
			out: accountingapi.TransactionResponse{
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
			in: accountingapi.TransactionRequest{
				ToAccountID: 2,
				Amount:      50,
			},
			out: accountingapi.TransactionResponse{
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

			ctl := gomock.NewController(t)
			service := mockservice.NewMockAccountService(ctl)

			mockInData := domain.Transaction{
				FromAccountID: test.in.FromAccountID,
				ToAccountID:   test.in.ToAccountID,
				Amount:        test.in.Amount,
			}

			mockOutData := &domain.Transaction{
				ID:            test.out.ID,
				FromAccountID: test.out.FromAccountID,
				ToAccountID:   test.out.ToAccountID,
				Amount:        test.out.Amount,
				Valid:         test.out.Valid,
			}

			service.EXPECT().InsertTransaction(mockInData).Return(mockOutData, nil)

			handler := accountinghandler.NewHandler(accountinghandler.Params{
				Logger:         logger.NewLogger(os.Stdout, test.name),
				AccountService: service,
			})

			reqBody, err := v1.Encode(test.in)
			if err != nil {
				t.Fatal(err)
			}

			resp := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPost, "/v1/transactions", bytes.NewBuffer(reqBody))

			handler.ServeHTTP(resp, req)

			if resp.Code != http.StatusOK {
				t.Errorf("StatusCode: %d", resp.Code)
			}

			respData := accountingapi.TransactionResponse{}
			err = accountingapi.DecodeJSON(resp.Body, &respData)
			require.NoError(t, err)

			if respData.ID != test.out.ID {
				t.Errorf("ID: Expected: %v, Got: %v", test.out.ID, respData.ID)
			}

			if respData.FromAccountID != test.out.FromAccountID {
				t.Errorf("FromAccountID: Expected: %v, Got: %v", test.out.FromAccountID, respData.FromAccountID)
			}

			if respData.ToAccountID != test.out.ToAccountID {
				t.Errorf("ToAccountID: Expected: %v, Got: %v", test.out.ToAccountID, respData.ToAccountID)
			}

			if respData.Amount != test.out.Amount {
				t.Errorf("Amount: Expected: %v, Got: %v", test.out.Amount, respData.Amount)
			}

			if respData.Valid != test.out.Valid {
				t.Errorf("Valid: Expected: %v, Got: %v", test.out.Valid, respData.Valid)
			}
		})
	}
}
