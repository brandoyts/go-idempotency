package postgres

import (
	idempotencyKey "github.com/brandoyts/go-idempotency/internal/core/idempotency_key"
	"github.com/brandoyts/go-idempotency/internal/infrastructure/db"
)

type IdempotencyKeyRepository struct {
	db db.DB
}

func NewIdempotencyKeyRepository(db db.DB) *IdempotencyKeyRepository {
	return &IdempotencyKeyRepository{db: db}
}

func (r *IdempotencyKeyRepository) FindOne(idempotentKey string) (*idempotencyKey.IdempotencyKey, error) {
	var o idempotencyKey.IdempotencyKey
	if err := r.db.First(&o, "idempotency_key = ?", idempotentKey).Error; err != nil {
		return nil, err
	}
	return &o, nil
}

func (r *IdempotencyKeyRepository) Create(o *idempotencyKey.IdempotencyKey) (*idempotencyKey.IdempotencyKey, error) {
	if err := r.db.Create(o).Error; err != nil {
		return nil, err
	}
	return o, nil
}
