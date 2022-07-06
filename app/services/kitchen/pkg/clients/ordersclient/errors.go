package ordersclient

import "errors"

var ErrOrderServiceFailed = errors.New("http request to order service resulted in error")
