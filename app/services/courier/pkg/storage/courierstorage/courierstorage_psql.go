package courierstorage

import (
	"courier/pkg/domain"
	"courier/pkg/service"
	"fmt"
	"strconv"

	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

// Params is the input parameter struct for the module that contains its dependencies
type Params struct {
	DB *sqlx.DB
}

type courierStorage struct {
	db *sqlx.DB
}

// NewCourierStorage constructs a new NewCourierStorage.
func NewCourierStorage(p Params) (service.CourierStorage, error) {
	courierStorageItem := &courierStorage{
		db: p.DB,
	}

	return courierStorageItem, nil
}

// InsertCourier inserts a new courier into the database.
func (c courierStorage) InsertCourier(courier domain.Courier) (*domain.Courier, error) {
	hashPass, err := bcrypt.GenerateFromPassword([]byte(courier.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("generating password hash: %w", err)
	}
	sql := `INSERT INTO
				courier
					(username, password, firstname, lastname, email, createdat, updatedat, phone, status, available)
			VALUES($1,$2,$3,$4,$5,now(),now(),$6,'active',true)
			returning *`

	newCourier := domain.Courier{}
	if err = c.db.QueryRow(sql, courier.Username, hashPass, courier.Firstname,
		courier.Lastname, courier.Email, courier.Phone).
		Scan(&newCourier.ID, &newCourier.Username, &newCourier.Password, &newCourier.Firstname,
			&newCourier.Lastname, &newCourier.Email, &newCourier.Createdat, &newCourier.Updatedat,
			&newCourier.Phone, &newCourier.Status, &newCourier.Available); err != nil {
		return &domain.Courier{}, err
	}

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

func (c courierStorage) UpdateCourier(courier domain.Courier) (*domain.Courier, error) {
	sql := `UPDATE 
				courier
			SET 
			    username = $1,
			    firstname = $2,
			    lastname = $3,
			  	email = $4,
			  	updatedat = now(),
			  	phone = $5
			    
			WHERE 
			    status = 'active'
			    AND 
			    id = $6
			returning *
	`
	var updatedCourier domain.Courier
	if err := c.db.QueryRow(sql, courier.Username, courier.Firstname, courier.Lastname,
		courier.Email, courier.Phone, courier.ID).
		Scan(&updatedCourier.ID, &updatedCourier.Username, &updatedCourier.Password, &updatedCourier.Firstname,
			&updatedCourier.Lastname, &updatedCourier.Email, &updatedCourier.Createdat, &updatedCourier.Updatedat,
			&updatedCourier.Phone, &updatedCourier.Status, &updatedCourier.Available); err != nil {
		return &domain.Courier{}, err
	}

	return &updatedCourier, nil
}

func (c courierStorage) UpdateCourierAvailable(id uint64, available bool) (*domain.Courier, error) {
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
	var updatedCourier domain.Courier
	if err := c.db.QueryRow(sql, id, available).
		Scan(&updatedCourier.ID, &updatedCourier.Username, &updatedCourier.Password, &updatedCourier.Firstname,
			&updatedCourier.Lastname, &updatedCourier.Email, &updatedCourier.Createdat, &updatedCourier.Updatedat,
			&updatedCourier.Phone, &updatedCourier.Status, &updatedCourier.Available); err != nil {
		return &domain.Courier{}, err
	}

	return &updatedCourier, nil
}

func (c courierStorage) GetAllCourier() ([]domain.Courier, error) {
	sql := `SELECT * FROM 
				courier
			WHERE status = 'active'
	`

	var allCourier []domain.Courier

	rows, err := c.db.Query(sql)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var courier domain.Courier
		if err := rows.Scan(&courier.ID, &courier.Username, &courier.Password, &courier.Firstname,
			&courier.Lastname, &courier.Email, &courier.Createdat, &courier.Updatedat,
			&courier.Phone, &courier.Status, &courier.Available); err != nil {
			break
		}
		allCourier = append(allCourier, courier)
	}

	return allCourier, nil
}

func (c courierStorage) GetCourier(id uint64, username, status string) (*domain.Courier, error) {
	sql := `SELECT * FROM 
				courier
	`
	where := `WHERE 1=1`

	if id != 0 {
		where = where + "AND id = " + strconv.Itoa(int(id))
	}

	if username != "" {
		where = where + "AND username = '" + username + "'"
	}

	if status != "" {
		where = where + "AND status = '" + status + "'"
	}
	sql = sql + where

	courier := domain.Courier{}

	if err := c.db.QueryRow(sql).Scan(&courier.ID, &courier.Username, &courier.Password, &courier.Firstname,
		&courier.Lastname, &courier.Email, &courier.Createdat, &courier.Updatedat,
		&courier.Phone, &courier.Status, &courier.Available); err != nil {
		return nil, err
	}

	return &courier, nil
}

func (c courierStorage) CleanDB() error {
	sql := `DELETE FROM
				courier
			WHERE 
				 1=1
	`
	if _, err := c.db.Exec(sql); err != nil {
		return err
	}

	return nil
}
