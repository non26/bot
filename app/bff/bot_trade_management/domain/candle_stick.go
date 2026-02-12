package domain

import (
	bnconstant "github.com/non26/tradepkg/pkg/bn/bn_constant"
)

type CandleStickDomain struct {
	Candle0                *CandleDomain
	Candle1                *CandleDomain
	Symbol                 string
	PositionSide           string
	AmountB                string
	AccountID              string
	BotID                  string
	ClientID               string
	NumberOfBarRestriction *CandleStickRestriction
}

func (c *CandleStickDomain) ToTradeDomain() *Trade {
	t := NewTrade()
	t.AccountID = c.AccountID
	t.Symbol = c.Symbol
	t.PositionSide = c.PositionSide
	t.Quantity = c.AmountB
	t.ClientID = c.ClientID
	t.BnClientID = t.CreateBnClientID(c.ClientID)
	t.BotId = c.BotID
	return t
}

func (c *CandleStickDomain) OpenLongPosition() bool {
	return c.Candle1.IsRedCandle() && c.Candle0.IsGreenCandle()
}

func (c *CandleStickDomain) CloseLongPosition() bool {
	return c.Candle1.IsGreenCandle() && c.Candle0.IsRedCandle()
}

func (c *CandleStickDomain) OpenShortPosition() bool {
	return c.Candle1.IsGreenCandle() && c.Candle0.IsRedCandle()
}

func (c *CandleStickDomain) CloseShortPosition() bool {
	return c.Candle1.IsRedCandle() && c.Candle0.IsGreenCandle()
}

func (c *CandleStickDomain) IsBothCandleIsGreen() bool {
	return c.Candle1.IsGreenCandle() && c.Candle0.IsGreenCandle()
}

func (c *CandleStickDomain) IsBothCandleIsRed() bool {
	return c.Candle1.IsRedCandle() && c.Candle0.IsRedCandle()
}

func (c *CandleStickDomain) IsLongPosition() bool {
	return c.PositionSide == bnconstant.LONG
}

func (c *CandleStickDomain) IsShortPosition() bool {
	return c.PositionSide == bnconstant.SHORT
}

func (c *CandleStickDomain) ToBotDomain() *BotDomain {
	return &BotDomain{
		BotID:    c.BotID,
		ClientID: c.ClientID,
	}
}
