package storage

import (
	"consumer/internal/models"
	"strconv"

	"github.com/jmoiron/sqlx"
)

// ConsumerStorage is the interface for the consumer storage.
type ConsumerStorage interface {
	InsertConsumer(consumer models.Consumer) (*models.Consumer, error)
	RemoveConsumer(id uint64) error
	UpdateConsumer(consumer models.Consumer) (*models.Consumer, error)
	GetAllConsumer() ([]*models.Consumer, error)
	GetConsumer(id uint64, phoneNumber, email, status string) (*models.Consumer, error)

	UpdateConsumerLocation(consumer models.ConsumerLocation) (*models.ConsumerLocation, error)
}

// Params is the input parameter struct for the module that contains its dependencies
type Params struct {
	DB *sqlx.DB
}

type consumerStorage struct {
	db *sqlx.DB
}

// NewConsumerStorage constructs a new NewConsumerStorage.
func NewConsumerStorage(p Params) (ConsumerStorage, error) {
	consumerStorageItem := &consumerStorage{
		db: p.DB,
	}

	return consumerStorageItem, nil
}

// InsertConsumer inserts a new consumer into the database create default location foe this user.
func (c consumerStorage) InsertConsumer(consumer models.Consumer) (*models.Consumer, error) {
	//todo: do it in transaction

	sql1 := `INSERT INTO
				consumer
					(country_code, phone_number, firstname, lastname, email, created_at, updated_at, status)
			VALUES($1,$2,$3,$4,$5,,now(),now(),$6,'active')
			returning *`

	var newConsumer models.Consumer
	err := c.db.QueryRow(sql1, consumer.CountryCode, consumer.PhoneNumber,
		consumer.Firstname, consumer.Lastname, consumer.Email, consumer.PhoneNumber).Scan(newConsumer.Fields()...)
	if err != nil {
		return &models.Consumer{}, err
	}

	//find better solution
	sql2 := `INSERT INTO consumer_location DEFAULT VALUES returning id;`

	var consumerLocationID uint64
	err = c.db.QueryRow(sql2).Scan(&consumerLocationID)
	if err != nil {
		return &models.Consumer{}, err
	}
	if consumerLocationID < 1 {
		return &models.Consumer{}, err
	}
	return &newConsumer, nil
}

func (c consumerStorage) RemoveConsumer(id uint64) error {
	sql := `UPDATE 
				consumer
			SET 
				status =  'nonactive',	
				updatedat = now()
			WHERE id = $1
			returning *
	`
	var removedConsumer models.Consumer
	if err := c.db.QueryRow(sql, id).Scan(removedConsumer.Fields()...); err != nil {
		return err
	}

	return nil
}

func (c consumerStorage) UpdateConsumer(consumer models.Consumer) (*models.Consumer, error) {
	sql := `UPDATE 
				consumer
			SET 
			    country_code = $1
			    phone_number = $2,
			    first_name = $3,
			    last_name = $4,
			  	email = $5,
			  	updated_at = now(),
			WHERE 
			    status = 'active'
			    AND 
			    id = $6
			returning *
	`
	var updatedConsumer models.Consumer
	if err := c.db.QueryRow(sql,
		consumer.CountryCode, consumer.PhoneNumber, consumer.Firstname, consumer.Lastname, consumer.Email,
		consumer.ID).Scan(updatedConsumer.Fields()...); err != nil {
		return &models.Consumer{}, err
	}

	return &updatedConsumer, nil
}

func (c consumerStorage) GetAllConsumer() ([]*models.Consumer, error) {
	sql := `SELECT * FROM 
				consumer
			WHERE status = 'active'
	`

	var allConsumer []*models.Consumer

	rows, err := c.db.Query(sql)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var consumer models.Consumer
		if err := rows.Scan(consumer.Fields()...); err != nil {
			break
		}
		allConsumer = append(allConsumer, &consumer)
	}

	return allConsumer, nil
}

func (c consumerStorage) GetConsumer(id uint64, phoneNumber, status, email string) (*models.Consumer, error) {
	sql := `SELECT * FROM 
				consumer
	`
	where := `WHERE id = 1=1`

	if id != 0 {
		where = where + "AND id =" + strconv.FormatInt(int64(id), 10)
	}

	if phoneNumber != "" {
		where = where + "AND phone_number = '" + phoneNumber + "'"
	}

	if email != "" {
		where = where + "AND emai = '" + email + "'"
	}

	if status != "" {
		where = where + "AND status = '" + status + "'"
	}

	sql = sql + where

	var consumer models.Consumer

	if err := c.db.QueryRow(sql).Scan(consumer.Fields()...); err != nil {
		return nil, err
	}

	return &consumer, nil
}

func (c consumerStorage) UpdateConsumerLocation(consumerLocation models.ConsumerLocation) (*models.ConsumerLocation, error) {

	sql := `UPDATE 
				consumer_location
			SET 
			    location_alt = $2
			    location_lat = $3,
			    country = $4,
			    city = $5,
			  	region = $6,
			  	street = $7,
			    home_number =$8,
			    floor = $9,
			    door = $10, 
			WHERE 
			    consumer_id = $1
			returning *
	`
	var updatedConsumerLocation models.ConsumerLocation

	if err := c.db.QueryRow(sql,
		consumerLocation.ConsumerID, consumerLocation.LocationAlt, consumerLocation.LocationLat,
		consumerLocation.Country, consumerLocation.City, consumerLocation.Region, consumerLocation.Street,
		consumerLocation.HomeNumber, consumerLocation.Floor, consumerLocation.Door).
		Scan(updatedConsumerLocation.Fields()...); err != nil {
		return &models.ConsumerLocation{}, err
	}

	return &updatedConsumerLocation, nil
}
