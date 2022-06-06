package courierservice

import (
	"database/sql"
	"github.com/nndergunov/deliveryApp/app/pkg/logger"
	"strconv"

	"accounting/pkg/domain"
)

// CourierService is the interface for the accounting service.
type CourierService interface {
	InsertNewCourierAccount(account domain.CourierAccount) (*domain.CourierAccount, error)
	GetCourierAccount(courierID string) (*domain.CourierAccount, error)
	DeleteCourierAccount(courierID string) (string, error)

	AddToBalanceCourierAccount(courierID string, account domain.CourierAccount) (*domain.CourierAccount, error)
	SubFromBalanceCourierAccount(courierID string, account domain.CourierAccount) (*domain.CourierAccount, error)
}

// Params is the input parameter struct for the module that contains its dependencies
type Params struct {
	Storage CourierStorage
	Logger  *logger.Logger
}

type Service struct {
	Storage CourierStorage
	logger  *logger.Logger
}

// NewService constructs a new NewService.
func NewService(p Params) *Service {
	ServiceItem := &Service{
		Storage: p.Storage,
		logger:  p.Logger,
	}

	return ServiceItem
}

func (c Service) InsertNewCourierAccount(account domain.CourierAccount) (*domain.CourierAccount, error) {

	if account.CourierID < 1 {
		return nil, errWrongCourierIDType
	}

	//check duplicate
	courierAccount, err := c.Storage.GetCourierAccountByID(account.CourierID)
	if err != nil && err != sql.ErrNoRows {
		c.logger.Println(err)
		return nil, systemErr
	}
	if courierAccount != nil {
		return nil, errCourierAccountExist
	}

	//insertNewCourierAccount
	newCourierAccount, err := c.Storage.InsertNewCourierAccount(account)
	if err != nil && err != sql.ErrNoRows {
		c.logger.Println(err)
		return nil, systemErr
	}

	return newCourierAccount, nil
}

func (c Service) GetCourierAccount(courierID string) (*domain.CourierAccount, error) {
	idUint, err := strconv.ParseUint(courierID, 10, 64)
	if err != nil {
		c.logger.Println(err)
		return nil, errWrongCourierIDType
	}

	courierAccount, err := c.Storage.GetCourierAccountByID(idUint)
	if err != nil && err != sql.ErrNoRows {
		c.logger.Println(err)
		return nil, systemErr
	}
	if courierAccount == nil {
		return nil, errCourierAccountNotFound
	}

	return courierAccount, nil
}

func (c Service) DeleteCourierAccount(courierID string) (string, error) {
	idUint, err := strconv.ParseUint(courierID, 10, 64)
	if err != nil {
		c.logger.Println(err)
		return "", errWrongCourierIDType
	}

	courierAccount, err := c.Storage.GetCourierAccountByID(idUint)
	if err != nil && err != sql.ErrNoRows {
		c.logger.Println(err)
		return "", systemErr
	}
	if courierAccount == nil {
		return "", errCourierAccountNotFound
	}

	err = c.Storage.DeleteCourierAccount(idUint)
	if err != nil && err != sql.ErrNoRows {
		c.logger.Println(err)
		return "", systemErr
	}

	return "courier account deleted", nil
}

func (c Service) AddToBalanceCourierAccount(courierID string, account domain.CourierAccount) (*domain.CourierAccount, error) {
	idUint, err := strconv.ParseUint(courierID, 10, 64)
	if err != nil {
		c.logger.Println(err)
		return nil, errWrongCourierIDType
	}
	account.CourierID = idUint

	if account.CourierID < 1 {
		return nil, errWrongCourierID
	}

	if account.Balance < 1 {
		return nil, errWrongAmount
	}

	courierAccount, err := c.Storage.GetCourierAccountByID(account.CourierID)
	if err != nil && err != sql.ErrNoRows {
		c.logger.Println(err)
		return nil, systemErr
	}
	if courierAccount == nil {
		return nil, errCourierAccountNotFound
	}

	courierAccountUpdated, err := c.Storage.AddToBalanceCourierAccount(account)
	if err != nil && err != sql.ErrNoRows {
		c.logger.Println(err)
		return nil, systemErr
	}

	return courierAccountUpdated, nil
}

func (c Service) SubFromBalanceCourierAccount(courierID string, account domain.CourierAccount) (*domain.CourierAccount, error) {
	idUint, err := strconv.ParseUint(courierID, 10, 64)
	if err != nil {
		c.logger.Println(err)
		return nil, errWrongCourierIDType
	}
	account.CourierID = idUint

	if account.CourierID < 1 {
		return nil, errWrongCourierID
	}

	if account.Balance < 1 {
		return nil, errWrongAmount
	}

	courierAccount, err := c.Storage.GetCourierAccountByID(account.CourierID)
	if err != nil && err != sql.ErrNoRows {
		c.logger.Println(err)
		return nil, systemErr
	}
	if courierAccount == nil {
		return nil, errCourierAccountNotFound
	}

	if courierAccount.Balance < account.Balance {
		return nil, errNotEnoughBalance
	}

	courierAccountUpdated, err := c.Storage.SubFromBalanceCourierAccount(account)
	if err != nil && err != sql.ErrNoRows {
		c.logger.Println(err)
		return nil, systemErr
	}

	return courierAccountUpdated, nil
}
