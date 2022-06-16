package domain

import "time"

type Consumer struct {
	ID        int
	Firstname string
	Lastname  string
	Email     string
	Phone     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type SearchParam map[string]string
