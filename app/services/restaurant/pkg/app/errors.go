package app

import "errors"

var (
	ErrIsNotInDatabase = errors.New("requested id is not in the database")
	ErrIsInDatabase    = errors.New("requested element is already in the database")
)
