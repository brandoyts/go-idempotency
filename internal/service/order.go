package service

import "github.com/brandoyts/go-idempotency/internal/core/order"

type OrderService struct {
	Repository order.OrderRepository
}

func NewOrderService(repository order.OrderRepository) *OrderService {
	return &OrderService{Repository: repository}
}

func (s *OrderService) GetAllOrders() ([]order.Order, error) {
	result, err := s.Repository.FindAll()
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *OrderService) GetOrderById(id uint64) (*order.Order, error) {
	result, err := s.Repository.FindOne(id)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *OrderService) CreateOrder(item *order.Order) (*order.Order, error) {
	result, err := s.Repository.Create(item)
	if err != nil {
		return nil, err
	}

	return result, nil
}
