package domain

import "time"

type ConsumerAccount struct {
	ConsumerID int
	Balance    int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
