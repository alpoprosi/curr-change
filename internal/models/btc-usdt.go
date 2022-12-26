package models

import "time"

type BTCUSDT struct {
	ID    int64
	Value float32
	Time  time.Time
}

func (*BTCUSDT) TableName() string {
	return "btc-usdt"
}
