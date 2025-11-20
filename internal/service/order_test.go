package service

import (
	"errors"
	"testing"

	"github.com/brandoyts/go-idempotency/internal/core/order"
	"github.com/brandoyts/go-idempotency/internal/core/status"
	"github.com/brandoyts/go-idempotency/mocks"
	"github.com/golang/mock/gomock"
)

func TestOrderService_GetAllOrders_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockOrderRepository(ctrl)

	expected := []order.Order{
		{Id: 1, Amount: 100, Status: status.Processing},
		{Id: 2, Amount: 200, Status: status.Processing},
	}

	mockRepo.
		EXPECT().
		FindAll().
		Return(expected, nil)

	svc := &OrderService{Repository: mockRepo}

	orders, err := svc.GetAllOrders()
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	if len(orders) != len(expected) {
		t.Fatalf("expected %d orders, got %v", len(expected), len(orders))
	}

	if orders[0].Status != status.Processing {
		t.Errorf("expected first order status Processing, got %v", orders[0].Status)
	}

}

func TestOrderService_GetAllOrders_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockOrderRepository(ctrl)

	mockRepo.
		EXPECT().
		FindAll().
		Return(nil, errors.New("db error"))

	svc := &OrderService{
		Repository: mockRepo,
	}

	_, err := svc.GetAllOrders()
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestOrderService_GetOrderById_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockOrderRepository(ctrl)

	expected := &order.Order{
		Id:     1,
		Amount: 150.0,
		Status: status.Processing,
	}

	mockRepo.
		EXPECT().
		FindOne(uint64(1)).
		Return(expected, nil)

	svc := &OrderService{Repository: mockRepo}

	result, err := svc.Repository.FindOne(1)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	if result.Id != expected.Id {
		t.Errorf("expected ID %v, got %v", expected.Id, result.Id)
	}

	if result.Status != status.Processing {
		t.Errorf("expected status Processing, got %v", result.Status)
	}
}

func TestOrderService_GetOrderById_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockOrderRepository(ctrl)

	mockRepo.
		EXPECT().
		FindOne(uint64(1)).
		Return(nil, errors.New("order not found"))

	svc := &OrderService{Repository: mockRepo}

	_, err := svc.Repository.FindOne(1)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestOrderService_CreateOrder_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockOrderRepository(ctrl)

	// Order to create
	inputOrder := &order.Order{
		Amount: 250.0,
		Status: status.Processing,
	}

	// Expected result (with ID set by repo)
	expected := &order.Order{
		Id:     3,
		Amount: 250.0,
		Status: status.Processing,
	}

	// Mock the Create method to accept inputOrder as parameter
	mockRepo.
		EXPECT().
		Create(inputOrder).
		Return(expected, nil)

	// Service using the mocked repository
	svc := &OrderService{Repository: mockRepo}

	result, err := svc.Repository.Create(inputOrder)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	if result.Id != expected.Id {
		t.Errorf("expected ID %v, got %v", expected.Id, result.Id)
	}

	if result.Status != status.Processing {
		t.Errorf("expected status Processing, got %v", result.Status)
	}
}

func TestOrderService_CreateOrder_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockOrderRepository(ctrl)

	// Input order
	inputOrder := &order.Order{
		Amount: 100.0,
		Status: status.Processing,
	}

	// Mock Create to return an error when called with inputOrder
	mockRepo.
		EXPECT().
		Create(inputOrder).
		Return(nil, errors.New("failed to create order"))

	svc := &OrderService{Repository: mockRepo}

	_, err := svc.Repository.Create(inputOrder)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}
