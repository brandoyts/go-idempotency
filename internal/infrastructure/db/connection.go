package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Credentials struct {
	Host         string
	User         string
	Password     string
	DatabaseName string
	Port         string
}

func New(credentials *Credentials) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%v user=%v dbname=%v password=%v port=%v", credentials.Host, credentials.User, credentials.DatabaseName, credentials.Password, credentials.Port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
