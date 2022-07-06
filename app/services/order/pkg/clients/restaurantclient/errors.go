package restaurantclient

import "errors"

var ErrRestaurantFail = errors.New("http request to restaurant service resulted in error")
