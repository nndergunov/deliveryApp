package models

import (
	"time"
)

type Courier struct {
	ID        int       `json:"id" yaml:"id"`
	Username  string    `json:"username" yaml:"username"`
	Password  string    `json:"password" yaml:"password"`
	Firstname string    `json:"firstname" yaml:"firstname"`
	Lastname  string    `json:"lastname,omitempty" yaml:"lastname,omitempty"`
	Email     string    `json:"email" yaml:"email"`
	Createdat time.Time `json:"createdat" yaml:"createdat"`
	Updatedat time.Time `json:"updatedat" yaml:"updatedat"`
	Phone     string    `json:"phone,omitempty" yaml:"phone,omitempty"`
	Status    string    `json:"status" yaml:"status"`
	Available bool      `json:"available" yaml:"available"`
}

// Fields will return all fields of this type
func (c *Courier) Fields() []interface{} {
	return []interface{}{
		&c.ID,
		&c.Username,
		&c.Password,
		&c.Firstname,
		&c.Lastname,
		&c.Email,
		&c.Createdat,
		&c.Updatedat,
		&c.Phone,
		&c.Status,
		&c.Available,
	}
}
