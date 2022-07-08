package accountingservice_test

import (
	"os"
	"strconv"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/nndergunov/deliveryApp/app/pkg/logger"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

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
			storage.EXPECT().GetAccountListByParam(domain.SearchParam{"user_id": strconv.Itoa(test.in.UserID), "user_type": "'" + test.in.UserType + "'"}).Return(nil, nil)

			service := accountingservice.NewService(accountingservice.Params{Storage: storage, Logger: logger.NewLogger(os.Stdout, "service: ")})
			newAccount, err := service.InsertNewAccount(test.in)
			require.NoError(t, err)

			assert.Equal(t, test.out, newAccount)
		})
	}
}

func TestGetAccountByID(t *testing.T) {
	t.Parallel()

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

			assert.Equal(t, test.out, result)
		})
	}
}

func TestGetAccountList(t *testing.T) {
	t.Parallel()

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

			assert.Equal(t, test.out, resultList)
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
				Valid:         true,
			},
		},
	}

	for _, currentTest := range tests {
		test := currentTest

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

			assert.Equal(t, test.out, result)
		})
	}
}
