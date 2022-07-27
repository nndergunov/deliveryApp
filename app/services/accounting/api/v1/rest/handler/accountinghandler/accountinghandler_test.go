package accountinghandler_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/nndergunov/deliveryApp/app/services/accounting/api/v1/rest/accountingapi"

	"github.com/nndergunov/deliveryApp/app/pkg/api/v1"
	"github.com/nndergunov/deliveryApp/app/pkg/logger"

	"github.com/nndergunov/deliveryApp/app/services/accounting/api/v1/rest/handler/accountinghandler"
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
				Logger:  logger.NewLogger(os.Stdout, test.name),
				Service: service,
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

			assert.Equal(t, test.out, respData)
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
				Logger:  logger.NewLogger(os.Stdout, test.name),
				Service: service,
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

			assert.Equal(t, test.out, respData)
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
				Logger:  logger.NewLogger(os.Stdout, test.name),
				Service: service,
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

			assert.Equal(t, test.out, respDataList)
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
				Logger:  logger.NewLogger(os.Stdout, test.name),
				Service: service,
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

			assert.Equal(t, test.out, accountDeletedResponse)
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
				Valid:         true,
			},
		},
	}

	for _, currentTest := range tests {
		test := currentTest

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
				Logger:  logger.NewLogger(os.Stdout, test.name),
				Service: service,
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

			assert.Equal(t, test.out, respData)
		})
	}
}
