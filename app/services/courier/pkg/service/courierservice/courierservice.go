package courierservice

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/nndergunov/deliveryApp/app/services/courier/pkg/domain"
	"github.com/nndergunov/deliveryApp/app/services/courier/pkg/service"

	"github.com/nndergunov/deliveryApp/app/pkg/logger"
)

// CourierService is the interface for the user service.
type CourierService interface {
	InsertCourier(courier domain.Courier) (*domain.Courier, error)
	DeleteCourier(id string) (data string, err error)
	UpdateCourier(courier domain.Courier, id string) (*domain.Courier, error)
	UpdateCourierAvailable(id, available string) (*domain.Courier, error)
	GetCourierList(params domain.SearchParam) ([]domain.Courier, error)
	GetCourier(id string) (*domain.Courier, error)

	InsertLocation(courier domain.Location, userID string) (*domain.Location, error)
	UpdateLocation(courier domain.Location, id string) (*domain.Location, error)
	GetLocation(userID string) (*domain.Location, error)
	GetLocationList(param domain.SearchParam) ([]domain.Location, error)
}

// Params is the input parameter struct for the module that contains its dependencies
type Params struct {
	CourierStorage service.CourierStorage
	Logger         *logger.Logger
}

type courierService struct {
	courierStorage service.CourierStorage
	logger         *logger.Logger
}

// NewCourierService constructs a new NewCourierService.
func NewCourierService(p Params) CourierService {
	courierServiceItem := &courierService{
		courierStorage: p.CourierStorage,
		logger:         p.Logger,
	}

	return courierServiceItem
}

// InsertCourier prepare and send data to courierStorage service.
func (c *courierService) InsertCourier(courier domain.Courier) (*domain.Courier, error) {
	if len(courier.Username) < 4 || len(courier.Password) < 8 {
		return nil, fmt.Errorf("username or password don't meet requirement")
	}
	param := domain.SearchParam{}
	param["username"] = courier.Username

	foundCourier, err := c.courierStorage.GetCourierDuplicateByParam(param)
	if err != nil && err != sql.ErrNoRows {
		c.logger.Println(err)
		return nil, systemErr
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

// DeleteCourier prepare courier data for removing.
func (c *courierService) DeleteCourier(id string) (data string, err error) {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.logger.Println(err)
		return "", errWrongCourierIDType
	}

	foundCourier, err := c.courierStorage.GetCourierByID(idInt)
	if err != nil && err != sql.ErrNoRows {
		c.logger.Println(err)
		return "", systemErr
	}
	if foundCourier == nil {
		return "", errCourierWithIDNotFound
	}

	if err = c.courierStorage.DeleteCourier(idInt); err != nil {
		c.logger.Println(err)
		return "", err
	}

	if err = c.courierStorage.DeleteLocation(idInt); err != nil {
		c.logger.Println(err)
		return "", err
	}

	return "courier deleted", nil
}

// UpdateCourier prepare data for updating.
func (c *courierService) UpdateCourier(courier domain.Courier, id string) (*domain.Courier, error) {
	// todo: if updating phone number or email send otp first and then update

	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.logger.Println(err)
		return nil, errWrongCourierIDType
	}

	param := domain.SearchParam{}
	param["id"] = id
	// check duplicate by:

	// username
	if courier.Username != "" {

		param["username"] = courier.Username

		foundCourier, err := c.courierStorage.GetCourierDuplicateByParam(param)
		if err != nil && err != sql.ErrNoRows {
			c.logger.Println(err)
			return nil, nil
		}
		if foundCourier != nil {
			return nil, fmt.Errorf("courier with this username already exist")
		}
	}
	// email
	if courier.Email != "" {

		param["email"] = courier.Email

		foundCourier, err := c.courierStorage.GetCourierDuplicateByParam(param)
		if err != nil && err != sql.ErrNoRows {
			c.logger.Println(err)
			return nil, systemErr
		}
		if foundCourier != nil {
			return nil, fmt.Errorf("courier with this email already exist")
		}
	}
	// phone
	if courier.Phone != "" {

		param["phone"] = courier.Phone

		foundCourier, err := c.courierStorage.GetCourierDuplicateByParam(param)
		if err != nil && err != sql.ErrNoRows {
			c.logger.Println(err)
			return nil, systemErr
		}
		if foundCourier != nil {
			return nil, fmt.Errorf("courier with this phone already exist")
		}
	}

	courier.ID = idInt

	updatedCourier, err := c.courierStorage.UpdateCourier(courier)
	if err != nil && err != sql.ErrNoRows {
		c.logger.Println(err)
		return nil, systemErr
	}

	return updatedCourier, nil
}

// UpdateCourierAvailable prepare data for updating.
func (c *courierService) UpdateCourierAvailable(id, available string) (*domain.Courier, error) {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.logger.Println(err)
		return nil, errWrongCourierIDType
	}

	availableBool, err := strconv.ParseBool(available)
	if err != nil {
		c.logger.Println(err)
		return nil, fmt.Errorf("wrong available type")
	}

	foundCourier, err := c.courierStorage.GetCourierByID(idInt)
	if err != nil && err != sql.ErrNoRows {
		c.logger.Println(err)
		return nil, systemErr
	}
	if foundCourier == nil {
		return nil, errCourierWithIDNotFound
	}

	updatedCourier, err := c.courierStorage.UpdateCourierAvailable(idInt, availableBool)
	if err != nil && err == sql.ErrNoRows {
		c.logger.Println(err)
		return nil, fmt.Errorf("couldn't update courier")
	}

	if err != nil && err != sql.ErrNoRows {
		return nil, nil
	}

	return updatedCourier, nil
}

// GetCourierList prepare data to get it from courierStorage.
func (c *courierService) GetCourierList(param domain.SearchParam) ([]domain.Courier, error) {
	availableStr := param["available"]
	if availableStr != "" {
		if _, err := strconv.ParseBool(availableStr); err != nil {
			c.logger.Println(err)
			return nil, fmt.Errorf("wrong available type")

		}
	}

	courierList, err := c.courierStorage.GetCourierList(param)
	if err != nil {
		c.logger.Println(err)
		return nil, systemErr
	}

	return courierList, nil
}

// GetCourier prepare data to get it from courierStorage.
func (c *courierService) GetCourier(id string) (*domain.Courier, error) {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.logger.Println(err)
		return nil, errWrongCourierIDType
	}

	courier, err := c.courierStorage.GetCourierByID(idInt)
	if err != nil && err == sql.ErrNoRows {
		return nil, errCourierWithIDNotFound
	}

	if err != nil && err != sql.ErrNoRows {
		c.logger.Println(err)
		return nil, systemErr
	}

	return courier, nil
}

// InsertLocation prepare and send data to courierStorage service.
func (c *courierService) InsertLocation(location domain.Location, userID string) (*domain.Location, error) {
	userIDInt, err := strconv.Atoi(userID)
	if err != nil {
		c.logger.Println(err)
		return nil, errWrongCourierIDType
	}

	foundCourier, err := c.courierStorage.GetCourierByID(userIDInt)
	if err != nil && err != sql.ErrNoRows {
		c.logger.Println(err)
		return nil, systemErr
	}
	if foundCourier == nil {
		return nil, errCourierWithIDNotFound
	}

	foundLocation, err := c.courierStorage.GetLocation(userIDInt)
	if err != nil && err != sql.ErrNoRows {
		c.logger.Println(err)
		return nil, systemErr
	}

	location.UserID = userIDInt

	if foundLocation != nil {
		return nil, fmt.Errorf("courier location already exist: please update old one")
	}

	newLocation, err := c.courierStorage.InsertLocation(location)
	if err != nil {
		c.logger.Println(err)
		return nil, err
	}

	return newLocation, nil
}

// UpdateLocation prepare data for updating.
func (c *courierService) UpdateLocation(location domain.Location, courierID string) (*domain.Location, error) {
	cidInt, err := strconv.Atoi(courierID)
	if err != nil {
		c.logger.Println(err)
		return nil, errWrongCourierIDType
	}

	location.UserID = cidInt
	gotLocation, err := c.courierStorage.GetLocation(location.UserID)
	if err != nil {
		c.logger.Println(err)
		return nil, fmt.Errorf("couldn't find location")
	}

	if gotLocation == nil {
		return nil, fmt.Errorf("couldn't find location")
	}

	updatedLocation, err := c.courierStorage.UpdateLocation(location)
	if err != nil && err == sql.ErrNoRows {
		c.logger.Println(err)
		return nil, fmt.Errorf("couldn't update location")
	}

	if err != nil && err != sql.ErrNoRows {
		c.logger.Println(err)
		return nil, systemErr
	}

	return updatedLocation, nil
}

// GetLocation prepare data to get it from customerStorage.
func (c *courierService) GetLocation(userID string) (*domain.Location, error) {
	userIDInt, err := strconv.Atoi(userID)
	if err != nil {
		c.logger.Println(err)
		return nil, errWrongCourierIDType
	}

	location, err := c.courierStorage.GetLocation(userIDInt)
	if err != nil && err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil && err != sql.ErrNoRows {
		c.logger.Println(err)
		return nil, systemErr
	}

	return location, nil
}

// GetLocationList prepare data to get it from courierStorage.
func (c *courierService) GetLocationList(param domain.SearchParam) ([]domain.Location, error) {
	locationList, err := c.courierStorage.GetLocationList(param)
	if err != nil {
		c.logger.Println(err)
		return nil, systemErr
	}

	return locationList, nil
}
