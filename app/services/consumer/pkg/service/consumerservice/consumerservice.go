package consumerservice

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/nndergunov/deliveryApp/app/pkg/logger"

	"github.com/nndergunov/deliveryApp/app/services/consumer/pkg/domain"
)

// ConsumerService is the interface for the user service.
type ConsumerService interface {
	InsertConsumer(consumer domain.Consumer) (*domain.Consumer, error)
	DeleteConsumer(id string) (data string, err error)
	UpdateConsumer(consumer domain.Consumer, id string) (*domain.Consumer, error)
	GetAllConsumer() ([]domain.Consumer, error)
	GetConsumer(id string) (*domain.Consumer, error)

	InsertLocation(location domain.Location, id string) (*domain.Location, error)
	UpdateLocation(location domain.Location, id string) (*domain.Location, error)
	GetLocation(id string) (*domain.Location, error)
}

// Params is the input parameter struct for the module that contains its dependencies
type Params struct {
	Storage ConsumerStorage
	Logger  *logger.Logger
}

type service struct {
	storage ConsumerStorage
	logger  *logger.Logger
}

// NewService constructs a new NewService.
func NewService(p Params) ConsumerService {
	consumerServiceItem := &service{
		storage: p.Storage,
		logger:  p.Logger,
	}

	return consumerServiceItem
}

// InsertConsumer prepare and send data to storage service.
func (c *service) InsertConsumer(consumer domain.Consumer) (*domain.Consumer, error) {
	if consumer.Phone == "" && consumer.Email == "" {
		return nil, fmt.Errorf("wrong phone or email")
	}
	param := domain.SearchParam{}
	param["email"] = consumer.Email
	param["phone"] = consumer.Phone
	foundConsumer, err := c.storage.GetConsumerDuplicateByParam(param)
	if err != nil && err != sql.ErrNoRows {
		c.logger.Println(err)
		return nil, systemErr
	}
	if foundConsumer != nil {
		return nil, fmt.Errorf("consumer with this email or phone already exist")
	}

	newConsumer, err := c.storage.InsertConsumer(consumer)
	if err != nil {
		c.logger.Println(err)
		return nil, systemErr
	}

	return newConsumer, nil
}

// DeleteConsumer prepare consumer data for deleting.
func (c *service) DeleteConsumer(id string) (data string, err error) {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.logger.Println(err)
		return "", errWrongConsumerIDType
	}

	foundConsumer, err := c.storage.GetConsumerByID(idInt)
	if err != nil && err != sql.ErrNoRows {
		c.logger.Println(err)
		return "", systemErr
	}
	if foundConsumer == nil {
		return "", errConsumerWithIDNotFound
	}

	if err = c.storage.DeleteConsumer(idInt); err != nil {
		c.logger.Println(err)
		return "", systemErr
	}

	if err = c.storage.DeleteLocation(idInt); err != nil {
		c.logger.Println(err)
		return "", systemErr
	}

	return "Consumer deleted", nil
}

// UpdateConsumer prepare data for updating.
func (c *service) UpdateConsumer(consumer domain.Consumer, id string) (*domain.Consumer, error) {
	// todo: if updating phone number or email send otp first and then update

	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.logger.Println(err)
		return nil, errWrongConsumerIDType
	}

	param := domain.SearchParam{}
	param["id"] = id
	param["email"] = consumer.Email
	param["phone"] = consumer.Phone
	foundConsumer, err := c.storage.GetConsumerDuplicateByParam(param)
	if err != nil && err != sql.ErrNoRows {
		c.logger.Println(err)
		return nil, systemErr
	}
	if foundConsumer != nil {
		return nil, fmt.Errorf("consumer with this email or phone already exist")
	}

	consumer.ID = idInt

	updatedConsumer, err := c.storage.UpdateConsumer(consumer)

	if err != nil && err != sql.ErrNoRows {
		c.logger.Println(err)
		return nil, systemErr
	}

	return updatedConsumer, nil
}

// GetAllConsumer prepare data to get it from storage.
func (c *service) GetAllConsumer() ([]domain.Consumer, error) {
	allConsumer, err := c.storage.GetAllConsumer()
	if err != nil {
		c.logger.Println(err)
		return nil, systemErr
	}

	return allConsumer, nil
}

// GetConsumer prepare data to get it from customerStorage.
func (c *service) GetConsumer(id string) (*domain.Consumer, error) {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.logger.Println(err)
		return nil, errWrongConsumerIDType
	}

	consumer, err := c.storage.GetConsumerByID(idInt)
	if err != nil && err == sql.ErrNoRows {
		return nil, err
	}

	if err != nil && err != sql.ErrNoRows {
		c.logger.Println(err)
		return nil, systemErr
	}

	return consumer, nil
}

// InsertLocation prepare and send data to storage service.
func (c *service) InsertLocation(location domain.Location, userID string) (*domain.Location, error) {
	userIDInt, err := strconv.Atoi(userID)
	if err != nil {
		c.logger.Println(err)
		return nil, errWrongConsumerIDType
	}

	foundConsumer, err := c.storage.GetConsumerByID(userIDInt)
	if err != nil && err != sql.ErrNoRows {
		c.logger.Println(err)
		return nil, systemErr
	}
	if foundConsumer == nil {
		return nil, fmt.Errorf("couldn't find consumer with this id")
	}

	foundConsumerLocation, err := c.storage.GetLocation(userIDInt)
	if err != nil && err != sql.ErrNoRows {
		c.logger.Println(err)
		return nil, systemErr
	}

	location.UserID = userIDInt

	if foundConsumerLocation != nil {
		return nil, fmt.Errorf("location already exist: instead update old one")
	}

	newLocation, err := c.storage.InsertLocation(location)
	if err != nil {
		c.logger.Println(err)
		return nil, systemErr
	}

	return newLocation, nil
}

// UpdateLocation prepare data for updating.
func (c *service) UpdateLocation(location domain.Location, userID string) (*domain.Location, error) {
	userIDInt, err := strconv.Atoi(userID)
	if err != nil {
		c.logger.Println(err)
		return nil, errWrongConsumerIDType
	}

	location.UserID = userIDInt

	updatedLocation, err := c.storage.UpdateLocation(location)
	if err != nil && err != sql.ErrNoRows {
		c.logger.Println(err)
		return nil, systemErr
	}

	return updatedLocation, nil
}

// GetLocation prepare data to get it from customerStorage.
func (c *service) GetLocation(userID string) (*domain.Location, error) {
	userIDInt, err := strconv.Atoi(userID)
	if err != nil {
		c.logger.Println(err)
		return nil, errWrongConsumerIDType
	}

	location, err := c.storage.GetLocation(userIDInt)
	if err != nil && err != sql.ErrNoRows {
		c.logger.Println(err)
		return nil, systemErr
	}
	if err != nil && err == sql.ErrNoRows {
		return nil, nil
	}

	return location, nil
}
