package models

import (
	"time"
)

type NewCourierRequest struct {
	Username  string `json:"username" yaml:"username"`
	Password  string `json:"password" yaml:"password"`
	Firstname string `json:"firstname" yaml:"firstname"`
	Lastname  string `json:"lastname" yaml:"lastname"`
	Email     string `json:"email" yaml:"email"`
	Phone     string `json:"phone,omitempty" yaml:"phone,omitempty"`
}

type UpdateCourierRequest struct {
	Username  string `json:"username" yaml:"username"`
	Firstname string `json:"firstname" yaml:"firstname"`
	Lastname  string `json:"lastname" yaml:"lastname"`
	Email     string `json:"email" yaml:"email"`
	Phone     string `json:"phone,omitempty" yaml:"phone,omitempty"`
}

type CourierResponse struct {
	ID        int       `json:"id" yaml:"id"`
	Username  string    `json:"username" yaml:"username"`
	Password  string    `json:"password,omitempty" yaml:"password,omitempty"`
	Firstname string    `json:"firstname" yaml:"firstname"`
	Lastname  string    `json:"lastname,omitempty" yaml:"lastname,omitempty"`
	Email     string    `json:"email" yaml:"email"`
	Createdat time.Time `json:"createdat" yaml:"createdat"`
	Updatedat time.Time `json:"updatedat" yaml:"updatedat"`
	Phone     string    `json:"phone,omitempty" yaml:"phone,omitempty"`
	Status    string    `json:"status" yaml:"status"`
	Available bool      `json:"available" yaml:"available"`
}
