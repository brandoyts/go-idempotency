package main

import (
	"log"

	idempotencyKey "github.com/brandoyts/go-idempotency/internal/core/idempotency_key"
	"github.com/brandoyts/go-idempotency/internal/core/order"
	"gorm.io/gorm"
)

func migrate(db *gorm.DB) {
	entities := []interface{}{
		&order.Order{},
		&idempotencyKey.IdempotencyKey{},
	}

	for _, entity := range entities {
		if err := db.AutoMigrate(entity); err != nil {
			log.Fatalf("failed to migrate entity %T: %v", entity, err)
		}
	}

	log.Println("Database migration completed successfully!")
}
