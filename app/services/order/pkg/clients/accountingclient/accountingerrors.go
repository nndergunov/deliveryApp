package accountingclient

import "errors"

var ErrAccountingServiceFail = errors.New("request to accounting service resulted in HTTP error")
