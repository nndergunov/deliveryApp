package deliverystorage

import (
	"database/sql"

	"github.com/nndergunov/deliveryApp/app/services/delivery/pkg/domain"
)

// Params is the input parameter struct for the module that contains its dependencies
type Params struct {
	DB *sql.DB
}

type DeliveryStorage struct {
	db *sql.DB
}

func NewDeliveryStorage(p Params) *DeliveryStorage {
	return &DeliveryStorage{
		db: p.DB,
	}
}

// AssignOrder store assigned assignOrder to the order
func (c DeliveryStorage) AssignOrder(order domain.AssignOrder) (*domain.AssignOrder, error) {
	sql := `INSERT INTO
				delivery
					(order_id, courier_id)
			VALUES($1,$2)
			returning *`

	newAssignedCourier := domain.AssignOrder{}
	if err := c.db.QueryRow(sql, order.OrderID, order.CourierID).
		Scan(&newAssignedCourier.OrderID, &newAssignedCourier.CourierID); err != nil {
		return nil, err
	}

	return &newAssignedCourier, nil
}

// DeleteAssignedOrder store assigned assignOrder to the order
func (c DeliveryStorage) DeleteAssignedOrder(orderID int) error {
	sql := `DELETE FROM delivery
			WHERE order_id = $1`
	if _, err := c.db.Exec(sql, orderID); err != nil {
		return err
	}

	return nil
}
