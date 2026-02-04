package domain

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
	return c.Close < c.Open
}

func (c *CandleDomain) IsGreenCandle() bool {
	return c.Close > c.Open
}

func (c *CandleDomain) CandleColor() string {
	if c.IsRedCandle() {
		return REDCOLOR
	}
	return GREENCOLOR
}
