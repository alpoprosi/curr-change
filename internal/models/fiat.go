package models

import "time"

type Fiat struct {
	ID       int64
	Fiat_ID  string
	Value    float32
	CurrName string
	Date     time.Time
}

func (*Fiat) TableName() string {
	return "fiat"
}
