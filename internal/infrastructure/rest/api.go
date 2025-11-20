package rest

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/brandoyts/go-idempotency/internal/infrastructure/rest/handler"
	"github.com/brandoyts/go-idempotency/internal/infrastructure/rest/middleware"
	"github.com/brandoyts/go-idempotency/internal/infrastructure/rest/route"
	"github.com/brandoyts/go-idempotency/internal/service"
)

type Services struct {
	OrderService          *service.OrderService
	IdempotencyKeyService *service.IdempotencyKeyService
}

type Api struct {
	server     http.Server
	handler    *http.ServeMux
	services   *Services
	middleware []func(http.Handler) http.Handler
}

func New(services *Services) *Api {
	mux := http.NewServeMux()

	a := &Api{
		handler:  mux,
		services: services,
		server: http.Server{
			Addr:    ":4000",
			Handler: mux,
		},
	}

	idempotencyMiddleware := &middleware.IdempotencyMiddleware{
		OrderService:       services.OrderService,
		IdempotencyService: services.IdempotencyKeyService,
	}

	a.use(idempotencyMiddleware.Handler())

	a.registerHandler("/health", handler.Health)
	route.RegisterOrderRoutes(mux, services.OrderService, services.IdempotencyKeyService)

	return a
}

func (a *Api) Start() {
	if err := a.server.ListenAndServe(); err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			fmt.Printf("error running http server: %s\n", err)
		}
	}
}

func (a *Api) registerHandler(path string, handler http.HandlerFunc) {
	a.handler.Handle(path, handler)
}

func (a *Api) use(mw func(http.Handler) http.Handler) {
	a.middleware = append(a.middleware, mw)
}
