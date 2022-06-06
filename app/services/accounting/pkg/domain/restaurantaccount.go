package domain

import "time"

type RestaurantAccount struct {
	RestaurantID int
	Balance      int
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
