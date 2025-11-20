package postgres

import (
	"errors"
	"testing"

	"github.com/brandoyts/go-idempotency/internal/core/order"
	"github.com/brandoyts/go-idempotency/internal/infrastructure/db"
	"github.com/brandoyts/go-idempotency/mocks"
	"github.com/golang/mock/gomock"
)

func TestOrderRepository_FindAll_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDB := mocks.NewMockDB(ctrl)

	expected := []order.Order{
		{Id: 1, Amount: 100},
		{Id: 2, Amount: 200},
	}

	mockDB.
		EXPECT().
		Find(gomock.Any()).
		DoAndReturn(func(dest interface{}, _ ...interface{}) *db.DBResult {
			out := dest.(*[]order.Order)
			*out = expected
			return &db.DBResult{Error: nil}
		})

	repo := NewOrderRepository(mockDB)

	results, err := repo.FindAll()
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}

	if len(results) != 2 {
		t.Errorf("expected 2 orders, got %d", len(results))
	}
}

func TestOrderRepository_FindOne_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDB := mocks.NewMockDB(ctrl)

	expected := order.Order{Id: 1, Amount: 150}

	mockDB.
		EXPECT().
		First(gomock.Any(), "id = ?", expected.Id).
		DoAndReturn(func(dest interface{}, _ ...interface{}) *db.DBResult {
			out := dest.(*order.Order)
			*out = expected
			return &db.DBResult{Error: nil}
		})

	repo := NewOrderRepository(mockDB)

	result, err := repo.FindOne(1)
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}

	if result.Id != expected.Id {
		t.Errorf("expected id %v, got %v", expected.Id, result.Id)
	}
	if result.Amount != expected.Amount {
		t.Errorf("expected amount %v, got %v", expected.Amount, result.Amount)
	}
}

func TestOrderRepository_FindOne_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDB := mocks.NewMockDB(ctrl)

	mockDB.
		EXPECT().
		First(gomock.Any(), "id = ?", uint64(99)).
		DoAndReturn(func(dest interface{}, _ ...interface{}) *db.DBResult {
			return &db.DBResult{Error: errors.New("record not found")}
		})

	repo := NewOrderRepository(mockDB)

	_, err := repo.FindOne(99)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if err.Error() != "record not found" {
		t.Fatalf("expected 'record not found', got %v", err)
	}
}

func TestOrderRepository_Create_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDB := mocks.NewMockDB(ctrl)

	newOrder := &order.Order{Amount: 250}

	mockDB.
		EXPECT().
		Create(newOrder).
		DoAndReturn(func(dest interface{}) *db.DBResult {
			o := dest.(*order.Order)
			o.Id = 3
			return &db.DBResult{Error: nil}
		})

	repo := NewOrderRepository(mockDB)

	result, err := repo.Create(newOrder)
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}

	if result.Id != 3 {
		t.Errorf("expected id 3, got %v", result.Id)
	}
}

func TestOrderRepository_Create_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDB := mocks.NewMockDB(ctrl)

	newOrder := &order.Order{Amount: 999}

	mockDB.
		EXPECT().
		Create(newOrder).
		Return(&db.DBResult{Error: errors.New("insert failed")})

	repo := NewOrderRepository(mockDB)

	_, err := repo.Create(newOrder)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}
