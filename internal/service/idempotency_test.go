package service

import (
	"errors"
	"testing"

	idempotencyKey "github.com/brandoyts/go-idempotency/internal/core/idempotency_key"
	"github.com/brandoyts/go-idempotency/internal/core/status"
	"github.com/brandoyts/go-idempotency/mocks"
	"github.com/golang/mock/gomock"
)

func TestService_FindOne_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockIdempotencyKeyRepository(ctrl)

	expected := &idempotencyKey.IdempotencyKey{
		IdempotencyKey: "abc-123",
		OrderId:        1,
		Status:         status.Processing,
	}

	mockRepo.
		EXPECT().
		FindOne("abc-123").
		Return(expected, nil)

	svc := NewIdempotencyKeyService(mockRepo)

	result, err := svc.FindOne("abc-123")
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}

	if result.IdempotencyKey != expected.IdempotencyKey {
		t.Errorf("expected key %s, got %s", expected.IdempotencyKey, result.IdempotencyKey)
	}
}

func TestService_FindOne_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockIdempotencyKeyRepository(ctrl)

	mockRepo.
		EXPECT().
		FindOne("missing-key").
		Return(nil, errors.New("not found"))

	svc := NewIdempotencyKeyService(mockRepo)

	_, err := svc.FindOne("missing-key")
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestService_Create_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockIdempotencyKeyRepository(ctrl)

	newKey := &idempotencyKey.IdempotencyKey{
		IdempotencyKey: "new-123",
		OrderId:        1,
		Status:         status.Done,
	}

	mockRepo.
		EXPECT().
		Create(newKey).
		Return(newKey, nil)

	svc := NewIdempotencyKeyService(mockRepo)

	result, err := svc.Create(newKey)
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}

	if result.IdempotencyKey != "new-123" {
		t.Errorf("expected new-123, got %s", result.IdempotencyKey)
	}
}

func TestService_Create_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockIdempotencyKeyRepository(ctrl)

	newKey := &idempotencyKey.IdempotencyKey{
		IdempotencyKey: "fail-999",
	}

	mockRepo.
		EXPECT().
		Create(newKey).
		Return(nil, errors.New("insert failed"))

	svc := NewIdempotencyKeyService(mockRepo)

	_, err := svc.Create(newKey)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}
