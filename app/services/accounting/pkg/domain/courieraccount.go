package domain

import "time"

type CourierAccount struct {
	CourierID int
	Balance   int
	CreatedAt time.Time
	UpdatedAt time.Time
}
