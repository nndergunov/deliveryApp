package restaurantservice

import (
	"github.com/nndergunov/deliveryApp/app/pkg/logger"

	"database/sql"
	"strconv"

	"accounting/pkg/domain"
)

// RestaurantService is the interface for the accounting service.
type RestaurantService interface {
	InsertNewRestaurantAccount(account domain.RestaurantAccount) (*domain.RestaurantAccount, error)
	GetRestaurantAccount(restaurantID string) (*domain.RestaurantAccount, error)
	DeleteRestaurantAccount(restaurantID string) (string, error)

	AddToBalanceRestaurantAccount(restaurantID string, account domain.RestaurantAccount) (*domain.RestaurantAccount, error)
	SubFromBalanceRestaurantAccount(restaurantID string, account domain.RestaurantAccount) (*domain.RestaurantAccount, error)
}

// Params is the input parameter struct for the module that contains its dependencies
type Params struct {
	Storage RestaurantStorage
	Logger  *logger.Logger
}

type Service struct {
	Storage RestaurantStorage
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

func (c Service) InsertNewRestaurantAccount(account domain.RestaurantAccount) (*domain.RestaurantAccount, error) {

	if account.RestaurantID < 1 {
		return nil, errWrongRestaurantID
	}

	//check duplicate
	RestaurantAccount, err := c.Storage.GetRestaurantAccountByID(account.RestaurantID)
	if err != nil && err != sql.ErrNoRows {
		c.logger.Println(err)
		return nil, systemErr
	}
	if RestaurantAccount != nil {
		return nil, errRestaurantAccountExist
	}

	//insertNewRestaurantAccount
	newRestaurantAccount, err := c.Storage.InsertNewRestaurantAccount(account)
	if err != nil && err != sql.ErrNoRows {
		c.logger.Println(err)
		return nil, systemErr
	}

	return newRestaurantAccount, nil
}

func (c Service) GetRestaurantAccount(restaurantID string) (*domain.RestaurantAccount, error) {
	idInt, err := strconv.Atoi(restaurantID)
	if err != nil {
		c.logger.Println(err)
		return nil, errWrongRestaurantIDType
	}

	RestaurantAccount, err := c.Storage.GetRestaurantAccountByID(idInt)
	if err != nil && err != sql.ErrNoRows {
		c.logger.Println(err)
		return nil, systemErr
	}
	if RestaurantAccount == nil {
		return nil, errRestaurantAccountNotFound
	}

	return RestaurantAccount, nil
}

func (c Service) DeleteRestaurantAccount(restaurantID string) (string, error) {
	idInt, err := strconv.Atoi(restaurantID)
	if err != nil {
		c.logger.Println(err)
		return "", errWrongRestaurantIDType
	}

	RestaurantAccount, err := c.Storage.GetRestaurantAccountByID(idInt)
	if err != nil && err != sql.ErrNoRows {
		c.logger.Println(err)
		return "", systemErr
	}
	if RestaurantAccount == nil {
		return "", errRestaurantAccountNotFound
	}

	err = c.Storage.DeleteRestaurantAccount(idInt)
	if err != nil && err != sql.ErrNoRows {
		c.logger.Println(err)
		return "", systemErr
	}

	return "Restaurant account deleted", nil
}

func (c Service) AddToBalanceRestaurantAccount(restaurantID string, account domain.RestaurantAccount) (*domain.RestaurantAccount, error) {
	idInt, err := strconv.Atoi(restaurantID)
	if err != nil {
		c.logger.Println(err)
		return nil, errWrongRestaurantIDType
	}
	account.RestaurantID = idInt

	if account.RestaurantID < 1 {
		return nil, errWrongRestaurantID
	}

	if account.Balance < 1 {
		return nil, errWrongAmount
	}

	RestaurantAccount, err := c.Storage.GetRestaurantAccountByID(account.RestaurantID)
	if err != nil && err != sql.ErrNoRows {
		c.logger.Println(err)
		return nil, systemErr
	}
	if RestaurantAccount == nil {
		return nil, errRestaurantAccountNotFound
	}

	RestaurantAccountUpdated, err := c.Storage.AddToBalanceRestaurantAccount(account)
	if err != nil && err != sql.ErrNoRows {
		c.logger.Println(err)
		return nil, systemErr
	}

	return RestaurantAccountUpdated, nil
}

func (c Service) SubFromBalanceRestaurantAccount(restaurantID string, account domain.RestaurantAccount) (*domain.RestaurantAccount, error) {
	idInt, err := strconv.Atoi(restaurantID)
	if err != nil {
		c.logger.Println(err)
		return nil, errWrongRestaurantIDType
	}

	account.RestaurantID = idInt

	if account.RestaurantID < 1 {
		return nil, errWrongRestaurantID
	}

	if account.Balance < 1 {
		return nil, errWrongAmount
	}

	RestaurantAccount, err := c.Storage.GetRestaurantAccountByID(account.RestaurantID)
	if err != nil && err != sql.ErrNoRows {
		c.logger.Println(err)
		return nil, systemErr
	}
	if RestaurantAccount == nil {
		return nil, errRestaurantAccountNotFound
	}

	if RestaurantAccount.Balance < account.Balance {
		return nil, errNotEnoughBalance
	}

	RestaurantAccountUpdated, err := c.Storage.SubFromBalanceRestaurantAccount(account)
	if err != nil && err != sql.ErrNoRows {
		c.logger.Println(err)
		return nil, systemErr
	}

	return RestaurantAccountUpdated, nil
}
