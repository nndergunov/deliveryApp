package domain

import "time"

type Courier struct {
	ID        uint64
	Username  string
	Password  string
	Firstname string
	Lastname  string
	Email     string
	Createdat time.Time
	Updatedat time.Time
	Phone     string
	Status    string
	Available bool
}
