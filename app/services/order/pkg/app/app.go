package app

import "github.com/nndergunov/deliveryApp/app/services/order/pkg/app/domain"

type App struct {
	orders map[int][]domain.Order // map of orders with a key that is a Restaurant ID
}

func NewApp() *App {
	orders := make(map[int][]domain.Order)

	return &App{orders: orders}
}

func (a *App) AddOrder(restaurantID int, order domain.Order) {
	a.orders[restaurantID] = append(a.orders[restaurantID], order)
}

func (a App) ReturnIncompleteOrders(restaurantID int) []domain.Order {
	var incOrders []domain.Order

	Complete := "Delivered"

	for _, order := range a.orders[restaurantID] {
		if order.Status != Complete {
			incOrders = append(incOrders, order)
		}
	}

	return incOrders
}

func (a *App) UpdateOrderStatus() {
	// TODO after updating global pkg
}
