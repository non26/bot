package domain

import "github.com/shopspring/decimal"

var REDCOLOR = "red"
var GREENCOLOR = "green"

type CandleDomain struct {
	Open  string
	Close string
}

func NewCandleDomain() *CandleDomain {
	return &CandleDomain{}
}

func (c *CandleDomain) IsRedCandle() bool {
	close, _ := decimal.NewFromString(c.Close)
	open, _ := decimal.NewFromString(c.Open)
	return close.LessThan(open)
}

func (c *CandleDomain) IsGreenCandle() bool {
	close, _ := decimal.NewFromString(c.Close)
	open, _ := decimal.NewFromString(c.Open)
	return close.GreaterThan(open)
}

func (c *CandleDomain) CandleColor() string {
	if c.IsRedCandle() {
		return REDCOLOR
	}
	return GREENCOLOR
}
