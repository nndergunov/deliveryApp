package restaurantclient

import "errors"

// ErrRestaurantFail returns if the request to the restaurant service resulted in
// an http error.
var ErrRestaurantFail = errors.New("http request to restaurant service resulted in error")
