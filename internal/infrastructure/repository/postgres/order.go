package postgres

import (
	"github.com/brandoyts/go-idempotency/internal/core/order"
	"github.com/brandoyts/go-idempotency/internal/infrastructure/db"
)

type OrderRepository struct {
	db db.DB
}

func NewOrderRepository(db db.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

func (or *OrderRepository) FindAll() ([]order.Order, error) {
	var orders []order.Order
	if err := or.db.Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}

func (r *OrderRepository) FindOne(id uint64) (*order.Order, error) {
	var o order.Order
	if err := r.db.First(&o, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &o, nil
}

func (r *OrderRepository) Create(o *order.Order) (*order.Order, error) {
	if err := r.db.Create(o).Error; err != nil {
		return nil, err
	}
	return o, nil
}
