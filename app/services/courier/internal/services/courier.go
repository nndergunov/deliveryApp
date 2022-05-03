package services

import (
	"database/sql"
	"net/http"
	"strconv"

	"courier/internal/models"
	"courier/internal/storage"

	"github.com/nndergunov/deliveryApp/app/pkg/logger"
)

// CourierService is the interface for the user services.
type CourierService interface {
	InsertNewCourier(courier models.Courier) (data any, statusCode int)
	RemoveCourier(id string) (data any, statusCode int)
	UpdateCourier(courier models.Courier) (data any, statusCode int)
	GetAllCourier() (data any, statusCode int)
	GetCourier(id string) (data any, statusCode int)
}

// Params is the input parameter struct for the module that contains its dependencies
type Params struct {
	CourierStorage storage.CourierStorage
	Logger         *logger.Logger
}

type courierService struct {
	courierStorage storage.CourierStorage
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

// InsertNewCourier prepare and send data to courierStorage services.
func (c courierService) InsertNewCourier(courier models.Courier) (data any, statusCode int) {

	if len(courier.Username) < 4 || len(courier.Password) < 8 {
		return "username or password don't meet requirement", http.StatusBadRequest
	}
	foundCourier, err := c.courierStorage.GetCourier(0, courier.Username, "active")
	if err != nil && err != sql.ErrNoRows {
		c.logger.Println(err)
		return "", http.StatusInternalServerError
	}
	if foundCourier != nil {
		return "courier with this username already exist", http.StatusBadRequest
	}

	newCourier, err := c.courierStorage.InsertCourier(courier)
	if err != nil {
		c.logger.Println(err)
		return err, http.StatusBadRequest
	}

	return newCourier, http.StatusOK
}

// RemoveCourier prepare courier data for removing.
func (c courierService) RemoveCourier(id string) (data any, statusCode int) {
	idUint, err := strconv.ParseUint(string(id), 10, 64)
	if err != nil {
		c.logger.Println(err)
		return "wrong id type", http.StatusBadRequest
	}

	foundCourier, err := c.courierStorage.GetCourier(idUint, "", "active")
	if err != nil && err != sql.ErrNoRows {
		c.logger.Println(err)
		return "", http.StatusInternalServerError
	}
	if foundCourier == nil {
		return "wrong id", http.StatusBadRequest
	}

	if err = c.courierStorage.RemoveCourier(idUint); err != nil {
		c.logger.Println(err)
		return err, http.StatusInternalServerError
	}

	return "Courier removed", http.StatusOK
}

// UpdateCourier prepare data for updating.
func (c courierService) UpdateCourier(courier models.Courier) (data any, statusCode int) {
	foundCourier, err := c.courierStorage.GetCourier(0, courier.Username, "")
	if err != nil && err != sql.ErrNoRows {
		c.logger.Println(err)
		return "", http.StatusInternalServerError
	}
	if foundCourier != nil {
		return "courier with this username already exist", http.StatusBadRequest
	}

	updatedCourier, err := c.courierStorage.UpdateCourier(courier)
	if err != nil && err == sql.ErrNoRows {
		c.logger.Println(err)
		return "this courier doesn't exist", http.StatusOK
	}

	if err != nil && err != sql.ErrNoRows {
		return "", http.StatusInternalServerError
	}

	return updatedCourier, http.StatusOK
}

// GetAllCourier prepare data to get it from courierStorage.
func (c courierService) GetAllCourier() (data any, statusCode int) {
	allCourier, err := c.courierStorage.GetAllCourier()
	if err != nil {
		c.logger.Println(err)
		return err, http.StatusInternalServerError
	}

	return allCourier, http.StatusOK
}

// GetCourier prepare data to get it from courierStorage.
func (c courierService) GetCourier(id string) (data any, statusCode int) {
	idUint, err := strconv.ParseUint(string(id), 10, 64)
	if err != nil {
		c.logger.Println(err)
		return "wrong id type", http.StatusBadRequest
	}

	courier, err := c.courierStorage.GetCourier(idUint, "", "active")
	if err != nil && err == sql.ErrNoRows {
		return "no courier found", http.StatusOK
	}
	if err != nil && err != sql.ErrNoRows {
		c.logger.Println(err)
		return "", http.StatusInternalServerError
	}

	return courier, http.StatusOK
}
