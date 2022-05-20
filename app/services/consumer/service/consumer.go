package service

import (
	"consumer/domain"
	"database/sql"
	"fmt"
	"strconv"

	"github.com/nndergunov/deliveryApp/app/pkg/logger"
)

// Params is the input parameter struct for the module that contains its dependencies
type Params struct {
	ConsumerStorage ConsumerStorage
	Logger          *logger.Logger
}

type consumerService struct {
	consumerStorage ConsumerStorage
	logger          *logger.Logger
}

// NewConsumerService constructs a new NewConsumerService.
func NewConsumerService(p Params) (ConsumerService, error) {
	consumerServiceItem := &consumerService{
		consumerStorage: p.ConsumerStorage,
		logger:          p.Logger,
	}

	return consumerServiceItem, nil
}

// InsertConsumer prepare and send data to consumerStorage service.
func (c *consumerService) InsertConsumer(consumer domain.Consumer) (*domain.Consumer, error) {

	if consumer.Phone == "" && consumer.Email == "" {
		return nil, fmt.Errorf("wrong phone or email credential received")
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
		c.logger.Println(err)
		return nil, fmt.Errorf("couldn't update consumer")
	}

	if err != nil && err != sql.ErrNoRows {
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
		return nil, fmt.Errorf("no consumer found")
	}
	if err != nil && err != sql.ErrNoRows {
		c.logger.Println(err)
		return nil, nil
	}

	return consumer, nil

}

// UpdateConsumerLocation prepare data for updating.
func (c *consumerService) UpdateConsumerLocation(consumerLocation domain.ConsumerLocation, consumerID string) (*domain.ConsumerLocation, error) {
	cidUint, err := strconv.ParseUint(string(consumerID), 10, 64)
	if err != nil {
		c.logger.Println(err)
		return nil, fmt.Errorf("wrong consumerID type")
	}

	foundConsumer, err := c.consumerStorage.GetConsumer(cidUint, "", "")
	if err != nil && err != sql.ErrNoRows {
		c.logger.Println(err)
		return nil, nil
	}
	if foundConsumer == nil {
		return nil, fmt.Errorf("consumer with this id not found")
	}

	consumerLocation.ID = foundConsumer.ConsumerLocation.ID

	updatedConsumerLocation, err := c.consumerStorage.UpdateConsumerLocation(consumerLocation)
	if err != nil && err == sql.ErrNoRows {
		c.logger.Println(err)
		return nil, fmt.Errorf("couldn't update consumerLocation")
	}

	if err != nil && err != sql.ErrNoRows {
		return nil, nil
	}

	return updatedConsumerLocation, nil
}
