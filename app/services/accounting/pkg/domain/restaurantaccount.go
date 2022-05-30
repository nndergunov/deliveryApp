package domain

import "time"

type RestaurantAccount struct {
	RestaurantID uint64
	Balance      int64
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
