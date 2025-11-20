package service

import (
	idempotencyKey "github.com/brandoyts/go-idempotency/internal/core/idempotency_key"
)

type IdempotencyKeyService struct {
	Repository idempotencyKey.IdempotencyKeyRepository
}

func NewIdempotencyKeyService(r idempotencyKey.IdempotencyKeyRepository) *IdempotencyKeyService {
	return &IdempotencyKeyService{
		Repository: r,
	}
}

func (s *IdempotencyKeyService) FindOne(key string) (*idempotencyKey.IdempotencyKey, error) {
	return s.Repository.FindOne(key)
}

func (s *IdempotencyKeyService) Create(record *idempotencyKey.IdempotencyKey) (*idempotencyKey.IdempotencyKey, error) {
	return s.Repository.Create(record)
}
