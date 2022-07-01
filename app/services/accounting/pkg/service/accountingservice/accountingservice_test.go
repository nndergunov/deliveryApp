package accountingservice_test

import (
	"github.com/golang/mock/gomock"
	"github.com/nndergunov/deliveryApp/app/pkg/logger"
	"github.com/stretchr/testify/require"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/nndergunov/deliveryApp/app/services/accounting/pkg/domain"
	mockstorage "github.com/nndergunov/deliveryApp/app/services/accounting/pkg/mocks"
	"github.com/nndergunov/deliveryApp/app/services/accounting/pkg/service/accountingservice"
)

func TestInsertNewAccount(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		in   domain.Account
		out  *domain.Account
	}{
		{
			"Insert new account test",
			domain.Account{
				UserID:   1,
				UserType: "courier",
			},
			&domain.Account{
				UserID:   1,
				UserType: "courier",
				Balance:  0,
			},
		},
	}

	for _, currentTest := range tests {
		test := currentTest

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			ctl := gomock.NewController(t)
			storage := mockstorage.NewMockAccountStorage(ctl)
			storage.EXPECT().InsertNewAccount(test.in).Return(test.out, nil).Times(1)
			storage.EXPECT().GetAccountListByParam(domain.SearchParam{"user_id": strconv.Itoa(test.in.UserID), "user_type": test.in.UserType}).Return(nil, nil)

			service := accountingservice.NewService(accountingservice.Params{Storage: storage, Logger: logger.NewLogger(os.Stdout, "service: ")})
			newAccount, err := service.InsertNewAccount(test.in)
			require.NoError(t, err)

			if newAccount.UserID != test.out.UserID {
				t.Errorf("UserID: Expected: %v, Got: %v", test.out.UserID, newAccount.UserID)
			}

			if newAccount.UserType != test.out.UserType {
				t.Errorf("UserID: Expected: %v, Got: %v", test.out.UserID, newAccount.UserID)
			}

			if newAccount.Balance != test.out.Balance {
				t.Errorf("Balance: Expected: %v, Got: %v", test.out.Balance, newAccount.Balance)
			}
		})
	}
}

func TestGetAccountByID(t *testing.T) {
	tests := []struct {
		name string
		in   int
		out  *domain.Account
	}{
		{
			"Get account by id",
			1,
			&domain.Account{
				UserID:   1,
				UserType: "courier",
			},
		},
	}

	for _, currentTest := range tests {
		test := currentTest

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			ctl := gomock.NewController(t)
			storage := mockstorage.NewMockAccountStorage(ctl)

			storage.EXPECT().GetAccountByID(test.in).Return(test.out, nil)

			service := accountingservice.NewService(accountingservice.Params{Storage: storage, Logger: logger.NewLogger(os.Stdout, "service: ")})

			result, err := service.GetAccountByID(strconv.Itoa(test.in))
			require.NoError(t, err)

			if result.UserID != test.out.UserID {
				t.Errorf("UserID: Expected: %v, Got: %v", test.out.UserID, result.UserID)
			}

			if result.UserType != test.out.UserType {
				t.Errorf("UserID: Expected: %v, Got: %v", test.out.UserID, result.UserID)
			}

			if result.Balance != test.out.Balance {
				t.Errorf("Balance: Expected: %v, Got: %v", test.out.Balance, result.Balance)
			}
		})
	}
}

func TestGetAccountList(t *testing.T) {
	tests := []struct {
		name string
		in   domain.SearchParam
		out  []domain.Account
	}{
		{
			"TestGetAccountList",
			domain.SearchParam{"user_id": "1", "user_type": "courier"},
			[]domain.Account{
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
	}

	for _, currentTest := range tests {
		test := currentTest

		t.Run(test.name, func(t *testing.T) {

			ctl := gomock.NewController(t)
			storage := mockstorage.NewMockAccountStorage(ctl)

			storage.EXPECT().GetAccountListByParam(test.in).Return(test.out, nil)

			service := accountingservice.NewService(accountingservice.Params{Storage: storage, Logger: logger.NewLogger(os.Stdout, "service: ")})

			resultList, err := service.GetAccountListByParam(test.in)
			require.NoError(t, err)

			if len(resultList) != len(test.out) {
				t.Errorf("len: Expected: %v, Got: %v", len(test.out), len(resultList))
			}
			result1 := resultList[0]
			test1 := test.out[0]

			if result1.ID != test1.ID {
				t.Errorf("ID: Expected: %v, Got: %v", test1.ID, result1.ID)
			}
			if result1.UserID != test1.UserID {
				t.Errorf("UserID: Expected: %v, Got: %v", test1.UserID, result1.UserID)
			}
			if result1.UserType != test1.UserType {
				t.Errorf("UserType: Expected: %v, Got: %v", test1.UserType, result1.UserType)
			}
			if result1.Balance != test1.Balance {
				t.Errorf("Balance: Expected: %v, Got: %v", test1.Balance, result1.Balance)
			}

			result2 := resultList[1]
			test2 := test.out[1]

			if result1.ID != test1.ID {
				t.Errorf("ID: Expected: %v, Got: %v", test1.ID, result1.ID)
			}
			if result2.UserID != test2.UserID {
				t.Errorf("UserID: Expected: %v, Got: %v", test2.UserID, result2.UserID)
			}
			if result2.UserType != test2.UserType {
				t.Errorf("UserType: Expected: %v, Got: %v", test2.UserType, result2.UserType)
			}
			if result2.Balance != test2.Balance {
				t.Errorf("Balance: Expected: %v, Got: %v", test2.Balance, result2.Balance)
			}
		})
	}
}

func TestInsertTransactions(t *testing.T) {
	t.Parallel()
	type test struct {
		name string
		in   domain.Transaction
		out  *domain.Transaction
	}
	tests := []test{
		{
			name: "from account to account",
			in: domain.Transaction{
				FromAccountID: 1,
				ToAccountID:   2,
				Amount:        50,
			},
			out: &domain.Transaction{
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
			in: domain.Transaction{
				FromAccountID: 1,
				Amount:        50,
			},
			out: &domain.Transaction{
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
			in: domain.Transaction{
				ToAccountID: 2,
				Amount:      50,
			},
			out: &domain.Transaction{
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

			ctl := gomock.NewController(t)
			storage := mockstorage.NewMockAccountStorage(ctl)

			fromAccount := &domain.Account{
				ID:       1,
				UserID:   1,
				UserType: "restaurant",
				Balance:  50,
			}

			toAccount := &domain.Account{
				ID:       2,
				UserID:   1,
				UserType: "courier",
				Balance:  0,
			}

			storage.EXPECT().GetAccountByID(test.in.FromAccountID).Return(fromAccount, nil)
			storage.EXPECT().GetAccountByID(test.in.ToAccountID).Return(toAccount, nil)
			storage.EXPECT().SubFromAccountBalance(test.in).Return(test.out, nil)
			storage.EXPECT().AddToAccountBalance(test.in).Return(test.out, nil)
			storage.EXPECT().InsertTransaction(test.in).Return(test.out, nil)

			service := accountingservice.NewService(accountingservice.Params{Storage: storage, Logger: logger.NewLogger(os.Stdout, "service: ")})

			result, err := service.InsertTransaction(test.in)
			require.NoError(t, err)

			if result.ID != test.out.ID {
				t.Errorf("ID: Expected: %v, Got: %v", test.out.ID, result.ID)
			}
			if result.FromAccountID != test.out.FromAccountID {
				t.Errorf("FromAccountID: Expected: %v, Got: %v", test.out.FromAccountID, result.FromAccountID)
			}
			if result.ToAccountID != test.out.ToAccountID {
				t.Errorf("ToAccountID: Expected: %v, Got: %v", test.out.ToAccountID, result.ToAccountID)
			}
			if result.Amount != test.out.Amount {
				t.Errorf("Amount: Expected: %v, Got: %v", test.out.Amount, result.Amount)
			}
			if result.CreatedAt != test.out.CreatedAt {
				t.Errorf("CreatedAt: Expected: %v, Got: %v", test.out.CreatedAt, result.CreatedAt)
			}
			if result.UpdatedAt != test.out.UpdatedAt {
				t.Errorf("UpdatedAt: Expected: %v, Got: %v", test.out.UpdatedAt, result.UpdatedAt)
			}
		})
	}
}
