package order

import "gorm.io/gorm"

type Schema struct {
	gorm.Model
	Order
}
