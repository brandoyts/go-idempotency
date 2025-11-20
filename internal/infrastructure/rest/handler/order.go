package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	idempotencyKey "github.com/brandoyts/go-idempotency/internal/core/idempotency_key"
	"github.com/brandoyts/go-idempotency/internal/core/order"
	"github.com/brandoyts/go-idempotency/internal/service"
)

type OrderHandler struct {
	orderService          *service.OrderService
	idempotencyKeyService *service.IdempotencyKeyService
}

func NewOrderHandler(orderService *service.OrderService, idempotencyKeyService *service.IdempotencyKeyService) *OrderHandler {
	return &OrderHandler{
		orderService:          orderService,
		idempotencyKeyService: idempotencyKeyService,
	}
}

func (h *OrderHandler) ListOrders(w http.ResponseWriter, r *http.Request) {
	orders, err := h.orderService.GetAllOrders()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(orders)
}

func (h *OrderHandler) GetOrder(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "missing id parameter", http.StatusBadRequest)
		return
	}

	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		http.Error(w, "invalid id parameter", http.StatusBadRequest)
		return
	}

	o, err := h.orderService.GetOrderById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(o)
}

func (h *OrderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var o order.Order
	if err := json.NewDecoder(r.Body).Decode(&o); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createdOrder, err := h.orderService.CreateOrder(&o)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	idempotentKey := r.Header.Get("Idempotency-Key")

	_, err = h.idempotencyKeyService.Create(&idempotencyKey.IdempotencyKey{
		IdempotencyKey: idempotentKey,
		OrderId:        createdOrder.Id,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdOrder)
}
