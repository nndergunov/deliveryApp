package restaurantgrp

import "errors"

var errNoRestaurantIDParam = errors.New("restaurant_id param is not found")
var errIncorrectInputData = errors.New("incorrect input data")
var systemErr = errors.New("system error")
