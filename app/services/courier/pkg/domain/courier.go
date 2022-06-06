package domain

import "time"

type Courier struct {
	ID        int
	Username  string
	Password  string
	Firstname string
	Lastname  string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
	Phone     string
	Available bool
}

type SearchParam map[string]string
