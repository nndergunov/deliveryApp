package courierservice

import (
	"database/sql"
	"fmt"
	"strconv"

	"courier/pkg/domain"
	"courier/pkg/service"

	"github.com/nndergunov/deliveryApp/app/pkg/logger"
)

// CourierService is the interface for the user service.
type CourierService interface {
	InsertCourier(courier domain.Courier) (*domain.Courier, error)
	DeleteCourier(id string) (data any, err error)
	UpdateCourier(courier domain.Courier, id string) (*domain.Courier, error)
	UpdateCourierAvailable(id, available string) (*domain.Courier, error)
	GetAllCourier(params map[string]string) ([]domain.Courier, error)
	GetCourier(id string) (*domain.Courier, error)

	InsertCourierLocation(courier domain.CourierLocation, id string) (*domain.CourierLocation, error)
	UpdateCourierLocation(courier domain.CourierLocation, id string) (*domain.CourierLocation, error)
	GetCourierLocation(id string) (*domain.CourierLocation, error)
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
func (c *courierService) DeleteCourier(id string) (data any, err error) {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.logger.Println(err)
		return nil, errWrongCourierIDType
	}

	foundCourier, err := c.courierStorage.GetCourierByID(idInt)
	if err != nil && err != sql.ErrNoRows {
		c.logger.Println(err)
		return nil, systemErr
	}
	if foundCourier == nil {
		return nil, errCourierWithIDNotFound
	}

	if err = c.courierStorage.DeleteCourier(idInt); err != nil {
		c.logger.Println(err)
		return nil, err
	}

	if err = c.courierStorage.DeleteCourierLocation(idInt); err != nil {
		c.logger.Println(err)
		return nil, err
	}

	return "courier deleted", nil
}

// UpdateCourier prepare data for updating.
func (c *courierService) UpdateCourier(courier domain.Courier, id string) (*domain.Courier, error) {
	//todo: if updating phone number or email send otp first and then update

	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.logger.Println(err)
		return nil, errWrongCourierIDType
	}

	param := domain.SearchParam{}
	param["id"] = id
	//check duplicate by:

	//username
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
	//email
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
	//phone
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

// GetAllCourier prepare data to get it from courierStorage.
func (c *courierService) GetAllCourier(param map[string]string) ([]domain.Courier, error) {

	availableStr := param["available"]
	if availableStr != "" {
		if _, err := strconv.ParseBool(availableStr); err != nil {
			c.logger.Println(err)
			return nil, fmt.Errorf("wrong available type")

		}
	}

	allCourier, err := c.courierStorage.GetAllCourier(param)
	if err != nil {
		c.logger.Println(err)
		return nil, systemErr
	}

	return allCourier, nil
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

// InsertCourierLocation prepare and send data to courierStorage service.
func (c *courierService) InsertCourierLocation(courierLocation domain.CourierLocation, id string) (*domain.CourierLocation, error) {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.logger.Println(err)
		return nil, errWrongCourierIDType
	}

	foundCourier, err := c.courierStorage.GetCourierByID(idInt)
	if err != nil && err != sql.ErrNoRows {
		c.logger.Println(err)
		return nil, systemErr
	}
	if foundCourier == nil {
		return nil, errCourierWithIDNotFound
	}

	foundCourierLocation, err := c.courierStorage.GetCourierLocation(idInt)
	if err != nil && err != sql.ErrNoRows {
		c.logger.Println(err)
		return nil, systemErr
	}

	courierLocation.CourierID = idInt

	if foundCourierLocation != nil {
		return nil, fmt.Errorf("courier location already exist: please update old one")
	}

	newCourierLocation, err := c.courierStorage.InsertCourierLocation(courierLocation)
	if err != nil {
		c.logger.Println(err)
		return nil, err
	}

	return newCourierLocation, nil
}

// UpdateCourierLocation prepare data for updating.
func (c *courierService) UpdateCourierLocation(courierLocation domain.CourierLocation, courierID string) (*domain.CourierLocation, error) {
	cidInt, err := strconv.Atoi(courierID)
	if err != nil {
		c.logger.Println(err)
		return nil, errWrongCourierIDType
	}

	courierLocation.CourierID = cidInt

	updatedCourierLocation, err := c.courierStorage.UpdateCourierLocation(courierLocation)
	if err != nil && err == sql.ErrNoRows {
		c.logger.Println(err)
		return nil, fmt.Errorf("couldn't update courierLocation")
	}

	if err != nil && err != sql.ErrNoRows {
		c.logger.Println(err)
		return nil, systemErr
	}

	return updatedCourierLocation, nil
}

// GetCourierLocation prepare data to get it from customerStorage.
func (c *courierService) GetCourierLocation(id string) (*domain.CourierLocation, error) {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.logger.Println(err)
		return nil, errWrongCourierIDType
	}

	courierLocation, err := c.courierStorage.GetCourierLocation(idInt)
	if err != nil && err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil && err != sql.ErrNoRows {
		c.logger.Println(err)
		return nil, systemErr
	}

	return courierLocation, nil

}
