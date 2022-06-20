package accountingapi

import "time"

type AccountResponse struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	UserType  string    `json:"user_type"`
	Balance   int       `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
type AccountListResponse struct {
	AccountList []AccountResponse
}
type TransactionResponse struct {
	ID            int       `json:"id"`
	FromAccountID int       `json:"from_account_id"`
	ToAccountID   int       `json:"to_account_id"`
	Amount        int       `json:"amount"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updatedAt"`
	Valid         bool      `json:"valid"`
}
