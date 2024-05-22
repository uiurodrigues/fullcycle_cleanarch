package database

import (
	"database/sql"

	"github.com/fullcycle_cleanarch/internal/entity"
)

type OrderRepository struct {
	Db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{Db: db}
}

func (r *OrderRepository) Save(order *entity.Order) error {
	stmt, err := r.Db.Prepare("INSERT INTO orders (id, price, tax, final_price) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(order.ID, order.Price, order.Tax, order.FinalPrice)
	if err != nil {
		return err
	}
	return nil
}

func (r *OrderRepository) GetTotal() (int, error) {
	var total int
	err := r.Db.QueryRow("Select count(*) from orders").Scan(&total)
	if err != nil {
		return 0, err
	}
	return total, nil
}

func (r *OrderRepository) List() ([]entity.Order, error) {
	rows, err := r.Db.Query("Select * from orders")
	if err != nil {
		return []entity.Order{}, err
	}
	defer rows.Close()

	var ordersList []entity.Order
	for rows.Next() {
		var order entity.Order
		err := rows.Scan(&order.ID, &order.Price, &order.Tax, &order.FinalPrice)
		if err != nil {
			return []entity.Order{}, err
		}
		ordersList = append(ordersList, order)
	}
	if err = rows.Err(); err != nil {
		return []entity.Order{}, err
	}

	return ordersList, nil
}
