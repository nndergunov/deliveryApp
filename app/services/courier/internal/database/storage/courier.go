package storage

import (
	"fmt"

	"courier/internal/models"

	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

// CourierStorage is the interface for the courier storage.
type CourierStorage interface {
	InsertCourier(courier models.NewCourierRequest) (*models.CourierResponse, error)
	RemoveCourier(id uint64) error
	UpdateCourier(courier models.UpdateCourierRequest, id uint64) (*models.CourierResponse, error)
	UpdateCourierAvailabe(id uint64, available bool) (*models.CourierResponse, error)
	GetAllCourier() ([]*models.CourierResponse, error)
	GetCourier(id uint64, username, status string) (*models.CourierResponse, error)
}

// Params is the input parameter struct for the module that contains its dependencies
type Params struct {
	DB *sqlx.DB
}

type courierStorage struct {
	db *sqlx.DB
}

// NewCourierStorage constructs a new NewCourierStorage.
func NewCourierStorage(p Params) (CourierStorage, error) {
	courierStorageItem := &courierStorage{
		db: p.DB,
	}

	return courierStorageItem, nil
}

// InsertCourier inserts a new courier into the database.
func (c courierStorage) InsertCourier(courier models.NewCourierRequest) (*models.CourierResponse, error) {
	hashPass, err := bcrypt.GenerateFromPassword([]byte(courier.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("generating password hash: %w", err)
	}
	sql := `INSERT INTO
				courier
					(username, password, firstname, lastname, email, createdat, updatedat, phone, status, available)
			VALUES($1,$2,$3,$4,$5,now(),now(),$6,'active',true)
			returning *`

	var newCourier models.CourierResponse
	if err = c.db.QueryRow(sql, courier.Username, hashPass, courier.Firstname,
		courier.Lastname, courier.Email, courier.Phone).
		Scan(newCourier.ID, newCourier.Username, newCourier.Password, newCourier.Firstname, newCourier.Lastname,
			newCourier.Email, newCourier.Createdat, newCourier.Updatedat, newCourier.Status, newCourier.Available); err != nil {
		return &models.CourierResponse{}, err
	}
	newCourier.Password = ""

	return &newCourier, nil
}

func (c courierStorage) RemoveCourier(id uint64) error {
	sql := `UPDATE 
				courier
			SET 
				status =  'nonactive',	
				available = false,
				updatedat = now()
			WHERE id = $1
	`
	if _, err := c.db.Exec(sql, id); err != nil {
		return err
	}

	return nil
}

func (c courierStorage) UpdateCourier(courier models.UpdateCourierRequest, id uint64) (*models.CourierResponse, error) {
	sql := `UPDATE 
				courier
			SET 
			    username = $1,
			    firstname = $2,
			    lastname = $3,
			  	email = $4,
			  	updatedat = now(),
			  	phone = $5,
			    
			WHERE 
			    status = 'active'
			    AND 
			    id = $6
			returning *
	`
	var updatedCourier models.CourierResponse
	if err := c.db.QueryRow(sql, courier.Username, courier.Firstname, courier.Lastname,
		courier.Email, courier.Phone, id).
		Scan(updatedCourier.ID, updatedCourier.Username, updatedCourier.Password, updatedCourier.Firstname,
			updatedCourier.Lastname, updatedCourier.Email, updatedCourier.Createdat, updatedCourier.Updatedat,
			updatedCourier.Status, updatedCourier.Available); err != nil {
		return &models.CourierResponse{}, err
	}

	updatedCourier.Password = ""

	return &updatedCourier, nil
}

func (c courierStorage) UpdateCourierAvailabe(id uint64, available bool) (*models.CourierResponse, error) {
	sql := `UPDATE 
				courier
			SET 
			    available = $2
			WHERE 
			    status = 'active'
			    AND 
			    id = $1
			returning *
	`
	var updatedCourier models.CourierResponse
	if err := c.db.QueryRow(sql, id, available).
		Scan(updatedCourier.ID, updatedCourier.Username, updatedCourier.Password, updatedCourier.Firstname,
			updatedCourier.Lastname, updatedCourier.Email, updatedCourier.Createdat, updatedCourier.Updatedat,
			updatedCourier.Status, updatedCourier.Available); err != nil {
		return &models.CourierResponse{}, err
	}

	updatedCourier.Password = ""

	return &updatedCourier, nil
}

func (c courierStorage) GetAllCourier() ([]*models.CourierResponse, error) {
	sql := `SELECT * FROM 
				courier
			WHERE status = 'active'
	`

	var allCourier []*models.CourierResponse

	rows, err := c.db.Query(sql)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var courier models.CourierResponse
		if err := rows.Scan(courier.ID, courier.Username, courier.Password, courier.Firstname,
			courier.Lastname, courier.Email, courier.Createdat, courier.Updatedat,
			courier.Status, courier.Available); err != nil {
			break
		}
		courier.Password = ""
		allCourier = append(allCourier, &courier)
	}

	return allCourier, nil
}

func (c courierStorage) GetCourier(id uint64, username, status string) (*models.CourierResponse, error) {
	sql := `SELECT * FROM 
				courier
	`
	where := `WHERE (id = $1 OR username = $2)`

	if status != "" {
		where = where + "AND status = '" + status + "'"
	}
	sql = sql + where

	courier := models.CourierResponse{}

	if err := c.db.QueryRow(sql, id, username).Scan(courier.ID, courier.Username, courier.Password, courier.Firstname,
		courier.Lastname, courier.Email, courier.Createdat, courier.Updatedat,
		courier.Status, courier.Available); err != nil {
		return nil, err
	}

	courier.Password = ""

	return &courier, nil
}
