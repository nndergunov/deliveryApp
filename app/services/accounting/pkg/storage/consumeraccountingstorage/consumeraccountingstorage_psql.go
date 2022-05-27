package consumeraccountingstorage

import (
	"accounting/pkg/domain"
	"database/sql"
)

// Params is the input parameter struct for the module that contains its dependencies
type Params struct {
	DB *sql.DB
}

type ConsumerAccountingStorage struct {
	db *sql.DB
}

func NewConsumerAccountingStorage(p Params) *ConsumerAccountingStorage {
	return &ConsumerAccountingStorage{
		db: p.DB,
	}
}

func (c ConsumerAccountingStorage) InsertNewConsumerAccount(consumerAccount domain.ConsumerAccount) (*domain.ConsumerAccount, error) {
	sql := `INSERT INTO
				consumer_account
					(consumer_id, created_at, updated_at)
			VALUES($1,now(),now())
			returning *`

	newCourier := domain.ConsumerAccount{}
	if err := c.db.QueryRow(sql, consumerAccount.ConsumerID).
		Scan(&newCourier.ConsumerID, &newCourier.Balance, &newCourier.CreatedAt, &newCourier.UpdatedAt); err != nil {
		return &domain.ConsumerAccount{}, err
	}

	return &newCourier, nil
}

func (c ConsumerAccountingStorage) GetConsumerAccountByID(id uint64) (*domain.ConsumerAccount, error) {
	sql := `SELECT * FROM 
				consumer_account
			WHERE
				consumer_id = $1
	`
	consumerAccount := domain.ConsumerAccount{}

	if err := c.db.QueryRow(sql, id).Scan(&consumerAccount.ConsumerID, &consumerAccount.Balance, &consumerAccount.CreatedAt,
		&consumerAccount.UpdatedAt); err != nil {
		return nil, err
	}

	return &consumerAccount, nil
}

func (c ConsumerAccountingStorage) DeleteConsumerAccount(consumerID uint64) error {
	sql := `DELETE FROM 
				consumer_account
			WHERE consumer_id = $1
	`
	if _, err := c.db.Exec(sql, consumerID); err != nil {
		return err
	}

	return nil
}

func (c ConsumerAccountingStorage) AddToConsumerAccount(account domain.ConsumerAccount) (*domain.ConsumerAccount, error) {
	sql := `UPDATE consumer_account
				    SET 
				        balance = (SELECT balance FROM 
						consumer_account) + $2

				WHERE consumer_id = $1
			returning *`

	newCourier := domain.ConsumerAccount{}
	if err := c.db.QueryRow(sql, account.ConsumerID, account.Balance).
		Scan(&newCourier.ConsumerID, &newCourier.Balance, &newCourier.CreatedAt, &newCourier.UpdatedAt); err != nil {
		return &domain.ConsumerAccount{}, err
	}

	return &newCourier, nil
}

func (c ConsumerAccountingStorage) SubFromConsumerAccount(account domain.ConsumerAccount) (*domain.ConsumerAccount, error) {
	sql := `UPDATE consumer_account
				    SET 
				        balance = (SELECT balance FROM 
						consumer_account) - $2

				WHERE consumer_id = $1
			returning *`

	newCourier := domain.ConsumerAccount{}
	if err := c.db.QueryRow(sql, account.ConsumerID, account.Balance).
		Scan(&newCourier.ConsumerID, &newCourier.Balance, &newCourier.CreatedAt, &newCourier.UpdatedAt); err != nil {
		return &domain.ConsumerAccount{}, err
	}

	return &newCourier, nil
}
