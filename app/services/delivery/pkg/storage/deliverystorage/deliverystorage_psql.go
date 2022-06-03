package deliverystorage

import (
	"database/sql"
	"delivery/pkg/domain"
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

// AssignCourier store assigned courier to the order
func (c DeliveryStorage) AssignCourier(courierID int, orderID int) (*domain.AssignedCourier, error) {
	sql := `INSERT INTO
				delivery
					(order_id, courier_id, delivered)
			VALUES($1,$2, false)
			returning *`

	newAssignedCourier := domain.AssignedCourier{}
	if err := c.db.QueryRow(sql, courierID, orderID).
		Scan(&newAssignedCourier.OrderID, &newAssignedCourier.CourierID); err != nil {
		return nil, err
	}

	return &newAssignedCourier, nil
}
