package accountservice

import "accounting/pkg/domain"

// AccountStorage is the interface for the accounting storage.
type AccountStorage interface {
	InsertNewAccount(account domain.Account) (*domain.Account, error)
	GetAccountByID(id int) (*domain.Account, error)
	GetAccountListByParam(params domain.SearchParam) ([]domain.Account, error)
	DeleteAccount(consumerID int) error

	AddToAccountBalance(tr domain.Transaction) (*domain.Transaction, error)
	SubFromAccountBalance(tr domain.Transaction) (*domain.Transaction, error)
	Transact(tr domain.Transaction) (*domain.Transaction, error)
}
