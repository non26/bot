package domain

import (
	"github.com/shopspring/decimal"
)

type HeikinAshiDomain struct {
	Candle0      *CandleDomain
	Candle1      *CandleDomain
	Open         string
	Close        string
	BotID        string
	ClientID     string
	Symbol       string
	PositionSide string
	// Side         string
	AmountBase string
	AccountID  string
	// TemplateID   string
}

func (h *HeikinAshiDomain) IsGrrenCandle() bool {
	_close, err := decimal.NewFromString(h.Close)
	if err != nil {
		return false
	}

	_open, err := decimal.NewFromString(h.Open)
	if err != nil {
		return false
	}

	return _close.GreaterThanOrEqual(_open)
}

func (h *HeikinAshiDomain) IsRedCandle() bool {
	_close, err := decimal.NewFromString(h.Close)
	if err != nil {
		return false
	}

	_open, err := decimal.NewFromString(h.Open)
	if err != nil {
		return false
	}
	return _close.LessThanOrEqual(_open)
}

func (h *HeikinAshiDomain) ToTradeDomain() *Trade {
	t := NewTrade()
	t.AccountID = h.AccountID
	t.Symbol = h.Symbol
	t.PositionSide = h.PositionSide
	// t.Side = h.Side
	t.Quantity = h.AmountBase
	t.ClientID = h.ClientID
	t.BnClientID = t.CreateBnClientID(h.ClientID)
	t.BotId = h.BotID
	// t.TemplateID = h.TemplateID
	// t.TradeType = h.TradeType
	return t
}

func (h *HeikinAshiDomain) OpenLongPosition() bool {
	return h.Candle1.IsRedCandle() && h.Candle0.IsGreenCandle()
}

func (h *HeikinAshiDomain) CloseLongPosition() bool {
	return h.Candle1.IsGreenCandle() && h.Candle0.IsRedCandle()
}

func (h *HeikinAshiDomain) OpenShortPosition() bool {
	return h.Candle1.IsGreenCandle() && h.Candle0.IsRedCandle()
}

func (h *HeikinAshiDomain) CloseShortPosition() bool {
	return h.Candle1.IsRedCandle() && h.Candle0.IsGreenCandle()
}

func (h *HeikinAshiDomain) ToBotDomain() *BotDomain {
	return &BotDomain{
		BotID:    h.BotID,
		ClientID: h.ClientID,
	}
}
