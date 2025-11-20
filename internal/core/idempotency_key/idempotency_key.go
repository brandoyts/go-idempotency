package idempotencyKey

import (
	"time"

	"github.com/brandoyts/go-idempotency/internal/core/status"
)

type IdempotencyKey struct {
	Id             uint64 `gorm:"primaryKey"`
	IdempotencyKey string
	OrderId        uint64
	Status         status.Status
	Ttl            time.Time
}
