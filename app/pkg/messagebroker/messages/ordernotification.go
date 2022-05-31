package messages

const (
	// CreatedChange is a notification message that shows that the order was created.
	CreatedChange = "order created"
	// UpdatedChange is a notification message that shows that the order was updated.
	UpdatedChange = "order updated"
	// StatusUpdatedChange is a notification message that shows that the status of an order was updated.
	StatusUpdatedChange = "order status updated"
)

type OrderNotification struct {
	Data string
}
