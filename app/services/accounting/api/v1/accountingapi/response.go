package accountingapi

import "time"

// AccountResponse contains information about the account.
// swagger:model
type AccountResponse struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	UserType  string    `json:"user_type"`
	Balance   float64   `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// AccountListResponse contains information about the accounts.
// swagger:model
type AccountListResponse struct {
	AccountList []AccountResponse
}

// TransactionResponse contains information about the transaction.
// swagger:model
type TransactionResponse struct {
	ID            int       `json:"id"`
	FromAccountID int       `json:"from_account_id"`
	ToAccountID   int       `json:"to_account_id"`
	Amount        float64   `json:"amount"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updatedAt"`
	Valid         bool      `json:"valid"`
}
