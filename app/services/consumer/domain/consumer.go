package domain

import "time"

type Consumer struct {
	ID               uint64
	Firstname        string
	Lastname         string
	Email            string
	Phone            string
	Createdat        time.Time
	Updatedat        time.Time
	ConsumerLocation ConsumerLocation
}
