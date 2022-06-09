package accountingapi

type NewAccountRequest struct {
	UserID   int    `json:"userID"`
	UserType string `json:"userType"`
}

type TransactionRequest struct {
	FromAccountID int `json:"FromAccountID"`
	ToAccountID   int `json:"ToAccountID"`
	Amount        int `json:"Amount"`
}

type SubBalanceConsumerAccountRequest struct {
	Amount int `json:"amount"`
}
