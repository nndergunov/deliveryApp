package accountstorage

import (
	"database/sql"

	"github.com/nndergunov/delivryApp/app/services/accounting/pkg/domain"
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

func (c Storage) InsertNewAccount(account domain.Account) (*domain.Account, error) {
	sql := `INSERT INTO account
    			(user_id, user_type, created_at, updated_at)
			VALUES ($1, $2, now(), now())
			returning *`

	newAccount := domain.Account{}
	if err := c.db.QueryRow(sql, account.UserID, account.UserType).
		Scan(&newAccount.ID, &newAccount.UserID, &newAccount.UserType, &newAccount.Balance, &newAccount.CreatedAt, &newAccount.UpdatedAt); err != nil {
		return &domain.Account{}, err
	}

	return &newAccount, nil
}

func (c Storage) GetAccountByID(id int) (*domain.Account, error) {
	sql := `SELECT * 
			FROM account
			WHERE id = $1;`

	account := domain.Account{}

	if err := c.db.QueryRow(sql, id).Scan(&account.ID, &account.UserID, &account.UserType, &account.Balance,
		&account.CreatedAt, &account.UpdatedAt); err != nil {
		return nil, err
	}

	return &account, nil
}

func (c Storage) GetAccountListByParam(param domain.SearchParam) ([]domain.Account, error) {
	sql := `SELECT * 
			FROM account`

	where := " WHERE 1=1"

	userID := param["user_id"]
	if userID != "" {
		where = where + " AND user_id = " + userID + ""
	}

	userType := param["user_type"]
	if userType != "" {
		where = where + " AND user_type = '" + userType + "'"
	}

	sql = sql + where

	var accountList []domain.Account

	rows, err := c.db.Query(sql)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var account domain.Account
		if err := rows.Scan(&account.ID, &account.UserID, &account.UserType, &account.Balance,
			&account.CreatedAt, &account.UpdatedAt); err != nil {
			break
		}
		accountList = append(accountList, account)
	}

	return accountList, nil
}

func (c Storage) DeleteAccount(id int) error {
	sql := `DELETE 
			FROM account
			WHERE id = $1
	`
	if _, err := c.db.Exec(sql, id); err != nil {
		return err
	}

	return nil
}

func (c Storage) AddToAccountBalance(tr domain.Transaction) (*domain.Transaction, error) {
	var err error
	tx, err := c.db.Begin()
	if err != nil {
		return nil, err
	}

	defer func() {
		err = tx.Rollback()
	}()

	sql := `
			UPDATE account
    		SET balance = (SELECT balance FROM account WHERE id = $1) + $2
    		WHERE id = $1;`

	if _, err := tx.Exec(sql, tr.ToAccountID, tr.Amount); err != nil {
		return nil, err
	}

	sql2 := `INSERT
			 INTO transactions (from_account_id, to_account_id, amount, created_at, updated_at, valid)
			 VALUES (0, $1, $2, now(), now(), true)
			 RETURNING *;`

	newTr := domain.Transaction{}
	err = tx.QueryRow(sql2, tr.ToAccountID, tr.Amount).
		Scan(&newTr.ID, &newTr.FromAccountID, &newTr.ToAccountID, &newTr.Amount, &newTr.CreatedAt, &newTr.UpdatedAt, &newTr.Valid)
	if err != nil {
		return nil, err
	}

	if err = tx.Commit(); err != nil {
		return nil, err
	}

	return &newTr, err
}

func (c Storage) SubFromAccountBalance(tr domain.Transaction) (*domain.Transaction, error) {
	var err error
	tx, err := c.db.Begin()
	if err != nil {
		return nil, err
	}

	defer func() {
		err = tx.Rollback()
	}()

	sql := `UPDATE account
    		SET balance = (SELECT balance FROM account WHERE id = $1) - $2
    		WHERE id = $1;`

	if _, err := tx.Exec(sql, tr.FromAccountID, tr.Amount); err != nil {
		return nil, err
	}

	sql2 := `INSERT
			 INTO transactions (from_account_id, to_account_id, amount, created_at, updated_at, valid)
			 VALUES ($1, 0, $2, now(), now(), true)
			 RETURNING *;`

	newTr := domain.Transaction{}
	err = tx.QueryRow(sql2, tr.FromAccountID, tr.Amount).
		Scan(&newTr.ID, &newTr.FromAccountID, &newTr.ToAccountID, &newTr.Amount, &newTr.CreatedAt, &newTr.UpdatedAt, &newTr.Valid)
	if err != nil {
		return nil, err
	}

	if err = tx.Commit(); err != nil {
		return nil, err
	}

	return &newTr, err
}

func (c Storage) Transact(tr domain.Transaction) (*domain.Transaction, error) {
	var err error
	tx, err := c.db.Begin()
	if err != nil {
		return nil, err
	}

	defer func() {
		err = tx.Rollback()
	}()

	sql := `UPDATE account
			SET balance = (SELECT balance FROM account WHERE id = $1) - $2
			WHERE id = $1;`

	if _, err := tx.Exec(sql, tr.FromAccountID, tr.Amount); err != nil {
		return nil, err
	}

	sql2 := `UPDATE account
			 SET balance = (SELECT balance FROM account WHERE id = $1) + $2
			 WHERE id = $1;`

	if _, err := tx.Exec(sql2, tr.ToAccountID, tr.Amount); err != nil {
		return nil, err
	}

	sql3 := `INSERT
			 INTO transactions (from_account_id, to_account_id, amount, created_at, updated_at, valid)
			 VALUES ($1, $2, $3, now(), now(), true)
			 RETURNING *;`

	newTr := domain.Transaction{}
	err = tx.QueryRow(sql3, tr.FromAccountID, tr.ToAccountID, tr.Amount).
		Scan(&newTr.ID, &newTr.FromAccountID, &newTr.ToAccountID, &newTr.Amount, &newTr.CreatedAt, &newTr.UpdatedAt, &newTr.Valid)
	if err != nil {
		return nil, err
	}

	if err = tx.Commit(); err != nil {
		return nil, err
	}

	return &newTr, nil
}
