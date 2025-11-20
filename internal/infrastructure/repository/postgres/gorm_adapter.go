package postgres

import (
	"github.com/brandoyts/go-idempotency/internal/infrastructure/db"
	"gorm.io/gorm"
)

type GormDBAdapter struct {
	db *gorm.DB
}

func NewGormDBAdapter(db *gorm.DB) *GormDBAdapter {
	return &GormDBAdapter{db: db}
}

func (g *GormDBAdapter) Find(dest interface{}, conds ...interface{}) *db.DBResult {
	result := g.db.Find(dest, conds...)
	return &db.DBResult{Error: result.Error}
}

func (g *GormDBAdapter) First(dest interface{}, conds ...interface{}) *db.DBResult {
	result := g.db.First(dest, conds...)
	return &db.DBResult{Error: result.Error}
}

func (g *GormDBAdapter) Create(value interface{}) *db.DBResult {
	result := g.db.Create(value)
	return &db.DBResult{Error: result.Error}
}
