package accountingstorage_test

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/nndergunov/deliveryApp/app/services/accounting/pkg/db/dbtest"
	"github.com/nndergunov/deliveryApp/app/services/accounting/pkg/docker"
	"github.com/nndergunov/deliveryApp/app/services/accounting/pkg/domain"
	"github.com/nndergunov/deliveryApp/app/services/accounting/pkg/storage/accountingstorage"
)

// equalAccount - function to compare selected fields from struct. Fields which needed
func equalAccount(t *testing.T, get domain.Account, want domain.Account) {
	if get.UserID != want.UserID {
		t.Errorf("UserID: Expected: %v, Got: %v", want.UserID, get.UserID)
	}

	if get.UserType != want.UserType {
		t.Errorf("UserType: Expected: %s, Got: %s", want.UserType, get.UserType)
	}

	if get.Balance != want.Balance {
		t.Errorf("Balance: Expected: %v, Got: %v", want.Balance, get.Balance)
	}
}

// equalTr - function to compare selected fields from struct. Fields which needed
func equalTr(t *testing.T, get domain.Transaction, want domain.Transaction) {
	if get.FromAccountID != want.FromAccountID {
		t.Errorf("FromAccountID: Expected: %v, Got: %v", want.FromAccountID, get.FromAccountID)
	}

	if get.ToAccountID != want.ToAccountID {
		t.Errorf("ToAccountID: Expected: %v, Got: %v", want.ToAccountID, get.ToAccountID)
	}

	if get.Amount != want.Amount {
		t.Errorf("Amount: Expected: %v, Got: %v", want.Amount, get.Amount)
	}
}

var c *docker.Container

func TestMain(m *testing.M) {
	var err error
	c, err = dbtest.StartDB()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer dbtest.StopDB(c)

	m.Run()
}

func TestInsertAccount(t *testing.T) {
	tests := []struct {
		name string
		in   domain.Account
		out  domain.Account
	}{
		{
			name: "test_insert_account",
			in: domain.Account{
				UserID:   1,
				UserType: "courier",
			},
			out: domain.Account{
				UserID:   1,
				UserType: "courier",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			database, teardown := dbtest.NewUnit(t, c, test.name)
			t.Cleanup(teardown)

			s := accountingstorage.NewStorage(accountingstorage.Params{DB: database})

			resp, err := s.InsertNewAccount(test.in)
			require.NoError(t, err)
			require.NotNil(t, resp)

			equalAccount(t, *resp, test.out)
		})
	}
}

func TestGetAccountByID(t *testing.T) {
	tests := []struct {
		name string
		in   domain.Account
		out  domain.Account
	}{
		{
			name: "test_get_account",
			in: domain.Account{
				UserID:   1,
				UserType: "courier",
			},

			out: domain.Account{
				UserID:   1,
				UserType: "courier",
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			database, teardown := dbtest.NewUnit(t, c, test.name)
			t.Cleanup(teardown)

			s := accountingstorage.NewStorage(accountingstorage.Params{DB: database})

			resp, err := s.InsertNewAccount(test.in)
			require.NoError(t, err)
			require.NotNil(t, resp)
			respGet, err := s.GetAccountByID(resp.ID)
			require.NoError(t, err)

			equalAccount(t, *respGet, test.out)
		})
	}
}

func TestGetAccountListByParam(t *testing.T) {
	tests := []struct {
		name string
		in   []domain.Account
		out  []domain.Account
	}{
		{
			name: "test_get_account_list_by_param",
			in: []domain.Account{
				{
					UserID:   1,
					UserType: "courier",
				},
				{
					UserID:   2,
					UserType: "consumer",
				},
			},
			out: []domain.Account{
				{
					UserID:   1,
					UserType: "courier",
				},
				{
					UserID:   2,
					UserType: "consumer",
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			database, teardown := dbtest.NewUnit(t, c, test.name)
			t.Cleanup(teardown)

			s := accountingstorage.NewStorage(accountingstorage.Params{DB: database})

			userIds := []string{}
			userTypes := []string{}

			for _, in := range test.in {
				_, err := s.InsertNewAccount(in)
				require.NoError(t, err)

				userIds = append(userIds, strconv.Itoa(in.UserID))
				userTypes = append(userTypes, "'"+in.UserType+"'")
			}

			param := domain.SearchParam{}
			param["user_id"] = strings.Join(userIds, ",")
			param["user_type"] = strings.Join(userTypes, ",")

			respList, err := s.GetAccountListByParam(param)
			require.NoError(t, err)
			require.NotNil(t, respList)

			for i := 0; i <= len(respList)-1; i++ {
				equalAccount(t, test.out[i], respList[i])
			}
		})
	}
}

func TestAddToAccountBalance(t *testing.T) {
	tests := []struct {
		name string
		in   domain.Transaction
		out  domain.Transaction
	}{
		{
			name: "test_add_to_account_balance",
			in: domain.Transaction{
				ToAccountID: 2,
				Amount:      50,
			},

			out: domain.Transaction{
				FromAccountID: 0,
				ToAccountID:   2,
				Amount:        50,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			database, teardown := dbtest.NewUnit(t, c, test.name)
			t.Cleanup(teardown)

			s := accountingstorage.NewStorage(accountingstorage.Params{DB: database})

			respData, err := s.AddToAccountBalance(test.in)
			if err != nil {
				t.Fatal(err)
			}
			equalTr(t, *respData, test.out)
		})
	}
}

func TestSubFromAccountBalance(t *testing.T) {
	tests := []struct {
		name string
		in   domain.Transaction
		out  domain.Transaction
	}{
		{
			name: "test_sub_to_account_balance",
			in: domain.Transaction{
				FromAccountID: 1,
				Amount:        50,
			},

			out: domain.Transaction{
				FromAccountID: 1,
				ToAccountID:   0,
				Amount:        50,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			database, teardown := dbtest.NewUnit(t, c, test.name)
			t.Cleanup(teardown)

			s := accountingstorage.NewStorage(accountingstorage.Params{DB: database})

			respData, err := s.SubFromAccountBalance(test.in)
			if err != nil {
				t.Fatal(err)
			}
			equalTr(t, *respData, test.out)
		})
	}
}

func TestInsertTransaction(t *testing.T) {
	tests := []struct {
		name string
		in   domain.Transaction
		out  domain.Transaction
	}{
		{
			name: "test_insert_transaction",
			in: domain.Transaction{
				FromAccountID: 1,
				ToAccountID:   2,
				Amount:        50,
			},

			out: domain.Transaction{
				FromAccountID: 1,
				ToAccountID:   2,
				Amount:        50,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			database, teardown := dbtest.NewUnit(t, c, test.name)
			t.Cleanup(teardown)

			s := accountingstorage.NewStorage(accountingstorage.Params{DB: database})

			respData, err := s.InsertTransaction(test.in)
			if err != nil {
				t.Fatal(err)
			}
			equalTr(t, *respData, test.out)
		})
	}
}
