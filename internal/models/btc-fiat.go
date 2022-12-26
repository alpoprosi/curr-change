package models

type BTCFiat struct{}

func (*BTCFiat) TableName() string {
	return "btc-fiat"
}
