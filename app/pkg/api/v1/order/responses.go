package order

type ReturnOrders struct {
	Orders []struct {
		FromUserID       int
		DeliveryLocation string
		SpecialRequests  string
		Order            map[string]int
	}
}

type ReturnOrderStatus struct {
	Status string
}
