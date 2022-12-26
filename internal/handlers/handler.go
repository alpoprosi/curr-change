package handlers

import (
	"time"

	"github.com/alpoprosi/curr-change/internal/db"
	"github.com/alpoprosi/curr-change/internal/models"
)

// DB interface defines database methods.
type DB interface {
	HistoryBTCUSDT(date *time.Time, currName string, offset int) ([]models.BTCUSDT, int64, error)

	HistoryBTCFiat(date *time.Time, currName string, offset int) ([]models.BTCFiat, int64, error)

	HistoryFiat(date *time.Time, currName string, offset int) ([]models.Fiat, int64, error)
}

// type check.
var _ DB = (*db.DB)(nil)

// Handler is the main structure that contains the database, and API handlers.
type Handler struct {
	db DB
}

// NewHandler returns new instance of Handler.
func NewHandler(db DB) Handler {
	return Handler{db: db}
}
