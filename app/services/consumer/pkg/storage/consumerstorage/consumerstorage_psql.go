package consumerstorage

import (
	"consumer/pkg/domain"
	"database/sql"
	"strconv"
)

// Params is the input parameter struct for the module that contains its dependencies
type Params struct {
	DB *sql.DB
}

type ConsumerStorage struct {
	db *sql.DB
}

func NewConsumerStorage(p Params) *ConsumerStorage {
	return &ConsumerStorage{
		db: p.DB,
	}
}

// InsertConsumer inserts a new consumer into the database.
func (c ConsumerStorage) InsertConsumer(consumer domain.Consumer) (*domain.Consumer, error) {

	sql := `
    INSERT INTO
        consumer (firstname, lastname, email, phone, created_at, updated_at)
        VALUES ($1, $2, $3, $4, now(), now())
        returning *
`
	var newConsumer domain.Consumer
	err := c.db.QueryRow(sql,
		consumer.Firstname, consumer.Lastname, consumer.Email, consumer.Phone).
		Scan(&newConsumer.ID, &newConsumer.Firstname, &newConsumer.Lastname, &newConsumer.Email,
			&newConsumer.Phone, &newConsumer.CreatedAt, &newConsumer.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &newConsumer, nil
}

func (c ConsumerStorage) DeleteConsumer(id uint64) error {
	sql := `
    DELETE
        FROM
            consumer
            WHERE id = $1
            returning id
;`
	if _, err := c.db.Exec(sql, id); err != nil {
		return err
	}
	return nil
}

func (c ConsumerStorage) UpdateConsumer(consumer domain.Consumer) (*domain.Consumer, error) {

	sql := `
    UPDATE
        consumer
            SET
                firstname = $1,
                lastname = $2,
                email = $3,
                phone = $4,
                updated_at = now()
            WHERE
                id = $5
            returning *`

	var updatedConsumer domain.Consumer
	if err := c.db.QueryRow(sql,
		consumer.Firstname, consumer.Lastname, consumer.Email, consumer.Phone, consumer.ID).
		Scan(&updatedConsumer.ID, &updatedConsumer.Firstname, &updatedConsumer.Lastname, &updatedConsumer.Email,
			&updatedConsumer.Phone, &updatedConsumer.CreatedAt, &updatedConsumer.UpdatedAt); err != nil {
		return nil, err
	}

	return &updatedConsumer, nil
}

func (c ConsumerStorage) GetAllConsumer() ([]domain.Consumer, error) {
	sql := `SELECT
				id, firstname, lastname, email, phone, created_at, updated_at
			FROM 
				consumer
			`

	var allConsumer []domain.Consumer

	rows, err := c.db.Query(sql)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var consumer domain.Consumer
		if err := rows.Scan(&consumer.ID, &consumer.Firstname, &consumer.Lastname, &consumer.Email, &consumer.Phone,
			&consumer.CreatedAt, &consumer.UpdatedAt); err != nil {
			return nil, err
		}
		allConsumer = append(allConsumer, consumer)
	}

	return allConsumer, nil
}

func (c ConsumerStorage) GetConsumerByID(id uint64) (*domain.Consumer, error) {
	sql := `SELECT
				id, firstname, lastname, email, phone, created_at, updated_at
			FROM 
				consumer 
			WHERE id = $1
	`
	var consumer domain.Consumer

	if err := c.db.QueryRow(sql, id).Scan(&consumer.ID, &consumer.Firstname, &consumer.Lastname,
		&consumer.Email, &consumer.Phone, &consumer.CreatedAt, &consumer.UpdatedAt); err != nil {
		return nil, err
	}

	return &consumer, nil
}

func (c ConsumerStorage) GetConsumerDuplicateByParam(param domain.SearchParam) (*domain.Consumer, error) {
	sql := `SELECT * FROM 
				consumer
	`
	where := "WHERE 1=1"

	id := param["id"]
	if id != "" {
		where = where + " AND id != " + id + ""
	}

	email := param["email"]
	if email != "" {
		where = where + " AND email = '" + email + "' "
	}

	phone := param["phone"]
	if phone != "" {
		where = where + " AND phone = '" + phone + "' "
	}

	sql = sql + where

	consumer := domain.Consumer{}

	if err := c.db.QueryRow(sql).Scan(&consumer.ID, &consumer.Firstname,
		&consumer.Lastname, &consumer.Email, &consumer.Phone, &consumer.CreatedAt, &consumer.UpdatedAt); err != nil {
		return nil, err
	}

	return &consumer, nil
}

func (c ConsumerStorage) CleanConsumerTable() error {
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

// InsertConsumerLocation inserts a new consumer into the database.
func (c ConsumerStorage) InsertConsumerLocation(consumerLocation domain.ConsumerLocation) (*domain.ConsumerLocation, error) {

	sql := `
   INSERT INTO
    consumer_location (consumer_id, altitude, longitude, country, city, region, street, home_number, floor, door)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
returning *
`
	var newConsumerLocation domain.ConsumerLocation
	err := c.db.QueryRow(sql, &consumerLocation.ConsumerID, &consumerLocation.Altitude,
		&consumerLocation.Longitude, &consumerLocation.Country, &consumerLocation.City,
		&consumerLocation.Region, &consumerLocation.Street, &consumerLocation.HomeNumber,
		&consumerLocation.Floor, &consumerLocation.Door).
		Scan(&newConsumerLocation.ConsumerID, &newConsumerLocation.Altitude,
			&newConsumerLocation.Longitude, &newConsumerLocation.Country, &newConsumerLocation.City,
			&newConsumerLocation.Region, &newConsumerLocation.Street, &newConsumerLocation.HomeNumber,
			&newConsumerLocation.Floor, &newConsumerLocation.Door)
	if err != nil {
		return nil, err
	}

	return &newConsumerLocation, nil
}

func (c ConsumerStorage) DeleteConsumerLocation(consumerID uint64) error {
	sql := `
    DELETE
FROM consumer_location
WHERE consumer_id = $1
;`
	if _, err := c.db.Exec(sql, consumerID); err != nil {
		return err
	}
	return nil
}
func (c ConsumerStorage) GetConsumerLocation(id uint64) (*domain.ConsumerLocation, error) {
	sql := `SELECT
				consumer_id, altitude, longitude, country, city, region, street, home_number, floor, door
			FROM 
				consumer_location 
	`
	where := `WHERE 1=1`

	if id != 0 {
		where = where + "AND consumer_id =" + strconv.FormatInt(int64(id), 10)
	}

	sql = sql + where

	var consumerLocation domain.ConsumerLocation

	if err := c.db.QueryRow(sql).Scan(&consumerLocation.ConsumerID, &consumerLocation.Altitude,
		&consumerLocation.Longitude, &consumerLocation.Country, &consumerLocation.City,
		&consumerLocation.Region, &consumerLocation.Street, &consumerLocation.HomeNumber,
		&consumerLocation.Floor, &consumerLocation.Door); err != nil {
		return nil, err
	}

	return &consumerLocation, nil
}

func (c ConsumerStorage) UpdateConsumerLocation(consumerLocation domain.ConsumerLocation) (*domain.ConsumerLocation, error) {

	sql := `UPDATE 
				consumer_location
			SET 
			    altitude = $1,
			    longitude = $2,
			    country = $3,
			    city = $4,
			  	region = $5,
			  	street = $6,
			    home_number =$7,
			    floor = $8,
			    door = $9 
			WHERE 
			    consumer_id = $10
			returning *
	`
	var updatedConsumerLocation domain.ConsumerLocation

	if err := c.db.QueryRow(sql,
		consumerLocation.Altitude, consumerLocation.Longitude,
		consumerLocation.Country, consumerLocation.City, consumerLocation.Region, consumerLocation.Street,
		consumerLocation.HomeNumber, consumerLocation.Floor, consumerLocation.Door, consumerLocation.ConsumerID).
		Scan(&updatedConsumerLocation.ConsumerID, &updatedConsumerLocation.Altitude, &updatedConsumerLocation.Longitude,
			&updatedConsumerLocation.Country, &updatedConsumerLocation.City, &updatedConsumerLocation.Region,
			&updatedConsumerLocation.Street, &updatedConsumerLocation.HomeNumber,
			&updatedConsumerLocation.Floor, &updatedConsumerLocation.Door); err != nil {
		return nil, err
	}

	return &updatedConsumerLocation, nil
}
