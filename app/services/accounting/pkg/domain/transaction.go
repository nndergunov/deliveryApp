package domain

import "time"

type Transaction struct {
	ID            int
	FromAccountID int
	ToAccountID   int
	Amount        float64
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Valid         bool
}
