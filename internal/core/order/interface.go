package order

//go:generate mockgen  -source=interface.go -destination=../../../mocks/order_repository_mock.go -package=mocks
type OrderRepository interface {
	FindAll() ([]Order, error)
	FindOne(id uint64) (*Order, error)
	Create(item *Order) (*Order, error)
}
