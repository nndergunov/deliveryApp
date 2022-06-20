package courierstorage

import (
	"database/sql"
	"fmt"
	"strconv"

	"golang.org/x/crypto/bcrypt"

	"github.com/nndergunov/deliveryApp/app/services/courier/pkg/domain"
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

func (c CourierStorage) GetCourierList(param domain.SearchParam) ([]domain.Courier, error) {
	sql := `SELECT * FROM 
				courier
`
	where := " WHERE 1=1"

	available, ok := param["available"]
	if ok && available != "" {
		where = where + " AND available = " + available + ""
	}

	sql = sql + where

	var courierList []domain.Courier

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
		courierList = append(courierList, courier)
	}

	return courierList, nil
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

// InsertLocation inserts a new courier into the database.
func (c CourierStorage) InsertLocation(location domain.Location) (*domain.Location, error) {
	sql := `INSERT 
			INTO
			    location (user_id, latitude, longitude, country, city, region, street, home_number, floor, door)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
			returning *
`
	var newLocation domain.Location
	err := c.db.QueryRow(sql, &location.UserID, &location.Latitude,
		&location.Longitude, &location.Country, &location.City,
		&location.Region, &location.Street, &location.HomeNumber,
		&location.Floor, &location.Door).
		Scan(&newLocation.UserID, &newLocation.Latitude,
			&newLocation.Longitude, &newLocation.Country, &newLocation.City,
			&newLocation.Region, &newLocation.Street, &newLocation.HomeNumber,
			&newLocation.Floor, &newLocation.Door)
	if err != nil {
		return nil, err
	}

	return &newLocation, nil
}

func (c CourierStorage) DeleteLocation(courierID int) error {
	sql := `DELETE
			FROM location
			WHERE user_id = $1;`

	if _, err := c.db.Exec(sql, courierID); err != nil {
		return err
	}
	return nil
}

func (c CourierStorage) GetLocation(userID int) (*domain.Location, error) {
	sql := `SELECT
				user_id, latitude, longitude, country, city, region, street, home_number, floor, door
			FROM 
				location 
	`
	where := ` WHERE 1=1`

	if userID != 0 {
		where = where + "AND user_id =" + strconv.FormatInt(int64(userID), 10)
	}

	sql = sql + where

	var location domain.Location

	if err := c.db.QueryRow(sql).Scan(&location.UserID, &location.Latitude,
		&location.Longitude, &location.Country, &location.City,
		&location.Region, &location.Street, &location.HomeNumber,
		&location.Floor, &location.Door); err != nil {
		return nil, err
	}

	return &location, nil
}

func (c CourierStorage) UpdateLocation(location domain.Location) (*domain.Location, error) {
	sql := `UPDATE 
				location
			SET 
			    latitude = $1,
			    longitude = $2,
			    country = $3,
			    city = $4,
			  	region = $5,
			  	street = $6,
			    home_number =$7,
			    floor = $8,
			    door = $9 
			WHERE 
			    user_id = $10
			returning *
	`
	var updatedLocation domain.Location

	if err := c.db.QueryRow(sql,
		location.Latitude, location.Longitude,
		location.Country, location.City, location.Region, location.Street,
		location.HomeNumber, location.Floor, location.Door, location.UserID).
		Scan(&updatedLocation.UserID, &updatedLocation.Latitude, &updatedLocation.Longitude,
			&updatedLocation.Country, &updatedLocation.City, &updatedLocation.Region,
			&updatedLocation.Street, &updatedLocation.HomeNumber,
			&updatedLocation.Floor, &updatedLocation.Door); err != nil {
		return nil, err
	}

	return &updatedLocation, nil
}

func (c CourierStorage) GetLocationList(param domain.SearchParam) ([]domain.Location, error) {
	sql := `SELECT * FROM 
				location
`
	where := " WHERE 1=1"

	city, ok := param["city"]
	if ok && city != "" {
		where = where + " AND city = '" + city + "'"
	}

	sql = sql + where

	var locationList []domain.Location

	rows, err := c.db.Query(sql)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var location domain.Location
		if err := rows.Scan(&location.UserID, &location.Latitude,
			&location.Longitude, &location.Country, &location.City,
			&location.Region, &location.Street, &location.HomeNumber,
			&location.Floor, &location.Door); err != nil {
			break
		}
		locationList = append(locationList, location)
	}

	return locationList, nil
}

func (c CourierStorage) CleanLocationTable() error {
	sql := `DELETE FROM
				location
			WHERE 
				 1=1
	`
	if _, err := c.db.Exec(sql); err != nil {
		return err
	}

	return nil
}
