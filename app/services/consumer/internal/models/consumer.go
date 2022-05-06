package models

import (
	"time"
)

type Consumer struct {
	ID          uint64    `json:"id" yaml:"id"`
	CountryCode string    `json:"countryCode"`
	PhoneNumber string    `json:"phoneNumber" yaml:"phoneNumber"`
	Firstname   string    `json:"firstname" yaml:"firstname"`
	Lastname    string    `json:"lastname,omitempty" yaml:"lastname,omitempty"`
	Email       string    `json:"email" yaml:"email"`
	CreatedAt   time.Time `json:"createdAt" yaml:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt" yaml:"updatedAt"`
	Status      string    `json:"status" yaml:"status"`
}

// Fields will return all fields of this type
func (c *Consumer) Fields() []interface{} {
	return []interface{}{
		&c.ID,
		&c.CountryCode,
		&c.PhoneNumber,
		&c.Firstname,
		&c.Lastname,
		&c.Email,
		&c.Status,
	}
}
