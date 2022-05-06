package services

import (
	"consumer/internal/database/storage"
	"database/sql"
	"net/http"
	"strconv"

	"consumer/internal/models"
	"github.com/nndergunov/deliveryApp/app/pkg/logger"
)

// ConsumerService is the interface for the user services.
type ConsumerService interface {
	InsertNewConsumer(consumer models.Consumer) (data any, statusCode int)
	RemoveConsumer(id string) (data any, statusCode int)
	UpdateConsumer(consumer models.Consumer) (data any, statusCode int)
	GetAllConsumer() (data any, statusCode int)
	GetConsumer(id string) (data any, statusCode int)

	UpdateConsumerLocation(consumerLocation models.ConsumerLocation) (data any, statusCode int)
}

// Params is the input parameter struct for the module that contains its dependencies
type Params struct {
	CustomerStorage storage.ConsumerStorage
	Logger          *logger.Logger
}

type customerService struct {
	customerStorage storage.ConsumerStorage
	logger          *logger.Logger
}

// NewCustomerService constructs a new NewCustomerService.
func NewCustomerService(p Params) (ConsumerService, error) {
	consumerServiceItem := &customerService{
		customerStorage: p.CustomerStorage,
		logger:          p.Logger,
	}

	return consumerServiceItem, nil
}

// InsertNewConsumer prepare and send data to consumerStorage services.
func (c *customerService) InsertNewConsumer(consumer models.Consumer) (data any, statusCode int) {

	if consumer.PhoneNumber == "" || consumer.Email == "" {
		return "wrong consumer credential received", http.StatusBadRequest
	}

	if consumer.PhoneNumber != "" && consumer.CountryCode == "" {
		return "wrong consumer credential received", http.StatusBadRequest
	}
	//todo: 1.check countryCode is exist in map 2. phone/email - regex

	foundConsumer, err := c.customerStorage.GetConsumer(0, consumer.PhoneNumber, consumer.Email, "active")
	if err != nil && err != sql.ErrNoRows {
		c.logger.Println(err)
		return "", http.StatusInternalServerError
	}
	if foundConsumer != nil {
		return "consumer with this credential already exist", http.StatusBadRequest
	}

	newConsumer, err := c.customerStorage.InsertConsumer(consumer)
	if err != nil {
		c.logger.Println(err)
		return err, http.StatusBadRequest
	}

	return newConsumer, http.StatusOK
}

// RemoveConsumer prepare consumer data for removing.
func (c *customerService) RemoveConsumer(id string) (data any, statusCode int) {
	idUint, err := strconv.ParseUint(string(id), 10, 64)
	if err != nil {
		c.logger.Println(err)
		return "wrong id type", http.StatusBadRequest
	}

	foundConsumer, err := c.customerStorage.GetConsumer(idUint, "", "", "active")
	if err != nil && err != sql.ErrNoRows {
		c.logger.Println(err)
		return "", http.StatusInternalServerError
	}
	if foundConsumer == nil {
		return "consumer not found", http.StatusBadRequest
	}

	if err = c.customerStorage.RemoveConsumer(idUint); err != nil {
		c.logger.Println(err)
		return err, http.StatusInternalServerError
	}

	return "Consumer removed", http.StatusOK
}

// UpdateConsumer prepare data for updating.
func (c *customerService) UpdateConsumer(consumer models.Consumer) (data any, statusCode int) {
	//todo: if updating phone number or email send otp first and then update

	updatedConsumer, err := c.customerStorage.UpdateConsumer(consumer)

	if err != nil && err != sql.ErrNoRows {
		return "", http.StatusInternalServerError
	}

	return updatedConsumer, http.StatusOK
}

// GetAllConsumer prepare data to get it from customerStorage.
func (c *customerService) GetAllConsumer() (data any, statusCode int) {
	allConsumer, err := c.customerStorage.GetAllConsumer()
	if err != nil {
		c.logger.Println(err)
		return err, http.StatusInternalServerError
	}

	return allConsumer, http.StatusOK
}

// GetConsumer prepare data to get it from customerStorage.
func (c *customerService) GetConsumer(id string) (data any, statusCode int) {
	idUint, err := strconv.ParseUint(string(id), 10, 64)
	if err != nil {
		c.logger.Println(err)
		return "wrong id type", http.StatusBadRequest
	}

	consumer, err := c.customerStorage.GetConsumer(idUint, "", "", "active")
	if err != nil && err == sql.ErrNoRows {
		return "no consumer found", http.StatusOK
	}
	if err != nil && err != sql.ErrNoRows {
		c.logger.Println(err)
		return "", http.StatusInternalServerError
	}

	return consumer, http.StatusOK
}

// UpdateConsumerLocation prepare data for updating.
func (c *customerService) UpdateConsumerLocation(consumerLocation models.ConsumerLocation) (data any, statusCode int) {
	//todo: check input data

	updatedConsumerLocation, err := c.customerStorage.UpdateConsumerLocation(consumerLocation)
	if err != nil && err != sql.ErrNoRows {
		return "couldn't create consumer location", http.StatusInternalServerError
	}

	return updatedConsumerLocation, http.StatusOK
}
