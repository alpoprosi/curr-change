package db

import (
	"fmt"
	"time"

	"github.com/alpoprosi/curr-change/internal/models"
)

func (db *DB) SaveFiat(f []models.Fiat) error {
	var errs []error
	for i, fiat := range f {
		err := db.Where(models.Fiat{FiatID: fiat.FiatID}).
			Order("date DESC").
			FirstOrCreate(&f[i]).Error
		if err != nil {
			aerr := fmt.Errorf("first or creating fiat %s: %w", f[i].FiatID, err)
			errs = append(errs, aerr)
		}
	}

	return collectErrors(errs)
}

func (db *DB) HistoryFiat(date *time.Time, currName string, offset int) ([]models.Fiat, int64, error) {
	var f []models.Fiat

	q := db.Model(&models.Fiat{})

	if date != nil {
		q = q.Where("date = ?", date)
	}

	if currName != "" {
		q = q.Where("curr_name = ?", currName)
	}

	c := int64(0)
	err := q.Offset(offset).Count(&c).Find(&f).Error
	if err != nil {
		return nil, 0, fmt.Errorf("selecting fiat: %w", err)
	}

	return f, c, nil
}
