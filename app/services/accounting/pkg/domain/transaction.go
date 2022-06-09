package domain

import "time"

type Transaction struct {
	ID            int
	FromAccountID int
	ToAccountID   int
	Amount        int
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Valid         bool
}
