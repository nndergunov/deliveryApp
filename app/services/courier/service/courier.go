package service

import (
	"courier/domain"
	"database/sql"
	"fmt"
	"strconv"

	"github.com/nndergunov/deliveryApp/app/pkg/logger"
)

// Params is the input parameter struct for the module that contains its dependencies
type Params struct {
	CourierStorage CourierStorage
	Logger         *logger.Logger
}

type courierService struct {
	courierStorage CourierStorage
	logger         *logger.Logger
}

// NewCourierService constructs a new NewCourierService.
func NewCourierService(p Params) (CourierService, error) {
	courierServiceItem := &courierService{
		courierStorage: p.CourierStorage,
		logger:         p.Logger,
	}

	return courierServiceItem, nil
}

// InsertCourier prepare and send data to courierStorage service.
func (c *courierService) InsertCourier(courier domain.Courier) (*domain.Courier, error) {

	if len(courier.Username) < 4 || len(courier.Password) < 8 {
		return nil, fmt.Errorf("username or password don't meet requirement")
	}
	foundCourier, err := c.courierStorage.GetCourier(0, courier.Username, "")
	if err != nil && err != sql.ErrNoRows {
		c.logger.Println(err)
		return nil, nil
	}
	if foundCourier != nil {
		return nil, fmt.Errorf("courier with this username already exist")
	}

	newCourier, err := c.courierStorage.InsertCourier(courier)
	if err != nil {
		c.logger.Println(err)
		return nil, err
	}

	return newCourier, nil
}

// RemoveCourier prepare courier data for removing.
func (c *courierService) RemoveCourier(id string) (data any, err error) {
	idUint, err := strconv.ParseUint(string(id), 10, 64)
	if err != nil {
		c.logger.Println(err)
		return nil, fmt.Errorf("wrong id type")
	}

	foundCourier, err := c.courierStorage.GetCourier(idUint, "", "active")
	if err != nil && err != sql.ErrNoRows {
		c.logger.Println(err)
		return nil, nil
	}
	if foundCourier == nil {
		return nil, fmt.Errorf("active courier with this id not found")
	}

	if err = c.courierStorage.RemoveCourier(idUint); err != nil {
		c.logger.Println(err)
		return nil, err
	}

	return "courier removed", nil
}

// UpdateCourier prepare data for updating.
func (c *courierService) UpdateCourier(courier domain.Courier, id string) (*domain.Courier, error) {
	idUint, err := strconv.ParseUint(string(id), 10, 64)
	if err != nil {
		c.logger.Println(err)
		return nil, fmt.Errorf("wrong id type")
	}

	foundCourier, err := c.courierStorage.GetCourier(0, courier.Username, "")
	if err != nil && err != sql.ErrNoRows {
		c.logger.Println(err)
		return nil, nil
	}
	if foundCourier != nil {
		return nil, fmt.Errorf("courier with this username already exist")
	}

	courier.ID = idUint

	updatedCourier, err := c.courierStorage.UpdateCourier(courier)
	if err != nil && err == sql.ErrNoRows {
		c.logger.Println(err)
		return nil, fmt.Errorf("couldn't update courier")
	}

	if err != nil && err != sql.ErrNoRows {
		return nil, nil
	}

	return updatedCourier, nil
}

// UpdateCourierAvailable prepare data for updating.
func (c *courierService) UpdateCourierAvailable(id, available string) (*domain.Courier, error) {
	idUint, err := strconv.ParseUint(string(id), 10, 64)
	if err != nil {
		c.logger.Println(err)
		return nil, fmt.Errorf("wrong id type")
	}

	availableBool, err := strconv.ParseBool(available)
	if err != nil {
		c.logger.Println(err)
		return nil, fmt.Errorf("wrong available type")
	}

	updatedCourier, err := c.courierStorage.UpdateCourierAvailable(idUint, availableBool)
	if err != nil && err == sql.ErrNoRows {
		c.logger.Println(err)
		return nil, fmt.Errorf("couldn't update courier")
	}

	if err != nil && err != sql.ErrNoRows {
		return nil, nil
	}

	return updatedCourier, nil
}

// GetAllCourier prepare data to get it from courierStorage.
func (c *courierService) GetAllCourier() ([]domain.Courier, error) {
	allCourier, err := c.courierStorage.GetAllCourier()
	if err != nil {
		c.logger.Println(err)
		return nil, nil
	}

	return allCourier, nil
}

// GetCourier prepare data to get it from courierStorage.
func (c *courierService) GetCourier(id string) (*domain.Courier, error) {
	idUint, err := strconv.ParseUint(string(id), 10, 64)
	if err != nil {
		c.logger.Println(err)
		return nil, fmt.Errorf("wrong id type")
	}

	courier, err := c.courierStorage.GetCourier(idUint, "", "active")
	if err != nil && err == sql.ErrNoRows {
		return nil, fmt.Errorf("no courier found")
	}
	if err != nil && err != sql.ErrNoRows {
		c.logger.Println(err)
		return nil, nil
	}

	return courier, nil
}
