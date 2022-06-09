package accountingapi

type NewAccountRequest struct {
	UserID   int    `json:"UserID"`
	UserType string `json:"UserType"`
}

type TransactionRequest struct {
	FromAccountID int `json:"FromAccountID"`
	ToAccountID   int `json:"ToAccountID"`
	Amount        int `json:"Amount"`
}

type SubBalanceConsumerAccountRequest struct {
	Amount int `json:"Amount"`
}
