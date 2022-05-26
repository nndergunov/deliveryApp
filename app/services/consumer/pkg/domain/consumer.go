package domain

import "time"

type Consumer struct {
	ID        uint64
	Firstname string
	Lastname  string
	Email     string
	Phone     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type SearchParam map[string]string
