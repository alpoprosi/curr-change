package models

import "time"

type Fiat struct {
	ID       int64
	FiatID  string
	Value    float32
	CurrName string
	Date     time.Time
}

func (*Fiat) TableName() string {
	return "fiat"
}
