package consumerservice

import (
	"consumer/pkg/domain"
	"consumer/pkg/service"
	"database/sql"
	"fmt"
	"strconv"

	"github.com/nndergunov/deliveryApp/app/pkg/logger"
)

// ConsumerService is the interface for the user service.
type ConsumerService interface {
	InsertConsumer(consumer domain.Consumer) (*domain.Consumer, error)
	DeleteConsumer(id string) (data any, err error)
	UpdateConsumer(consumer domain.Consumer, id string) (*domain.Consumer, error)
	GetAllConsumer() ([]domain.Consumer, error)
	GetConsumer(id string) (*domain.Consumer, error)

	InsertConsumerLocation(consumer domain.ConsumerLocation, id string) (*domain.ConsumerLocation, error)
	UpdateConsumerLocation(consumer domain.ConsumerLocation, id string) (*domain.ConsumerLocation, error)
	GetConsumerLocation(id string) (*domain.ConsumerLocation, error)
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
func (c *consumerService) DeleteConsumer(id string) (data any, err error) {
	idUint, err := strconv.ParseUint(string(id), 10, 64)
	if err != nil {
		c.logger.Println(err)
		return nil, errWrongConsumerIDType
	}

	foundConsumer, err := c.consumerStorage.GetConsumerByID(idUint)
	if err != nil && err != sql.ErrNoRows {
		c.logger.Println(err)
		return nil, systemErr
	}
	if foundConsumer == nil {
		return nil, errConsumerWithIDNotFound
	}

	if err = c.consumerStorage.DeleteConsumer(idUint); err != nil {
		c.logger.Println(err)
		return nil, systemErr
	}

	if err = c.consumerStorage.DeleteConsumerLocation(idUint); err != nil {
		c.logger.Println(err)
		return nil, systemErr
	}

	return "Consumer deleted", nil
}

// UpdateConsumer prepare data for updating.
func (c *consumerService) UpdateConsumer(consumer domain.Consumer, id string) (*domain.Consumer, error) {
	//todo: if updating phone number or email send otp first and then update

	idUint, err := strconv.ParseUint(string(id), 10, 64)
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

	consumer.ID = idUint

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
	idUint, err := strconv.ParseUint(string(id), 10, 64)
	if err != nil {
		c.logger.Println(err)
		return nil, errWrongConsumerIDType
	}

	consumer, err := c.consumerStorage.GetConsumerByID(idUint)
	if err != nil && err == sql.ErrNoRows {
		return &domain.Consumer{}, nil
	}

	if err != nil && err != sql.ErrNoRows {
		c.logger.Println(err)
		return nil, systemErr
	}

	return consumer, nil

}

// InsertConsumerLocation prepare and send data to consumerStorage service.
func (c *consumerService) InsertConsumerLocation(consumerLocation domain.ConsumerLocation, id string) (*domain.ConsumerLocation, error) {
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.logger.Println(err)
		return nil, errWrongConsumerIDType
	}

	foundConsumer, err := c.consumerStorage.GetConsumerByID(idUint)
	if err != nil && err != sql.ErrNoRows {
		c.logger.Println(err)
		return nil, systemErr
	}
	if foundConsumer == nil {
		return nil, fmt.Errorf("couldn't find consumer with this id")
	}

	foundConsumerLocation, err := c.consumerStorage.GetConsumerLocation(idUint)
	if err != nil && err != sql.ErrNoRows {
		c.logger.Println(err)
		return nil, systemErr
	}

	consumerLocation.ConsumerID = idUint

	if foundConsumerLocation != nil {
		return nil, fmt.Errorf("consumer location already exist: instead update old one")
	}

	newConsumerLocation, err := c.consumerStorage.InsertConsumerLocation(consumerLocation)
	if err != nil {
		c.logger.Println(err)
		return nil, systemErr
	}

	return newConsumerLocation, nil
}

// UpdateConsumerLocation prepare data for updating.
func (c *consumerService) UpdateConsumerLocation(consumerLocation domain.ConsumerLocation, consumerID string) (*domain.ConsumerLocation, error) {
	cidUint, err := strconv.ParseUint(string(consumerID), 10, 64)
	if err != nil {
		c.logger.Println(err)
		return nil, errWrongConsumerIDType
	}

	consumerLocation.ConsumerID = cidUint

	updatedConsumerLocation, err := c.consumerStorage.UpdateConsumerLocation(consumerLocation)
	if err != nil && err != sql.ErrNoRows {
		c.logger.Println(err)
		return nil, systemErr
	}

	return updatedConsumerLocation, nil
}

// GetConsumerLocation prepare data to get it from customerStorage.
func (c *consumerService) GetConsumerLocation(id string) (*domain.ConsumerLocation, error) {
	idUint, err := strconv.ParseUint(string(id), 10, 64)
	if err != nil {
		c.logger.Println(err)
		return nil, errWrongConsumerIDType
	}

	consumerLocation, err := c.consumerStorage.GetConsumerLocation(idUint)
	if err != nil && err != sql.ErrNoRows {
		c.logger.Println(err)
		return nil, systemErr
	}

	return consumerLocation, nil

}
