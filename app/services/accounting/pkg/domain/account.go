package domain

import "time"

type Account struct {
	ID        int
	UserID    int
	UserType  string
	Balance   int
	CreatedAt time.Time
	UpdatedAt time.Time
}
