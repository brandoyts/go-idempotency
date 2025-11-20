package middleware

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/brandoyts/go-idempotency/internal/service"
)

type IdempotencyMiddleware struct {
	IdempotencyService *service.IdempotencyKeyService
	OrderService       *service.OrderService
}

// Request payload for /orders/charge
type ChargeRequest struct {
	OrderID uint64 `json:"order_id"`
}

// Middleware function
func (m *IdempotencyMiddleware) Handler() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			if r.URL.Path != "/orders/create" {
				next.ServeHTTP(w, r)
				return
			}

			// Check Idempotency-Key header
			key := r.Header.Get("Idempotency-Key")
			if key == "" {
				writeJSON(w, http.StatusBadRequest, map[string]string{"error": "Idempotency-Key header is required"})
				return
			}

			// Read request body
			bodyBytes, err := io.ReadAll(r.Body)
			if err != nil {
				writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid request body"})
				return
			}
			// Restore body for next handler
			r.Body = io.NopCloser(bytes.NewReader(bodyBytes))

			// Check if Idempotency-Key exists
			existingKey, err := m.IdempotencyService.FindOne(key)
			if err == nil && existingKey != nil {
				// Return stored status
				writeJSON(w, http.StatusOK, map[string]interface{}{
					"order_id": existingKey.OrderId,
					"status":   existingKey.Status,
				})
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

// Helper to write JSON
func writeJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}
