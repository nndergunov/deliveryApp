package db

import "errors"

var errExistsInDatabase = errors.New("current entry already exists in this database")
