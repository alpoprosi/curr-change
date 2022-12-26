package db

import (
	"fmt"

	gorm_logrus "github.com/onrik/gorm-logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	*gorm.DB
}

func NewDB(dsl string) (*DB, error) {
	db, err := gorm.Open(
		postgres.Open(dsl),
		&gorm.Config{Logger: gorm_logrus.New()},
	)
	if err != nil {
		return nil, fmt.Errorf("opening database dsl: %w", err)
	}

	return &DB{db}, nil
}
