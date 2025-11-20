package idempotencyKey

//go:generate mockgen  -source=interface.go -destination=../../../mocks/idempotency_key_repository_mock.go -package=mocks
type IdempotencyKeyRepository interface {
	FindOne(id string) (*IdempotencyKey, error)
	Create(item *IdempotencyKey) (*IdempotencyKey, error)
}
