package db

import (
	"database/sql"
	"errors"
	"fmt"

	// Postgres driver.
	_ "github.com/lib/pq"
	"github.com/nndergunov/deliveryApp/app/services/order/pkg/db/internal/models"
	"github.com/nndergunov/deliveryApp/app/services/order/pkg/domain"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type Database struct {
	db *sql.DB
}

func NewDatabase(dbURL string) (*Database, error) {
	database, err := sql.Open("postgres", dbURL)
	if err != nil {
		return nil, fmt.Errorf("NewDatabase: %w", err)
	}

	return &Database{
		db: database,
	}, nil
}

func (d Database) getOrderID(order domain.Order) (int, error) {
	dbOrder, err := models.Orders(qm.Where("customer_id=? and restaurant_id=? and order_items=? and status=?",
		order.FromUserID, order.RestaurantID, intArrToInt64Arr(order.OrderItems), order.Status.Status)).One(d.db)
	if err != nil {
		return 0, fmt.Errorf("getOrderID: %w", err)
	}

	return dbOrder.ID, nil
}

func (d Database) getOrderByID(orderID int) (*models.Order, error) {
	dbOrder, err := models.Orders(qm.Where("id=?", orderID)).One(d.db)
	if err != nil {
		return nil, fmt.Errorf("getOrderByID: %w", err)
	}

	return dbOrder, nil
}

func (d Database) GetAllOrders(params *domain.SearchParameters) ([]domain.Order, error) {
	var mods []qm.QueryMod

	if params != nil {
		if params.FromRestaurantID != nil {
			mods = append(mods, qm.Where("restaurant_id=?", params.FromRestaurantID))
		}

		if params.Statuses != nil {
			convertedStatuses := make([]interface{}, len(params.Statuses))
			for index, num := range params.Statuses {
				convertedStatuses[index] = num
			}

			mods = append(mods, qm.WhereIn("status in ?", convertedStatuses...))
		}

		if params.ExcludeStatuses != nil {
			convertedExclStatuses := make([]interface{}, len(params.ExcludeStatuses))
			for index, num := range params.ExcludeStatuses {
				convertedExclStatuses[index] = num
			}

			mods = append(mods, qm.WhereNotIn("status not in ?", convertedExclStatuses...))
		}
	}

	dbOrders, err := models.Orders(mods...).All(d.db)
	if err != nil {
		return nil, fmt.Errorf("GetAllOrders: %w", err)
	}

	orders := make([]domain.Order, 0, len(dbOrders))

	for _, dbOrder := range dbOrders {
		order := domain.Order{
			OrderID:      dbOrder.ID,
			FromUserID:   dbOrder.CustomerID,
			RestaurantID: dbOrder.RestaurantID,
			OrderItems:   int64ArrToIntArr(dbOrder.OrderItems),
			Status: domain.OrderStatus{
				OrderID: dbOrder.ID,
				Status:  dbOrder.Status,
			},
		}

		orders = append(orders, order)
	}

	return orders, nil
}

func (d Database) InsertOrder(order domain.Order) (int, error) {
	order.Status.Status = "initialOrder created"

	_, err := d.getOrderID(order)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return 0, fmt.Errorf("InsertOrder: %w", err)
		}
	} else {
		return 0, fmt.Errorf("InsertOrder: %w", errExistsInDatabase)
	}

	var dbOrder models.Order

	dbOrder.CustomerID = order.FromUserID
	dbOrder.RestaurantID = order.RestaurantID
	dbOrder.OrderItems = intArrToInt64Arr(order.OrderItems)
	dbOrder.Status = order.Status.Status

	err = dbOrder.Insert(d.db, boil.Infer())
	if err != nil {
		return 0, fmt.Errorf("InsertOrder: %w", err)
	}

	orderID, err := d.getOrderID(order)
	if err != nil {
		return 0, fmt.Errorf("InsertOrder: %w", err)
	}

	return orderID, nil
}

func (d Database) GetOrder(orderID int) (*domain.Order, error) {
	dbOrder, err := d.getOrderByID(orderID)
	if err != nil {
		return nil, fmt.Errorf("GetOrder: %w", err)
	}

	return &domain.Order{
		OrderID:      dbOrder.ID,
		FromUserID:   dbOrder.CustomerID,
		RestaurantID: dbOrder.RestaurantID,
		OrderItems:   int64ArrToIntArr(dbOrder.OrderItems),
		Status: domain.OrderStatus{
			OrderID: dbOrder.ID,
			Status:  dbOrder.Status,
		},
	}, nil
}

func (d Database) UpdateOrder(order domain.Order) error {
	dbOrder, err := d.getOrderByID(order.OrderID)
	if err != nil {
		return fmt.Errorf("UpdateOrder: %w", err)
	}

	dbOrder.CustomerID = order.FromUserID
	dbOrder.RestaurantID = order.RestaurantID
	dbOrder.OrderItems = intArrToInt64Arr(order.OrderItems)

	_, err = dbOrder.Update(d.db, boil.Infer())
	if err != nil {
		return fmt.Errorf("UpdateOrder: %w", err)
	}

	return nil
}

func (d Database) DeleteOrder(orderID int) error {
	dbOrder, err := d.getOrderByID(orderID)
	if err != nil {
		return fmt.Errorf("DeleteOrder: %w", err)
	}

	_, err = dbOrder.Delete(d.db)
	if err != nil {
		return fmt.Errorf("DeleteOrder: %w", err)
	}

	return nil
}

func (d Database) UpdateOrderStatus(orderID int, status string) error {
	dbOrder, err := d.getOrderByID(orderID)
	if err != nil {
		return fmt.Errorf("UpdateOrderStatus: %w", err)
	}

	dbOrder.Status = status

	_, err = dbOrder.Update(d.db, boil.Infer())
	if err != nil {
		return fmt.Errorf("UpdateOrderStatus: %w", err)
	}

	return nil
}
