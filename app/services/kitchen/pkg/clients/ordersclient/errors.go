package ordersclient

import "errors"

// ErrOrderServiceFailed returns when a request to the order service results in
// http error.
var ErrOrderServiceFailed = errors.New("http request to order service resulted in error")
