package accountingclient

import "errors"

// ErrAccountingServiceFail returns is the request to accounting service results
// in http error.
var ErrAccountingServiceFail = errors.New("request to accounting service resulted in HTTP error")
