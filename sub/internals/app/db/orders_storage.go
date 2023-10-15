package db

import (
	"database/sql"
)

type OrdersStorage struct {
	db *sql.DB
}

func NewOrdersStorage(db *sql.DB) *OrdersStorage {
	ordersStorage := new(OrdersStorage)
	ordersStorage.db = db
	return ordersStorage
}

func (o *OrdersStorage) GetOrderByID(id string) (result string, err error) {
	query := "SELECT id, order_details FROM orders WHERE id = $1"

	err = o.db.QueryRow(query, id).Scan(&id, &result)

	if err != nil {
		return result, err
	}
	return result, nil
}

func (o *OrdersStorage) AddOrderToDB(id string, orderDetails []byte) error {
	query := "INSERT INTO orders(id, order_details) VALUES($1, $2)"

	_, err := o.db.Exec(query, id, orderDetails)

	if err != nil {
		return err
	}

	return nil
}
