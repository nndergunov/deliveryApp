package accountservice

import (
	"database/sql"
	"errors"
	"strconv"

	"github.com/nndergunov/deliveryApp/app/pkg/logger"

	"accounting/pkg/domain"
)

// AccountService is the interface for the accounting service.
type AccountService interface {
	InsertNewAccount(account domain.Account) (*domain.Account, error)
	GetAccountByID(ID string) (*domain.Account, error)
	GetAccountListByParam(param domain.SearchParam) ([]domain.Account, error)
	DeleteAccount(id string) (string, error)

	Transact(transaction domain.Transaction) (*domain.Transaction, error)
}

// Params is the input parameter struct for the module that contains its dependencies
type Params struct {
	Storage AccountStorage
	Logger  *logger.Logger
}

type Service struct {
	storage AccountStorage
	logger  *logger.Logger
}

// NewService constructs a new NewService.
func NewService(p Params) *Service {
	ServiceItem := &Service{
		storage: p.Storage,
		logger:  p.Logger,
	}

	return ServiceItem
}

func (c Service) InsertNewAccount(account domain.Account) (*domain.Account, error) {
	if account.UserID < 1 {
		return nil, errWrongUserID
	}

	if account.UserType == "" {
		return nil, errWrongUserType
	}
	_, ok := domain.UserType[account.UserType]
	if !ok {
		return nil, errWrongUserType
	}

	// check duplicate
	param := domain.SearchParam{}
	param["user_id"] = strconv.Itoa(account.UserID)
	param["user_type"] = account.UserType

	gotAccountList, err := c.storage.GetAccountListByParam(param)
	if err != nil && err != sql.ErrNoRows {
		c.logger.Println(err)
		return nil, systemErr
	}
	if len(gotAccountList) > 1 {
		return nil, errMaxNumberOfAccount
	}

	// insertAccount
	newAccount, err := c.storage.InsertNewAccount(account)
	if err != nil && err != sql.ErrNoRows {
		c.logger.Println(err)
		return nil, systemErr
	}

	return newAccount, nil
}

func (c Service) GetAccountByID(id string) (*domain.Account, error) {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.logger.Println(err)
		return nil, errWrongUserType
	}

	account, err := c.storage.GetAccountByID(idInt)
	if err != nil && err != sql.ErrNoRows {
		c.logger.Println(err)
		return nil, systemErr
	}
	if account == nil {
		return nil, errAccountNotFound
	}

	return account, nil
}

func (c Service) GetAccountListByParam(param domain.SearchParam) ([]domain.Account, error) {
	_, err := strconv.Atoi(param["user_id"])
	if err != nil {
		c.logger.Println(err)
		return nil, errWrongUserID
	}

	accountList, err := c.storage.GetAccountListByParam(param)
	if err != nil && err != sql.ErrNoRows {
		c.logger.Println(err)
		return nil, systemErr
	}
	if accountList == nil {
		return nil, errAccountNotFound
	}

	return accountList, nil
}

func (c Service) DeleteAccount(id string) (string, error) {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.logger.Println(err)
		return "", errWrongUserType
	}

	gotAccount, err := c.storage.GetAccountByID(idInt)
	if err != nil && err != sql.ErrNoRows {
		c.logger.Println(err)
		return "", systemErr
	}
	if gotAccount == nil {
		return "", errAccountNotFound
	}

	err = c.storage.DeleteAccount(idInt)
	if err != nil && err != sql.ErrNoRows {
		c.logger.Println(err)
		return "", systemErr
	}

	return "account deleted", nil
}

func (c Service) Transact(transaction domain.Transaction) (*domain.Transaction, error) {
	if transaction.Amount < 1 {
		return nil, errWrongAmount
	}

	if transaction.FromAccountID == 0 && transaction.ToAccountID != 0 {
		// add to balance
		toAccount, err := c.storage.GetAccountByID(transaction.ToAccountID)
		if err != nil && err != sql.ErrNoRows {
			c.logger.Println(err)
			return nil, systemErr
		}
		if toAccount == nil {
			return nil, errAccountNotFound
		}

		savedTransaction, err := c.storage.AddToAccountBalance(transaction)
		if err != nil && err != sql.ErrNoRows {
			c.logger.Println(err)
			return nil, systemErr
		}

		return savedTransaction, nil
	}

	if transaction.FromAccountID != 0 && transaction.ToAccountID == 0 {
		// sub from balance
		fromAccount, err := c.storage.GetAccountByID(transaction.FromAccountID)
		if err != nil && err != sql.ErrNoRows {
			c.logger.Println(err)
			return nil, systemErr
		}
		if fromAccount == nil {
			return nil, errAccountNotFound
		}

		if fromAccount.Balance < transaction.Amount {
			return nil, errWrongAmount
		}

		savedTransaction, err := c.storage.SubFromAccountBalance(transaction)
		if err != nil && err != sql.ErrNoRows {
			c.logger.Println(err)
			return nil, systemErr
		}

		return savedTransaction, nil
	}

	if transaction.FromAccountID != 0 && transaction.ToAccountID != 0 {
		// sub from balance and add to balance
		fromAccount, err := c.storage.GetAccountByID(transaction.FromAccountID)
		if err != nil && err != sql.ErrNoRows {
			c.logger.Println(err)
			return nil, systemErr
		}
		if fromAccount == nil {
			return nil, errAccountNotFound
		}

		if fromAccount.Balance < transaction.Amount {
			return nil, errWrongAmount
		}

		toAccount, err := c.storage.GetAccountByID(transaction.ToAccountID)
		if err != nil && err != sql.ErrNoRows {
			c.logger.Println(err)
			return nil, systemErr
		}
		if toAccount == nil {
			return nil, errAccountNotFound
		}

		savedTransaction, err := c.storage.Transact(transaction)
		if err != nil && err != sql.ErrNoRows {
			c.logger.Println(err)
			return nil, systemErr
		}

		return savedTransaction, nil
	}

	return nil, errors.New("wrong operation")
}
