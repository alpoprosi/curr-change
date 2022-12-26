package db

import (
	"time"

	"github.com/alpoprosi/curr-change/internal/models"
)

func (db *DB) SaveBTCFiat(_ []models.BTCFiat) error {
	return nil
}

func (db *DB) HistoryBTCFiat(
	date *time.Time,
	currName string,
	offset int,
) ([]models.BTCFiat, int64, error) {
	return nil, 0, nil
}
