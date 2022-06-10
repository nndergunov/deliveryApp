package accountingapi

type NewAccountRequest struct {
	UserID   int    `json:"user_id"`
	UserType string `json:"user_type"`
}

type TransactionRequest struct {
	FromAccountID int `json:"from_account_id"`
	ToAccountID   int `json:"to_account_id"`
	Amount        int `json:"amount"`
}

type SubBalanceConsumerAccountRequest struct {
	Amount int `json:"amount"`
}
