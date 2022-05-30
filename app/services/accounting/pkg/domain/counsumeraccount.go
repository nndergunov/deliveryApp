package domain

import "time"

type ConsumerAccount struct {
	ConsumerID uint64
	Balance    int64
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type SearchParam map[string]string
