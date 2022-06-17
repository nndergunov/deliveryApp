package domain

import "time"

type Account struct {
	ID        int
	UserID    int
	UserType  string
	Balance   float64
	CreatedAt time.Time
	UpdatedAt time.Time
}

type SearchParam map[string]string
