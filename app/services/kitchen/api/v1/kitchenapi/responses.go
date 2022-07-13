package kitchenapi

// Tasks contains information about what number of what dishes must be prepared.
// swagger:model
type Tasks struct {
	Tasks map[int]int
}
