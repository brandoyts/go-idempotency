package idempotencyKey

import "gorm.io/gorm"

type Schema struct {
	gorm.Model
	IdempotencyKey
}
