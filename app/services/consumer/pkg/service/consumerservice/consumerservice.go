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
	GetConsumerLocation(id string) (*domain.ConsumerLocation, error)
	UpdateConsumerLocation(consumer domain.ConsumerLocation, id string) (*domain.ConsumerLocation, error)
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

	foundConsumer, err := c.consumerStorage.GetConsumer(0, consumer.Phone, consumer.Email)
	if err != nil && err != sql.ErrNoRows {
		c.logger.Println(err)
		return nil, nil
	}
	if foundConsumer != nil {
		return nil, fmt.Errorf("consumer with this email or phone already exist")
	}

	newConsumer, err := c.consumerStorage.InsertConsumer(consumer)
	if err != nil {
		c.logger.Println(err)
		return nil, err
	}

	return newConsumer, nil
}

// DeleteConsumer prepare consumer data for deleting.
func (c *consumerService) DeleteConsumer(id string) (data any, err error) {
	idUint, err := strconv.ParseUint(string(id), 10, 64)
	if err != nil {
		c.logger.Println(err)
		return nil, fmt.Errorf("wrong id type")
	}

	foundConsumer, err := c.consumerStorage.GetConsumer(idUint, "", "")
	if err != nil && err != sql.ErrNoRows {
		c.logger.Println(err)
		return nil, nil
	}
	if foundConsumer == nil {
		return nil, fmt.Errorf("consumer not found")
	}

	if err = c.consumerStorage.DeleteConsumer(idUint); err != nil {
		c.logger.Println(err)
		return nil, err
	}

	if err = c.consumerStorage.DeleteConsumerLocation(idUint); err != nil {
		c.logger.Println(err)
		return nil, err
	}

	return "Consumer deleted", nil
}

// UpdateConsumer prepare data for updating.
func (c *consumerService) UpdateConsumer(consumer domain.Consumer, id string) (*domain.Consumer, error) {
	//todo: if updating phone number or email send otp first and then update

	idUint, err := strconv.ParseUint(string(id), 10, 64)
	if err != nil {
		c.logger.Println(err)
		return nil, fmt.Errorf("wrong id type")
	}

	foundConsumer, err := c.consumerStorage.GetConsumer(0, consumer.Phone, consumer.Email)
	if err != nil && err != sql.ErrNoRows {
		c.logger.Println(err)
		return nil, nil
	}
	if foundConsumer != nil {
		return nil, fmt.Errorf("consumer with this email or phone already exist")
	}

	consumer.ID = idUint

	updatedConsumer, err := c.consumerStorage.UpdateConsumer(consumer)
	if err != nil && err == sql.ErrNoRows {
		return nil, fmt.Errorf("couldn't update consumer")
	}

	if err != nil && err != sql.ErrNoRows {
		c.logger.Println(err)
		return nil, nil
	}

	return updatedConsumer, nil
}

// GetAllConsumer prepare data to get it from consumerStorage.
func (c *consumerService) GetAllConsumer() ([]domain.Consumer, error) {
	allConsumer, err := c.consumerStorage.GetAllConsumer()
	if err != nil {
		c.logger.Println(err)
		return nil, nil
	}

	return allConsumer, nil
}

// GetConsumer prepare data to get it from customerStorage.
func (c *consumerService) GetConsumer(id string) (*domain.Consumer, error) {
	idUint, err := strconv.ParseUint(string(id), 10, 64)
	if err != nil {
		c.logger.Println(err)
		return nil, fmt.Errorf("wrong id type")
	}

	consumer, err := c.consumerStorage.GetConsumer(idUint, "", "")
	if err != nil && err == sql.ErrNoRows {
		return &domain.Consumer{}, nil
	}

	if err != nil && err != sql.ErrNoRows {
		c.logger.Println(err)
		return nil, nil
	}

	return consumer, nil

}

// InsertConsumerLocation prepare and send data to consumerStorage service.
func (c *consumerService) InsertConsumerLocation(consumerLocation domain.ConsumerLocation, id string) (*domain.ConsumerLocation, error) {
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.logger.Println(err)
		return nil, fmt.Errorf("wrong consumer_id type")
	}

	foundConsumer, err := c.consumerStorage.GetConsumer(idUint, "", "")
	if err != nil && err != sql.ErrNoRows {
		c.logger.Println(err)
		return nil, nil
	}
	if foundConsumer != nil {
		return nil, fmt.Errorf("couldn't find consumer with this id")
	}

	foundConsumerLocation, err := c.consumerStorage.GetConsumerLocation(idUint)
	if err != nil && err != sql.ErrNoRows {
		c.logger.Println(err)
		return nil, nil
	}
	if foundConsumerLocation != nil {
		return nil, fmt.Errorf("consumer location already exist: please update old one")
	}

	newConsumerLocation, err := c.consumerStorage.InsertConsumerLocation(consumerLocation)
	if err != nil {
		c.logger.Println(err)
		return nil, err
	}

	return newConsumerLocation, nil
}

// GetConsumerLocation prepare data to get it from customerStorage.
func (c *consumerService) GetConsumerLocation(id string) (*domain.ConsumerLocation, error) {
	idUint, err := strconv.ParseUint(string(id), 10, 64)
	if err != nil {
		c.logger.Println(err)
		return nil, fmt.Errorf("wrong consumer_id type")
	}

	consumerLocation, err := c.consumerStorage.GetConsumerLocation(idUint)
	if err != nil && err == sql.ErrNoRows {
		return &domain.ConsumerLocation{}, nil
	}

	if err != nil && err != sql.ErrNoRows {
		c.logger.Println(err)
		return nil, nil
	}

	return consumerLocation, nil

}

// UpdateConsumerLocation prepare data for updating.
func (c *consumerService) UpdateConsumerLocation(consumerLocation domain.ConsumerLocation, consumerID string) (*domain.ConsumerLocation, error) {
	cidUint, err := strconv.ParseUint(string(consumerID), 10, 64)
	if err != nil {
		c.logger.Println(err)
		return nil, fmt.Errorf("wrong consumer_id type")
	}

	consumerLocation.ConsumerID = cidUint

	updatedConsumerLocation, err := c.consumerStorage.UpdateConsumerLocation(consumerLocation)
	if err != nil && err == sql.ErrNoRows {
		c.logger.Println(err)
		return nil, fmt.Errorf("couldn't update consumerLocation")
	}

	if err != nil && err != sql.ErrNoRows {
		c.logger.Println(err)
		return nil, nil
	}

	return updatedConsumerLocation, nil
}
