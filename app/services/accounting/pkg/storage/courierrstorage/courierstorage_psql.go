package courierstorage

import (
	"accounting/pkg/domain"
	"database/sql"
)

// Params is the input parameter struct for the module that contains its dependencies
type Params struct {
	DB *sql.DB
}

type Storage struct {
	db *sql.DB
}

func NewStorage(p Params) *Storage {
	return &Storage{
		db: p.DB,
	}
}

func (c Storage) InsertNewCourierAccount(courierAccount domain.CourierAccount) (*domain.CourierAccount, error) {
	sql := `INSERT INTO
				courier_account
					(courier_id, created_at, updated_at)
			VALUES($1,now(),now())
			returning *`

	newCourier := domain.CourierAccount{}
	if err := c.db.QueryRow(sql, courierAccount.CourierID).
		Scan(&newCourier.CourierID, &newCourier.Balance, &newCourier.CreatedAt, &newCourier.UpdatedAt); err != nil {
		return &domain.CourierAccount{}, err
	}

	return &newCourier, nil
}

func (c Storage) GetCourierAccountByID(id uint64) (*domain.CourierAccount, error) {
	sql := `SELECT * FROM 
				courier_account
			WHERE
				courier_id = $1
	`
	courierAccount := domain.CourierAccount{}

	if err := c.db.QueryRow(sql, id).Scan(&courierAccount.CourierID, &courierAccount.Balance, &courierAccount.CreatedAt,
		&courierAccount.UpdatedAt); err != nil {
		return nil, err
	}

	return &courierAccount, nil
}

func (c Storage) DeleteCourierAccount(courierID uint64) error {
	sql := `DELETE FROM 
				courier_account
			WHERE courier_id = $1
	`
	if _, err := c.db.Exec(sql, courierID); err != nil {
		return err
	}

	return nil
}

func (c Storage) AddToBalanceCourierAccount(account domain.CourierAccount) (*domain.CourierAccount, error) {
	sql := `UPDATE courier_account
				    SET 
				        balance = (SELECT balance FROM 
						courier_account) + $2

				WHERE courier_id = $1
			returning *`

	newCourier := domain.CourierAccount{}
	if err := c.db.QueryRow(sql, account.CourierID, account.Balance).
		Scan(&newCourier.CourierID, &newCourier.Balance, &newCourier.CreatedAt, &newCourier.UpdatedAt); err != nil {
		return &domain.CourierAccount{}, err
	}

	return &newCourier, nil
}

func (c Storage) SubFromBalanceCourierAccount(account domain.CourierAccount) (*domain.CourierAccount, error) {
	sql := `UPDATE courier_account
				    SET 
				        balance = (SELECT balance FROM 
						courier_account) - $2

				WHERE courier_id = $1
			returning *`

	newCourier := domain.CourierAccount{}
	if err := c.db.QueryRow(sql, account.CourierID, account.Balance).
		Scan(&newCourier.CourierID, &newCourier.Balance, &newCourier.CreatedAt, &newCourier.UpdatedAt); err != nil {
		return &domain.CourierAccount{}, err
	}

	return &newCourier, nil
}
