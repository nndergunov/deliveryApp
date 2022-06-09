package accountingapi

import "time"

type AccountResponse struct {
	ID        int       `json:"ID"`
	UserID    int       `json:"UserID"`
	UserType  string    `json:"UserType"`
	Balance   int       `json:"Balance"`
	CreatedAt time.Time `json:"CreatedAt"`
	UpdatedAt time.Time `json:"UpdatedAt"`
}

type TransactionResponse struct {
	ID            int       `json:"ID"`
	FromAccountID int       `json:"FromAccountID"`
	ToAccountID   int       `json:"ToAccountID"`
	Amount        int       `json:"Amount"`
	CreatedAt     time.Time `json:"CreatedAt"`
	UpdatedAt     time.Time `json:"UpdatedAt"`
	Valid         bool      `json:"Valid"`
}
