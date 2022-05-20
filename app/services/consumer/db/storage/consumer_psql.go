package storage

import (
	"consumer/domain"
	"consumer/service"
	"strconv"

	"github.com/jmoiron/sqlx"
)

// Params is the input parameter struct for the module that contains its dependencies
type Params struct {
	DB *sqlx.DB
}

type consumerStorage struct {
	db *sqlx.DB
}

// NewConsumerStorage constructs a new NewConsumerStorage.
func NewConsumerStorage(p Params) (service.ConsumerStorage, error) {
	consumerStorageItem := &consumerStorage{
		db: p.DB,
	}

	return consumerStorageItem, nil
}

// InsertConsumer inserts a new consumer into the database.
func (c consumerStorage) InsertConsumer(consumer domain.Consumer) (*domain.Consumer, error) {

	//find better solution
	//create default location for consumer
	sql1 := `INSERT INTO 
    			consumer_location 
    		DEFAULT VALUE 
    			returning id;`

	var consumerLocationID uint64
	err := c.db.QueryRow(sql1).Scan(&consumerLocationID)
	if err != nil {
		return nil, err
	}

	sql2 := `INSERT INTO
				consumer
					(firstname, lastname, email, phone, created_at, updated_at, location_id)
			VALUES($1,$2,$3,$4,now(),now(), $5)
			returning *`

	var newConsumer domain.Consumer
	err = c.db.QueryRow(sql2,
		consumer.Firstname, consumer.Lastname, consumer.Email, consumer.Phone, consumerLocationID).
		Scan(&newConsumer.ID, &newConsumer.Firstname, &newConsumer.Lastname, &newConsumer.Email, &newConsumer.Phone,
			&newConsumer.Createdat, &newConsumer.Updatedat, &newConsumer.ConsumerLocation.ID)
	if err != nil {
		return &domain.Consumer{}, err
	}

	return &newConsumer, nil
}

func (c consumerStorage) DeleteConsumer(id uint64) error {
	sql := `DELETE FROM 
				consumer
			WHERE id = $1
			CASCADE 
	`
	if _, err := c.db.Exec(sql, id); err != nil {
		return err
	}
	return nil
}

func (c consumerStorage) UpdateConsumer(consumer domain.Consumer) (*domain.Consumer, error) {
	//make it in one query
	sql := `UPDATE 
				consumer
			SET 
			    firstname = $1,
			    lastname = $2,
			  	email = $3,
			    phone = $4,
			  	updated_at = now(),
			WHERE 
			    id = $5
			returning *
	`
	var updatedConsumer domain.Consumer
	if err := c.db.QueryRow(sql,
		consumer.Firstname, consumer.Lastname, consumer.Email, consumer.Phone, consumer.ID).
		Scan(&updatedConsumer.ID, &updatedConsumer.Firstname, &updatedConsumer.Lastname, &updatedConsumer.Email,
			&updatedConsumer.Phone, &updatedConsumer.Createdat, &updatedConsumer.Updatedat,
			&updatedConsumer.ConsumerLocation.ID); err != nil {
		return nil, err
	}

	sql2 := `SELECT * FROM
				consumer_location
			WHERE 
			    id = $1
	`
	var consumerLocation domain.ConsumerLocation
	if err := c.db.QueryRow(sql2, updatedConsumer.ConsumerLocation.ID).
		Scan(&consumerLocation.ID, &consumerLocation.LocationAlt, &consumerLocation.LocationLat,
			&consumerLocation.Country, &consumerLocation.City, &consumerLocation.Region,
			&consumerLocation.Street, &consumerLocation.HomeNumber, &consumerLocation.Door,
			&consumerLocation.Floor); err != nil {
		return nil, err
	}

	updatedConsumer.ConsumerLocation = consumerLocation

	return &updatedConsumer, nil
}

func (c consumerStorage) GetAllConsumer() ([]domain.Consumer, error) {
	sql := `SELECT
				c.id, c.first_name, c.last_name, c.email, c.phone, c.created_at, c.updated_at,
       			cl.ID, cl.location_alt, cl.location_lat, cl.country, cl.city, cl.region, cl.street, 
       			cl.home_number, cl.floor, cl.door	
			FROM 
				consumer c
			INNER JOIN 
				consumer_location cl
			ON 
			    c.location_id = cl.id
	`

	var allConsumer []domain.Consumer

	rows, err := c.db.Query(sql)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var consumer domain.Consumer
		if err := rows.Scan(&consumer.ID, &consumer.Firstname, &consumer.Lastname, &consumer.Email,
			&consumer.Phone, &consumer.Createdat, &consumer.Updatedat,
			&consumer.ConsumerLocation.ID, &consumer.ConsumerLocation.LocationAlt, &consumer.ConsumerLocation.LocationLat,
			&consumer.ConsumerLocation.Country, &consumer.ConsumerLocation.City, &consumer.ConsumerLocation.Region,
			&consumer.ConsumerLocation.Street, &consumer.ConsumerLocation.HomeNumber, &consumer.ConsumerLocation.Door,
			&consumer.ConsumerLocation.Floor); err != nil {
			break
		}
		allConsumer = append(allConsumer, consumer)
	}

	return allConsumer, nil
}

func (c consumerStorage) GetConsumer(id uint64, phone, email string) (*domain.Consumer, error) {
	sql := `SELECT
				c.id, c.first_name, c.last_name, c.email, c.phone, c.created_at, c.updated_at,
       			cl.ID, cl.location_alt, cl.location_lat, cl.country, cl.city, cl.region, cl.street, 
       			cl.home_number, cl.floor, cl.door	
			FROM 
				consumer c
	`
	where := `WHERE id = 1=1`

	if id != 0 {
		where = where + "AND id =" + strconv.FormatInt(int64(id), 10)
	}

	if phone != "" {
		where = where + "AND phone = '" + phone + "'"
	}

	if email != "" {
		where = where + "AND email = '" + email + "'"
	}

	sql = sql + where
	sql = sql + `
				INNER JOIN 
					consumer_location cl
				ON 
			    	c.location_id = cl.id`

	var consumer domain.Consumer

	if err := c.db.QueryRow(sql).Scan(&consumer.ID, &consumer.Firstname, &consumer.Lastname, &consumer.Email,
		&consumer.Phone, &consumer.Createdat, &consumer.Updatedat,
		&consumer.ConsumerLocation.ID, &consumer.ConsumerLocation.LocationAlt, &consumer.ConsumerLocation.LocationLat,
		&consumer.ConsumerLocation.Country, &consumer.ConsumerLocation.City, &consumer.ConsumerLocation.Region,
		&consumer.ConsumerLocation.Street, &consumer.ConsumerLocation.HomeNumber, &consumer.ConsumerLocation.Door,
		&consumer.ConsumerLocation.Floor); err != nil {
		return nil, err
	}

	return &consumer, nil
}

func (c consumerStorage) CleanConsumerTable() error {
	sql := `DELETE FROM
				consumer
			WHERE 
				 1=1
	`
	if _, err := c.db.Exec(sql); err != nil {
		return err
	}

	return nil
}

func (c consumerStorage) UpdateConsumerLocation(consumerLocation domain.ConsumerLocation) (*domain.ConsumerLocation, error) {

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
			    id = $1
			returning *
	`
	var updatedConsumerLocation domain.ConsumerLocation

	if err := c.db.QueryRow(sql,
		consumerLocation.ID, consumerLocation.LocationAlt, consumerLocation.LocationLat,
		consumerLocation.Country, consumerLocation.City, consumerLocation.Region, consumerLocation.Street,
		consumerLocation.HomeNumber, consumerLocation.Floor, consumerLocation.Door).
		Scan(&updatedConsumerLocation.ID, &updatedConsumerLocation.LocationAlt, &updatedConsumerLocation.LocationLat,
			&updatedConsumerLocation.Country, &updatedConsumerLocation.City, &updatedConsumerLocation.Region,
			&updatedConsumerLocation.Street, &updatedConsumerLocation.HomeNumber,
			&updatedConsumerLocation.Floor, &updatedConsumerLocation.Door); err != nil {
		return nil, err
	}

	return &updatedConsumerLocation, nil
}

func (c consumerStorage) CleanConsumerLocationTable() error {
	sql := `DELETE FROM
				consumer_location
			WHERE 
				 1=1
	`
	if _, err := c.db.Exec(sql); err != nil {
		return err
	}

	return nil
}
