package db

import "github.com/friendsofgo/errors"

var (
	ErrExistsInDatabase       = errors.New("current entry already exists in this database")
	ErrCouldNotCheckExistence = errors.New("could not check if current entry already exists")
	ErrNotExistsInDatabase    = errors.New("could not retrieve entry from database")
)
