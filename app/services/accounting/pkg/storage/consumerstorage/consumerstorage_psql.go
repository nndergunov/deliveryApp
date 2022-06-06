package consumerstorage

import (
	"database/sql"

	"accounting/pkg/domain"
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

func (c Storage) InsertNewConsumerAccount(consumerAccount domain.ConsumerAccount) (*domain.ConsumerAccount, error) {

	sql := `INSERT INTO consumer_account
    			(consumer_id, created_at, updated_at)
			VALUES ($1, now(), now())
			returning *`

	newCourier := domain.ConsumerAccount{}
	if err := c.db.QueryRow(sql, consumerAccount.ConsumerID).
		Scan(&newCourier.ConsumerID, &newCourier.Balance, &newCourier.CreatedAt, &newCourier.UpdatedAt); err != nil {
		return &domain.ConsumerAccount{}, err
	}

	return &newCourier, nil
}

func (c Storage) GetConsumerAccountByID(id uint64) (*domain.ConsumerAccount, error) {

	sql := `SELECT * FROM 
				consumer_account
			WHERE
				consumer_id = $1`

	consumerAccount := domain.ConsumerAccount{}

	if err := c.db.QueryRow(sql, id).Scan(&consumerAccount.ConsumerID, &consumerAccount.Balance, &consumerAccount.CreatedAt,
		&consumerAccount.UpdatedAt); err != nil {
		return nil, err
	}

	return &consumerAccount, nil
}

func (c Storage) DeleteConsumerAccount(consumerID uint64) error {
	sql := `DELETE FROM 
				consumer_account
			WHERE consumer_id = $1
	`
	if _, err := c.db.Exec(sql, consumerID); err != nil {
		return err
	}

	return nil
}

func (c Storage) AddToBalanceConsumerAccount(account domain.ConsumerAccount) (*domain.ConsumerAccount, error) {
	sql := `UPDATE consumer_account
			SET balance = (SELECT balance FROM consumer_account) + $2
			WHERE consumer_id = $1
			returning *`

	newCourier := domain.ConsumerAccount{}
	if err := c.db.QueryRow(sql, account.ConsumerID, account.Balance).
		Scan(&newCourier.ConsumerID, &newCourier.Balance, &newCourier.CreatedAt, &newCourier.UpdatedAt); err != nil {
		return &domain.ConsumerAccount{}, err
	}

	return &newCourier, nil
}

func (c Storage) SubFromBalanceConsumerAccount(account domain.ConsumerAccount) (*domain.ConsumerAccount, error) {
	sql := `UPDATE consumer_account
			SET balance = (SELECT balance FROM consumer_account) - $2
			WHERE consumer_id = $1
			returning *`

	newCourier := domain.ConsumerAccount{}
	if err := c.db.QueryRow(sql, account.ConsumerID, account.Balance).
		Scan(&newCourier.ConsumerID, &newCourier.Balance, &newCourier.CreatedAt, &newCourier.UpdatedAt); err != nil {
		return &domain.ConsumerAccount{}, err
	}

	return &newCourier, nil
}
