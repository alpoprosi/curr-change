package db

import (
	"fmt"
	"time"

	"github.com/alpoprosi/curr-change/internal/models"
)

func (db *DB) SaveBTCUSDT(bu models.BTCUSDT) error {
	err := db.Order("time DESC").
		FirstOrCreate(&bu).Error
	if err != nil {
		return fmt.Errorf("saving btc-usdt: %w", err)
	}

	return nil
}

func (db *DB) HistoryBTCUSDT(
	date *time.Time,
	currName string,
	offset int,
) ([]models.BTCUSDT, int64, error) {
	return nil, 0, nil
}
