package domain

import "time"

type CourierAccount struct {
	CourierID uint64
	Balance    int64
	CreatedAt  time.Time
	UpdatedAt  time.Time
}