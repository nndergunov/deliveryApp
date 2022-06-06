package restaurantstorage

import (
	"database/sql"

	"accounting/pkg/domain"
)

// Params is the input parameter struct for the module that contains its dependencies
type Params struct {
	DB *sql.DB
}

type RestaurantStorage struct {
	db *sql.DB
}

func NewStorage(p Params) *RestaurantStorage {
	return &RestaurantStorage{
		db: p.DB,
	}
}

func (c RestaurantStorage) InsertNewRestaurantAccount(restaurantAccount domain.RestaurantAccount) (*domain.RestaurantAccount, error) {

	sql := `INSERT INTO restaurant_account
    			(restaurant_id, created_at, updated_at)
			VALUES ($1, now(), now())
			returning *`

	newRestaurant := domain.RestaurantAccount{}
	if err := c.db.QueryRow(sql, restaurantAccount.RestaurantID).
		Scan(&newRestaurant.RestaurantID, &newRestaurant.Balance, &newRestaurant.CreatedAt, &newRestaurant.UpdatedAt); err != nil {
		return &domain.RestaurantAccount{}, err
	}

	return &newRestaurant, nil
}

func (c RestaurantStorage) GetRestaurantAccountByID(id uint64) (*domain.RestaurantAccount, error) {

	sql := `SELECT *
			FROM restaurant_account
			WHERE restaurant_id = $1`

	RestaurantAccount := domain.RestaurantAccount{}

	if err := c.db.QueryRow(sql, id).Scan(&RestaurantAccount.RestaurantID, &RestaurantAccount.Balance, &RestaurantAccount.CreatedAt,
		&RestaurantAccount.UpdatedAt); err != nil {
		return nil, err
	}

	return &RestaurantAccount, nil
}

func (c RestaurantStorage) DeleteRestaurantAccount(RestaurantID uint64) error {

	sql := `DELETE 
			FROM restaurant_account
			WHERE restaurant_id = $1
	`
	if _, err := c.db.Exec(sql, RestaurantID); err != nil {
		return err
	}

	return nil
}

func (c RestaurantStorage) AddToBalanceRestaurantAccount(account domain.RestaurantAccount) (*domain.RestaurantAccount, error) {

	sql := `UPDATE Restaurant_account
			SET balance = (SELECT balance FROM restaurant_account) + $2
			WHERE restaurant_id = $1
			returning *`

	newRestaurant := domain.RestaurantAccount{}
	if err := c.db.QueryRow(sql, account.RestaurantID, account.Balance).
		Scan(&newRestaurant.RestaurantID, &newRestaurant.Balance, &newRestaurant.CreatedAt, &newRestaurant.UpdatedAt); err != nil {
		return &domain.RestaurantAccount{}, err
	}

	return &newRestaurant, nil
}

func (c RestaurantStorage) SubFromBalanceRestaurantAccount(account domain.RestaurantAccount) (*domain.RestaurantAccount, error) {

	sql := `UPDATE restaurant_account
			SET balance = (SELECT balance FROM restaurant_account) - $2
			WHERE restaurant_id = $1
			returning *`

	newRestaurant := domain.RestaurantAccount{}
	if err := c.db.QueryRow(sql, account.RestaurantID, account.Balance).
		Scan(&newRestaurant.RestaurantID, &newRestaurant.Balance, &newRestaurant.CreatedAt, &newRestaurant.UpdatedAt); err != nil {
		return &domain.RestaurantAccount{}, err
	}

	return &newRestaurant, nil
}
