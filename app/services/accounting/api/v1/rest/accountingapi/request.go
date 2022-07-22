package accountingapi

// NewAccountRequest contains information about the account.
// swagger:model
type NewAccountRequest struct {
	// required: true
	UserID int `json:"user_id"`
	// required: true
	UserType string `json:"user_type"`
}

// TransactionRequest contains information about the transaction.
// swagger:model
type TransactionRequest struct {
	FromAccountID int `json:"from_account_id"`
	ToAccountID   int `json:"to_account_id"`
	// required: true
	Amount float64 `json:"amount"`
}
