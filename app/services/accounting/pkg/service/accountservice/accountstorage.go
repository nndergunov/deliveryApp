package accountservice

import "accounting/pkg/domain"

// ConsumerStorage is the interface for the accounting storage.
type ConsumerStorage interface {
	InsertNewAccount(account domain.Account) (*domain.Account, error)
	GetAccountByUserID(userID int, userType string) (*domain.Account, error)
	GetAccount(id int) (*domain.Account, error)
	DeleteAccount(consumerID int) error

	AddToAccountBalance(tr domain.Transaction) (*domain.Transaction, error)
	SubFromAccountBalance(tr domain.Transaction) (*domain.Transaction, error)
	Transact(tr domain.Transaction) (*domain.Transaction, error)
}
