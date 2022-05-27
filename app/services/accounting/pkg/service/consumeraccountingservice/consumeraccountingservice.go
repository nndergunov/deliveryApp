package consumeraccountingservice

import (
	"accounting/pkg/domain"
	"accounting/pkg/service"
	"database/sql"
	"github.com/nndergunov/deliveryApp/app/pkg/logger"
	"strconv"
)

// Params is the input parameter struct for the module that contains its dependencies
type Params struct {
	ConsumerAccountingStorage service.AccountingStorage
	Logger                    *logger.Logger
}

type consumerAccountingService struct {
	AccountingStorage service.AccountingStorage
	logger            *logger.Logger
}

// NewConsumerAccountingService constructs a new NewConsumerAccountingService.
func NewConsumerAccountingService(p Params) service.AccountingService {
	consumerAccountingServiceItem := &consumerAccountingService{
		AccountingStorage: p.ConsumerAccountingStorage,
		logger:            p.Logger,
	}

	return consumerAccountingServiceItem
}

func (c consumerAccountingService) InsertNewConsumerAccount(account domain.ConsumerAccount) (*domain.ConsumerAccount, error) {

	//check duplicate
	consumerAccount, err := c.AccountingStorage.GetConsumerAccountByID(account.ConsumerID)
	if err != nil && err != sql.ErrNoRows {
		c.logger.Println(err)
		return nil, systemErr
	}
	if consumerAccount != nil {
		return nil, errConsumerAccountExist
	}

	//insertNewConsumerAccount
	newConsumerAccount, err := c.AccountingStorage.InsertNewConsumerAccount(account)
	if err != nil && err != sql.ErrNoRows {
		c.logger.Println(err)
		return nil, systemErr
	}

	return newConsumerAccount, nil
}

func (c consumerAccountingService) GetConsumerAccount(consumerID string) (*domain.ConsumerAccount, error) {
	idUint, err := strconv.ParseUint(consumerID, 10, 64)
	if err != nil {
		c.logger.Println(err)
		return nil, errWrongConsumerIDType
	}

	consumerAccount, err := c.AccountingStorage.GetConsumerAccountByID(idUint)
	if err != nil && err != sql.ErrNoRows {
		c.logger.Println(err)
		return nil, systemErr
	}
	if consumerAccount == nil {
		return nil, errConsumerAccountNotFound
	}

	return consumerAccount, nil
}

func (c consumerAccountingService) DeleteConsumerAccount(consumerID string) (string, error) {
	idUint, err := strconv.ParseUint(consumerID, 10, 64)
	if err != nil {
		c.logger.Println(err)
		return "", errWrongConsumerIDType
	}

	err = c.AccountingStorage.DeleteConsumerAccount(idUint)
	if err != nil && err != sql.ErrNoRows {
		c.logger.Println(err)
		return "", systemErr
	}

	return "consumer account deleted", nil
}

func (c consumerAccountingService) AddToConsumerAccount(account domain.ConsumerAccount) (*domain.ConsumerAccount, error) {

	if account.Balance < 0 {
		return nil, errWrongAmount
	}

	consumerAccount, err := c.AccountingStorage.GetConsumerAccountByID(account.ConsumerID)
	if err != nil && err != sql.ErrNoRows {
		c.logger.Println(err)
		return nil, systemErr
	}
	if consumerAccount == nil {
		return nil, errConsumerAccountNotFound
	}

	consumerAccountUpdated, err := c.AccountingStorage.AddToConsumerAccount(account)
	if err != nil && err != sql.ErrNoRows {
		c.logger.Println(err)
		return nil, systemErr
	}

	return consumerAccountUpdated, nil
}

func (c consumerAccountingService) SubFromConsumerAccount(account domain.ConsumerAccount) (*domain.ConsumerAccount, error) {

	if account.Balance < 0 {
		return nil, errWrongAmount
	}

	consumerAccount, err := c.AccountingStorage.GetConsumerAccountByID(account.ConsumerID)
	if err != nil && err != sql.ErrNoRows {
		c.logger.Println(err)
		return nil, systemErr
	}
	if consumerAccount == nil {
		return nil, errConsumerAccountNotFound
	}

	if consumerAccount.Balance < account.Balance {
		return nil, errNotEnoughBalance
	}

	consumerAccountUpdated, err := c.AccountingStorage.SubToConsumerAccount(account)
	if err != nil && err != sql.ErrNoRows {
		c.logger.Println(err)
		return nil, systemErr
	}

	return consumerAccountUpdated, nil
}
