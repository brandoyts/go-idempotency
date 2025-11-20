package main

import (
	"log"
	"os"

	"github.com/brandoyts/go-idempotency/internal/infrastructure/db"
	"github.com/brandoyts/go-idempotency/internal/infrastructure/repository/postgres"
	"github.com/brandoyts/go-idempotency/internal/infrastructure/rest"
	"github.com/brandoyts/go-idempotency/internal/service"
)

func main() {
	dbCredentials := &db.Credentials{
		Host:         os.Getenv("DB_HOST"),
		User:         os.Getenv("DB_USER"),
		Password:     os.Getenv("DB_PASSWORD"),
		DatabaseName: os.Getenv("DB_NAME"),
		Port:         os.Getenv("DB_PORT"),
	}

	gormDb, err := db.New(dbCredentials)
	if err != nil {
		log.Fatal("error connecting to db:", err)
	}

	adapter := postgres.NewGormDBAdapter(gormDb)

	migrate(gormDb)

	orderRepository := postgres.NewOrderRepository(adapter)
	orderService := service.NewOrderService(orderRepository)

	idempotencyKeyRepository := postgres.NewIdempotencyKeyRepository(adapter)
	idempotencyKeyService := service.NewIdempotencyKeyService(idempotencyKeyRepository)

	services := &rest.Services{
		OrderService:          orderService,
		IdempotencyKeyService: idempotencyKeyService,
	}

	server := rest.New(services)
	server.Start()
}
