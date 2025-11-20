package route

import (
	"net/http"

	"github.com/brandoyts/go-idempotency/internal/infrastructure/rest/handler"
	"github.com/brandoyts/go-idempotency/internal/infrastructure/rest/middleware"
	"github.com/brandoyts/go-idempotency/internal/service"
)

func RegisterOrderRoutes(
	mux *http.ServeMux,
	orderService *service.OrderService,
	idemService *service.IdempotencyKeyService,
) {
	h := handler.NewOrderHandler(orderService, idemService)

	// Middleware instance
	idemMiddleware := &middleware.IdempotencyMiddleware{
		OrderService:       orderService,
		IdempotencyService: idemService,
	}

	// /orders route (GET=List, POST=Create)
	mux.HandleFunc("/orders", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			h.ListOrders(w, r)
		case http.MethodPost:
			// Wrap POST with Idempotency middleware
			idemMiddleware.Handler()(http.HandlerFunc(h.CreateOrder)).ServeHTTP(w, r)
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// /orders/get route (GET by query param: /orders/get?id=...)
	mux.HandleFunc("/orders/get", h.GetOrder)

	mux.HandleFunc("/orders/create", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
		// Wrap Charge handler with Idempotency middleware
		idemMiddleware.Handler()(http.HandlerFunc(h.CreateOrder)).ServeHTTP(w, r)
	})
}
