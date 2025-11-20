package postgres

import (
	"errors"
	"testing"

	idempotencyKey "github.com/brandoyts/go-idempotency/internal/core/idempotency_key"
	"github.com/brandoyts/go-idempotency/internal/core/status"
	"github.com/brandoyts/go-idempotency/internal/infrastructure/db"
	"github.com/brandoyts/go-idempotency/mocks"
	"github.com/golang/mock/gomock"
)

func TestIdempotencyKeyRepository_FindOne_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDB := mocks.NewMockDB(ctrl)

	expected := &idempotencyKey.IdempotencyKey{
		IdempotencyKey: "abc-123",
		OrderId:        1,
		Status:         status.Processing,
	}

	mockDB.
		EXPECT().
		First(gomock.Any(), "idempotency_key = ?", expected.IdempotencyKey).
		DoAndReturn(func(dest interface{}, conds ...interface{}) *db.DBResult {
			*dest.(*idempotencyKey.IdempotencyKey) = *expected
			return &db.DBResult{Error: nil}
		})

	repo := NewIdempotencyKeyRepository(mockDB)

	result, err := repo.FindOne(expected.IdempotencyKey)
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}

	if result.IdempotencyKey != expected.IdempotencyKey {
		t.Errorf("expected %s, got %s", expected.IdempotencyKey, result.IdempotencyKey)
	}
}

func TestIdempotencyKeyRepository_FindOne_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDB := mocks.NewMockDB(ctrl)

	mockDB.
		EXPECT().
		First(gomock.Any(), "idempotency_key = ?", "missing-key").
		Return(&db.DBResult{Error: errors.New("record not found")})

	repo := NewIdempotencyKeyRepository(mockDB)

	_, err := repo.FindOne("missing-key")
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestIdempotencyKeyRepository_Create_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDB := mocks.NewMockDB(ctrl)

	newKey := &idempotencyKey.IdempotencyKey{
		IdempotencyKey: "new-123",
		OrderId:        2,
		Status:         status.Done,
	}

	mockDB.
		EXPECT().
		Create(newKey).
		Return(&db.DBResult{Error: nil})

	repo := NewIdempotencyKeyRepository(mockDB)

	result, err := repo.Create(newKey)
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}

	if result.IdempotencyKey != "new-123" {
		t.Errorf("expected new-123, got %s", result.IdempotencyKey)
	}
}

func TestIdempotencyKeyRepository_Create_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDB := mocks.NewMockDB(ctrl)

	newKey := &idempotencyKey.IdempotencyKey{
		IdempotencyKey: "fail-999",
	}

	mockDB.
		EXPECT().
		Create(newKey).
		Return(&db.DBResult{Error: errors.New("insert failed")})

	repo := NewIdempotencyKeyRepository(mockDB)

	_, err := repo.Create(newKey)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}
