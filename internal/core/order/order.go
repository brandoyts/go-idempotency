package order

import "github.com/brandoyts/go-idempotency/internal/core/status"

type Order struct {
	Id     uint64 `gorm:"primaryKey"`
	Amount float32
	Status status.Status
}
