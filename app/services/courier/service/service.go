package service

import (
	"courier/models"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/nndergunov/deliveryApp/app/pkg/logger"
	"golang.org/x/crypto/bcrypt"
)

// CourierService is the interface for the user service.
type CourierService interface {
	Insert(courier models.Courier) (*models.Courier, error)
	Remove(id uint64) error
	Update(courier models.Courier) (*models.Courier, error)
	GetAll() ([]*models.Courier, error)
	Get(id uint64, username string) (*models.Courier, error)
}

// Params is the input parameter struct for the module that contains its dependencies
type Params struct {
	DB     *sqlx.DB
	Logger *logger.Logger
}

type courierService struct {
	db     *sqlx.DB
	logger *logger.Logger
}

// NewCourierService constructs a new NewCourierService.
func NewCourierService(p Params) (CourierService, error) {

	courierServiceItem := &courierService{
		db:     p.DB,
		logger: p.Logger,
	}

	return courierServiceItem, nil

}

// Insert inserts a new courier into the database.
func (c courierService) Insert(courier models.Courier) (*models.Courier, error) {

	hashPass, err := bcrypt.GenerateFromPassword([]byte(courier.Password), bcrypt.DefaultCost)
	if err != nil {
		return &models.Courier{}, fmt.Errorf("generating password hash: %w", err)
	}
	sql := `INSERT INTO
				courier
					(username, password, firstname, lastname, email, createdat, updatedat, phone, status, available)
			VALUES($1,$2,$3,$4,$5,now(),now(),$6,'active',true)
			returning *`

	var newCourier models.Courier
	err = c.db.QueryRow(sql, courier.Username, hashPass, courier.Firstname,
		courier.Lastname, courier.Email, courier.Phone).Scan(newCourier.Fields()...)
	if err != nil {
		return &models.Courier{}, err
	}
	newCourier.Password = ""

	return &newCourier, nil
}

func (c courierService) Remove(id uint64) error {
	sql := `UPDATE 
				courier
			SET 
				status =  'nonactive',	
				available = false,
				updatedat = now()
			WHERE id = $1
			returning *
	`
	var removedCourier models.Courier
	if err := c.db.QueryRow(sql, id).Scan(removedCourier.Fields()...); err != nil {
		return err
	}

	removedCourier.Password = ""
	c.logger.Println("Removed courier:", removedCourier)

	return nil
}
func (c courierService) Update(courier models.Courier) (*models.Courier, error) {
	sql := `UPDATE 
				courier
			SET 
			    username = $1,
			    firstname = $2,
			    lastname = $3,
			  	email = $4,
			  	updatedat = now(),
			  	phone = $5,
				available = $6
			    
			WHERE 
			    status = 'active'
			    AND 
			    id = $7
			returning *
	`
	var updatedCourier models.Courier
	if err := c.db.QueryRow(sql, courier.Username, courier.Firstname, courier.Lastname,
		courier.Email, courier.Phone, courier.Available,
		courier.ID).Scan(updatedCourier.Fields()...); err != nil {

		return &models.Courier{}, err
	}

	updatedCourier.Password = ""

	return &updatedCourier, nil
}
func (c courierService) GetAll() ([]*models.Courier, error) {
	sql := `SELECT * FROM 
				courier
			WHERE status = 'active'
	`

	var allCourier []*models.Courier

	rows, err := c.db.Query(sql)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var courier models.Courier
		if err := rows.Scan(courier.Fields()...); err != nil {
			c.logger.Println(err)
			break
		}
		courier.Password = ""
		allCourier = append(allCourier, &courier)
	}

	return allCourier, nil
}

func (c courierService) Get(id uint64, username string) (*models.Courier, error) {
	sql := `SELECT * FROM 
				courier
			WHERE 
			      status = 'active'
			AND 
			      (id = $1
			OR 
			      username = $2)
	`

	var courier models.Courier

	if err := c.db.QueryRow(sql, id, username).Scan(courier.Fields()...); err != nil {
		return nil, err
	}

	courier.Password = ""

	return &courier, nil
}
