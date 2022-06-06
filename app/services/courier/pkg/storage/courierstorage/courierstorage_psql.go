package courierstorage

import (
	"database/sql"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"strconv"

	"courier/pkg/domain"
)

// Params is the input parameter struct for the module that contains its dependencies
type Params struct {
	DB *sql.DB
}

type CourierStorage struct {
	db *sql.DB
}

func NewCourierStorage(p Params) *CourierStorage {
	return &CourierStorage{
		db: p.DB,
	}
}

// InsertCourier inserts a new courier into the database.
func (c CourierStorage) InsertCourier(courier domain.Courier) (*domain.Courier, error) {
	hashPass, err := bcrypt.GenerateFromPassword([]byte(courier.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("generating password hash: %w", err)
	}
	sql := `INSERT INTO
				courier
					(username, password, firstname, lastname, email, created_at, updated_at, phone, available)
			VALUES($1,$2,$3,$4,$5,now(),now(),$6,true)
			returning *`

	newCourier := domain.Courier{}
	if err = c.db.QueryRow(sql, courier.Username, hashPass, courier.Firstname,
		courier.Lastname, courier.Email, courier.Phone).
		Scan(&newCourier.ID, &newCourier.Username, &newCourier.Password, &newCourier.Firstname,
			&newCourier.Lastname, &newCourier.Email, &newCourier.CreatedAt, &newCourier.UpdatedAt,
			&newCourier.Phone, &newCourier.Available); err != nil {
		return &domain.Courier{}, err
	}

	return &newCourier, nil
}

func (c CourierStorage) DeleteCourier(id int) error {
	sql := `DELETE FROM 
				courier
			WHERE id = $1
	`
	if _, err := c.db.Exec(sql, id); err != nil {
		return err
	}

	return nil
}

func (c CourierStorage) UpdateCourier(courier domain.Courier) (*domain.Courier, error) {
	sql := `UPDATE 
				courier
			SET 
			    username = $1,
			    firstname = $2,
			    lastname = $3,
			  	email = $4,
			  	updated_at = now(),
			  	phone = $5
			    
			WHERE 
			    id = $6
			returning *
	`
	var updatedCourier domain.Courier
	if err := c.db.QueryRow(sql, courier.Username, courier.Firstname, courier.Lastname,
		courier.Email, courier.Phone, courier.ID).
		Scan(&updatedCourier.ID, &updatedCourier.Username, &updatedCourier.Password, &updatedCourier.Firstname,
			&updatedCourier.Lastname, &updatedCourier.Email, &updatedCourier.CreatedAt, &updatedCourier.UpdatedAt,
			&updatedCourier.Phone, &updatedCourier.Available); err != nil {
		return &domain.Courier{}, err
	}

	return &updatedCourier, nil
}

func (c CourierStorage) UpdateCourierAvailable(id int, available bool) (*domain.Courier, error) {
	sql := `UPDATE 
				courier
			SET 
			    available = $2
			WHERE 
			    id = $1
			returning *
	`
	var updatedCourier domain.Courier
	if err := c.db.QueryRow(sql, id, available).
		Scan(&updatedCourier.ID, &updatedCourier.Username, &updatedCourier.Password, &updatedCourier.Firstname,
			&updatedCourier.Lastname, &updatedCourier.Email, &updatedCourier.CreatedAt, &updatedCourier.UpdatedAt,
			&updatedCourier.Phone, &updatedCourier.Available); err != nil {
		return &domain.Courier{}, err
	}

	return &updatedCourier, nil
}

func (c CourierStorage) GetAllCourier(param domain.SearchParam) ([]domain.Courier, error) {
	sql := `SELECT * FROM 
				courier
`
	where := "WHERE 1=1"

	available := param["available"]
	if available != "" {
		where = where + " AND available = " + available + ""
	}

	sql = sql + where

	var allCourier []domain.Courier

	rows, err := c.db.Query(sql)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var courier domain.Courier
		if err := rows.Scan(&courier.ID, &courier.Username, &courier.Password, &courier.Firstname,
			&courier.Lastname, &courier.Email, &courier.CreatedAt, &courier.UpdatedAt,
			&courier.Phone, &courier.Available); err != nil {
			break
		}
		allCourier = append(allCourier, courier)
	}

	return allCourier, nil
}

func (c CourierStorage) GetCourierByID(id int) (*domain.Courier, error) {
	sql := `SELECT * FROM 
				courier
			WHERE
				id = $1
	`
	courier := domain.Courier{}

	if err := c.db.QueryRow(sql, id).Scan(&courier.ID, &courier.Username, &courier.Password, &courier.Firstname,
		&courier.Lastname, &courier.Email, &courier.CreatedAt, &courier.UpdatedAt,
		&courier.Phone, &courier.Available); err != nil {
		return nil, err
	}

	return &courier, nil
}

func (c CourierStorage) GetCourierDuplicateByParam(param domain.SearchParam) (*domain.Courier, error) {
	sql := `SELECT * FROM 
				courier
	`
	where := "WHERE 1=1"

	id := param["id"]
	if id != "" {
		where = where + " AND id != " + id + ""
	}

	username := param["username"]
	if username != "" {
		where = where + " AND username = '" + username + "'"
	}

	email := param["email"]
	if email != "" {
		where = where + " AND email = '" + email + "' "
	}

	phone := param["phone"]
	if email != "" {
		where = where + " AND phone = '" + phone + "' "
	}

	sql = sql + where

	courier := domain.Courier{}

	if err := c.db.QueryRow(sql).Scan(&courier.ID, &courier.Username, &courier.Password, &courier.Firstname,
		&courier.Lastname, &courier.Email, &courier.CreatedAt, &courier.UpdatedAt,
		&courier.Phone, &courier.Available); err != nil {
		return nil, err
	}

	return &courier, nil
}

func (c CourierStorage) CleanCourierTable() error {
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

// InsertCourierLocation inserts a new courier into the database.
func (c CourierStorage) InsertCourierLocation(courierLocation domain.CourierLocation) (*domain.CourierLocation, error) {

	sql := `
   INSERT INTO
    courier_location (courier_id, altitude, longitude, country, city, region, street, home_number, floor, door)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
returning *
`
	var newCourierLocation domain.CourierLocation
	err := c.db.QueryRow(sql, &courierLocation.CourierID, &courierLocation.Altitude,
		&courierLocation.Longitude, &courierLocation.Country, &courierLocation.City,
		&courierLocation.Region, &courierLocation.Street, &courierLocation.HomeNumber,
		&courierLocation.Floor, &courierLocation.Door).
		Scan(&newCourierLocation.CourierID, &newCourierLocation.Altitude,
			&newCourierLocation.Longitude, &newCourierLocation.Country, &newCourierLocation.City,
			&newCourierLocation.Region, &newCourierLocation.Street, &newCourierLocation.HomeNumber,
			&newCourierLocation.Floor, &newCourierLocation.Door)
	if err != nil {
		return nil, err
	}

	return &newCourierLocation, nil
}

func (c CourierStorage) DeleteCourierLocation(courierID int) error {
	sql := `
    DELETE
FROM courier_location
WHERE courier_id = $1
;`
	if _, err := c.db.Exec(sql, courierID); err != nil {
		return err
	}
	return nil
}

func (c CourierStorage) GetCourierLocation(id int) (*domain.CourierLocation, error) {
	sql := `SELECT
				courier_id, altitude, longitude, country, city, region, street, home_number, floor, door
			FROM 
				courier_location 
	`
	where := `WHERE 1=1`

	if id != 0 {
		where = where + "AND courier_id =" + strconv.FormatInt(int64(id), 10)
	}

	sql = sql + where

	var courierLocation domain.CourierLocation

	if err := c.db.QueryRow(sql).Scan(&courierLocation.CourierID, &courierLocation.Altitude,
		&courierLocation.Longitude, &courierLocation.Country, &courierLocation.City,
		&courierLocation.Region, &courierLocation.Street, &courierLocation.HomeNumber,
		&courierLocation.Floor, &courierLocation.Door); err != nil {
		return nil, err
	}

	return &courierLocation, nil
}

func (c CourierStorage) UpdateCourierLocation(courierLocation domain.CourierLocation) (*domain.CourierLocation, error) {

	sql := `UPDATE 
				courier_location
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
			    courier_id = $10
			returning *
	`
	var updatedCurierLocation domain.CourierLocation

	if err := c.db.QueryRow(sql,
		courierLocation.Altitude, courierLocation.Longitude,
		courierLocation.Country, courierLocation.City, courierLocation.Region, courierLocation.Street,
		courierLocation.HomeNumber, courierLocation.Floor, courierLocation.Door, courierLocation.CourierID).
		Scan(&updatedCurierLocation.CourierID, &updatedCurierLocation.Altitude, &updatedCurierLocation.Longitude,
			&updatedCurierLocation.Country, &updatedCurierLocation.City, &updatedCurierLocation.Region,
			&updatedCurierLocation.Street, &updatedCurierLocation.HomeNumber,
			&updatedCurierLocation.Floor, &updatedCurierLocation.Door); err != nil {
		return nil, err
	}

	return &updatedCurierLocation, nil
}
