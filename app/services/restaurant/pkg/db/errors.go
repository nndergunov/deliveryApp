package db

import "github.com/friendsofgo/errors"

var (
	errExistsInDatabase    = errors.New("current entry already exists in this database")
	errNotExistsInDatabase = errors.New("could not retrieve entry from database")
)
