package service

import "github.com/friendsofgo/errors"

var ErrItemIsNotInMenu = errors.New("item with such id was not found in the required restaurant")
