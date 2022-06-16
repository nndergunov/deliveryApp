package consumerservice

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/nndergunov/deliveryApp/app/services/consumer/pkg/domain"
	"github.com/nndergunov/deliveryApp/app/services/consumer/pkg/service"

	"github.com/nndergunov/deliveryApp/app/pkg/logger"
)

// ConsumerService is the interface for the user service.
type ConsumerService interface {
	InsertConsumer(consumer domain.Consumer) (*domain.Consumer, error)
	DeleteConsumer(id string) (data string, err error)
	UpdateConsumer(consumer domain.Consumer, id string) (*domain.Consumer, error)
	GetAllConsumer() ([]domain.Consumer, error)
	GetConsumer(id string) (*domain.Consumer, error)

	InsertLocation(consumer domain.Location, id string) (*domain.Location, error)
	UpdateLocation(consumer domain.Location, id string) (*domain.Location, error)
	GetLocation(id string) (*domain.Location, error)
}

// Params is the input parameter struct for the module that contains its dependencies
type Params struct {
	ConsumerStorage service.ConsumerStorage
	Logger          *logger.Logger
}

type consumerService struct {
	consumerStorage service.ConsumerStorage
	logger          *logger.Logger
}

// NewConsumerService constructs a new NewConsumerService.
func NewConsumerService(p Params) ConsumerService {
	consumerServiceItem := &consumerService{
		consumerStorage: p.ConsumerStorage,
		logger:          p.Logger,
	}

	return consumerServiceItem
}

// InsertConsumer prepare and send data to consumerStorage service.
func (c *consumerService) InsertConsumer(consumer domain.Consumer) (*domain.Consumer, error) {
	if consumer.Phone == "" && consumer.Email == "" {
		return nil, fmt.Errorf("wrong phone or email")
	}
	param := domain.SearchParam{}
	param["email"] = consumer.Email
	param["phone"] = consumer.Phone
	foundConsumer, err := c.consumerStorage.GetConsumerDuplicateByParam(param)
	if err != nil && err != sql.ErrNoRows {
		c.logger.Println(err)
		return nil, systemErr
	}
	if foundConsumer != nil {
		return nil, fmt.Errorf("consumer with this email or phone already exist")
	}

	newConsumer, err := c.consumerStorage.InsertConsumer(consumer)
	if err != nil {
		c.logger.Println(err)
		return nil, systemErr
	}

	return newConsumer, nil
}

// DeleteConsumer prepare consumer data for deleting.
func (c *consumerService) DeleteConsumer(id string) (data string, err error) {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.logger.Println(err)
		return "", errWrongConsumerIDType
	}

	foundConsumer, err := c.consumerStorage.GetConsumerByID(idInt)
	if err != nil && err != sql.ErrNoRows {
		c.logger.Println(err)
		return "", systemErr
	}
	if foundConsumer == nil {
		return "", errConsumerWithIDNotFound
	}

	if err = c.consumerStorage.DeleteConsumer(idInt); err != nil {
		c.logger.Println(err)
		return "", systemErr
	}

	if err = c.consumerStorage.DeleteLocation(idInt); err != nil {
		c.logger.Println(err)
		return "", systemErr
	}

	return "Consumer deleted", nil
}

// UpdateConsumer prepare data for updating.
func (c *consumerService) UpdateConsumer(consumer domain.Consumer, id string) (*domain.Consumer, error) {
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
	foundConsumer, err := c.consumerStorage.GetConsumerDuplicateByParam(param)
	if err != nil && err != sql.ErrNoRows {
		c.logger.Println(err)
		return nil, systemErr
	}
	if foundConsumer != nil {
		return nil, fmt.Errorf("consumer with this email or phone already exist")
	}

	consumer.ID = idInt

	updatedConsumer, err := c.consumerStorage.UpdateConsumer(consumer)

	if err != nil && err != sql.ErrNoRows {
		c.logger.Println(err)
		return nil, systemErr
	}

	return updatedConsumer, nil
}

// GetAllConsumer prepare data to get it from consumerStorage.
func (c *consumerService) GetAllConsumer() ([]domain.Consumer, error) {
	allConsumer, err := c.consumerStorage.GetAllConsumer()
	if err != nil {
		c.logger.Println(err)
		return nil, systemErr
	}

	return allConsumer, nil
}

// GetConsumer prepare data to get it from customerStorage.
func (c *consumerService) GetConsumer(id string) (*domain.Consumer, error) {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.logger.Println(err)
		return nil, errWrongConsumerIDType
	}

	consumer, err := c.consumerStorage.GetConsumerByID(idInt)
	if err != nil && err == sql.ErrNoRows {
		return nil, err
	}

	if err != nil && err != sql.ErrNoRows {
		c.logger.Println(err)
		return nil, systemErr
	}

	return consumer, nil
}

// InsertLocation prepare and send data to consumerStorage service.
func (c *consumerService) InsertLocation(location domain.Location, userID string) (*domain.Location, error) {
	userIDInt, err := strconv.Atoi(userID)
	if err != nil {
		c.logger.Println(err)
		return nil, errWrongConsumerIDType
	}

	foundConsumer, err := c.consumerStorage.GetConsumerByID(userIDInt)
	if err != nil && err != sql.ErrNoRows {
		c.logger.Println(err)
		return nil, systemErr
	}
	if foundConsumer == nil {
		return nil, fmt.Errorf("couldn't find consumer with this id")
	}

	foundConsumerLocation, err := c.consumerStorage.GetLocation(userIDInt)
	if err != nil && err != sql.ErrNoRows {
		c.logger.Println(err)
		return nil, systemErr
	}

	location.UserID = userIDInt

	if foundConsumerLocation != nil {
		return nil, fmt.Errorf("location already exist: instead update old one")
	}

	newLocation, err := c.consumerStorage.InsertLocation(location)
	if err != nil {
		c.logger.Println(err)
		return nil, systemErr
	}

	return newLocation, nil
}

// UpdateLocation prepare data for updating.
func (c *consumerService) UpdateLocation(location domain.Location, userID string) (*domain.Location, error) {
	userIDInt, err := strconv.Atoi(userID)
	if err != nil {
		c.logger.Println(err)
		return nil, errWrongConsumerIDType
	}

	location.UserID = userIDInt

	updatedLocation, err := c.consumerStorage.UpdateLocation(location)
	if err != nil && err != sql.ErrNoRows {
		c.logger.Println(err)
		return nil, systemErr
	}

	return updatedLocation, nil
}

// GetLocation prepare data to get it from customerStorage.
func (c *consumerService) GetLocation(userID string) (*domain.Location, error) {
	userIDInt, err := strconv.Atoi(userID)
	if err != nil {
		c.logger.Println(err)
		return nil, errWrongConsumerIDType
	}

	location, err := c.consumerStorage.GetLocation(userIDInt)
	if err != nil && err != sql.ErrNoRows {
		c.logger.Println(err)
		return nil, systemErr
	}

	return location, nil
}
